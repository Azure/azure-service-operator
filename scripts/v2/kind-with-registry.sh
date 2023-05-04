#!/bin/sh
set -o errexit

# create registry container unless it already exists
reg_name='kind-registry'
reg_port='5000'
running="$(docker inspect -f '{{.State.Running}}' "${reg_name}" 2>/dev/null || true)"
if [ "${running}" != 'true' ]; then
  docker run \
    -d --restart=always -p "127.0.0.1:${reg_port}:5000" --name "${reg_name}" \
    registry:2
fi

# create a cluster with the local registry enabled in containerd
if [ -z "${SERVICE_ACCOUNT_ISSUER}" ]; then
cat <<EOF | kind create cluster --name "${KIND_CLUSTER_NAME}" --config=-
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
containerdConfigPatches:
- |-
  [plugins."io.containerd.grpc.v1.cri".registry.mirrors."localhost:${reg_port}"]
    endpoint = ["http://${reg_name}:${reg_port}"]
EOF
else
  echo "ISSUER: ${SERVICE_ACCOUNT_ISSUER}"
  echo "SERVICE_ACCOUNT_KEY_FILE: ${SERVICE_ACCOUNT_KEY_FILE}"
  echo "SERVICE_ACCOUNT_SIGNING_KEY_FILE: ${SERVICE_ACCOUNT_SIGNING_KEY_FILE}"

# Refer to https://github.com/Azure/azure-workload-identity/blob/2f92b47789ff94ba2578f73a0368589f8672f5c4/docs/book/src/development.md
# and https://github.com/Azure/azure-workload-identity/blob/main/scripts/create-kind-cluster.sh#L58-L91 for where this came from
  cat <<EOF | kind create cluster --name "${KIND_CLUSTER_NAME}" --config=-
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
nodes:
- role: control-plane
  extraMounts:
  - hostPath: ${SERVICE_ACCOUNT_KEY_FILE}
    containerPath: /etc/kubernetes/pki/sa.pub
  - hostPath: ${SERVICE_ACCOUNT_SIGNING_KEY_FILE}
    containerPath: /etc/kubernetes/pki/sa.key
  kubeadmConfigPatches:
  - |
    kind: InitConfiguration
    nodeRegistration:
    taints:
    - key: "kubeadmNode"
      value: "master"
      effect: "NoSchedule"
  - |
    kind: ClusterConfiguration
    apiServer:
      extraArgs:
        service-account-issuer: ${SERVICE_ACCOUNT_ISSUER}
        service-account-key-file: /etc/kubernetes/pki/sa.pub
        service-account-signing-key-file: /etc/kubernetes/pki/sa.key
    controllerManager:
      extraArgs:
        service-account-private-key-file: /etc/kubernetes/pki/sa.key
containerdConfigPatches:
- |-
  [plugins."io.containerd.grpc.v1.cri".registry.mirrors."localhost:${reg_port}"]
    endpoint = ["http://${reg_name}:${reg_port}"]
EOF
fi

# connect the registry to the cluster network
# (the network may already be connected)
docker network connect "kind" "${reg_name}" || true

# Document the local registry
# https://github.com/kubernetes/enhancements/tree/master/keps/sig-cluster-lifecycle/generic/1755-communicating-a-local-registry
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: ConfigMap
metadata:
  name: local-registry-hosting
  namespace: kube-public
data:
  localRegistryHosting.v1: |
    host: "localhost:${reg_port}"
    help: "https://kind.sigs.k8s.io/docs/user/local-registry/"
EOF

