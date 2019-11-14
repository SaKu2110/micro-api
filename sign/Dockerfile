FROM golang:latest

WORKDIR /go/src/github.com/SaKu2110/micro-api/sign
COPY ./ ./
RUN go get github.com/gin-gonic/gin
RUN go get github.com/jinzhu/gorm

ENV DB_USER signuser
ENV DB_PASS tokatonton
ENV DB_IP mysql
ENV DB_PORT 3306
ENV DB_NAME sign

EXPOSE 9000
CMD ["go", "run","/go/src/github.com/SaKu2110/micro-api/sign/main.go"]