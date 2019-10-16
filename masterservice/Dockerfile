FROM golang:1.11.2
COPY . /go/src/github.com/sathishkumar64/funbook/masterservice
WORKDIR /go/src/github.com/sathishkumar64/funbook/masterservice
RUN go install -ldflags="-s -w" ./cmd/...