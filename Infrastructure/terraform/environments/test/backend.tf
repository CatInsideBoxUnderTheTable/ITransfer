terraform {
  backend "s3" {
    bucket         = "terraform-remote-management-catinsideboxunderthetable-testenv"
    dynamodb_table = "terraform-remote-management"
    key            = "terraform.tfstate"
    region         = "eu-central-1"
    encrypt        = true
  }
}