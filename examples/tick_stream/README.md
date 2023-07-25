#Deriv Ticks Stream

This Go script demonstrates how to use the Deriv API to subscribe to the ticks stream for a specified symbol. The script uses the `deriv-api` and `deriv-api/schema` packages to interact with the Deriv API and send requests to subscribe to the ticks stream.

The script first creates a new `DerivAPI` instance and authorizes the connection using the provided API token. It then creates a new ticks subscription request with a specified symbol.

The script then sends the subscription request to the Deriv API and waits for a response. Once a response is received, the script logs the initial tick for the specified symbol.

The script then subscribes to the ticks stream and waits for new ticks to be received. Once a new tick is received, the script logs the symbol and quote for the tick.

Note that this script is provided as an example and may require modifications to work with your specific environment and use case. It is important to thoroughly test any code that interacts with the Deriv API before using it with real funds.