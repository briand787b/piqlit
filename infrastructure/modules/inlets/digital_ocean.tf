provider digitalocean {
    token = var.digitalocean_token
    version = "~> 1.11"
}

data "digitalocean_account" "account" {}