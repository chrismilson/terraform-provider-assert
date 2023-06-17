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
  for_each = toset(["a", "b"])

  condition     = local.value
  error_message = "This is an error!"
}

data "assert" "warn" {
  for_each = toset(["a", "b"])

  condition       = local.value
  warning_message = "This is a warning!"
}
