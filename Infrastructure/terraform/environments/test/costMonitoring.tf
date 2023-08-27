module "cost_monitoring" {
  source = "../../modules/costMonitoring"

  aws_region   = var.aws_region
  default_tags = var.default_tags

  environment_name      = var.environment_name
  notification_receiver = var.email_notification_subscriber

  overall_spending_monitoring = {
    maxAllowedPriceInUsd = 20
  }
}