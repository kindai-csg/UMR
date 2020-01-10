FROM golang:1.13.5-alpine3.10

RUN apk add git && \
    go get -v github.com/gin-gonic/gin

RUN go get -v gopkg.in/ldap.v2

CMD ["go", "run", "main.go"]
