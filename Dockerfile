FROM golang:1.20 AS builder
ADD . /src
WORKDIR /src
RUN go build -o inn cmd/inn.go
FROM debian:bullseye
COPY --from=builder --chown=root:root --chmod=0755 /src/inn /app/
ENV GIN_MODE=release
EXPOSE 8080
CMD	/app/inn
