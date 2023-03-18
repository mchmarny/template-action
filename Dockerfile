FROM golang:buster AS builder
WORKDIR /src/
COPY . /src/
ARG VERSION
ENV VERSION ${VERSION}
RUN CGO_ENABLED=0 go build -trimpath -ldflags="\
    -w -s -X main.version=$VERSION \
	-extldflags '-static'" \
    -mod vendor -o app cmd/action/main.go

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /src/app /app
ENTRYPOINT ["/app"]
