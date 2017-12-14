# go-onion

Attempting clean architecture with golang. Each services will be split into a folder with the suffix `-service` (e.g. `food-service`).

Each folder will contain the following files:

1. `model.go` which contains business logic and model validation. Calls the `Store` which is an API to data source such as Database or external HTTP Request.
2. `store.go` which is basically the Data Access Layer. Connects to the Database and does not contain any business logic. Provides an interface that allows itself to be mocked.
3. `route.go` which is the HTTP Endpoint for the service. Is responsible for constructing the request params for the service and returning the response or error after calling the service. Dumb endpoint, as it does not contain any business logic. Makes it easier to swap transport (GraphQL, gRPC, cron, etc) since the core logic is contained in the `Model`.
4. `main.go` which binds the `Model`, `Store` and `Route` together through dependency injection.
5. `interface.go` which contains the interface for each component and also define the request and response structs.