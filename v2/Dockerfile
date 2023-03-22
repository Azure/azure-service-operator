# Note: This Dockerfile assumes that the binary has been built using the top-level Taskfile.yml

# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY ./bin/aso-controller .
COPY ./out/crds ./crds
USER nonroot:nonroot
ENTRYPOINT ["/aso-controller"]
