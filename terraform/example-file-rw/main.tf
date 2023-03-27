terraform {
  required_providers {
    null = {
      source  = "hashicorp/null"
      version = "3.1.0"
    }

    ## add provider for http call
    http = {
      source  = "hashicorp/http"
      version = "3.2.1"
    }
  }
}

## create a null resource that creates test.txt file with hello world
resource "null_resource" "test" {
  provisioner "local-exec" {
    command = "echo 'hello world' > test.txt"
  }
}
