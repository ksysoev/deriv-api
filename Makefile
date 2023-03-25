# .PHONY: all
.DEFAULT_GOAL := generate_all

generate_all: clone generate-calls prepare-schema generate-schema clean

clone:
	git clone https://github.com/binary-com/deriv-developers-portal.git

generate-calls: 
	go run bin/generate_calls.go

prepare-schema:
	cd deriv-developers-portal/config/v3/ && \
	mkdir schema && \
	find . -type f | grep -v example | grep -v schema | sed 'p; s/^\./.\/schema/; s/\/send//; s/\/receive/_resp/g' | xargs -n2 cp -f && \
	sed  's/9999999999999999999/9223372036854775807/g' schema/transaction_resp.json > schema/transaction_resp_tmp.json && \
	mv schema/transaction_resp_tmp.json schema/transaction_resp.json && \
	sed  's/"5000"/5000/g' schema/ticks_history.json > schema/ticks_history_tmp.json && \
	mv schema/ticks_history_tmp.json schema/ticks_history.json

generate-schema:
	cd deriv-developers-portal/config/v3/schema && \
	gojsonschema -p deriv *.json > ../../../schema.go

clean:
	rm -rf deriv-developers-portal

test:
	go test -v  .

coverage:
	go test -covermode=count -coverprofile=coverage.out ./...
	cat coverage.out | grep -v "/schema.go" | grep -v "/calls.go" | grep -v "subscription_calls.go" > coverage.final.out
	go tool cover -func=coverage.final.out
	rm coverage.out coverqage.final.out
