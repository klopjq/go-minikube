FROM golang:1.16
COPY app /go/src/app
WORKDIR /go/src/app
RUN go mod init
RUN go mod tidy
RUN go build -o server .
EXPOSE 3000
CMD ["/go/src/app/server"]