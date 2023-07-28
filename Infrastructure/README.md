# INFRASTRUCTURE SETUP
## STATE MANAGEMENT
For state management, s3 bucket and dynamoDB table is used. For more details, check 'remoteState' module.

## MODULES
| MODULE NAME | DESCRIPTION |
|:-----------:|-------------|
| 'stateManagement' | Module used for remote state configuration. Take note that backend.tf is set up in root catalog |

## ENVIRONMENTS
### TEST ENVIRONMENT
Test environment is created for development purposes.