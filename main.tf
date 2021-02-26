terraform {
  backend "remote" {
    organization = "devopsperu-demo"

    workspaces {
      name = "app1-infra-provisioning"
    }
  }
}

module "storage" {
    source  = "app.terraform.io/devopsperu-demo/storage/azurerm"
    version = "1.0.2"
  # insert required variables here
    account_tier = "Standard"
    account_replication_type = "LRS"
    type = var.type
    stage = var.stage
}

module "cdn" {
  source  = "app.terraform.io/devopsperu-demo/cdn/azurerm"
  version = "1.0.2"
  static_endpoint = replace(replace(module.storage.primaryWebEndpoint,"https://",""),"/","")
  type = var.type
  stage = var.stage
}