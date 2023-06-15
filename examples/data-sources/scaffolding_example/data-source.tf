terraform {
  required_providers {
    check = {
      source = "chrismilson/check"
    }
  }
}

locals {
  value = false
}

data "check" "something" {
  assert {
    condition     = local.value
    error_message = "This is an error!"
  }

  assert {
    condition       = "two" == "one"
    warning_message = "This is a warning!"
  }
}
