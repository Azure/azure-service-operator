variable "service_principal_app_id" {
  type        = string
}

variable "service_principal_client_secret" {
  type        = string
}

variable "resource_group" {
  description = "Name of the resource group"
  type        = string
  default     = "azure-sql-operator"
}

variable "location" {
  description = "Cloud resource location"
  type        = string
  default     = "westus2"
}