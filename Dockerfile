FROM golang:latest AS builder

WORKDIR $GOPATH/src/github.com/rwlist/rwcore
COPY . .

RUN go get -d -v .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /app .

FROM scratch
COPY --from=builder /app ./

ENTRYPOINT ["./app"]
