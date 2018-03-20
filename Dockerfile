FROM golang:1.10-alpine as builder

RUN apk update && apk upgrade && apk add --no-cache git pcre-dev gcc musl-dev

ADD . /go/src/github.com/fabbricadigitale/scimd
WORKDIR /go/src/github.com/fabbricadigitale/scimd

RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure -vendor-only

ARG LDFLAGS
RUN GOOS=linux go build -ldflags "${LDFLAGS} -extldflags -static" -a  -o /tmp/scimd .

FROM scratch
COPY --from=builder /tmp/scimd /scimd
ENTRYPOINT ["/scimd"]