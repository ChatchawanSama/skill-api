FROM golang:1.25 AS builder
WORKDIR /app
COPY . ./
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -o main ./cmd/main.go

FROM alpine:3.20
WORKDIR /app
RUN apk --no-cache add ca-certificates tzdata libc6-compat
ENV TZ=Asia/Bangkok
COPY --from=builder /app/main .
ENTRYPOINT ["./main"]
