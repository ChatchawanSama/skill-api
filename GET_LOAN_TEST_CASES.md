# Get Loan API Test Cases



## 1. Get Loan by Application ID


### 1.1 Existing application

Request:

```text
GET http://127.0.0.1:8080/api/v1/loans/3fa85f64-5717-4562-b3fc-2c963f66afa6
```

Expected status:

```text
200 OK
```

Expected response:

```json
{
  "applicationId": "3fa85f64-5717-4562-b3fc-2c963f66afa6",
  "fullName": "Somkanit Jitsanook",
  "monthlyIncome": 15000,
  "loanAmount": 100000,
  "loanPurpose": "home",
  "age": 25,
  "phoneNumber": "0851234567",
  "email": "demo@example.com",
  "eligible": true,
  "reason": "Eligible under base rules",
  "timestamp": "2026-08-18T15:00:00+07:00"
}
```


### 1.2 Application not found

Request:

```text
GET http://127.0.0.1:8080/api/v1/loans/00000000-0000-0000-0000-000000000000
```

Expected status:

```text
404 Not Found
```

Expected response:

```json
{
  "message": "Loan application not found",
  "reason": "applicationId not found: 00000000-0000-0000-0000-000000000000"
}
```

## 2. Get All Loans

### 2.1 No query parameters

Request:

```text
GET http://127.0.0.1:8080/api/v1/loans
```

Expected behavior:

```text
page defaults to 1
limit defaults to 10
```

Expected status:

```text
200 OK
```

Example response:

```json
{
  "applications": [
    {
      "applicationId": "3fa85f64-5717-4562-b3fc-2c963f66afa6",
      "fullName": "Somkanit Jitsanook",
      "monthlyIncome": 15000,
      "loanAmount": 100000,
      "loanPurpose": "home",
      "age": 25,
      "phoneNumber": "0851234567",
      "email": "demo@example.com",
      "eligible": true,
      "reason": "Eligible under base rules",
      "timestamp": "2026-08-18T15:00:00+07:00"
    }
  ],
  "page": 1,
  "totalPages": 3
}
```

### 2.2 Page and limit

Request:

```text
GET http://127.0.0.1:8080/api/v1/loans?page=2&limit=5
```

Expected status:

```text
200 OK
```

Example response:

```json
{
  "applications": [
    {
      "applicationId": "7c9e6679-7425-40de-944b-e07fc1f90ae7",
      "fullName": "Example Applicant",
      "monthlyIncome": 20000,
      "loanAmount": 100000,
      "loanPurpose": "car",
      "age": 30,
      "phoneNumber": "0810000001",
      "email": "applicant@example.com",
      "eligible": true,
      "reason": "Eligible under base rules",
      "timestamp": "2026-08-18T14:00:00+07:00"
    }
  ],
  "page": 2,
  "totalPages": 5
}
```

### 2.3 Eligible applications only

Request:

```text
GET http://127.0.0.1:8080/api/v1/loans?eligible=true
```

Expected status:

```text
200 OK
```

Expected response layout:

```json
{
  "applications": [
    {
      "applicationId": "3fa85f64-5717-4562-b3fc-2c963f66afa6",
      "fullName": "Somkanit Jitsanook",
      "monthlyIncome": 15000,
      "loanAmount": 100000,
      "loanPurpose": "home",
      "age": 25,
      "phoneNumber": "0851234567",
      "email": "demo@example.com",
      "eligible": true,
      "reason": "Eligible under base rules",
      "timestamp": "2026-08-18T15:00:00+07:00"
    }
  ],
  "page": 1,
  "totalPages": 1
}
```

### 2.4 Ineligible applications only

Request:

```text
GET http://127.0.0.1:8080/api/v1/loans?eligible=false
```

Expected status:

```text
200 OK
```

Expected response layout:

```json
{
  "applications": [
    {
      "applicationId": "8f14e45f-ea32-4c6a-bf0a-123456789abc",
      "fullName": "Example Applicant",
      "monthlyIncome": 5000,
      "loanAmount": 10000,
      "loanPurpose": "home",
      "age": 25,
      "phoneNumber": "0810000002",
      "email": "ineligible@example.com",
      "eligible": false,
      "reason": "Monthly income is insufficient",
      "timestamp": "2026-08-18T15:05:00+07:00"
    }
  ],
  "page": 1,
  "totalPages": 1
}
```

Every returned application should have `eligible: false`.

### 2.5 Filter by loan purpose

Request:

```text
GET http://127.0.0.1:8080/api/v1/loans?purpose=home
```

Expected status:

```text
200 OK
```

Expected response layout:

```json
{
  "applications": [
    {
      "applicationId": "3fa85f64-5717-4562-b3fc-2c963f66afa6",
      "fullName": "Somkanit Jitsanook",
      "monthlyIncome": 15000,
      "loanAmount": 100000,
      "loanPurpose": "home",
      "age": 25,
      "phoneNumber": "0851234567",
      "email": "demo@example.com",
      "eligible": true,
      "reason": "Eligible under base rules",
      "timestamp": "2026-08-18T15:00:00+07:00"
    }
  ],
  "page": 1,
  "totalPages": 1
}
```

### 2.6 Combined filters and pagination

Request:

```text
GET http://127.0.0.1:8080/api/v1/loans?page=1&limit=5&eligible=true&purpose=home
```

Expected status:

```text
200 OK
```

Expected response layout:

```json
{
  "applications": [
    {
      "applicationId": "3fa85f64-5717-4562-b3fc-2c963f66afa6",
      "fullName": "Somkanit Jitsanook",
      "monthlyIncome": 15000,
      "loanAmount": 100000,
      "loanPurpose": "home",
      "age": 25,
      "phoneNumber": "0851234567",
      "email": "demo@example.com",
      "eligible": true,
      "reason": "Eligible under base rules",
      "timestamp": "2026-08-18T15:00:00+07:00"
    }
  ],
  "page": 1,
  "totalPages": 1
}
```


### 2.7 No matching applications

Request:

```text
GET http://127.0.0.1:8080/api/v1/loans?purpose=purpose-with-no-records
```

Expected status:

```text
200 OK
```

Expected response:

```json
{
  "applications": [],
  "page": 1,
  "totalPages": 0
}
```

### 2.8 Page beyond available results

Request:

```text
GET http://127.0.0.1:8080/api/v1/loans?page=100&limit=10
```

Expected status:

```text
200 OK
```

Example response:

```json
{
  "applications": [],
  "page": 100,
  "totalPages": 3
}
```


### 2.9 Invalid page value

Request:

```text
GET http://127.0.0.1:8080/api/v1/loans?page=hello
```

Expected status:

```text
400 Bad Request
```

Expected response layout:

```json
{
  "message": "Invalid query parameters",
  "reason": "<Echo query binding error>"
}
```

### 2.10 Invalid eligible value

Request:

```text
GET http://127.0.0.1:8080/api/v1/loans?eligible=maybe
```

Expected status:

```text
400 Bad Request
```

Expected response layout:

```json
{
  "message": "Invalid query parameters",
  "reason": "<Echo query binding error>"
}
```
