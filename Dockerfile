### Multi-stage build
FROM golang:1.8.3-alpine3.6 as build

RUN apk --no-cache add git curl openssh

COPY keys/id_rsa /root/.ssh/id_rsa
RUN chmod 700 /root/.ssh/id_rsa && \
    echo -e "Host github.com\n\tStrictHostKeyChecking no\n" >> /root/.ssh/config && \
    git config --global url."ssh://git@github.com:".insteadOf "https://github.com"

RUN go get -u github.com/goadesign/goa/... && \
    go get -u gopkg.in/mgo.v2 && \
    go get -u golang.org/x/crypto/bcrypt

RUN go get -u github.com/JormungandrK/microservice-tools; exit 0

RUN go get -u github.com/JormungandrK/microservice-security/...; exit 0

COPY . /go/src/github.com/JormungandrK/microservice-user-profile
RUN go install github.com/JormungandrK/microservice-user-profile


### Main
FROM alpine:3.6

COPY --from=build /go/bin/microservice-user-profile /usr/local/bin/microservice-user-profile
COPY config.json /config.json
EXPOSE 8080

ENV SERVICE_CONFIG_FILE="config.json"
ENV API_GATEWAY_URL="http://localhost:8001"

CMD ["/usr/local/bin/microservice-user-profile"]
