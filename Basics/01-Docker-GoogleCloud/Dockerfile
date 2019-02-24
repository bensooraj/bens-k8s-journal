FROM golang:1.8
RUN mkdir /go_app 
COPY . /go_app/
WORKDIR /go_app
RUN go build -o hello-world .
EXPOSE 8000
CMD ["/go_app/hello-world"]