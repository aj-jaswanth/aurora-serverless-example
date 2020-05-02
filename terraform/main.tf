locals {
  tags = {
    Custodian = "jaswanth"
  }
  region = "ap-southeast-1"
}

provider "aws" {
  region = local.region
}
