FROM golang:alpine as builder
RUN apk update
RUN apk add -U --no-cache ca-certificates && update-ca-certificates
WORKDIR /app/cmd
COPY . /app
EXPOSE 8080
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o temperday

FROM scratch
WORKDIR /app/cmd
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/cmd/temperday .
COPY --from=builder /app/cmd/.env .
EXPOSE 8080
ENTRYPOINT [ "./temperday" ]