resource "aws_budgets_budget" "overall_spending_budget" {
  name = "overall-spending-monitoring"
  limit_amount = var.overall_spending_monitoring.maxAllowedPriceInUsd
  budget_type  = "COST"
  limit_unit   = "USD"
  time_unit    = "MONTHLY"

  notification {
    comparison_operator        = "GREATER_THAN"
    notification_type          = "ACTUAL"
    threshold                  = 75
    threshold_type             = "PERCENTAGE"
    subscriber_email_addresses = [var.notification_receiver]
  }
}