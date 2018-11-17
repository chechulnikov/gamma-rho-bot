FROM golang:1.11

RUN mkdir /app
ADD ./src /app/
WORKDIR /app

RUN go build -o grammar-bot .
RUN chmod +x grammar-bot

CMD ["/app/grammar-bot"]
