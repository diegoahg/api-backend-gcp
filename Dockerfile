FROM golang:alpine AS build
RUN apk --no-cache add gcc g++ make git
WORKDIR /go/src/app
COPY . .
RUN go build -o goapp cmd/main.go 

FROM alpine:3.17
RUN apk --no-cache add ca-certificates
WORKDIR /usr/bin
COPY --from=build /go/src/app/goapp /usr/local/bin/
EXPOSE 8080
ENTRYPOINT /usr/local/bin/goapp --port 8080