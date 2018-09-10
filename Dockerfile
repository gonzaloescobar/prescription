FROM golang:1.8-alpine
RUN apk add --no-cache bash git openssh
ADD . /go/src/app
RUN go get -t -v ./...
RUN go install app

FROM alpine:latest
COPY --from=0 /go/bin/app .
ENV PORT 8080
CMD ["./app"]s