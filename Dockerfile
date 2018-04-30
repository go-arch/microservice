FROM golang:latest
RUN mkdir -p /go/src/students
ADD . /go/src/students/
WORKDIR /go/src/students/
RUN go get ./
RUN go build -o main .
CMD ["/go/src/students/main"]