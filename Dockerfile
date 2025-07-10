FROM golang:1.24.4 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o product-sorter ./cmd/product-sorter

FROM gcr.io/distroless/static-debian12

USER nonroot:nonroot

COPY --from=builder /app/product-sorter /product-sorter

ENTRYPOINT ["/product-sorter"]