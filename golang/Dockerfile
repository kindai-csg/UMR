FROM golang:1.13.5-alpine3.10

ENV GOPATH /.go

RUN apk add git

RUN go get -v github.com/gin-gonic/gin
RUN go get -v gopkg.in/ldap.v2
RUN go get -v github.com/garyburd/redigo/redis
RUN go get -v github.com/golang/mock/gomock
RUN go get -v golang.org/x/tools/go/packages
RUN go get -v github.com/google/uuid
RUN go install github.com/golang/mock/mockgen
RUN go get -v gopkg.in/go-playground/validator.v9
RUN go get -v github.com/go-sql-driver/mysql
RUN go get -v github.com/dgrijalva/jwt-go

CMD ["go", "run", "main.go"]
