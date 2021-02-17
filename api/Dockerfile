### Build
FROM golang:1.16.0-alpine as build

LABEL org.opencontainers.image.source="https://github.com/kostaspt/domane"

WORKDIR /go/src/app

# Install system dependencies
RUN apk add --no-cache make

# Install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Generate production build
COPY . .
RUN make build

### Serve
FROM alpine:3.13.1
WORKDIR /app

COPY --from=build /go/src/app/bin/domane /app

EXPOSE $API_PORT

CMD ["sh" , "-c", "./domane start -p ${API_PORT}"]