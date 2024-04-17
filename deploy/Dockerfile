FROM golang:1.22.2-alpine as builder

ARG SERVICE

RUN apk update && apk add --no-cache bash curl tzdata
RUN apk add

RUN mkdir /build
WORKDIR /build
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN go build -o bin/main ./cmd/$SERVICE

FROM alpine:latest
COPY --from=builder /build/bin /bin/
COPY --from=builder /build/.env .
COPY --from=builder /build/api/swagger/swagger-config.yaml /api/swagger/

ENTRYPOINT ["/bin/main"]