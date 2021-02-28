#build stage
FROM golang:alpine AS builder
WORKDIR /go/src/app
COPY . .
RUN go install -v
RUN go build -o app
ENTRYPOINT ./app
LABEL Name=govote Version=0.0.1
EXPOSE 3000
