FROM golang:1.12.9
CMD /bin/bash
WORKDIR /go/src/github.com/jigintern/Foodmates-server
RUN go get -u github.com/gin-gonic/gin \
  && go get github.com/jinzhu/gorm \
  && go get github.com/go-sql-driver/mysql \
  && go get github.com/joho/godotenv

CMD ["go", "run", "main.go"]