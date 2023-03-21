#!/bin/sh
git clone https://github.com/binary-com/deriv-developers-portal.git
cd deriv-developers-portal/config/v3/
ls | go run ../../../bin/generate_calls.go > ../../../calls.go
mkdir schema
find . -type f | grep -v example | grep -v schema |  sed 'p; s/^\./.\/schema/; s/\/send//; s/\/receive/_resp/g' | xargs -n2 cp -f

# Patch the schema for transaction response... Live without types so much easier
sed  's/9999999999999999999/9223372036854775807/g' schema/transaction_resp.json > schema/transaction_resp_tmp.json
mv schema/transaction_resp_tmp.json schema/transaction_resp.json
sed  's/"5000"/5000/g' schema/ticks_history.json > schema/ticks_history_tmp.json
mv schema/ticks_history_tmp.json schema/ticks_history.json

gojsonschema -p deriv schema/*.json > ../../../schema.go 
cd ../../../
rm -rf deriv-developers-portal schema