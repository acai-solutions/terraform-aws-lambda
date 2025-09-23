# ACAI AWS Lambda Module
# Copyright (C) 2024, 2025 ACAI GmbH
# Licensed under AGPL v3
#
# Visit https://www.acai.gmbh or https://docs.acai.gmbh for more information.
# 
# For full license text, see LICENSE file in repository root.


output "lambda_name" {
  value = module.use_case_5_lambda.lambda.name
}

output "lambda_arn" {
  value = module.use_case_5_lambda.lambda.arn
}

output "use_case_5_lambda" {
  value = module.use_case_5_lambda
}

output "use_case_5_lambda_result" {
  value = jsondecode(aws_lambda_invocation.use_case_5_lambda.result)
}

