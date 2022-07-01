### Build
FROM golang:1.18.3-alpine as build

LABEL org.opencontainers.image.source="https://github.com/kostaspt/domane"

WORKDIR /go/src/app

# Install system dependencies
RUN apk add --no-cache make git

# Install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Generate production build
COPY . .
RUN make build

### Serve
FROM alpine:3.13.5
WORKDIR /app

COPY --from=build /go/src/app/bin/start /app

EXPOSE $API_PORT

CMD ["sh" , "-c", "./start -p ${API_PORT}"]