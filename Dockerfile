# Build the server binary
FROM golang:1.23.3-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o build/app ./cmd/server

# The final image
FROM scratch
COPY --from=builder /app/build/app /app

ENV PORT=80
EXPOSE ${PORT}

CMD ["/app"]
