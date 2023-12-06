FROM golang AS builder

COPY . /app
WORKDIR /app

RUN go build -o http-tester .

FROM debian:stable
COPY --from=builder /app/http-tester /usr/local/bin/http-tester

WORKDIR /
ENTRYPOINT ["/usr/local/bin/http-tester"]
