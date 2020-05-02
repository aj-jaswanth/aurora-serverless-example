resource "aws_lambda_function" "default" {
  function_name    = "hello-lambda"
  runtime          = "go1.x"
  handler          = "hello-lambda"
  role             = aws_iam_role.lambda.arn
  filename         = "../hello-lambda/function.zip"
  source_code_hash = filebase64sha256("../hello-lambda/function.zip")
  environment {
    variables = {
      DB_USER = aws_rds_cluster.default.master_username
      DB_PWD  = aws_rds_cluster.default.master_password
      DB_HOST = aws_rds_cluster.default.endpoint
      DB_NAME = aws_rds_cluster.default.database_name
    }
  }
  timeout = 60
  vpc_config {
    subnet_ids         = [aws_subnet.private1.id, aws_subnet.private2.id]
    security_group_ids = [aws_security_group.default.id]
  }

  tags = local.tags
}
