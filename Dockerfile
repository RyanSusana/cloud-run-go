FROM golang

COPY . /go

WORKDIR /go/src

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["main"]