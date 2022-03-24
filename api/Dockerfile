FROM golang:1.16-buster AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /api .

FROM alpine:latest

WORKDIR /

COPY --from=build /api /api

EXPOSE 3000

ENTRYPOINT ["/api"]