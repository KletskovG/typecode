FROM golang:1.14-buster AS deps
EXPOSE 8081
WORKDIR /


COPY ./src/go.mod ./
COPY ./src/go.sum ./
RUN go mod download -x

FROM deps AS builder
COPY ./src ./
RUN go build -v -o /bin/server ./main.go

FROM debian:buster-slim
RUN set -x && apt-get update && \
  DEBIAN_FRONTEND=noninteractive apt-get install -y ca-certificates && \
  rm -rf /var/lib/apt/lists/*

WORKDIR /app
COPY --from=builder /bin/server ./

CMD ["./server"]