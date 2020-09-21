FROM golang:1.14-alpine

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN go build -o tfarbe .

ENTRYPOINT ["./tfarbe"]
