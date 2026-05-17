FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o portfolio-app .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/portfolio-app .


COPY --from=builder /app/templates ./templates
COPY --from=builder /app/static ./static

EXPOSE 8000

CMD ["./portfolio-app"]
