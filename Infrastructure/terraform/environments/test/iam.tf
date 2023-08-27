module "iam" {
  source = "../../modules/iam"

  aws_region                         = var.aws_region
  default_tags                       = var.default_tags
  environment_name                   = var.environment_name
  bucket_console_users               = ["default"]
  transfer_bucket_arn                = module.transfer_bucket.transfer_bucket_arn
  transfer_bucket_encryption_key_arn = module.transfer_bucket.transfer_bucket_encryption_key_arn
}