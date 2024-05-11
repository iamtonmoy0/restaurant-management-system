# Restaurant Management System

The Restaurant Management System is a Go-based application designed to streamline restaurant operations. It provides functionalities to manage food items, menus, tables, orders, invoices, and user accounts.

## Features

- **Food Management**: Add, update, delete food items with their prices and images.
- **Menu Management**: Create and manage menus with categories and date ranges.
- **Table Management**: Track table availability and number of guests.
- **Order Management**: Place orders for customers and track order status.
- **Invoice Management**: Generate invoices and manage payment status.
- **User Management**: Register users, manage their accounts, and handle authentication.

## Technologies Used

- **Go (Golang)**: The backend of the application is developed in Go, leveraging its concurrency features and performance.
- **MongoDB**: MongoDB is used as the database to store restaurant data.
- **Gin**: Gin is used as the HTTP web framework for building the RESTful API.
- **JWT Authentication**: JSON Web Tokens (JWT) are used for user authentication.
- **Validator**: Input validation is implemented using the validator package.
- **MongoDB Go Driver**: The official MongoDB Go driver is used to interact with the MongoDB database.

## Installation

1. **Clone the Repository**:

    ```bash
    git clone https://github.com/iamtonmoy0/restaurant-management-system.git
    ```

2. **Install Dependencies**:

    ```bash
    go mod tidy
    ```

3. **Run the Application**:

    ```bash
    go run main.go
    ```

## Usage

- **API Documentation**: Explore the API endpoints and their usage in the [API Documentation](api-doc.md).

## Contributors

- [Tonmoy](https://github.com/iamtonmoy0)

## License

This project is licensed under the [MIT License](LICENSE).
