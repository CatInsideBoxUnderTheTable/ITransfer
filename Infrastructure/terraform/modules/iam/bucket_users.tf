locals {
  formatted_names = [for v in var.bucket_console_users : format("bucket-user-%s", v)]
}

resource "aws_iam_user" "bucket_users" {
  for_each = toset(local.formatted_names)

  name = each.value
  path = "/users/uploaders/"
}

resource "aws_iam_user_group_membership" "upload_bucket_group_membership" {
  depends_on = [aws_iam_user.bucket_users]
  for_each   = toset(local.formatted_names)

  groups = [aws_iam_group.bucket_uploaders.name]
  user   = each.value
}