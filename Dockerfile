FROM golang:1.21-alpine AS build

WORKDIR /
COPY . ./

RUN go build -o ./main ./cmd/main.go

RUN chmod +x ./main

CMD ["./main"]
