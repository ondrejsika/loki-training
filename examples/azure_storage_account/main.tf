locals {
  storage_account_name    = "loki${random_string.suffix.result}"
  storage_container_names = ["loki"]
  resource_group_name     = "loki-storage"
  location                = "germanywestcentral"
}

resource "random_string" "suffix" {
  length  = 10
  lower   = false
  upper   = false
  special = false
}


resource "azurerm_resource_group" "loki" {
  name     = local.resource_group_name
  location = local.location
}

resource "azurerm_storage_account" "loki" {
  name                     = local.storage_account_name
  resource_group_name      = azurerm_resource_group.loki.name
  location                 = azurerm_resource_group.loki.location
  account_tier             = "Standard"
  account_replication_type = "LRS"
}

resource "azurerm_storage_container" "loki" {
  for_each = toset(local.storage_container_names)

  name                  = each.value
  storage_account_id    = azurerm_storage_account.loki.id
  container_access_type = "private"
}


output "storage_account_name" {
  value = local.storage_account_name
}

output "storage_container_name" {
  value = local.storage_container_names
}

output "primary_access_key" {
  value     = azurerm_storage_account.loki.primary_access_key
  sensitive = true
}
