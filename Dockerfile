# Build the books binary
FROM golang:1.15 as builder

WORKDIR /workspace
# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

# Copy the go source
COPY main.go main.go


# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -o books main.go

# Use distroless as minimal base image to package the books binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM gcr.io/distroless/static:nonroot
LABEL maintainer="Antonio Tolentino - atcjunior@gmail.com"
WORKDIR /
COPY --from=builder /workspace/books .
USER 65532:65532

EXPOSE 8080

ENTRYPOINT ["/books"]
