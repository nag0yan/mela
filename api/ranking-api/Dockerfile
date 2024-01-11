FROM golang:1.21.5
WORKDIR /usr/src/app
COPY go.mod .
COPY go.sum .
COPY *.go .
RUN go build -o ./bin
EXPOSE 8080
CMD ["./bin"]