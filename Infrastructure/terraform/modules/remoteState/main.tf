provider "aws" {
  default_tags {
    tags = {
      Module  = "RemoteState"
      Purpose = "StateManagement"
    }
  }
}