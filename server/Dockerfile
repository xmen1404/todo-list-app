FROM golang:1.17.6

WORKDIR /app

COPY . .
RUN go get .

EXPOSE 8000

CMD ["go", "run", "./server.go"]