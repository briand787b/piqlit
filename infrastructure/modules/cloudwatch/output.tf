output "cloudwatch_short_term_log_group" {
    value = aws_cloudwatch_log_group.short_term_logs.name
}