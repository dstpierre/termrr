package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"sort"

	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/customer"
)

type monthstats struct {
	revenue   float64
	customers int
	key       string
}

func main() {
	key := flag.String("key", "", "this is your stripe key, you may use an environment variable STRIPE_KEY.")
	flag.Parse()

	apiKey := os.Getenv("STRIPE_KEY")

	if *key != "" {
		apiKey = *key
	}

	if apiKey == "" {
		flag.Usage()
		return
	}

	stripe.Key = apiKey
	stripe.LogLevel = 1

	var mrr float64

	p := &stripe.CustomerListParams{}
	p.Filters.AddFilter("limit", "", "50")

	months := make(map[string]monthstats)

	customers := customer.List(p)
	for customers.Next() {
		c := customers.Customer()

		if c.Subs.Count > 0 {
			subs := c.Subs.Values

			var revenue float64

			for _, s := range subs {
				if s.Plan.Interval == "year" {
					revenue += float64(s.Plan.Amount) / 12.0
				} else {
					revenue += float64(s.Plan.Amount * s.Quantity)
				}

				if c.Discount != nil && c.Discount.Coupon != nil {
					revenue = applyCoupon(revenue, c.Discount.Coupon)
				}

				if s.Discount != nil && s.Discount.Coupon != nil {
					revenue = applyCoupon(revenue, s.Discount.Coupon)
				}
			}

			mrr += revenue / 100.0

			// logging monthly stats
			d := time.Unix(c.Created, 0)
			monthKey := d.Format("2006-01")

			if stats, ok := months[monthKey]; ok {
				stats.revenue += revenue / 100.0
				stats.customers++
				months[monthKey] = stats
			} else {
				stats = monthstats{customers: 1, revenue: revenue / 100.0, key: monthKey}
				months[monthKey] = stats
			}
		} else {
			log.Println("this customer has no subs")
		}
	}

	var keys []string
	for _, v := range months {
		keys = append(keys, v.key)
	}

	sort.Strings(keys)

	fmt.Printf("MRR is\t%.2f\n", mrr)
	fmt.Println("Month over month stats\n=====================================")
	for i := len(keys) - 1; i >= 0; i-- {
		k := keys[i]
		fmt.Println(k, fmt.Sprintf("New customers: %d", months[k].customers), fmt.Sprintf("MRR: %.2f", months[k].revenue))
	}
}

func applyCoupon(v float64, coupon *stripe.Coupon) float64 {
	if coupon.Amount > 0 {
		return v - float64(coupon.Amount)
	} else if coupon.Percent > 0 {
		return v - (float64(coupon.Percent) / 100.0)
	}
	return v
}
