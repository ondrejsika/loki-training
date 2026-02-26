terraform {
  required_providers {
    azurerm = {
      source = "hashicorp/azurerm"
    }
    random = {
      source = "hashicorp/random"
    }
  }
}

provider "azurerm" {
  features {}
  subscription_id = "5768238c-1ecd-49ab-83cc-b09bf70a7bff" // SikaLabs DEV
}
