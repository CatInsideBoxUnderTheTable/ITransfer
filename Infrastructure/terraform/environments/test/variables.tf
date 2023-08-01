variable "aws_region" {
  type    = string
  default = "eu-central-1"
}

variable "environment_name" {
  type    = string
  default = "testenv"
}

variable "default_tags" {
  type = map(string)
  default = {
    Owner       = "CatInsideBoxUnderTheTable"
    Environment = "TestEnv"
    ManagedBy   = "Terraform"
  }
}

variable "email_notification_subscriber" {
  type = string
}

A