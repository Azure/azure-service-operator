#! /bin/bash 
set -e
set -x

GO111MODULE="on" go get sigs.k8s.io/kind@v0.6.1

# Clear down pkg file
rm -rf /go/pkg && rm -rf /go/src