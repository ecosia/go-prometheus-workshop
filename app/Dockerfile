FROM golang
WORKDIR /go/src/github.com/ecosia/go-prometheus-workshop
COPY . .
RUN go get -v .
RUN go build .
CMD ./go-prometheus-workshop
