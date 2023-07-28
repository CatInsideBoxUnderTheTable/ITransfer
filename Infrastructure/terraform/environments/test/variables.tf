variable "aws_region" {
  type = string
  default = "eu-central-1"
}

variable "environment_name" {
  type = string
  default = "testEnv"
}

variable "default_tags" {
  type = map(string)
  default = {
    Owner       = "CatInsideBoxUnderTheTable"
    Environment = "Test"
  }
}

