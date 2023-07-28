variable "aws_region" {
  type = string
  default = "eu-central-1"
}
variable "default_tags" {
  type = map(string)
  default = {
    Owner       = "CatInsideBoxUnderTheTable"
    Environment = "Test"
  }
}