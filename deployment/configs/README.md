# Environment Files for Deployment

This folder contains environment configuration files for different services. Please ensure to update the following placeholders in each `.env` file with the correct values before deploying your application:

---

## `frontend-gateway.env`

This file contains configuration for the frontend gateway service.

- **`<frontend_host>`**: The hostname or IP address of the frontend service.
- **`<frontend_port>`**: The port number for the frontend service (default is 80).
- **`<exchange_name>`**: The name of the user exchange in RabbitMQ.
- **`<rabbitmq_user>`**: The username for accessing RabbitMQ.
- **`<rabbitmq_password>`**: The password for accessing RabbitMQ.
- **`<rabbitmq_host>`**: The hostname or IP address of the RabbitMQ service.
- **`<rabbitmq_port>`**: The port number for RabbitMQ (default is 5672).

---

## `mongo-express.env`

This file contains configuration for Mongo Express (a web-based MongoDB admin interface).

- **`<mongodb_admin_user>`**: The username for MongoDB admin access (default is "root").
- **`<mongodb_admin_password>`**: The password for MongoDB admin access (default is "password").
- **`<mongodb_host>`**: The hostname or IP address of the MongoDB service.
- **`<mongodb_port>`**: The port number for MongoDB (default is 27017).
- **`<true_or_false>`**: Whether to enable basic authentication for Mongo Express (`true` or `false`).

---

## `mongodb.env`

This file contains configuration for initializing the MongoDB root user and password.

- **`<mongodb_root_user>`**: The root username for MongoDB (default is "root").
- **`<mongodb_root_password>`**: The root password for MongoDB (default is "password").

---

## `network-info.env`

This file contains network-related configuration for different services in the project.

- **`<frontend_gateway_project_name>`**: The name of the frontend gateway service project (e.g., "frontend-gateway").
- **`<frontend_gateway_ip>`**: The IP address or hostname of the frontend gateway service.
- **`<frontend_gateway_port>`**: The port number for the frontend gateway service (default is 5000).
  
- **`<user_service_project_name>`**: The name of the user service project (e.g., "user-service").
- **`<user_service_ip>`**: The IP address or hostname of the user service.
- **`<user_service_port>`**: The port number for the user service (default is 5001).

- **`<video_manager_project_name>`**: The name of the video manager project (e.g., "video-manager").
- **`<video_manager_ip>`**: The IP address or hostname of the video manager service.
- **`<video_manager_port>`**: The port number for the video manager service (default is 5002).

- **`<frontend_ip>`**: The IP address or hostname of the frontend service.
- **`<frontend_port>`**: The port number for the frontend service (default is 4200).

---

## `rabbitmq.env`

This file contains configuration for RabbitMQ user credentials.

- **`<rabbitmq_user>`**: The default RabbitMQ username (default is "admin").
- **`<rabbitmq_password>`**: The default RabbitMQ password (default is "admin").

---

## `user-service.env`

This file contains configuration for the user service.

- **`<db_name>`**: The name of the database for the user service (default is "users").
- **`<db_user>`**: The username for accessing the MongoDB database (default is "root").
- **`<db_password>`**: The password for the MongoDB database (default is "password").
- **`<mongodb_host>`**: The hostname or IP address of the MongoDB service.
- **`<rabbitmq_user>`**: The username for accessing RabbitMQ (default is "admin").
- **`<rabbitmq_password>`**: The password for accessing RabbitMQ (default is "admin").
- **`<rabbitmq_host>`**: The hostname or IP address of the RabbitMQ service.
- **`<rabbitmq_port>`**: The port number for RabbitMQ (default is 5672).

---

Make sure to update each `.env` file with the appropriate values specific to your deployment environment.
