resource "aws_dynamodb_table" "terraform_dynamodb_state_locks" {
  name         = var.state_dynamoDb_table_name
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "LockID"

  attribute {
    name = "LockID"
    type = "S"
  }

  server_side_encryption {
    enabled     = true
    kms_key_arn = aws_kms_key.terraform_state_dynamodb_kms_key.arn
  }
}

resource "aws_kms_key" "terraform_state_dynamodb_kms_key" {
  description             = "key used to encrypt dynamodb for state management"
  deletion_window_in_days = 7
  enable_key_rotation     = true
}

resource "aws_kms_alias" "terraform_state_dynamodb_kms_key_alias" {
  name          = format("alias/%s-kms-key", var.state_dynamoDb_table_name)
  target_key_id = aws_kms_key.terraform_state_dynamodb_kms_key.key_id
}
