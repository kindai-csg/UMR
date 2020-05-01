FROM golang:1.13.10-buster


ENV GOPATH /.go

RUN apt-get -y update && \
    apt-get -y install git unzip

RUN curl -OL https://github.com/google/protobuf/releases/download/v3.11.4/protoc-3.11.4-linux-x86_64.zip && \
    unzip protoc-3.11.4-linux-x86_64.zip -d protoc3 && \
    mv protoc3/bin/* /bin/ && \
    rm -rf protoc-3.9.0-linux-x86_64.zip protoc3 && \
    export PATH=$PATH:$GOPATH/bin >> .bashrc

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
RUN go get -v github.com/BurntSushi/toml
RUN go get -v google.golang.org/grpc 
RUN go get -v github.com/golang/protobuf/protoc-gen-go

CMD ["go", "run", "main.go"]
