terraform {
  required_providers {
    assert = {
      source = "chrismilson/assert"
    }
  }
}

locals {
  environments = {
    dev = {
      size = "small"
    }
    prod = {
      size = "large"
    }
    bad = {
      size = "sundried tomato"
    }
  }

  valid_environment_sizes = ["small", "large"]
}

data "assert" "error" {
  for_each = local.environments

  severity  = "error"
  condition = contains(local.valid_environment_sizes, each.value.size)
  summary   = "Unrecognised Environment Size"
  detail    = "The environment size must be one of [${join(", ", local.valid_environment_sizes)}]. got '${each.value.size}'."
}
