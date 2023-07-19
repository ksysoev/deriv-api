# Deriv Api Client

[![Tests](https://github.com/ksysoev/deriv-api/actions/workflows/main.yml/badge.svg)](https://github.com/ksysoev/deriv-api/actions/workflows/main.yml)
[![codecov](https://codecov.io/gh/ksysoev/deriv-api/branch/main/graph/badge.svg?token=V89HZZ5L9H)](https://codecov.io/gh/ksysoev/deriv-api)
[![Schema Updated](https://github.com/ksysoev/deriv-api/actions/workflows/schema.yml/badge.svg)](https://github.com/ksysoev/deriv-api/actions/workflows/schema.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/ksysoev/deriv-api)](https://goreportcard.com/report/github.com/ksysoev/deriv-api)
[![Go Reference](https://pkg.go.dev/badge/github.com/ksysoev/deriv-api.svg)](https://pkg.go.dev/github.com/ksysoev/deriv-api)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

This is unofficial library for working with [Deriv API](https://api.deriv.com) 

## Installation

To install deriv-api, use `go get`:

```
go get github.com/ksysoev/deriv-api
```

## Usage

### Subscribing on tick stream

```golang
import (
	"fmt"
	"log"

	"github.com/ksysoev/deriv-api"
)


api, err := deriv.NewDerivAPI("wss://ws.binaryws.com/websockets/v3", 1, "en", "http://example.com/")

if err != nil {
    log.Fatal(err)
}

defer api.Disconnect()

resp, sub, err := api.SubscribeTicks(deriv.Ticks{Ticks: "R_50"})

if err != nil {
    log.Fatal(err)
    return
}

fmt.Println("Symbol: ", *resp.Tick.Symbol, "Quote: ", *resp.Tick.Quote)

for tick := range sub.Stream {
    fmt.Println("Symbol: ", *tick.Tick.Symbol, "Quote: ", *tick.Tick.Quote)
}
```

### Buying a contract

```golang
import (
	"log"

	"github.com/ksysoev/deriv-api"
)

apiToken := "YOU_API_TOKEN_HERE"
api, err := deriv.NewDerivAPI("wss://ws.binaryws.com/websockets/v3", 1, "en", "https://www.binary.com")

if err != nil {
    log.Fatal(err)
}

defer api.Disconnect()

// First, we need to authorize the connection
reqAuth := deriv.Authorize{Authorize: apiToken}
_, err = api.Authorize(reqAuth)

if err != nil {
    log.Fatal(err)
}

amount := 100.0
barrier := "+0.001"
duration := 5
basis := deriv.ProposalBasisPayout

reqProp := deriv.Proposal{
    Proposal:     1,
    Amount:       &amount,
    Barrier:      &barrier,
    Basis:        &basis,
    ContractType: deriv.ProposalContractTypeCALL,
    Currency:     "USD",
    Duration:     &duration,
    DurationUnit: deriv.ProposalDurationUnitT,
    Symbol:       "R_50",
}

// Send a proposal request
proposal, err := api.Proposal(reqProp)

if err != nil {
    log.Fatal(err)
}

buyReq := deriv.Buy{
    Buy:   proposal.Proposal.Id,
    Price: 100.0,
}

_, buySub, err := api.SubscribeBuy(buyReq)

if err != nil {
    log.Fatal(err)
}

for proposalOpenContract := range buySub.Stream {
    log.Println("Current Spot: ", *proposalOpenContract.ProposalOpenContract.CurrentSpot)
    if *proposalOpenContract.ProposalOpenContract.IsSold == 1 {
        log.Println("Contract Result: ", proposalOpenContract.ProposalOpenContract.Status.Value)
        buySub.Forget()
        break
    }
}
```