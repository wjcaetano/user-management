# User Management Microservice

The **User Management Microservice** is a core component of a scalable, microservices-based platform designed to handle user registration, authentication, and profile management. Built with **Golang**, it follows **clean architecture principles** and leverages **Kafka** for event-driven communication.

---

## Features

- **User Registration**: Secure registration with data validation.
- **Authentication**: Implements JWT-based authentication.
- **Profile Management**: Role-based access for managing user profiles.
- **Event-Driven Design**: Kafka integration for asynchronous communication.
- **RESTful API**: Clean and well-documented endpoints.
- **Database Integration**: Uses MySQL with GORM for database operations.
- **Testing**: Includes unit and integration tests.
- **Dockerized Setup**: Fully containerized for easy deployment.

---

## Technologies

- **Language**: Golang
- **Database**: MySQL (via GORM)
- **Message Broker**: Apache Kafka
- **Framework**: Echo Framework
- **Authentication**: JSON Web Tokens (JWT)
- **Containerization**: Docker & Docker Compose
- **Documentation**: Swagger

---

## Getting Started

### Prerequisites

Ensure you have the following installed:
- [Golang](https://golang.org/dl/)
- [Docker](https://www.docker.com/products/docker-desktop)
- [Docker Compose](https://docs.docker.com/compose/)
- [MySQL](https://www.mysql.com/)
- [Kafka](https://kafka.apache.org/)

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/user-management.git
   cd user-management
   ```

2. Build the Docker containers:
   ```bash
   make fs
   ```

4. Start the application:
   ```bash
   go run cmd/api/main.go
   ```

### Environment Variables

Create a `.env` file in the root directory and configure the following variables:

```env
DATABASE_URL=mysql://user:password@tcp(localhost:3306)/user_management
JWT_SECRET=your_jwt_secret
KAFKA_BROKERS=localhost:9092
``` 

---

## Usage

### API Endpoints

#### Authentication
- **POST** `/auth/login`: Authenticate a user and return a JWT.

#### User Management
- **POST** `/users`: Register a new user.
- **GET** `/users`: List all users (admin only).
- **PUT** `/users/{id}`: Update user details.
- **DELETE** `/users/{id}`: Delete a user.

### Kafka Topics
- `user-created`: Publishes events for newly created users.
- `user-updated`: Publishes events for updated user data.
- `user-deleted`: Publishes events for deleted users.

---

## Testing

Run unit tests with:
```bash
go test ./...
```

---

## Contributing

Contributions are welcome! Please follow these steps:
1. Fork the repository.
2. Create a feature branch.
3. Commit your changes.
4. Submit a pull request.

---

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

## Contact

For questions or support, please contact: [caetano.willian@gmail.com](mailto:caetano.willian@gmail.com)
