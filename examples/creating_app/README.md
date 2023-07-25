# Create Deriv App

This Go script demonstrates how to create a new Deriv app using the Deriv API. The script uses the `deriv-api` and `deriv-api/schema` packages to interact with the Deriv API and send requests to create a new app.

To use this script, you will need to replace the `ApiToken` constant with your own API token. You can obtain an API token by logging in to your Deriv account and generating a new token in the API tokens section.

The script first creates a new `DerivAPI` instance and authorizes the connection using the provided API token. It then creates a new app registration request with a name and a set of scopes. The script also includes optional fields for specifying a redirect URI, verification URI, markup percentage, and general app info.

Finally, the script sends the app registration request to the Deriv API and logs the ID of the newly created app. You can customize the app registration request and the optional fields to suit your specific needs.

Note that this script is provided as an example and may require modifications to work with your specific environment and use case.