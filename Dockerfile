FROM golang:1.22.5-alpine

WORKDIR /app

COPY ./задание/ .
COPY wait-for-it.sh /app/wait-for-it.sh
RUN chmod +x /app/wait-for-it.sh

RUN apk add --no-cache gcc musl-dev bash

RUN go mod download

RUN CGO_ENABLED=1 go build -o main ./cmd

CMD ["sh", "-c", "/app/wait-for-it.sh db:5432 -- ./main"]
