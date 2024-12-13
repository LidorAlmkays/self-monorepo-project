# Gateway Service

The **Gateway Service** acts as the central entry point for all requests from the frontend, ensuring seamless routing to the appropriate backend services. Its primary role is to route requests to the correct backend service based on the endpoint, and it does this by coordinating with other services such as the authentication service to validate user credentials.

## Features

- **Request Routing**: Forwards requests to the correct backend service based on the endpoint.
- **Coordination with Authentication Service**: Relies on the authentication service to validate user credentials before forwarding requests.
- **Centralized Access Point**: Serves as a single interface for frontend applications to interact with backend services.

## Request Flow

1. **Frontend Request**: The frontend sends a request to the gateway.
2. **Authentication Check**: The gateway communicates with the authentication service to validate user credentials.
3. **Routing**: Once authenticated, the gateway routes the request to the appropriate backend service.
4. **Response**: The backend service processes the request and sends the response back through the gateway to the frontend.

## Example Endpoints

- `/api/auth/login` -> Routes to `authService`.
- `/api/user/profile` -> Routes to `userService`.
- `/api/product/list` -> Routes to `productService`.

## Security

- **Coordination with Authentication Service**: The gateway relies on a dedicated authentication service to ensure that all requests are authenticated.
- **Request Validation**: Invalid requests are rejected, and users are prompted to reauthenticate.

## License
This project is licensed under the MIT License.

