# Loan Application Requests

Send each body separately to:

```text
POST http://127.0.0.1:8080/api/v1/loans
```

Use the header `Content-Type: application/json`.

## Application 1 — Eligible, home

```json
{
  "fullName": "Anan Chaiyasit",
  "monthlyIncome": 25000,
  "loanAmount": 150000,
  "loanPurpose": "home",
  "age": 30,
  "phoneNumber": "0810000001",
  "email": "anan01@example.com"
}
```

## Application 2 — Eligible, car

```json
{
  "fullName": "Benja Srisuk",
  "monthlyIncome": 30000,
  "loanAmount": 200000,
  "loanPurpose": "car",
  "age": 35,
  "phoneNumber": "0810000002",
  "email": "benja02@example.com"
}
```

## Application 3 — Eligible, education

```json
{
  "fullName": "Chaiwat Boonmee",
  "monthlyIncome": 18000,
  "loanAmount": 80000,
  "loanPurpose": "education",
  "age": 27,
  "phoneNumber": "0810000003",
  "email": "chaiwat03@example.com"
}
```

## Application 4 — Eligible, personal

```json
{
  "fullName": "Darunee Kanjana",
  "monthlyIncome": 22000,
  "loanAmount": 100000,
  "loanPurpose": "personal",
  "age": 40,
  "phoneNumber": "0810000004",
  "email": "darunee04@example.com"
}
```

## Application 5 — Ineligible, insufficient income

```json
{
  "fullName": "Ekkachai Prasert",
  "monthlyIncome": 5000,
  "loanAmount": 20000,
  "loanPurpose": "home",
  "age": 29,
  "phoneNumber": "0810000005",
  "email": "ekkachai05@example.com"
}
```

## Application 6 — Ineligible, age below range

```json
{
  "fullName": "Fahsai Wongsa",
  "monthlyIncome": 16000,
  "loanAmount": 50000,
  "loanPurpose": "education",
  "age": 19,
  "phoneNumber": "0810000006",
  "email": "fahsai06@example.com"
}
```

## Application 7 — Ineligible, age above range

```json
{
  "fullName": "Garn Sombat",
  "monthlyIncome": 40000,
  "loanAmount": 120000,
  "loanPurpose": "personal",
  "age": 61,
  "phoneNumber": "0810000007",
  "email": "garn07@example.com"
}
```

## Application 8 — Ineligible, business purpose

```json
{
  "fullName": "Hathai Rattanakul",
  "monthlyIncome": 50000,
  "loanAmount": 300000,
  "loanPurpose": "business",
  "age": 38,
  "phoneNumber": "0810000008",
  "email": "hathai08@example.com"
}
```

## Application 9 — Ineligible, excessive amount

```json
{
  "fullName": "Itthipol Kaewdee",
  "monthlyIncome": 10000,
  "loanAmount": 120001,
  "loanPurpose": "car",
  "age": 32,
  "phoneNumber": "0810000009",
  "email": "itthipol09@example.com"
}
```

## Application 10 — Eligible, exact income boundary

```json
{
  "fullName": "Jintana Saelim",
  "monthlyIncome": 10000,
  "loanAmount": 120000,
  "loanPurpose": "home",
  "age": 20,
  "phoneNumber": "0810000010",
  "email": "jintana10@example.com"
}
```

## Application 11 — Eligible, upper age boundary

```json
{
  "fullName": "Kamon Thongchai",
  "monthlyIncome": 28000,
  "loanAmount": 200000,
  "loanPurpose": "personal",
  "age": 60,
  "phoneNumber": "0810000011",
  "email": "kamon11@example.com"
}
```

## Application 12 — Ineligible, insufficient income

```json
{
  "fullName": "Lalita Charoen",
  "monthlyIncome": 9000,
  "loanAmount": 40000,
  "loanPurpose": "education",
  "age": 24,
  "phoneNumber": "0810000012",
  "email": "lalita12@example.com"
}
```

## Application 13 — Eligible, home

```json
{
  "fullName": "Manop Kittikul",
  "monthlyIncome": 45000,
  "loanAmount": 400000,
  "loanPurpose": "home",
  "age": 45,
  "phoneNumber": "0810000013",
  "email": "manop13@example.com"
}
```

## Application 14 — Ineligible, business purpose

```json
{
  "fullName": "Naree Sukjai",
  "monthlyIncome": 35000,
  "loanAmount": 100000,
  "loanPurpose": "business",
  "age": 33,
  "phoneNumber": "0810000014",
  "email": "naree14@example.com"
}
```

## Application 15 — Eligible, car

```json
{
  "fullName": "Oranuch Maneerat",
  "monthlyIncome": 32000,
  "loanAmount": 250000,
  "loanPurpose": "car",
  "age": 37,
  "phoneNumber": "0810000015",
  "email": "oranuch15@example.com"
}
```

## Application 16 — Ineligible, excessive amount

```json
{
  "fullName": "Preecha Noichan",
  "monthlyIncome": 20000,
  "loanAmount": 250000,
  "loanPurpose": "home",
  "age": 42,
  "phoneNumber": "0810000016",
  "email": "preecha16@example.com"
}
```

## Application 17 — Eligible, education

```json
{
  "fullName": "Quinn Pattanakul",
  "monthlyIncome": 17000,
  "loanAmount": 90000,
  "loanPurpose": "education",
  "age": 26,
  "phoneNumber": "0810000017",
  "email": "quinn17@example.com"
}
```

## Application 18 — Ineligible, insufficient income

```json
{
  "fullName": "Ratri Phromma",
  "monthlyIncome": 7500,
  "loanAmount": 15000,
  "loanPurpose": "personal",
  "age": 36,
  "phoneNumber": "0810000018",
  "email": "ratri18@example.com"
}
```

## Application 19 — Eligible, personal

```json
{
  "fullName": "Somchai Intarak",
  "monthlyIncome": 60000,
  "loanAmount": 500000,
  "loanPurpose": "personal",
  "age": 50,
  "phoneNumber": "0810000019",
  "email": "somchai19@example.com"
}
```

## Application 20 — Ineligible, age below range

```json
{
  "fullName": "Thida Rungsri",
  "monthlyIncome": 12000,
  "loanAmount": 50000,
  "loanPurpose": "car",
  "age": 18,
  "phoneNumber": "0810000020",
  "email": "thida20@example.com"
}
```

## Application 21 — Eligible, home

```json
{
  "fullName": "Udom Lertchai",
  "monthlyIncome": 55000,
  "loanAmount": 600000,
  "loanPurpose": "home",
  "age": 48,
  "phoneNumber": "0810000021",
  "email": "udom21@example.com"
}
```

## Application 22 — Ineligible, business purpose

```json
{
  "fullName": "Varee Boonyong",
  "monthlyIncome": 27000,
  "loanAmount": 150000,
  "loanPurpose": "business",
  "age": 31,
  "phoneNumber": "0810000022",
  "email": "varee22@example.com"
}
```

## Application 23 — Eligible, car

```json
{
  "fullName": "Wichai Meechai",
  "monthlyIncome": 24000,
  "loanAmount": 200000,
  "loanPurpose": "car",
  "age": 39,
  "phoneNumber": "0810000023",
  "email": "wichai23@example.com"
}
```

## Application 24 — Eligible, education

```json
{
  "fullName": "Yada Sutham",
  "monthlyIncome": 19000,
  "loanAmount": 120000,
  "loanPurpose": "education",
  "age": 28,
  "phoneNumber": "0810000024",
  "email": "yada24@example.com"
}
```
