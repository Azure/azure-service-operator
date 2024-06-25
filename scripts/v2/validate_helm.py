#!/usr/local/bin/venv/bin/python3

import subprocess
from argparse import ArgumentParser
from deepdiff import DeepDiff
import logging
import yaml
import logging

logger = logging.getLogger()
logging.basicConfig(encoding='utf-8', level=logging.INFO)

def get_yaml(yaml_path):
    with open(yaml_path, "r") as f:
        aso_yaml_file = f.read()
        aso_yaml = yaml.full_load_all(aso_yaml_file)
        return aso_yaml

def get_helm_templates(helm_dir):
    helm_res = subprocess.check_output(
        f"helm template asov2 {helm_dir} --namespace=azureserviceoperator-system".split(" "))

    helm_resources = yaml.full_load_all(helm_res.decode("utf-8"))
    return helm_resources

def validate_helm(helm_dir, yaml_path):
    yaml = get_yaml(yaml_path)
    helm_templates = get_helm_templates(helm_dir)

    # We store the helm docs in a dictionary for easy lookup
    helm_resources = {}
    for resource in helm_templates:
        helm_resources[resource["metadata"]["name"]] = resource

    errors = []

    # We check against kustomize docs since helm may have more resources than kustomize. E.g Network Policies
    for resource in yaml:
        if resource["kind"] == "Namespace":
            continue

        resource_name = resource['metadata']['name']
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

    logger.info("Validation successful")


if __name__ == '__main__':

    args_parser = ArgumentParser()

    args_parser.add_argument("--helm-dir", type=str, help="path to the helm directory")
    args_parser.add_argument("--yaml-path", type=str, help="path to the aso yaml file. This should be a single YAML file containing all of the ASO resources")
    args = args_parser.parse_args()
    validate_helm(args.helm_dir, args.yaml_path)
