FROM golang:alpine AS builder

RUN mkdir /app
ADD . /app
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o fxbot .

FROM alpine

COPY --from=builder /app/fxbot /fxbot

ENTRYPOINT ["/fxbot"]
