FROM golang:1.13.5-alpine3.10

RUN apk add git && \
    go get -v github.com/gin-gonic/gin

CMD ["go", "run", "main.go"]
