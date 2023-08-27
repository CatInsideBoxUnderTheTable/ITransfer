module "transfer_bucket" {
  source = "../../modules/transferBucket"

  aws_region   = var.aws_region
  default_tags = var.default_tags

  bucket_object_lifetime = 3
  environment_name       = var.environment_name
}