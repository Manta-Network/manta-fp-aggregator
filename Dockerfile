FROM golang:1.23-alpine3.21 as builder

RUN apk add --no-cache make gcc musl-dev linux-headers git

WORKDIR /app/manta-fp-aggregator

COPY . .

RUN make build

FROM alpine:3.21

WORKDIR /app/manta-fp-aggregator

RUN apk add --no-cache ca-certificates busybox-extras

COPY --from=builder /app/manta-fp-aggregator/manta-fp-aggregator .

ENTRYPOINT ["./manta-fp-aggregator"]
