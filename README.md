# Go REST API with JWT Authentication

This repository contains a Go-based REST API that implements user registration, login, payments, balance top-ups, and transfers between users, all secured with JWT authentication. The application uses PostgreSQL as the database and GORM as the ORM library.

## API Overview

### 1. User Registration
Allows users to create an account.

#### How to Run:
```bash
POST /register
```

#### Request Body:
```json
{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "password123"
}
```

#### Response:
```json
{
  "status": "SUCCESS",
  "message": "User registered successfully"
}
```

### 2. User Login
Allows users to log in and receive a JWT token for authentication.

#### How to Run:
```bash
POST /login
```

#### Request Body:
```json
{
  "email": "john@example.com",
  "password": "password123"
}
```

#### Response:
```json
{
  "status": "SUCCESS",
  "access_token": "your-jwt-token"
}
```

### 3. Top Up Balance
Authenticated users can add funds to their account.

#### How to Run:
```bash
POST /topup
```

#### Request Headers:
```bash
Authorization: Bearer <JWT_TOKEN>
```

#### Request Body:
```json
{
  "amount": 50000
}
```

#### Response:
```json
{
  "status": "SUCCESS",
  "result": {
      "amount": 50000,
      "balance_before": 100000,
      "balance_after": 150000
  }
}
```

### 4. Payment
Authenticated users can make a payment, deducting from their balance.

#### How to Run:
```bash
POST /payment
```

#### Request Headers:
```bash
Authorization: Bearer <JWT_TOKEN>
```

#### Request Body:
```json
{
  "amount": 50000,
  "remarks": "Payment for services"
}
```

#### Response:
```json
{
  "status": "SUCCESS",
  "result": {
      "amount": 50000,
      "balance_before": 150000,
      "balance_after": 100000
  }
}
```

### 5. Transfer Money
Authenticated users can transfer funds to other users.

#### How to Run:
```bash
POST /transfer
```

#### Request Headers:
```bash
Authorization: Bearer <JWT_TOKEN>
```

#### Request Body:
```json
{
  "target_user": "uuid-of-target-user",
  "amount": 50000,
  "remarks": "Transfer to friend"
}
```

#### Response:
```json
{
  "status": "SUCCESS",
  "sender_transaction": {
      "amount": 50000,
      "balance_before": 100000,
      "balance_after": 50000
  },
  "receiver_transaction": {
      "amount": 50000,
      "balance_before": 20000,
      "balance_after": 70000
  }
}
```

### 6. View Transactions
Authenticated users can retrieve a list of their transactions.

#### How to Run:
```bash
GET /transactions
```

#### Request Headers:
```bash
Authorization: Bearer <JWT_TOKEN>
```

#### Response:
```json
{
  "status": "SUCCESS",
  "transactions": [
    {
      "id": "uuid",
      "user_id": "uuid",
      "amount": 50000,
      "transaction_type": "CREDIT",
      "balance_before": 100000,
      "balance_after": 150000,
      "created_at": "2024-10-06T14:20:00Z"
    },
    {
      "id": "uuid",
      "user_id": "uuid",
      "amount": 50000,
      "transaction_type": "DEBIT",
      "balance_before": 150000,
      "balance_after": 100000,
      "created_at": "2024-10-06T14:30:00Z"
    }
  ]
}
```

### 7. Update Profile
Authenticated users can update their profile details.

#### How to Run:
```bash
PUT /profile
```

#### Request Headers:
```bash
Authorization: Bearer <JWT_TOKEN>
```

#### Request Body:
```json
{
  "name": "New Name",
  "email": "newemail@example.com"
}
```

#### Response:
```json
{
  "status": "SUCCESS",
  "message": "Profile updated successfully"
}
```

---

## How to Run the Application

1. Make sure you have [Go installed](https://golang.org/doc/install) and PostgreSQL set up.
2. Clone the repository:
   ```bash
   git clone https://github.com/your-repo/go-rest-api.git
   cd go-rest-api
   ```
3. Update your PostgreSQL configuration in `config/config.go`.
4. Run the application:
   ```bash
   go run main.go
   ```

---

## Folder Structure

```
├── controllers/
│   └── login.go
│   └── payment.go
│   └── profile.go
│   └── register.go
│   └── topup.go
│   └── transfer.go
│   └── transactions.go
├── middlewares/
│   └── jwt_auth.go
├── models/
│   └── transaction.go
│   └── user.go
├── config/
│   └── config.go
├── main.go
```

Each folder contains relevant code for implementing specific API functionalities.

## Authentication Flow

1. **Register**: Register a new user with the `/register` endpoint.
2. **Login**: Obtain a JWT token using the `/login` endpoint.
3. **Access Authenticated Routes**: Use the JWT token in the `Authorization: Bearer <token>` header to access protected endpoints such as `/topup`, `/payment`, `/transfer`, and `/transactions`.

---

## License

This project is licensed under the MIT License.