# مرحله اول: بیلد برای معماری ARM
FROM golang:1.17-alpine AS builder


RUN apk add --no-cache build-base


RUN apk add --no-cache \
    build-base \
    gcc \
    binutils \
    musl-dev 


ENV GO111MODULE=on
ENV CGO_ENABLED=1


WORKDIR /nakama/data


COPY go.mod go.sum ./


RUN go mod tidy


COPY . .


RUN GOARCH=arm64 go build -buildmode=plugin -o /nakama/data/modules/backend.so main.go


RUN chmod +x /nakama/data/modules/backend.so


FROM --platform=linux/arm64 heroiclabs/nakama:3.22.0 AS runtime


WORKDIR /nakama


COPY --from=builder /nakama/data/modules/backend.so /nakama/data/modules/backend.so
COPY --from=builder /nakama/data/local.yml /nakama/data/local.yml
COPY --from=builder /nakama/data/modules/*.json /nakama/data/modules/


RUN chmod +x /nakama/nakama


CMD ["/nakama/nakama", "migrate", "up", "--database.address", "postgres:localdb@postgres:5432/nakama"]
