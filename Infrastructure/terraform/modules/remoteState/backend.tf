terraform {
  backend "s3" {
    encrypt = true
    bucket = var.state_bucket_name
    dynamodb_table = var.state_dynamoDb_table_name
    key    = "terraform.tfstate"
    region = "eu-central-1"
  }
}