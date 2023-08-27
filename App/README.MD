# ITRANSFER

## WHAT THIS APPLICATION DO

This application takes user path to file and uploads pointed file to AWS s3 bucket.
After that, it generates presigned URL with TTL.
Anyone that have this URL can download this file as long as URL TTL did not pass.

## REQUIRED ENVIRONMENT VARIABLES

|       VARIABLE NAME       | DESCRIPTION                                                           |
|:-------------------------:|-----------------------------------------------------------------------|
|  AWS_STORAGE_BUCKET_NAME  | Name of your aws bucket.                                              |
| AWS_STORAGE_BUCKET_REGION | Region in which you provisioned infrastructure.                       |
|     AWS_CONSOLE_LOGIN     | Your IAM user(default prepared user: `bucket-user-default`) AccessKey | 

## COMMAND SWITCHES

| SWITCH | DESCRIPTION                                                                                                                                                                                  |
|:------:|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
|  `-n`  | Change file name (in bucket, not locally on yours machine). If file with this name is currently inside bucket, it will be overwritten.                                                       |
|  `-l`  | TTL - Define link lifespan. It is based on AWS pre-signed URL, more informations available [here.](https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/s3-example-presigned-urls.html) |

## CONSOLE PROMPTS

|      PROMPT       | MEANING                                     |
|:-----------------:|:--------------------------------------------|
| `Password prompt` | You should provide your AWS SecretAccessKey |