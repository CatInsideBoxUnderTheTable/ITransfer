module "remote_state" {
  source    = "./../../modules/remoteState"
  providers = { aws = aws }

  state_bucket_name         = "terraform-remote-management-bucket-CatInsideBoxUnderTheTable"
  state_dynamoDb_table_name = "terraform-remote-management"
}