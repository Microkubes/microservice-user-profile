### Multi-stage build
FROM golang:1.17.3-alpine3.15 as build

RUN apk --no-cache add git

COPY . /go/src/github.com/Microkubes/microservice-user-profile

RUN cd /go/src/github.com/Microkubes/microservice-user-profile && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go install


### Main
FROM alpine:3.15

COPY --from=build /go/src/github.com/Microkubes/microservice-user-profile/config.json /config.json
COPY --from=build /go/bin/microservice-user-profile /usr/local/bin/microservice-user-profile

EXPOSE 8080

CMD ["/usr/local/bin/microservice-user-profile"]
