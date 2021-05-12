# Build the manager binary
FROM golang:1.13.15 as builder

WORKDIR /workspace

# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

# Copy the go source
# COPY . ./
COPY main.go main.go
COPY apis/ apis/
COPY controllers/ controllers/
COPY pkg/ pkg/

# Build
# TODO: Use Makefile here -- right now it's awkward to do so because:
#   1. tools.mk is required for the makefile from the above directory, but Dockerfile can only look in its directory and below.
#   2. Having Dockerfile here but building it from above could work except that there's another Dockerfile and a .dockerignore
#      up above that break things. For now we just build by hand
# RUN make build

# TODO: Do we want CGO_ENALBED=0 and the other options below in the makefile?
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -o k8sinfra-controller main.go

# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=builder /workspace/k8sinfra-controller .
USER nonroot:nonroot
ENTRYPOINT ["/k8sinfra-controller"]
