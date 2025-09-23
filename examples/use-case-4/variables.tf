# ACAI AWS Lambda Module
# Copyright (C) 2024, 2025 ACAI GmbH
# Licensed under AGPL v3
#
# Visit https://www.acai.gmbh or https://docs.acai.gmbh for more information.
# 
# For full license text, see LICENSE file in repository root.


variable "function_name" {
  description = "Unique name for your Lambda Function."
  type        = string
  default     = "use_case_4"
}

variable "resource_tags" {
  type = map(string)
  default = {
    scope = "use_case_4"
  }
}
