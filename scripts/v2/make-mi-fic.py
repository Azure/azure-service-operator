#!/usr/bin/env python3

import subprocess
import json
import os
import random
import string
import argparse
import time
import datetime
import ast


def create_role_assignment(subscription_id: str, object_id: str) -> str:
    cmd = [
        "az",
        "role",
        "assignment",
        "create",
        "--assignee-object-id", object_id,
        "--assignee-principal-type", "ServicePrincipal",
        "--role", "Owner",
        "--subscription", subscription_id,
        "--scope", f"/subscriptions/{subscription_id}"
    ]
    result = subprocess.check_output(cmd, text=True)
    return json.loads(result)["id"]


def retry_create_role_assignment(subscription_id: str, object_id: str) -> str:
    start_time = datetime.datetime.utcnow()
    while True:
        try:
            return create_role_assignment(subscription_id, object_id)
        except subprocess.CalledProcessError:
            if datetime.datetime.utcnow() - start_time > datetime.timedelta(minutes=1):
                raise Exception("timed out waiting for role assignment creation")
            print("Retrying role assignment creation...")
            time.sleep(5)


def generate_random_string(length=12) -> str:
    return "".join(random.choices(string.ascii_lowercase + string.digits, k=length))


def get_az_identity_show_args(resource_group: str, identity_name: str, query: str) -> [str]:
    return [
        "az",
        "identity",
        "show",
        "--resource-group",
        resource_group,
        "--name",
        identity_name,
        "--query",
        "'{}'".format(query),
        "-otsv"]


def main():
    # Create the parser
    parser = argparse.ArgumentParser(
        description='Make a managed identity, Federated Identity Credentials, and assign associated permissions')

    # Add arguments
    parser.add_argument('-g', '--resource-group', type=str, help='Resource group name')
    parser.add_argument('-i', '--issuer', type=str, help='Issuer name')
    parser.add_argument('-d', '--directory', type=str, help='Directory path')
    parser.add_argument('-s', '--subject', type=str, help='Subject')

    # Parse arguments
    args = parser.parse_args()

    # Access arguments
    resource_group = args.resource_group
    issuer = args.issuer
    directory = args.directory
    subject = args.subject

    if not subject:
        subject = "system:serviceaccount:azureserviceoperator-system:azureserviceoperator-default"

    if not resource_group or not issuer or not directory:
        parser.print_help()
        exit(1)

    existing_identity = False

    subscription_id = os.getenv("AZURE_SUBSCRIPTION_ID")
    identity_name = os.getenv("AZURE_IDENTITY_NAME")
    if identity_name:
        existing_identity = True
        resource_group = os.getenv("AZURE_IDENTITY_RG")
    else:
        # TODO: Might almost be easier to use the Azure SDK here...
        identities = subprocess.check_output(
            ["az", "identity", "list", "--resource-group", resource_group, "--query", "[].name"])
        identities = identities.decode('utf-8')
        identities = ast.literal_eval(identities)
        if len(identities) > 0:
            identity_name = identities[0]
            print(f"An identity '{identity_name}' already exists, not creating another one")
            existing_identity = True

    if not existing_identity:
        identity_name = f"mi{generate_random_string()}"
        print(f"Generated new identity name: {identity_name}")
        subprocess.run(
            ["az", "identity", "create", "--name", identity_name, "--resource-group", resource_group],
            check=True)

    subprocess.run([
        "az",
        "identity",
        "federated-credential",
        "create",
        "--identity-name",
        identity_name,
        "--name",
        "fic",
        "--resource-group",
        resource_group,
        "--issuer",
        issuer,
        "--subject",
        subject,
        "--audiences",
        "api://AzureADTokenExchange"
    ], check=True)


    # need shell=True for the --query features
    user_assigned_client_id = subprocess.check_output(
        "az identity show --resource-group {} --name {} --query 'clientId' -otsv".format(
            resource_group,
            identity_name),
        text=True,
        shell=True)

    # need shell=True for the --query features
    user_assigned_object_id = subprocess.check_output(
        "az identity show --resource-group {} --name {} --query 'principalId' -otsv".format(
            resource_group,
            identity_name),
        text=True,
        shell=True)

    if not existing_identity:
        role_assignment_id = retry_create_role_assignment(subscription_id, user_assigned_object_id)
        with open(os.path.join(directory, "azure/roleassignmentid.txt"), "w") as f:
            f.write(role_assignment_id)
    else:
        with open(os.path.join(directory, "azure/fic.txt"), "w") as f:
            f.write("fic")

    with open(os.path.join(directory, "azure/miclientid.txt"), "w") as f:
        f.write(user_assigned_client_id)



if __name__ == "__main__":
    main()
