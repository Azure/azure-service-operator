#!/usr/bin/env python3
import subprocess

subprocess.run("pip install deepdiff".split(" "))
subprocess.run("pip install pyyaml".split(" "))

from argparse import ArgumentParser
from deepdiff import DeepDiff
import logging
import yaml
import logging

logger = logging.getLogger()
logging.basicConfig(encoding='utf-8', level=logging.INFO)

def get_aso_yaml_templates(aso_yaml_path):
    with open(aso_yaml_path, "r") as f:
        aso_yaml_file = f.read()
        aso_yaml = yaml.full_load_all(aso_yaml_file)
        return aso_yaml

def get_helm_templates(helm_dir):
    helm_res = subprocess.run(
        f"helm template asov2 {helm_dir} --namespace=azureserviceoperator-system".split(" "),
        capture_output=True)

    if helm_res.returncode != 0:
        error = helm_res.stderr.decode("utf-8")
        logger.error(f"helm template failed: \n{error}")
        exit(helm_res.returncode)

    helm_docs = yaml.full_load_all(helm_res.stdout.decode("utf-8"))
    return helm_docs

def validate_helm(helm_dir, aso_yaml_path):
    aso_yaml_templates = get_aso_yaml_templates(aso_yaml_path)
    helm_templates = get_helm_templates(helm_dir)

    # We store the helm docs in a dictionary for easy lookup
    helm_dict = {}
    for doc in helm_templates:
        helm_dict[doc["metadata"]["name"]] = doc

    error = False
    errors = []

    # We check against kustomize docs since helm may have more resources than kustomize. E.g Network Policies
    for doc in aso_yaml_templates:
        if doc["kind"] == "CustomResourceDefinition" or doc["kind"] == "Namespace":
            continue
        
        resource_name = doc['metadata']['name']
        if resource_name not in helm_dict:
            errors.append(f"Resource Kind: {doc['kind']}, Name:{doc['metadata']['name']} not found in helm")
            error = True

        diff = DeepDiff(doc, helm_dict[resource_name], ignore_order=True)
        if diff == {}:
            logger.info(f"{resource_name} matched")
        else: 
            error = True
            errors.append(f"Values in {resource_name} didn't match. \n {diff}")

    if error:
        for error in errors:
            logger.error(error)
        exit(1)

    logger.info("Validation successful")
    

if __name__ == '__main__':

    args_parser = ArgumentParser()

    args_parser.add_argument("--helm_dir", type=str, help="path to the output folder")
    args_parser.add_argument("--aso_yaml_path", type=str, help="path to the aso yaml file")
    args = args_parser.parse_args()

    validate_helm(args.helm_dir, args.aso_yaml_path)