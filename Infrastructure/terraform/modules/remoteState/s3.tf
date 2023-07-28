# after you create bucket by hand you will need to import it
# 'import module.remote_state.aws_s3_bucket.terraform_state_bucket "terraform-remote-management-catinsideboxunderthetable-testenv"'
resource "aws_s3_bucket" "terraform_state_bucket" {
  bucket        = var.state_bucket_name
  force_destroy = true
}

resource "aws_s3_bucket_versioning" "terraform_state_bucket_versioning" {
  bucket = aws_s3_bucket.terraform_state_bucket.id

  versioning_configuration {
    status = "Enabled"
  }
}

resource "aws_s3_bucket_public_access_block" "terraform_state_bucket_public_access_block" {
  bucket = aws_s3_bucket.terraform_state_bucket.id

  block_public_acls       = true
  ignore_public_acls      = true
  block_public_policy     = true
  restrict_public_buckets = true
}

resource "aws_s3_bucket_server_side_encryption_configuration" "terraform_state_bucket_encryption" {
  bucket = aws_s3_bucket.terraform_state_bucket.id

  rule {
    apply_server_side_encryption_by_default {
      kms_master_key_id = aws_kms_key.terraform_state_bucket_encryption_key.arn
      sse_algorithm     = "aws:kms"
    }
  }
}

resource "aws_kms_key" "terraform_state_bucket_encryption_key" {
  description             = "key used to encrypt bucket for state management"
  deletion_window_in_days = 7
  enable_key_rotation     = true
}

resource "aws_kms_alias" "terraform_state_bucket_encryption_key_alias" {
  name          = format("alias/%s-bucket-kms-key", var.state_bucket_name)
  target_key_id = aws_kms_key.terraform_state_bucket_encryption_key.key_id
}