FROM golang:alpine
COPY . /tmp/app
WORKDIR /tmp/app
RUN go mod download
RUN go mod verify
RUN go build -o ftx-bot /tmp/app/main.go

#Run in alpine directly for optimization
FROM alpine

RUN mkdir /app
COPY --from=0 /tmp/app/ftx-bot /app/ftx-bot

WORKDIR /app
EXPOSE 3000
CMD ["/app/ftx-bot"]
