#!/bin/sh
git clone https://github.com/binary-com/deriv-developers-portal.git
pwd
cd deriv-developers-portal/config/v3/
pwd
mkdir schema
find . -type f | grep -v example | grep -v schema |  sed 'p; s/^\./.\/schema/; s/\/send/_request/; s/\/receive/_response/g' | xargs -n2 cp -f

# Patch the schema for transaction response... Live without types so much easier
sed  's/9999999999999999999/9223372036854775807/g' schema/transaction_response.json > schema/transaction_response_tmp.json
mv schema/transaction_response_tmp.json schema/transaction_response.json
sed  's/"5000"/5000/g' schema/ticks_history_request.json > schema/ticks_history_request_tmp.json
mv schema/ticks_history_request_tmp.json schema/ticks_history_request.json

gojsonschema -p main schema/*.json > ../../../shema.go 
cd ../../../
rm -rf deriv-developers-portal schema