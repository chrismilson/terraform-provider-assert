terraform {
  required_providers {
    assert = {
      source = "chrismilson/assert"
    }
  }
}

locals {
  value = false
}

data "assert" "error" {
  severity  = "error"
  condition = local.value
  summary   = "Invalid Condition"
  detail    = "This is an error!"
}
