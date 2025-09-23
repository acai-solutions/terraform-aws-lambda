# ACAI AWS Lambda Module
# Copyright (C) 2024, 2025 ACAI GmbH
# Licensed under AGPL v3
#
# Visit https://www.acai.gmbh or https://docs.acai.gmbh for more information.
# 
# For full license text, see LICENSE file in repository root.


output "use_case_3_lambda1" {
  value = module.use_case_3_lambda1
}

output "use_case_3_lambda1_result" {
  value = jsondecode(aws_lambda_invocation.use_case_3_lambda1.result)
}

output "use_case_3_lambda2" {
  value = module.use_case_3_lambda2
}

output "use_case_3_lambda2_result" {
  value = jsondecode(aws_lambda_invocation.use_case_3_lambda2.result)
}
