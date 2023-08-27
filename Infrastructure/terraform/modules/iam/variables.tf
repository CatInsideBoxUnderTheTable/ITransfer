variable "default_tags" {
  type = map(string)
}

variable "aws_region" {
  type = string
}

variable "environment_name" {
  type = string
}

variable "transfer_bucket_arn" {
  type = string
}

variable "transfer_bucket_encryption_key_arn" {
  type = string
}

variable "bucket_console_users" {
  type        = set(string)
  description = "list of usernames that will be created with required permissions to use bucket. Remember, for security reasons, they will have to generate their access keys."
}