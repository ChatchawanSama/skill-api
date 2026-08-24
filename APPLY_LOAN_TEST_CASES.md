# Apply Loan JSON Test Cases

Endpoint: `POST /api/v1/loans`

Content-Type: `application/json`

## 1. Invalid Requests

Invalid requests should return HTTP `400 Bad Request`.

### 1.1 Invalid `fullName`

```json
{
  "fullName": "",
  "monthlyIncome": 15000,
  "loanAmount": 100000,
  "loanPurpose": "home",
  "age": 25,
  "phoneNumber": "0851234567",
  "email": "demo@example.com"
}
```

Expected response:

```json
{
  "message": "Invalid request body",
  "reason": "fullName is required or must contain 2–255 characters"
}
```

### 1.2 Invalid `monthlyIncome`

```json
{
  "fullName": "Somkanit Jitsanook",
  "monthlyIncome": 4999,
  "loanAmount": 100000,
  "loanPurpose": "home",
  "age": 25,
  "phoneNumber": "0851234567",
  "email": "demo@example.com"
}
```

Expected response:

```json
{
  "message": "Invalid request body",
  "reason": "monthlyIncome must be between 5000 and 5000000"
}
```

### 1.3 Invalid `loanAmount`

```json
{
  "fullName": "Somkanit Jitsanook",
  "monthlyIncome": 15000,
  "loanAmount": 999,
  "loanPurpose": "home",
  "age": 25,
  "phoneNumber": "0851234567",
  "email": "demo@example.com"
}
```

Expected response:

```json
{
  "message": "Invalid request body",
  "reason": "loanAmount must be between 1000 and 5000000"
}
```

### 1.4 Invalid `loanPurpose`

```json
{
  "fullName": "Somkanit Jitsanook",
  "monthlyIncome": 15000,
  "loanAmount": 100000,
  "loanPurpose": "holiday",
  "age": 25,
  "phoneNumber": "0851234567",
  "email": "demo@example.com"
}
```

Expected response:

```json
{
  "message": "Invalid request body",
  "reason": "loanPurpose must be one of education, home, car, business, personal"
}
```

### 1.5 Invalid `age`

```json
{
  "fullName": "Somkanit Jitsanook",
  "monthlyIncome": 15000,
  "loanAmount": 100000,
  "loanPurpose": "home",
  "age": 0,
  "phoneNumber": "0851234567",
  "email": "demo@example.com"
}
```

Expected response:

```json
{
  "message": "Invalid request body",
  "reason": "age must be more than 0"
}
```

### 1.6 Invalid `phoneNumber`

```json
{
  "fullName": "Somkanit Jitsanook",
  "monthlyIncome": 15000,
  "loanAmount": 100000,
  "loanPurpose": "home",
  "age": 25,
  "phoneNumber": "085123",
  "email": "demo@example.com"
}
```

Expected response:

```json
{
  "message": "Invalid request body",
  "reason": "phoneNumber must contain exactly 10 numeric digits"
}
```

### 1.7 Invalid `email`

```json
{
  "fullName": "Somkanit Jitsanook",
  "monthlyIncome": 15000,
  "loanAmount": 100000,
  "loanPurpose": "home",
  "age": 25,
  "phoneNumber": "0851234567",
  "email": "not-an-email"
}
```

Expected response:

```json
{
  "message": "Invalid request body",
  "reason": "email must be a valid email"
}
```

## 2. Valid Requests

Valid requests should return HTTP `200 OK`.

### 2.1 Eligible

```json
{
  "fullName": "Somkanit Jitsanook",
  "monthlyIncome": 15000,
  "loanAmount": 100000,
  "loanPurpose": "home",
  "age": 25,
  "phoneNumber": "0851234567",
  "email": "demo@example.com"
}
```

Expected response:

```json
{
  "applicationId": "3fa85f64-5717-4562-b3fc-2c963f66afa6",
  "eligible": true,
  "reason": "Eligible under base rules",
  "timestamp": "2026-08-14T15:00:00+07:00"
}
```

The `applicationId` and `timestamp` values are generated for each request, so the actual values will differ.

## 3. Valid but Ineligible Requests

These requests pass input validation and should return HTTP `200 OK` with `eligible: false`.

### 3.1 Insufficient monthly income

```json
{
  "fullName": "Somkanit Jitsanook",
  "monthlyIncome": 5000,
  "loanAmount": 10000,
  "loanPurpose": "home",
  "age": 25,
  "phoneNumber": "0851234567",
  "email": "demo@example.com"
}
```

Expected response:

```json
{
  "applicationId": "3fa85f64-5717-4562-b3fc-2c963f66afa6",
  "eligible": false,
  "reason": "Monthly income is insufficient",
  "timestamp": "2026-08-14T15:00:00+07:00"
}
```

### 3.2 Age outside eligible range

```json
{
  "fullName": "Somkanit Jitsanook",
  "monthlyIncome": 15000,
  "loanAmount": 100000,
  "loanPurpose": "home",
  "age": 19,
  "phoneNumber": "0851234567",
  "email": "demo@example.com"
}
```

Expected response:

```json
{
  "applicationId": "3fa85f64-5717-4562-b3fc-2c963f66afa6",
  "eligible": false,
  "reason": "Age not in range (must be between 20-60)",
  "timestamp": "2026-08-14T15:00:00+07:00"
}
```

### 3.3 Unsupported business purpose

```json
{
  "fullName": "Somkanit Jitsanook",
  "monthlyIncome": 15000,
  "loanAmount": 100000,
  "loanPurpose": "business",
  "age": 25,
  "phoneNumber": "0851234567",
  "email": "demo@example.com"
}
```

Expected response:

```json
{
  "applicationId": "3fa85f64-5717-4562-b3fc-2c963f66afa6",
  "eligible": false,
  "reason": "Business loans not supported",
  "timestamp": "2026-08-14T15:00:00+07:00"
}
```

### 3.4 Loan amount exceeds 12 months of income

```json
{
  "fullName": "Somkanit Jitsanook",
  "monthlyIncome": 15000,
  "loanAmount": 180001,
  "loanPurpose": "home",
  "age": 25,
  "phoneNumber": "0851234567",
  "email": "demo@example.com"
}
```

Expected response:

```json
{
  "applicationId": "3fa85f64-5717-4562-b3fc-2c963f66afa6",
  "eligible": false,
  "reason": "Loan amount cannot exceed 12 months of income",
  "timestamp": "2026-08-14T15:00:00+07:00"
}
```
