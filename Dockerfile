FROM golang:1.17.6-alpine

# to test
ENV CGO_ENABLED 0

RUN apk update && \
  apk --no-cache add git

WORKDIR /app
ADD . /app

RUN go mod tidy

RUN go install github.com/cosmtrek/air@v1.29.0 && \
  go install github.com/x-motemen/gore/cmd/gore@latest

# for hot reload
CMD ["air"]
