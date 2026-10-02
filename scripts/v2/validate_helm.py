#!/usr/local/bin/venv/bin/python3

import subprocess
from argparse import ArgumentParser
from deepdiff import DeepDiff
import logging
import yaml

WORKLOAD_IDENTITY_ROLE = "azureserviceoperator-workload-identity-token-creator-role"
WORKLOAD_IDENTITY_BINDING = "azureserviceoperator-workload-identity-token-creator-rolebinding"
WORKLOAD_IDENTITY_RESOURCES = {WORKLOAD_IDENTITY_ROLE, WORKLOAD_IDENTITY_BINDING}

logger = logging.getLogger()
logging.basicConfig(encoding='utf-8', level=logging.INFO)

def get_yaml(yaml_path):
    with open(yaml_path, "r") as f:
        aso_yaml_file = f.read()
        aso_yaml = yaml.safe_load_all(aso_yaml_file)
        return aso_yaml

def get_helm_templates(helm_dir, workload_identity_auth_mode="relaxed"):
    command = (
        f"helm template asov2 {helm_dir} --namespace=azureserviceoperator-system "
        f"--set workloadIdentityAuthMode={workload_identity_auth_mode}"
    )
    helm_res = subprocess.check_output(command.split(" "))

    return list(yaml.safe_load_all(helm_res.decode("utf-8")))

def resources_by_name(resources):
    return {resource["metadata"]["name"]: resource for resource in resources}

def validate_workload_identity_rbac(helm_dir):
    # Relaxed mode uses the controller's projected token and must not grant permission to mint tokens
    # for namespace ServiceAccounts.
    relaxed_resources = resources_by_name(get_helm_templates(helm_dir, "relaxed"))
    unexpected = WORKLOAD_IDENTITY_RESOURCES.intersection(relaxed_resources)
    if unexpected:
        raise AssertionError(f"relaxed mode rendered strict-only RBAC: {sorted(unexpected)}")

    # Strict mode needs a TokenRequest grant for the well-known default ServiceAccount name.
    strict_resources = resources_by_name(get_helm_templates(helm_dir, "strict"))
    role = strict_resources.get(WORKLOAD_IDENTITY_ROLE)
    binding = strict_resources.get(WORKLOAD_IDENTITY_BINDING)
    if role is None or binding is None:
        raise AssertionError("strict mode did not render workload identity TokenRequest RBAC")

    expected_rules = [{
        "apiGroups": [""],
        "resources": ["serviceaccounts/token"],
        "resourceNames": ["aso-workload"],
        "verbs": ["create"],
    }]
    # resourceNames is the security boundary here. The shipped role must never permit token
    # creation for arbitrary ServiceAccounts or custom names configured by individual secrets.
    if role.get("rules") != expected_rules:
        raise AssertionError(f"strict mode rendered unexpected TokenRequest rules: {role.get('rules')}")

    # Verify the grant is bound to the controller ServiceAccount in the release namespace.
    expected_subject = {
        "kind": "ServiceAccount",
        "name": "azureserviceoperator-default",
        "namespace": "azureserviceoperator-system",
    }
    if binding.get("subjects") != [expected_subject]:
        raise AssertionError(f"strict mode rendered unexpected TokenRequest binding: {binding.get('subjects')}")

def validate_helm(helm_dir, yaml_path):
    yaml = get_yaml(yaml_path)
    helm_templates = get_helm_templates(helm_dir)

    # We store the helm docs in a dictionary for easy lookup
    helm_resources = resources_by_name(helm_templates)

    errors = []

    # We check against kustomize docs since helm may have more resources than kustomize. E.g Network Policies
    for resource in yaml:
        if resource["kind"] == "Namespace":
            continue

        resource_name = resource['metadata']['name']
        # Kustomize always includes the exact-name TokenRequest RBAC, while Helm renders it only
        # for strict mode. validate_workload_identity_rbac() checks both Helm modes explicitly.
        if resource_name in WORKLOAD_IDENTITY_RESOURCES:
            continue
        if resource_name not in helm_resources:
            errors.append(f"Resource Kind: {resource['kind']}, Name:{resource['metadata']['name']} not found in helm")
            continue

        diff = DeepDiff(helm_resources[resource_name], resource, ignore_order=True)
        if diff == {}:
            logger.info(f"{resource_name} matched")
        else:
            errors.append(f"Values in {resource_name} didn't match. \n {diff}")

    if errors:
        for error in errors:
            logger.error(error)
        exit(1)

    validate_workload_identity_rbac(helm_dir)

    logger.info("Validation successful")


if __name__ == '__main__':

    args_parser = ArgumentParser()

    args_parser.add_argument("--helm-dir", type=str, help="path to the helm directory")
    args_parser.add_argument("--yaml-path", type=str, help="path to the aso yaml file. This should be a single YAML file containing all of the ASO resources")
    args = args_parser.parse_args()
    validate_helm(args.helm_dir, args.yaml_path)
