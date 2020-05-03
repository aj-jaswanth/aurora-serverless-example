resource "aws_rds_cluster" "default" {
  cluster_identifier     = "aurora-serverless-poc"
  copy_tags_to_snapshot  = true
  master_username        = "admin"
  master_password        = "ac0n3x72"
  skip_final_snapshot    = true
  apply_immediately      = true
  engine                 = "aurora"
  database_name          = "poc"
  engine_mode            = "serverless"
  engine_version         = "5.6.10a"
  db_subnet_group_name   = aws_db_subnet_group.default.name
  vpc_security_group_ids = [aws_security_group.default.id]
  scaling_configuration {
    auto_pause               = true
    max_capacity             = 2
    min_capacity             = 1
    seconds_until_auto_pause = 300
  }
  tags = local.tags
}

resource "aws_db_subnet_group" "default" {
  name       = "aurora-serverless-poc"
  subnet_ids = [aws_subnet.private1.id, aws_subnet.private2.id]
}