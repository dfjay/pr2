FROM --platform=$BUILDPLATFORM golang:1.24-alpine AS builder

ARG BUILDPLATFORM
ARG TARGETOS
ARG TARGETARCH

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o app .

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/app .

EXPOSE 8080
CMD ["./app"]
