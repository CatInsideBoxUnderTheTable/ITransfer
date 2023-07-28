variable "default_tags" {
  type = map(string)
}

variable "aws_region" {
  type = string
}

variable "environment_name" {
  type = string
}

variable "bucket_object_lifetime" {
  type = number
  description = "Defines how long objects stored in s3 should live. Measured in days"
}

