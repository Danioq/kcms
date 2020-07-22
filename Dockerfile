FROM golang:1.14

RUN mkdir /kcms
RUN git clone https://github.com/Danioq/kcms /kcms
WORKDIR /kcms

RUN go build -o main .
CMD ["/kcms/main"]
