# ITRANSFER

## PURPOSE

Main purpose of this application is to add user availability to share files via link (inspired by WeTransfer) from your
console, as well as provide some level of control of data.

## REQUIRED ENVIRONMENT VARIABLES

|       VARIABLE NAME       | DESCRIPTION                                                                                                                                                      |
|:-------------------------:|------------------------------------------------------------------------------------------------------------------------------------------------------------------|
|  AWS_STORAGE_BUCKET_NAME  | Name of your aws bucket.                                                                                                                                         |
| AWS_STORAGE_BUCKET_REGION | Region in wich you provisioned infrastructure.                                                                                                                   |
|   AWS_FILE_AUTH_PROFILE   | Name of your credential and configuration profile. More information available [here.](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-files.html) |

## COMMAND SWITCHES

| SWITCH | DESCRIPTION                                                                                                                                                                            |
|:------:|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
|  `-n`  | Define own unique file name. If file with this name is currently inside bucket, it will be overwritten.                                                                                |
|  `-l`  | Define link lifespan. It is based on AWS pre-signed URL, more informations available [here.](https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/s3-example-presigned-urls.html) |
