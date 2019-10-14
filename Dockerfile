# Build the manager binary
FROM golang:1.12.5 as builder

WORKDIR /workspace/
# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

# Copy the go source

COPY . ./

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -o manager main.go

# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM gcr.io/distroless/static:latest
ENV AZURE_CLIENT_ID "${AZURE_CLIENT_ID}"
ENV AZURE_CLIENT_SECRET "{AZURE_CLIENT_SECRET}"
ENV AZURE_SUBSCRIPTION_ID "${AZURE_SUBSCRIPTION_ID}"
ENV AZURE_TENANT_ID "${AZURE_TENANT_ID}"
ENV REQUEUE_AFTER "30"
WORKDIR /
COPY --from=builder /workspace/manager .
ENTRYPOINT ["/manager"]
