
# Table: aws_apigateway_rest_api_gateway_responses
A gateway response of a given response type and status code, with optional response parameters and mapping templates
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|rest_api_cq_id|uuid|Unique CloudQuery ID of aws_apigateway_rest_apis table (FK)|
|rest_api_id|text|The API's identifier|
|arn|text|The Amazon Resource Name (ARN) for the resource|
|default_response|boolean|A Boolean flag to indicate whether this GatewayResponse is the default gateway response (true) or not (false)|
|response_parameters|jsonb|Response parameters (paths, query strings and headers) of the GatewayResponse as a string-to-string map of key-value pairs|
|response_templates|jsonb|Response templates of the GatewayResponse as a string-to-string map of key-value pairs|
|response_type|text|The response type of the associated GatewayResponse|
|status_code|text|The HTTP status code for this GatewayResponse|
