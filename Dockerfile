# docker build --pull -t main:dev
FROM golang:alpine as builder
COPY . "${GOPATH}/src/package/app/"
WORKDIR "${GOPATH}/src/package/app"
RUN apk add --no-cache --upgrade \
		git \
		ca-certificates && \
	adduser -D -g '' app && \
	go get -d -v && \
	CGO_ENABLED=0 go build -a -o "/go/bin/main";

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /go/bin/main /go/bin/main
USER app
WORKDIR "/go/bin"
ENTRYPOINT ["/go/bin/main"]
