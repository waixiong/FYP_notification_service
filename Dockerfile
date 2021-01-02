FROM golang:1.13-alpine AS build_base

RUN apk add --no-cache git

WORKDIR /tmp/service

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY cmd ./cmd/
COPY pkg ./pkg/

RUN go build -o ./out/server ./cmd/server/main.go

FROM alpine:3.9
RUN apk add ca-certificates tzdata

COPY --from=build_base /tmp/service/out/server /cmd/server

# EXPOSE 80
# EXPOSE 443
EXPOSE 8120
EXPOSE 8121

COPY configs /configs/
COPY .env .

COPY third_party/mail/html /third_party/mail/html/

CMD ["/cmd/server"]