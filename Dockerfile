FROM golang:1.14

WORKDIR /go/src/app
RUN git clone https://github.com/Danioq/kcms .
# COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["app"]