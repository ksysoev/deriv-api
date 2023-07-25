# Deriv Binary Options Trading Bot

This Go script demonstrates how to use the Deriv API to create a simple trading bot for binary options. The script uses the `deriv-api` and `deriv-api/schema` packages to interact with the Deriv API and send requests to create a new proposal and buy a binary option.

To use this script, you will need to replace the `ApiToken` constant with your own API token. You can obtain an API token by logging in to your Deriv account and generating a new token in the API tokens section.

The script first creates a new `DerivAPI` instance and authorizes the connection using the provided API token. It then creates a new proposal request with a specified amount, barrier, duration, basis, contract type, currency, and symbol.

The script then sends the proposal request to the Deriv API and waits for a response. Once a proposal is received, the script creates a new buy request with the proposal ID and a specified price.

The script then subscribes to the buy request and waits for a response. Once the binary option is bought, the script logs the current spot and the contract result.

Note that this script is provided as an example and may require modifications to work with your specific environment and use case. It is important to thoroughly test any trading bot before using it with real funds.