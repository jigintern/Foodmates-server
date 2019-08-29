FROM golang:1.12.9
ENV GO111MODULE=on
CMD /bin/bash
WORKDIR /go/src/github.com/jigintern/Foodmates-server
COPY go.mod go.sum ./
RUN go mod download

COPY . .

CMD ["go", "run", "main.go"]
