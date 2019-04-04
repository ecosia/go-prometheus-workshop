FROM golang
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build .
CMD ./app
