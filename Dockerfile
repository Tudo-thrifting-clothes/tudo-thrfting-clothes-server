FROM golang:alpine AS builder

WORKDIR /build

COPY . .

RUN go mod download

RUN go build -o tudo-thrifting-server ./cmd/server

FROM scratch

COPY ./config /config

COPY --from=builder /build/tudo-thrifting-server /

CMD ["/tudo-thrifting-server", "config/production.yaml"]
