# INFRASTRUCTURE SETUP

## WHAT THIS INFRASTRUCTURE DO

This infrastructure spawns:
- state management resources
- required group and user with the least privileges
- s3 bucket with blocked public access and KMS encryption
- automatic stale objects deletion
- monitoring utilities(like spending alarms and so on)

## STATE MANAGEMENT

For state management, s3 bucket and dynamoDB table is used. For more details, check 'remoteState'
module.

## REQUIRED INPUTS

|         VARIABLE NEAME          | PURPOSE                                                                            |
|:-------------------------------:|:-----------------------------------------------------------------------------------|
| `email_notification_subscriber` | provides email address that will be used for budget & notification purposes by AWS |

## ENVIRONMENTS

### TEST ENVIRONMENT

Test environment is created for development purposes.