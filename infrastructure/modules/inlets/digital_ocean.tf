provider digitalocean {
    token = var.digitalocean_token
}

data "digitalocean_account" "account" {}