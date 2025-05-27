FROM golang:1.23-alpine3.21 as go_builder
WORKDIR /app
COPY . .
RUN apk update && apk add tzdata ca-certificates
RUN CGO_ENABLED=0 GOOS=linux go build -o app -ldflags "-w -s" .

FROM scratch
WORKDIR /app
COPY --from=go_builder /app/app .
COPY --from=go_builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=go_builder /usr/share/zoneinfo /usr/share/zoneinfo
ENV TZ="Europe/Budapest"
ENV GIN_MODE="release"

ENTRYPOINT ["./app"]
