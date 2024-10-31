## Example Usage

```
provider "splunk" {
  url                  = "localhost:8089"
  username             = "admin"
  password             = "changeme"
  insecure_skip_verify = true
}
```

Terraform 0.13 and later must add:
```
terraform {
   required_providers {
    splunk = {
      source  = "splunk/splunk"
    }
  }
}
```