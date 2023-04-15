# Also update GitHub Actions workflow when bumping
FROM --platform=${BUILDPLATFORM} docker.io/library/golang:1.20 AS builder
WORKDIR /src/
RUN GOARCH=amd64 go install golang.org/x/vuln/cmd/govulncheck@latest
COPY . .
RUN govulncheck ./...
ARG TARGETOS TARGETARCH TARGETVARIANT
RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg \
    CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH GOARM=${TARGETVARIANT#v} go build -ldflags='-extldflags=-static' -o /bin/service

FROM docker.io/library/alpine:latest
COPY --from=builder /bin/service /service
EXPOSE 8080
ENTRYPOINT ["./service"]
