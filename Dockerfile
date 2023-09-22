FROM golang:1.21 as builder

RUN mkdir /app
COPY go.mod go.sum *.go /app/

WORKDIR /app
RUN go build -o /bin/prometheus-unix-time-exporter main.go

FROM debian:latest

COPY --from=builder /bin/prometheus-unix-time-exporter /bin/prometheus-unix-time-exporter

ENTRYPOINT [ "/bin/prometheus-unix-time-exporter" ]