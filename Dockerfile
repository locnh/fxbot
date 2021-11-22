FROM golang:alpine as builder

RUN mkdir /app
ADD . /app
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o mastercardfxbot .

FROM alpine

COPY --from=builder /app/mastercardfxbot /mastercardfxbot

ENTRYPOINT ["/mastercardfxbot"]
