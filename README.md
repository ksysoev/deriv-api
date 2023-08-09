# Deriv Api Client

[![Tests](https://github.com/ksysoev/deriv-api/actions/workflows/main.yml/badge.svg)](https://github.com/ksysoev/deriv-api/actions/workflows/main.yml)
[![codecov](https://codecov.io/gh/ksysoev/deriv-api/branch/main/graph/badge.svg?token=V89HZZ5L9H)](https://codecov.io/gh/ksysoev/deriv-api)
[![Schema Updated](https://github.com/ksysoev/deriv-api/actions/workflows/schema.yml/badge.svg)](https://github.com/ksysoev/deriv-api/actions/workflows/schema.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/ksysoev/deriv-api)](https://goreportcard.com/report/github.com/ksysoev/deriv-api)
[![Go Reference](https://pkg.go.dev/badge/github.com/ksysoev/deriv-api.svg)](https://pkg.go.dev/github.com/ksysoev/deriv-api)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)


**This is unofficial library for working with [Deriv API](https://api.deriv.com)**

DerivAPI is a Go package that provides a client for the Deriv API. It allows you to connect to the Deriv API and make requests to retrieve data and perform actions on your or your client's Deriv account.

## Installation

To use DerivAPI in your Go project, you can install it using go get:

```
go get github.com/ksysoev/deriv-api
```

## Usage

### Tick stream

Here's an example of how to use DerivAPI to connect to the Deriv API and subscribe on market data updates:

```golang
import (
	"fmt"
	"log"

	"github.com/ksysoev/deriv-api"
	"github.com/ksysoev/deriv-api/schema"
)


api, err := deriv.NewDerivAPI("wss://ws.binaryws.com/websockets/v3", 1, "en", "http://example.com/")

if err != nil {
    log.Fatal(err)
}

defer api.Disconnect()

resp, sub, err := api.SubscribeTicks(schema.Ticks{Ticks: "R_50"})

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

Here's another example of how to use DerivAPI to buy contract and listen for updates for this contract.

```golang
import (
	"log"

	"github.com/ksysoev/deriv-api"
	"github.com/ksysoev/deriv-api/schema"
)

apiToken := "YOU_API_TOKEN_HERE"
api, err := deriv.NewDerivAPI("wss://ws.binaryws.com/websockets/v3", 1, "en", "https://www.binary.com")

if err != nil {
    log.Fatal(err)
}

defer api.Disconnect()

// First, we need to authorize the connection
reqAuth := schema.Authorize{Authorize: apiToken}
_, err = api.Authorize(reqAuth)

if err != nil {
    log.Fatal(err)
}

amount := 100.0
barrier := "+0.001"
duration := 5
basis := schema.ProposalBasisPayout

reqProp := schema.Proposal{
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

buyReq := schema.Buy{
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