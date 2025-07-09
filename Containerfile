FROM golang:1.23.4-alpine3.21 as builder

WORKDIR /app

COPY go.mod .
COPY go.sum .
COPY src/ src/
RUN go build -o bin/api .

COPY config/ config/

FROM alpine:3.21.0

WORKDIR /redirgo
COPY --from=builder /app/bin/ /redirgo/bin/

COPY config/ config/
COPY public/ public/

CMD ["/redirgo/bin/api"]
