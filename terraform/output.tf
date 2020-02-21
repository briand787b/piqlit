output "piqlit_http_clone_url" {
  value = module.codesuite.piqlit_http_clone_url
}

output "digitalocean_account_email" {
  value = module.inlets.digitalocean_account_email
}

output "cloudwatch_log_group_short_term_arn" {
  value = module.cloudwatch.cloudwatch_short_term_log_group.arn
}