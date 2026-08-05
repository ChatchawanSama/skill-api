# loan-api (ArisePreQ)

Loan pre-approval API scaffold (Go + Echo + MySQL). See `PROJECT.md` for full spec.

## Status: SKELETON ONLY

Wiring done. API business logic stubbed with `TODO:` markers. Students implement 3 API endpoints only.

## Structure

```
├── cmd/main.go                    # entrypoint (config + wire + serve)
├── config/                        # config.go, config.yaml, secret.env.example
├── internal/
│   ├── handlers/rest/
│   │   ├── router.go              # routes
│   │   ├── health_check.go        # /health /ready
│   │   └── loan.go                # TODO: 3 endpoint handlers (return 501 now)
│   ├── models/
│   │   ├── common.go              # response wrappers
│   │   └── loan.go                # DTOs + loan purpose list
│   ├── repositories/adaptor/
│   │   ├── adaptor.go             # repository interface
│   │   ├── entities/loan.go       # MySQL row struct
│   │   ├── mocks/                 # mockery output (go generate)
│   │   └── mysql/
│   │       ├── db.go              # sql.Open + ping
│   │       └── loan.go           # TODO: 3 repo methods (return errUnimplemented)
│   └── services/
│       ├── loan.go               # TODO: 3 service methods (return ErrNotImplemented)
│       ├── common.go             # errors + response helpers
│       └── mocks/                # mockery output (go generate)
├── _compose/compose.yaml          # MySQL docker compose
├── Dockerfile                     # DONE
├── Makefile                       # DONE
└── go.mod
```

## Layout mirrored from `orch-e-statement`

Layered: handler -> service -> adaptor (interface) -> mysql impl.
Mocks via `mockery` for unit tests. Config via `SECRET_*` env vars (see `config/secret.env.example`).

## Tasks for student

Implement 3 API endpoints per `PROJECT.md`. Fill `TODO:` stubs in:

1. `internal/repositories/adaptor/mysql/loan.go` — SQL against `loan_applications` table
2. `internal/services/loan.go` — eligibility rules, error mapping, pagination
3. `internal/handlers/rest/loan.go` — bind, validate, call service, format response

Endpoints:

- `POST /api/v1/loans` — apply for a loan (validate, evaluate eligibility, persist, return decision)
- `GET /api/v1/loans/:applicationId` — get loan status by id
- `GET /api/v1/loans?page=&limit=&eligible=&purpose=` — list loans with pagination

Eligibility rules (in service):

- `monthlyIncome >= 10000` else `"Monthly income is insufficient"`
- `age` 20..60 else `"Age not in range (must be between 20-60)"`
- `loanPurpose != "business"` else `"Business loans not supported"`
- `loanAmount <= 12 * monthlyIncome` else `"Loan amount cannot exceed 12 months of income"`
- all pass -> `eligible=true`, `reason="Eligible under base rules"`

Timestamps: Asia/Bangkok (UTC+7). Application IDs: UUID.

## Commands

```
make run          # APPENV=local go run -race cmd/main.go
make test         # go test -race -v ./...
make cover        # coverage report
make mock         # go generate mocks
make compose-up   # docker compose up -d (MySQL)
make compose-down # docker compose down
```
