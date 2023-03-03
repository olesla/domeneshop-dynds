FROM golang AS builder

COPY main.go /go/src/domenshop-dyndns/main.go
WORKDIR /go/src/domenshop-dyndns/
RUN go mod init
RUN go mod tidy
RUN CGO_ENABLED=0 go build -o /go/src/domenshop-dyndns/app .


FROM scratch

COPY --from=builder /go/src/domenshop-dyndns/app /app
COPY --from=builder /etc/ssl /etc/ssl

ENTRYPOINT ["/app"]
