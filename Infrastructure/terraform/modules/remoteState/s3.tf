resource "aws_s3_bucket" "terraform_s3_state" {
  bucket        = var.state_bucket_name
  force_destroy = true
}

resource "aws_s3_bucket_versioning" "terraform_s3_state_versioning" {
  bucket = aws_s3_bucket.terraform_s3_state.id

  versioning_configuration {
    status = "Enabled"
  }
}

resource "aws_s3_bucket_server_side_encryption_configuration" "terraform_s3_state_encryption" {
  bucket = aws_s3_bucket.terraform_s3_state.id

  rule {
    apply_server_side_encryption_by_default {
      kms_master_key_id = aws_kms_key.terraform_s3_state_kms_key.arn
      sse_algorithm     = "aws:kms"
    }
  }
}

resource "aws_kms_key" "terraform_dynamodb_state_locks_key" {
  description             = "key used to encrypt bucket for state management"
  deletion_window_in_days = 7
  enable_key_rotation     = true
}

resource "aws_kms_alias" "terraform_s3_state_kms_key_alias" {
  name          = format("alias/%s-bucket-kms-key", var.state_bucket_name)
  target_key_id = aws_kms_key.terraform_s3_state_kms_key.key_id
}