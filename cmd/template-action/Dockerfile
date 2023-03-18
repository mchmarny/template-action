
ARG BUILD_BASE=golang@sha256:57dbdd5c8fe24e357b15a4ed045b0b1607a99a1465b5101304ea39e11547be27
FROM $BUILD_BASE AS builder
WORKDIR /src/
COPY . /src/
ARG VERSION
ENV VERSION=${VERSION}
RUN CGO_ENABLED=0 go build -a -trimpath -ldflags="\
    -w -s -X main.version=$VERSION \
	-extldflags '-static'" \
    -mod vendor -o app cmd/template-action/main.go

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /src/app /app
ENTRYPOINT ["/app"]
