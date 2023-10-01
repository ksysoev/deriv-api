# .PHONY: all
.DEFAULT_GOAL := generate_all

generate_all: clone generate-calls prepare-schema generate-schema clean

clone:
	git clone https://github.com/binary-com/deriv-developers-portal.git

generate-calls: 
	go run bin/generate_calls.go

prepare-schema:
	cd deriv-developers-portal/ && \
	git apply ../schema.patch && \
	cd config/v3/ && \
	mkdir schema && \
	find . -type f | grep -v example | grep -v schema | sed 'p; s/^\./.\/schema/; s/\/send//; s/\/receive/_resp/; s/\.json//;' | xargs -n2 cp -f

generate-schema:
	cd deriv-developers-portal/config/v3/schema && \
	for file in *; do gojsonschema --tags json -p schema "$$file" > "../../../../schema/$${file%.*}.go"; done

clean:
	rm -rf deriv-developers-portal

test:
	go test -v --race .

coverage:
	go test -covermode=count -coverprofile=coverage.out .
	cat coverage.out | grep -v "/calls.go" | grep -v "subscription_calls.go" > coverage.final.out
	go tool cover -func=coverage.final.out
	rm coverage.out coverage.final.out
