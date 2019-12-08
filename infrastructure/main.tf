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

module "codesuite" {
    source = "./modules/codesuite"

    github_token = var.github_token
}