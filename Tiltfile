# -*- mode: Python -*-

# set defaults

settings = {
    "allowed_contexts": [
        "kind-k8sinfra"
    ],
    "deploy_cert_manager": True,
    "preload_images_for_kind": True,
    "kind_cluster_name": "k8sinfra",
    "cert_manager_version": "v0.11.0",
}

keys = ["AZURE_SUBSCRIPTION_ID", "AZURE_TENANT_ID", "AZURE_CLIENT_SECRET", "AZURE_CLIENT_ID"]

# global settings
settings.update(read_json(
    "tilt-settings.json",
    default={},
))

if settings.get("trigger_mode") == "manual":
    trigger_mode(TRIGGER_MODE_MANUAL)

if "allowed_contexts" in settings:
    allow_k8s_contexts(settings.get("allowed_contexts"))

if "default_registry" in settings:
    default_registry(settings.get("default_registry"))


# Prepull all the cert-manager images to your local environment and then load them directly into kind. This speeds up
# setup if you're repeatedly destroying and recreating your kind cluster, as it doesn't have to pull the images over
# the network each time.
def deploy_cert_manager():
    registry = "quay.io/jetstack"
    version = settings.get("cert_manager_version")
    images = ["cert-manager-controller", "cert-manager-cainjector", "cert-manager-webhook"]

    if settings.get("preload_images_for_kind"):
        for image in images:
            local("docker pull {}/{}:{}".format(registry, image, version))
            local("kind load docker-image --name {} {}/{}:{}".format(settings.get("kind_cluster_name"), registry, image,
                                                                     version))

    local("kubectl apply -f https://github.com/jetstack/cert-manager/releases/download/{}/cert-manager.yaml".format(
        version))

    # wait for the service to become available
    local("kubectl wait --for=condition=Available --timeout=300s apiservice v1beta1.webhook.cert-manager.io")


# Users may define their own Tilt customizations in tilt.d. This directory is excluded from git and these files will
# not be checked in to version control.
def include_user_tilt_files():
    user_tiltfiles = listdir("tilt.d")
    for f in user_tiltfiles:
        include(f)


tilt_helper_dockerfile_header = """
# Tilt image
FROM golang:1.13.8 as tilt-helper
# Support live reloading with Tilt
RUN wget --output-document /restart.sh --quiet https://raw.githubusercontent.com/windmilleng/rerun-process-wrapper/master/restart.sh  && \
    wget --output-document /start.sh --quiet https://raw.githubusercontent.com/windmilleng/rerun-process-wrapper/master/start.sh && \
    chmod +x /start.sh && chmod +x /restart.sh
"""

tilt_dockerfile_header = """
FROM gcr.io/distroless/base:debug as tilt
WORKDIR /
COPY --from=tilt-helper /start.sh .
COPY --from=tilt-helper /restart.sh .
COPY manager .
"""


b64_cmd = "source .env && echo ${{{}}} | tr -d '\n' | base64 | tr -d '\n'"
yaml_substitution = "${{{}}}"


def append_arg_for_container_in_deployment(yaml_stream, name, namespace, contains_image_name, args):
    for item in yaml_stream:
        if item["kind"] == "Deployment" and item.get("metadata").get("name") == name and item.get("metadata").get("namespace") == namespace:
            containers = item.get("spec").get("template").get("spec").get("containers")
            for container in containers:
                if contains_image_name in container.get("image"):
                    container.get("args").extend(args)


def fixup_yaml_empty_arrays(yaml_str):
    yaml_str = yaml_str.replace("conditions: null", "conditions: []")
    return yaml_str.replace("storedVersions: null", "storedVersions: []")


# Build k8s-infra and add feature gates
def k8s_infra():
    # Apply the kustomized yaml for this provider
    yaml = str(kustomize("./config/credentials"))
    substitutions = settings.get("kustomize_substitutions", {})
    for substitution in substitutions:
        value = substitutions[substitution]
        yaml = yaml.replace("${" + substitution + "}", value)

    for k in keys:
        b64_secret = str(local(b64_cmd.format(k)))
        if not b64_secret:
            fail("missing env var {}".format(k))
        yaml = yaml.replace("${" + k + "_B64}", b64_secret)

    # add extra_args if they are defined
    if settings.get("extra_args"):
        azure_extra_args = settings.get("extra_args").get("k8sinfra")
        if azure_extra_args:
            yaml_dict = decode_yaml_stream(yaml)
            append_arg_for_container_in_deployment(yaml_dict, "k8sinfra-controller-manager", "k8sinfra-system",
                                                   "k8sinfra-controller", azure_extra_args)
            yaml = str(encode_yaml_stream(yaml_dict))
            yaml = fixup_yaml_empty_arrays(yaml)

    # Set up a local_resource build of the provider's manager binary.
    local_resource(
        "manager",
        cmd='mkdir -p .tiltbuild;CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags \'-extldflags "-static"\' -o .tiltbuild/manager',
        deps=["./api", "./main.go", "./pkg", "./controllers", "./cloud", "./exp", "./internal"]
    )

    dockerfile_contents = "\n".join([
        tilt_helper_dockerfile_header,
        tilt_dockerfile_header,
    ])

    entrypoint = ["sh", "/start.sh", "/manager"]
    extra_args = settings.get("extra_args")
    if extra_args:
        entrypoint.extend(extra_args)

    # Set up an image build for the provider. The live update configuration syncs the output from the local_resource
    # build into the container.
    docker_build(
        ref="kind-registry:5000/fake/k8s-infra-controller",
        context="./.tiltbuild/",
        dockerfile_contents=dockerfile_contents,
        target="tilt",
        entrypoint=entrypoint,
        only="manager",
        live_update=[
            sync("./.tiltbuild/manager", "/manager"),
            run("sh /restart.sh"),
        ],
    )

    k8s_yaml(blob(yaml))


##############################
# Actual work happens here
##############################

include_user_tilt_files()

if settings.get("deploy_cert_manager"):
    deploy_cert_manager()

k8s_infra()
