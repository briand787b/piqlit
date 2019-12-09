provider digitalocean {
    token = var.digitalocean_token
    version = "~> 1.11"
}

data "digitalocean_account" "account" {}

# # Create a new Web Droplet in the nyc2 region
# resource "digitalocean_droplet" "inlet" {
#   image  = "ubuntu-18-04-x64"
#   name   = "inlet"
#   region = "nyc2"
#   size   = "s-1vcpu-1gb"
# }