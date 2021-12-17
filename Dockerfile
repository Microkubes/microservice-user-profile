### Multi-stage build
FROM golang:1.13.5-alpine3.10 as build

RUN apk --no-cache add git

COPY . /go/src/github.com/Microkubes/microservice-user-profile

RUN cd /go/src/github.com/Microkubes/microservice-user-profile && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go install


### Main
FROM alpine:3.10

COPY --from=build /go/src/github.com/Microkubes/microservice-user-profile/config.json /config.json
COPY --from=build /go/bin/microservice-user-profile /usr/local/bin/microservice-user-profile

EXPOSE 8080

CMD ["/usr/local/bin/microservice-user-profile"]
