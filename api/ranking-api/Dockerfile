FROM golang:1.21.5 AS builder
WORKDIR /usr/src/app/
COPY go.mod go.sum /usr/src/app/
RUN go mod download
COPY *.go /usr/src/app/
RUN CGO_ENABLED=0 go build -o /bin/app

FROM scratch AS runner
ENV GIN_MODE=release
COPY --from=builder /bin/app /bin/app
EXPOSE 8080
ENTRYPOINT ["/bin/app"]