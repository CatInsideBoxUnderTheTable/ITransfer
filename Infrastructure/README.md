# INFRASTRUCTURE SETUP

## STATE MANAGEMENT

For state management, s3 bucket and dynamoDB table is used. For more details, check 'remoteState' module.

## MODULES

|   MODULE NAME    | DESCRIPTION                                                                                                    |
|:----------------:|----------------------------------------------------------------------------------------------------------------|
|  'remoteState'   | Module used for remote state configuration. Take note that backend.tf is set up in root catalog                |
| 'transferBucket' | Module used for holding transferred files via s3 bucket. ACL rules are utilized to remove stale bucket objects |       

## ENVIRONMENTS

### TEST ENVIRONMENT

Test environment is created for development purposes.