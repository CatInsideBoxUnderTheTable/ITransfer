locals {
  bucket_name                        = "${var.environment_name}-transferred-files-storage"
  multipart_upload_max_lifetime_days = 2
}

resource "aws_s3_bucket" "transfer_bucket" {
  bucket        = local.bucket_name
  force_destroy = true
}

resource "aws_s3_bucket_lifecycle_configuration" "transfer_bucket_lifecycle" {
  bucket = aws_s3_bucket.transfer_bucket.id

  rule {
    id     = "CleanStaleObjects"
    status = "Enabled"

    abort_incomplete_multipart_upload {
      days_after_initiation = local.multipart_upload_max_lifetime_days
    }
    expiration {
      days = var.bucket_object_lifetime
    }
  }
}

resource "aws_s3_bucket_public_access_block" "transfer_bucket_public_access_block" {
  bucket = aws_s3_bucket.transfer_bucket.id

  block_public_acls       = true
  ignore_public_acls      = true
  block_public_policy     = true
  restrict_public_buckets = true
}

resource "aws_s3_bucket_server_side_encryption_configuration" "transfer_bucket_encryption" {
  bucket = aws_s3_bucket.transfer_bucket.id

  rule {
    apply_server_side_encryption_by_default {
      kms_master_key_id = aws_kms_key.transfer_bucket_encryption_key.arn
      sse_algorithm     = "aws:kms"
    }
  }
}

resource "aws_kms_key" "transfer_bucket_encryption_key" {
  description             = "key used to encrypt bucket for state management"
  deletion_window_in_days = 7
  enable_key_rotation     = true
}

resource "aws_kms_alias" "transfer_bucket_encryption_key_alias" {
  name          = format("alias/%s-bucket-kms-key", local.bucket_name)
  target_key_id = aws_kms_key.transfer_bucket_encryption_key.key_id
}