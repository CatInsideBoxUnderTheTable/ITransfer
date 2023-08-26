resource "aws_iam_group" "bucket_uploaders" {
  name = "bucket-upload-group"
  path = "/groups/"
}

resource "aws_iam_policy_attachment" "bucket_uploaders_policy_attachment" {
  name       = "bucket-upload-policy-attachment"
  policy_arn = aws_iam_policy.bucket_user_actions_policy.arn
  groups     = [aws_iam_group.bucket_uploaders.name]
}

resource "aws_iam_policy" "bucket_user_actions_policy" {
  name        = "bucket-upload-policy"
  policy      = data.aws_iam_policy_document.bucket_user_policy.json
  description = "allows for access to ITransfer upload bucket"
}

data "aws_iam_policy_document" "bucket_user_policy" {
  statement {
    effect    = "Allow"
    actions   = ["s3:GetObject", "s3:PutObject"]
    resources = [var.transfer_bucket_arn]
  }

  statement {
    effect    = "Allow"
    actions   = ["kms:Decrypt", "kms:Encrypt"]
    resources = [var.transfer_bucket_encryption_key_arn]
  }
}