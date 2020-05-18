# Tips & tricks for Kubebuilder

1. Kubebuilder documentation recommends you to install Kubebuilder using the following commands.

```shell
os=$(go env GOOS)
arch=$(go env GOARCH)

# download kubebuilder and extract it to tmp
curl -sL https://go.kubebuilder.io/dl/2.0.1/${os}/${arch} | tar -xz -C /tmp/

# move to a long-term location and put it on your path
# (you'll need to set the KUBEBUILDER_ASSETS env var if you put it somewhere else)
sudo mv /tmp/kubebuilder_2.0.1_${os}_${arch} /usr/local/kubebuilder
export PATH=$PATH:/usr/local/kubebuilder/bin
```

Sometimes, the `os` and `arch` do not resolve correctly on MacOS. You might need to substitute them manually in the `curl` command based on what `go env GOOS` and `go env GOARCH` return.

2. You need the environment variable `GO111MODULE` to be set to `on`. You can set this as follows.

```shell
export GO111MODULE=on
```

3. If you get errors with `make install`, check the kube version and make sure the server version is atleast 1.14

```shell
kubectl version
Client Version: version.Info{Major:"1", Minor:"13", GitVersion:"v1.13.2", GitCommit:"cff46ab41ff0bb44d8584413b598ad8360ec1def", GitTreeState:"clean", BuildDate:"2019-01-13T23:15:13Z", GoVersion:"go1.11.4", Compiler:"gc", Platform:"darwin/amd64"}
Server Version: version.Info{Major:"1", Minor:"14", GitVersion:"v1.14.3", GitCommit:"5e53fd6bc17c0dec8434817e69b04a25d8ae0ff0", GitTreeState:"clean", BuildDate:"2019-06-06T01:36:19Z", GoVersion:"go1.12.5", Compiler:"gc", Platform:"linux/amd64"}
```

4. If you get errors with `make deploy`, look at this issue if you see the error listed here - https://github.com/kubernetes-sigs/kustomize/issues/1373
