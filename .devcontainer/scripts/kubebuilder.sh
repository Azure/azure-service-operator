#! /bin/bash 
set -e
set -x

os=$(go env GOOS)
arch=$(go env GOARCH)
kb_version="2.2.0"

# download kubebuilder and extract it to tmp
curl -sL https://go.kubebuilder.io/dl/${kb_version}/${os}/${arch} | tar -xz -C /tmp/

# move to a long-term location and put it on your path
# (you'll need to set the KUBEBUILDER_ASSETS env var if you put it somewhere else)
mv /tmp/kubebuilder_${kb_version}_${os}_${arch} /usr/local/kubebuilder
export PATH=$PATH:/usr/local/kubebuilder/bin

# Clear down pkg file
rm -rf /go/pkg && rm -rf /go/src