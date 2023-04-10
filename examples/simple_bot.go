package main

import (
	"log"

	"github.com/ksysoev/deriv-api"
)

const ApiToken = "YOUR_API_TOKEN_HERE" // Replace with your API token

func main() {
	api, err := deriv.NewDerivAPI("wss://ws.binaryws.com/websockets/v3", 1, "en", "https://www.binary.com")

	if err != nil {
		log.Fatal(err)
	}

	defer api.Disconnect()

	// First, we need to authorize the connection
	reqAuth := deriv.Authorize{Authorize: ApiToken}
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
}
