# Apply Loan JSON Test Cases

Endpoint: `POST /api/v1/loans`

Content-Type: `application/json`

## 1. Invalid requests (HTTP 400)

All invalid requests return `"message": "Invalid request body"`.

### 1.1 Malformed JSON

```json
{"fullName":"Somkanit Jitsanook",}
```

```json
{
  "message": "Invalid request body",
  "reason": "<JSON syntax error returned by Echo>"
}
```

The exact syntax-error text can differ depending on the JSON parser version.

### 1.2 Missing all required fields

```json
{}
```

```json
{
  "message": "Invalid request body",
  "reason": "missing required fields: fullName, monthlyIncome, loanAmount, loanPurpose, age, phoneNumber, email"
}
```

### 1.3 Missing some required fields

This example is missing `fullName` and `monthlyIncome`.

```json
{
  "loanAmount": 100000,
  "loanPurpose": "home",
  "age": 25,
  "phoneNumber": "0851234567",
  "email": "demo@example.com"
}
```

```json
{
  "message": "Invalid request body",
  "reason": "missing required fields: fullName, monthlyIncome"
}
```

If fields are both missing and invalid, the current validator reports only the missing fields.

### 1.4 Invalid `fullName`

```json
{
  "fullName": "S",
  "monthlyIncome": 15000,
  "loanAmount": 100000,
  "loanPurpose": "home",
  "age": 25,
  "phoneNumber": "0851234567",
  "email": "demo@example.com"
}
```

```json
{"message":"Invalid request body","reason":"fullName must be a valid fullName"}
```

### 1.5 Invalid `monthlyIncome`

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

```json
{"message":"Invalid request body","reason":"monthlyIncome must be a valid monthlyIncome"}
```

### 1.6 Invalid `loanAmount`

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

```json
{"message":"Invalid request body","reason":"loanAmount must be a valid loanAmount"}
```

### 1.7 Invalid `loanPurpose`

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

```json
{"message":"Invalid request body","reason":"loanPurpose must be a valid loanPurpose"}
```

### 1.8 Invalid `age`

Use `-1`. An age of `0` is Go's zero value and fails `required`, so it is reported as missing instead.

```json
{
  "fullName": "Somkanit Jitsanook",
  "monthlyIncome": 15000,
  "loanAmount": 100000,
  "loanPurpose": "home",
  "age": -1,
  "phoneNumber": "0851234567",
  "email": "demo@example.com"
}
```

```json
{"message":"Invalid request body","reason":"age must be a valid age"}
```

### 1.9 Invalid `phoneNumber`

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

```json
{"message":"Invalid request body","reason":"phoneNumber must be a valid phoneNumber"}
```

### 1.10 Invalid `email`

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

```json
{"message":"Invalid request body","reason":"email must be a valid email"}
```

## 2. Valid and eligible (HTTP 200)

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

```json
{
  "applicationId": "3fa85f64-5717-4562-b3fc-2c963f66afa6",
  "eligible": true,
  "reason": "Eligible under base rules",
  "timestamp": "2026-08-26T15:00:00+07:00"
}
```

The UUID and timestamp are generated for each request, so their actual values will differ.

## 3. Valid but ineligible (HTTP 200)

These requests pass handler validation but return `"eligible": false`.

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

```json
{
  "applicationId": "3fa85f64-5717-4562-b3fc-2c963f66afa6",
  "eligible": false,
  "reason": "Monthly income is insufficient",
  "timestamp": "2026-08-26T15:00:00+07:00"
}
```

### 3.2 Age outside eligible range

Use the valid request above with `"age": 19`.

```json
{
  "applicationId": "3fa85f64-5717-4562-b3fc-2c963f66afa6",
  "eligible": false,
  "reason": "Age not in range (must be between 20-60)",
  "timestamp": "2026-08-26T15:00:00+07:00"
}
```

### 3.3 Unsupported business purpose

Use the valid request above with `"loanPurpose": "business"`.

```json
{
  "applicationId": "3fa85f64-5717-4562-b3fc-2c963f66afa6",
  "eligible": false,
  "reason": "Business loans not supported",
  "timestamp": "2026-08-26T15:00:00+07:00"
}
```

### 3.4 Loan amount exceeds 12 months of income

Use the valid request above with `"loanAmount": 180001`.

```json
{
  "applicationId": "3fa85f64-5717-4562-b3fc-2c963f66afa6",
  "eligible": false,
  "reason": "Loan amount cannot exceed 12 months of income",
  "timestamp": "2026-08-26T15:00:00+07:00"
}
```

## 4. Service or database error (HTTP 500)

After a valid request, a service or repository failure returns:

```json
{
  "message": "Unable to process loan application",
  "reason": "<service or database error>"
}
```
