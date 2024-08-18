package main

import (
	"context"
	"log"

	"github.com/ksysoev/deriv-api"
	"github.com/ksysoev/deriv-api/schema"
)

const ApiToken = "YOUR_API_TOKEN_HERE" // Replace with your API token

func main() {
	ctx := context.Background()
	api, err := deriv.NewDerivAPI("wss://ws.binaryws.com/websockets/v3", 1089, "en", "https://localhost/")

	if err != nil {
		log.Fatal(err)
	}

	defer api.Disconnect()

	// First, we need to authorize the connection
	reqAuth := schema.Authorize{Authorize: ApiToken}
	_, err = api.Authorize(ctx, reqAuth)

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
		ContractType: schema.ProposalContractTypeCALL,
		Currency:     "USD",
		Duration:     &duration,
		DurationUnit: schema.ProposalDurationUnitT,
		Symbol:       "R_50",
	}

	// Send a proposal request
	proposal, err := api.Proposal(ctx, reqProp)

	if err != nil {
		log.Fatal(err)
	}

	buyReq := schema.Buy{
		Buy:   proposal.Proposal.Id,
		Price: 100.0,
	}

	_, buySub, err := api.SubscribeBuy(ctx, buyReq)

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
}
