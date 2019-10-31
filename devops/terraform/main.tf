provider "azurerm" {
  version = "~>1.31.0"
}

provider "null" {
  version = "~>2.1.0"
}

provider "azuread" {
  version = "~>0.4.0"
}

provider "azurestack" {
  version = "=0.8.0"
}

data "azuread_service_principal" "cluster" {
  application_id = var.service_principal_app_id
}

data "azurerm_subscription" "primary" {}

data "azurerm_client_config" "auth" {}

locals {
  cluster_name = "aso-kubernetes"
  operator_namespace = "azureoperator-system"
}

// Generate a random string that we can use to uniquely identify resources
resource "random_pet" "build" {
  length  = 1
  keepers = {
    # Generate a new pet name each time we generate a new resource group
    resource_group_id = azurerm_resource_group.operator.id
  }
}

// Create Resource Group
resource "azurerm_resource_group" "operator" {
  name     = var.resource_group
  location = var.location
}

// Create Azure VNET
resource "azurerm_virtual_network" "cluster_pool" {
  name                = "aso-cluster-pool-vnet"
  resource_group_name = azurerm_resource_group.operator.name
  location            = azurerm_resource_group.operator.location
  address_space       = ["10.10.0.0/24"]
}

// Add Subnet to Azure VNET
resource "azurerm_subnet" "cluster_pool" {
  name                 = "aso-cluster-pool-subnet"
  resource_group_name  = azurerm_resource_group.operator.name
  virtual_network_name = azurerm_virtual_network.cluster_pool.name
  address_prefix       = "10.10.0.0/24"
  service_endpoints    = ["Microsoft.Sql", "Microsoft.Storage"]

  // Work around https://github.com/Azure/AKS/issues/400 - AKS updates route table on subnet, but
  // Terraform isn't aware of this change and tries to revert it.
  lifecycle {
    ignore_changes = ["route_table_id"]
  }
}

// Create Kubernetes Cluster
resource "azurerm_kubernetes_cluster" "operator" {
  name                = local.cluster_name
  resource_group_name = azurerm_resource_group.operator.name
  location            = azurerm_resource_group.operator.location
  dns_prefix          = "aso-kubernetes"

  agent_pool_profile {
    name            = "default"
    vm_size         = "Standard_D2_v2"
    os_disk_size_gb = 30
    vnet_subnet_id  = azurerm_subnet.cluster_pool.id
  }

  service_principal {
    client_id     = var.service_principal_app_id
    client_secret = var.service_principal_client_secret
  }
}

// Create a Managed Service Identity
resource "azurerm_user_assigned_identity" "manager_identity" {
  resource_group_name = azurerm_resource_group.operator.name
  location            = azurerm_resource_group.operator.location

  name = "aso-manager-identity"
}

// Create a role that can manage identities (with a scope of our manager identity) and give it to the AKS SP
resource "azurerm_role_assignment" "aad_role" {
  role_definition_name = "Managed Identity Operator"
  scope                = azurerm_user_assigned_identity.manager_identity.id
  principal_id         = data.azuread_service_principal.cluster.id
}

// Create a role that can manage azure resources in our subscription and give it to our manager identity
resource "azurerm_role_assignment" "manager_role" {
  role_definition_name = "Owner"
  scope                = data.azurerm_subscription.primary.id
  principal_id         = azurerm_user_assigned_identity.manager_identity.principal_id
}

// Setup Terraform Kubernetes provider to point at our new AKS cluster
provider "kubernetes" {
  host                   = "${azurerm_kubernetes_cluster.operator.kube_config.0.host}"
  client_certificate     = "${base64decode(azurerm_kubernetes_cluster.operator.kube_config.0.client_certificate)}"
  client_key             = "${base64decode(azurerm_kubernetes_cluster.operator.kube_config.0.client_key)}"
  cluster_ca_certificate = "${base64decode(azurerm_kubernetes_cluster.operator.kube_config.0.cluster_ca_certificate)}"
}

// Create namespace to deploy the operator into
resource "kubernetes_namespace" "operator-system" {
  metadata {
    name = local.operator_namespace
  }
}

// Create Kubernetes secrets
resource "kubernetes_secret" "operator" {
  metadata {
    name = "azureoperatorsettings"
    namespace = local.operator_namespace
  }

  data = {
    AZURE_CLIENT_ID       = data.azurerm_client_config.auth.client_id
    AZURE_CLIENT_SECRET   = var.service_principal_client_secret
    AZURE_SUBSCRIPTION_ID = data.azurerm_client_config.auth.subscription_id
    AZURE_TENANT_ID       = data.azurerm_client_config.auth.tenant_id
  }
}

// Give the default service account on the operator namespace "cluster-admin"
// access to the namespace where the operator deploys the resources by default.
resource "kubernetes_cluster_role_binding" "operator-admin" {
  metadata {
    name = "cluster-admin-aso"
  }
  role_ref {
    api_group = "rbac.authorization.k8s.io"
    kind      = "ClusterRole"
    name      = "cluster-admin"
  }
  subject {
    kind      = "User"
    name      = "admin"
    api_group = "rbac.authorization.k8s.io"
  }
  subject {
    kind      = "ServiceAccount"
    name      = "default"
    namespace = local.operator_namespace
  }
  subject {
    kind      = "Group"
    name      = "system:serviceaccount"
    api_group = "rbac.authorization.k8s.io"
  }
}

// Create an ACR to hold our images
resource "azurerm_container_registry" "acr" {
  name                = "asoregistry${random_pet.build.id}"
  resource_group_name = azurerm_resource_group.operator.name
  location            = azurerm_resource_group.operator.location
  sku                 = "Basic"
}

// Give the AKS SP access to pull images from our ACR
resource "azurerm_role_assignment" "image_role" {
  role_definition_name = "AcrPull"
  scope                = azurerm_container_registry.acr.id
  principal_id         = data.azuread_service_principal.cluster.id
}

// Setup kubeconfig to facilitate running kubectl outside of terraform
resource "local_file" "kubeconfig" {
  content  = azurerm_kubernetes_cluster.operator.kube_config_raw
  filename = pathexpand("~/.kube/config")
  directory_permission = "0600"
}

output "registry_url" {
  value = azurerm_container_registry.acr.login_server
}