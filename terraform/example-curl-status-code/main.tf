
terraform {
  required_providers {
    ## add provider for http call
    http = {
      source  = "hashicorp/http"
      version = "3.2.1"
    }
  }
}

## get curl from google.com
data "http" "curl" {
  url = "https://google.com"
}

## output the status code of the curl
output "status" {
  value = jsonencode(data.http.curl.status_code)
}

