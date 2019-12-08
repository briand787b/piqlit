provider "github" {
  token = var.github_token
  version = "~> 2.2"
  organization = "briand787b"
}

data "github_repository" "piqlit" {
    full_name = "briand787b/piqlit"
}
