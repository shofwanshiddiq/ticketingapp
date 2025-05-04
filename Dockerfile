# Stage 1 Build


FROM golang:1.23-alpine AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /server

# Stage 2: Deploy
FROM alpine:edge

WORKDIR /app

COPY .env .env

COPY --from=build /server /app/service

COPY --from=build /server /app/migrations 

# Install timezone data and certificates
RUN apk --no-cache add ca-certificates tzdata

# Set the entrypoint command
ENTRYPOINT ["/app/service"]