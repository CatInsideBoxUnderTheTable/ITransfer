resource "aws_budgets_budget" "storage_usage_budget" {
  name         = "${var.tracked_bucket_usage_monitoring.tagValue}-tag-usage-monitoring"
  limit_amount = var.tracked_bucket_usage_monitoring.maxAllowedSizeInGb
  time_unit    = "MONTHLY"
  budget_type  = "USAGE"
  limit_unit   = "GB"

  cost_filter {
    name   = "TagKeyValue"
    values = ["TagKey${"$"}${var.tracked_bucket_usage_monitoring.tagValue}"]
  }

  notification {
    comparison_operator        = "GREATER_THAN"
    notification_type          = "ACTUAL"
    threshold                  = 75
    threshold_type             = "PERCENTAGE"
    subscriber_email_addresses = [var.notification_receiver]
  }
}