FROM golang:1.19-bullseye AS builder

RUN apt update -qq \
    && apt upgrade -y -qq \
    && apt install -y -qq build-essential \
    && apt clean \
    && apt autoclean \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64  go build -a -installsuffix cgo -o /app/apigo ./cmd/api

FROM scratch

COPY --from=builder /app/apigo /apigo

EXPOSE 9000

ENTRYPOINT ["/apigo"]