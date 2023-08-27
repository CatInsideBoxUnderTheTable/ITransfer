variable "state_bucket_name" {
  type = string
}

variable "state_dynamoDb_table_name" {
  type = string
}

variable "default_tags" {
  type = map(string)
}

variable "aws_region" {
  type = string
}