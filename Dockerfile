### Multi-stage build
FROM jormungandrk/goa-build as build

COPY . /go/src/github.com/Microkubes/microservice-user-profile
RUN go install github.com/Microkubes/microservice-user-profile


### Main
FROM alpine:3.7

COPY --from=build /go/bin/microservice-user-profile /usr/local/bin/microservice-user-profile
EXPOSE 8080

ENV API_GATEWAY_URL="http://localhost:8001"

CMD ["/usr/local/bin/microservice-user-profile"]
