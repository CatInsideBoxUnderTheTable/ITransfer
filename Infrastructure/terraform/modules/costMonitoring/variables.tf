variable "default_tags" {
  type = map(string)
}

variable "aws_region" {
  type = string
}

variable "environment_name" {
  type = string
}

variable "notification_receiver" {
  type        = string
  description = "Email address of person which will receive notifications"
}

variable "overall_spending_monitoring" {
  type = object(
    {
      maxAllowedPriceInUsd =number
    }
  )
}