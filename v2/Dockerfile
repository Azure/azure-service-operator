ARG VERSION_FLAGS
ARG CONTROLLER_APP=aso-controller

# Build the manager binary
FROM golang:1.20 as builder

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
RUN CGO_ENABLED=0 GO111MODULE=on go build ${VERSION_FLAGS} -o ./bin/${CONTROLLER_APP} ./cmd/controller/

# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=builder /workspace/bin/${CONTROLLER_APP} .
COPY ./out/crds ./crds
USER nonroot:nonroot
ENTRYPOINT ["/aso-controller"]
