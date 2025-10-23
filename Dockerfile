FROM golang:alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 -ldflags="-s -w" go build -o mcp .

FROM gcr.io/distroless/static-debian12
COPY --from=build /app/mcp /mcp
CMD ["/mcp"]
