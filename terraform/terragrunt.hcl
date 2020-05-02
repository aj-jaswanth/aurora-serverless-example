remote_state {
  backend = "s3"
  config = {
    encrypt = true
    bucket = "aurora-poc-tf-state"
    key = "terraform.tfstate"
    region = "ap-southeast-1"
    dynamodb_table = "aurora-poc-tf-lock"
  }
  generate = {
    path = "backend.tf"
    if_exists = "overwrite_terragrunt"
  }
}