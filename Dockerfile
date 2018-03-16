FROM golang:1.10 as builder

ARG LDFLAGS
ARG GOOS=linux

ENV GOOS ${GOOS}

ADD . /go/src/github.com/fabbricadigitale/scimd
WORKDIR /go/src/github.com/fabbricadigitale/scimd
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure
RUN apt-get update && apt-get install -y libpcre++-dev
RUN go build -ldflags "${LDFLAGS} -extldflags -static" -o /tmp/scimd .

FROM scratch
COPY --from=builder /tmp/scimd /scimd
ENTRYPOINT ["/scimd"]