FROM golang:1.23 as builder

WORKDIR /app/manta-fp-aggregator

COPY . .

RUN make build

FROM debian:bookworm-slim

WORKDIR /app/manta-fp-aggregator

RUN apt-get update \
    && apt-get install -y --no-install-recommends ca-certificates net-tools \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/* /var/cache/apt/archives/*

COPY --from=builder /app/manta-fp-aggregator/manta-fp-aggregator .

ENTRYPOINT ["./manta-fp-aggregator"]
