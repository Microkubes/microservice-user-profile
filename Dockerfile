### Multi-stage build
FROM golang:1.10-alpine3.7 as build

RUN apk --no-cache add git

RUN go get -u -v github.com/goadesign/goa/... && \
    go get -u -v gopkg.in/mgo.v2 && \
    go get -u -v github.com/Microkubes/microservice-security/...

COPY . /go/src/github.com/Microkubes/microservice-user-profile

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go install github.com/Microkubes/microservice-user-profile


### Main
FROM scratch

ENV API_GATEWAY_URL="http://localhost:8001"

COPY --from=build /go/bin/microservice-user-profile /usr/local/bin/microservice-user-profile

EXPOSE 8080

CMD ["/usr/local/bin/microservice-user-profile"]
