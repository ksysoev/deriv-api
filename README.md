# Deriv Api Client

[![Tests](https://github.com/ksysoev/deriv-api/actions/workflows/main.yml/badge.svg)](https://github.com/ksysoev/deriv-api/actions/workflows/main.yml)

[![Schema Updated](https://github.com/ksysoev/deriv-api/actions/workflows/schema.yml/badge.svg)](https://github.com/ksysoev/deriv-api/actions/workflows/schema.yml)

This is unofficial library for working with [Deriv API](https://api.deriv.com) 

## Installation

To install deriv-api, use `go get`:

```
go get github.com/ksysoev/deriv-api
```

## Usage

```golang
api, err := deriv.NewDerivAPI("wss://ws.binaryws.com/websockets/v3", 1, "en", "http://example.com/")

if err != nil {
    log.Fatal(err)
}

defer api.Disconnect()

sub, err := api.SubscribeTicks(deriv.Ticks{Ticks: "R_50"})

if err != nil {
    log.Fatal(err)
    return
}

for tick := range sub.Stream {
    fmt.Println("Symbol: ", *tick.Tick.Symbol, "Quote: ", *tick.Tick.Quote)
}
```
