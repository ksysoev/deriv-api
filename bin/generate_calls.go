package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

const (
	// API_CALLS is the file name for the api calls
	API_CALLS = "./calls.go"

	// SUBSCRIPTION_CALLS is the file name for the subscription calls
	SUBSCRIPTION_CALLS = "./subscription_calls.go"

	// SCHEMA_PATH is the path to the schema files
	SCHEMA_PATH = "./deriv-developers-portal/config/v3/"
)

func main() {
	files, err := ioutil.ReadDir(SCHEMA_PATH)
	if err != nil {
		log.Fatal(err)
	}

	apiCalls := "// Code generated by bin/generate_call.go, DO NOT EDIT.\n\npackage deriv\n\n"
	subcriptionCalls := "// Code generated by bin/generate_call.go, DO NOT EDIT.\n\npackage deriv\n\n"

	for _, file := range files {
		name := file.Name()
		request, err := ioutil.ReadFile("deriv-developers-portal/config/v3/" + name + "/send.json")
		if err != nil {
			log.Fatal(err)
		}

		var requestSchema map[string]any
		err = json.Unmarshal(request, &requestSchema)

		if err != nil {
			log.Fatal(err)
		}

		title := CreateTitle(name)

		description, _ := requestSchema["description"].(string)

		apiCalls += CreateApiCallFunction(title, description)

		if HasSubscribe(requestSchema) {
			subcriptionCalls += CreateSubscriptionCallFunction(title, description)
		}
	}

	ioutil.WriteFile(API_CALLS, []byte(apiCalls), 0644)
	ioutil.WriteFile(SUBSCRIPTION_CALLS, []byte(subcriptionCalls), 0644)
}

func HasSubscribe(r map[string]any) bool {
	rawProperties, ok := r["properties"]
	if !ok {
		return false
	}

	properties, ok := rawProperties.(map[string]any)
	if !ok {
		return false
	}
	_, ok = properties["subscribe"]

	return ok
}

func CreateTitle(name string) string {
	var title string
	for _, s := range strings.Split(name, "_") {
		if s == "p2p" {
			s = "P2P"
		}
		title = title + strings.Title(s)
	}

	return title
}

func CreateApiCallFunction(title string, description string) string {
	signature := fmt.Sprintf("// %s %s\nfunc (a *DerivAPI) %s(r %s) (resp %sResp, err error)", title, description, title, title, title)

	body := fmt.Sprintf(
		`id := a.getNextRequestID()
	r.ReqId = &id
	err = a.SendRequest(id, r, &resp)
	return`)

	return fmt.Sprintf("%s {\n\t%s\n}\n\n", signature, body)
}

func CreateSubscriptionCallFunction(title string, description string) string {
	signature := fmt.Sprintf("// Subscribe%s %s\nfunc (a *DerivAPI) Subscribe%s(r %s) (s *Subsciption[%sResp], err error)", title, description, title, title, title)

	var body string
	if title == "Transaction" {
		body = fmt.Sprintf(
			`id := a.getNextRequestID()
	var f %sSubscribe = 1
	r.ReqId = &id
	r.Subscribe = f
	s = NewSubcription[%sResp](a)
	err = s.Start(id, r)
	return`, title, title)
	} else {
		body = fmt.Sprintf(
			`id := a.getNextRequestID()
	var f %sSubscribe = 1
	r.ReqId = &id
	r.Subscribe = &f
	s = NewSubcription[%sResp](a)
	err = s.Start(id, r)
	return`, title, title)
	}

	return fmt.Sprintf("%s {\n\t%s\n}\n\n", signature, body)
}