package main

import (
	"log"

	"github.com/ksysoev/deriv-api"
	"github.com/ksysoev/deriv-api/schema"
)

const ApiToken = "YOUR_API_TOKEN_HERE" // Replace with your API token

func main() {
	api, err := deriv.NewDerivAPI("wss://ws.derivws.com/websockets/v3", 1089, "en", "https://localhost/")

	if err != nil {
		log.Fatal(err)
	}

	defer api.Disconnect()

	// First, we need to authorize the connection
	reqAuth := schema.Authorize{Authorize: ApiToken}
	_, err = api.Authorize(reqAuth)

	if err != nil {
		log.Fatal(err)
	}

	// Now we can send request for creating a new app

	scopes := []schema.AppRegisterScopesElem{
		schema.AppRegisterScopesElemRead,
		schema.AppRegisterScopesElemTrade,
	}
	// Create a new app
	reqAppReg := schema.AppRegister{
		AppRegister: 1,
		Name:        "Example App",
		Scopes:      scopes,
	}

	// Optional fields

	// Redirect URI for OAuth2 authorization
	// This needed for client facing apps that user can login to
	// For backend scripts you can use just API tokens

	// redirectUri := "https://example.com/auth_callback"
	// reqAppReg.RedirectUri = &redirectUri

	// Verification URI for email verifications

	// verificationUri := "https://example.com/verify_callback"
	// reqAppReg.VerificationUri = &verificationUri

	// Markup percentage for app
	// This is the percentage that will be added to the payout

	// markupPercentage := 1.0 // 1%
	// reqAppReg.AppMarkupPercentage = &markupPercentage

	// General app info, also optional

	// Homepage link
	// homepage := "https://example.com"
	// reqAppReg.Homepage = &homepage

	// Appstore link
	// appstore := "https://itunes.apple.com/example"
	// reqAppReg.Appstore = &appstore

	// Google Playstore link
	// googleplay := "https://play.google.com/store/apps/details?id=example"
	// reqAppReg.Googleplay = &googleplay

	// Github link
	// github := "https://github.com/example/example"
	// reqAppReg.Github = &github

	resp, err := api.AppRegister(reqAppReg)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("App created with id %v", resp.AppRegister.AppId)
}
