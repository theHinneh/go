FROM golang:1.14.9-alpine AS builder
RUN mkdir /bin
ADD go.mod go.sum main.go ./todo/ ./database/ /bin/
WORKDIR /build
RUN go build -o bin/rangotodo -v .

FROM alpine
RUN adduser -S -D -H -h /app rango
USER rango

COPY --from=builder /bin/rangotodo /app/

WORKDIR /app

CMD ["./rangotodo"]