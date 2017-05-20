# termrr (for terminal MRR)

I wanted to have a quick answer to the question: "What's our current MRR" for [Roadmap](https://roadmap.space).

I kind of pass the majority of my time in the terminal, so I built this simple tool. This is a sample of the output:

```bash
MRR is  425.00
Month over month stats
=====================================
2017-05 New customers: 1 MRR: 25.00
2017-04 New customers: 1 MRR: 50.00
2017-03 New customers: 1 MRR: 55.00
2017-02 New customers: 2 MRR: 80.00
2017-01 New customers: 1 MRR: 95.00
2016-12 New customers: 2 MRR: 120.00
```

It's showing the MRR and month over month new customers and added MRR for that month.

## Installation

### Binary
You can download a binary for your platform in the [Releases](https://github.com/dstpierre/termrr/releases) page.

### If you have Go installed

```bash
go get https://github.com/dstpierre/termrr
```

## Usage

You only need to pass a [Stripe](https://stripe.com) key (live or test).

```bash
termrr -key stripe_key_here
```

You may also have an environment variable `STRIPE_KEY`.

```bash
export STRIPE_KEY=sk_test_1234567
termrr
```