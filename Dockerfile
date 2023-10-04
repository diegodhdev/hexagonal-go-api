FROM golang:alpine AS build

RUN apk add --update git
WORKDIR /go/src/github.com/diegodhdev/hexagonal-go-api
COPY . .
RUN CGO_ENABLED=0 go build -o /go/bin/codelytv-mooc-api requests/cmd/api/main.go

# Building image with the binary
FROM scratch
COPY --from=build /go/bin/codelytv-mooc-api /go/bin/codelytv-mooc-api
ENTRYPOINT ["/go/bin/codelytv-mooc-api"]
