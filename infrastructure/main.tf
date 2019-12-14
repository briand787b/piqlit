provider "aws" {
  profile = "sbox"
  version = "~> 2.0"
  region = "us-east-1"
}

// backend
terraform {
    backend "s3" {
        bucket  = "piqlit-terraform-state"
        key     = "terraform.tfstate"
        region  = "us-east-1"
        profile = "sbox"
    }
}

// data
data "aws_availability_zones" "available" {}

module "cloudwatch" {
    source = "./modules/cloudwatch"
}

module "codesuite" {
    source = "./modules/codesuite"

    github_token = var.github_token
    postman_api_key = var.postman_api_key
    postman_collection_id = var.postman_collection_id 
    codebuild_log_group = module.cloudwatch.cloudwatch_short_term_log_group
}

module "inlets" {
    source = "./modules/inlets"

    digitalocean_token = var.digitalocean_token
}