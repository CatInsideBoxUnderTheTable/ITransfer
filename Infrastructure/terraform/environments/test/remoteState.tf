module "remote_state" {
  source = "./../../modules/remoteState"

  aws_region   = var.aws_region
  default_tags = var.default_tags

  state_bucket_name         = "terraform-remote-management-catinsideboxunderthetable-testenv"
  state_dynamoDb_table_name = "terraform-remote-management"
}