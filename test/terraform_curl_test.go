package test

import (
	"crypto/tls"
	"testing"
	"time"

	http_helper "github.com/gruntwork-io/terratest/modules/http-helper"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

// TEST CURL STATUS CODE WITH TF OUTPUT
func TestTerraformCurl(t *testing.T) {
	t.Parallel()
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// Set the path to the Terraform code that will be tested.
		TerraformDir: "../terraform/example-curl-status-code",
	})

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	output := terraform.Output(t, terraformOptions, "status")
	assert.Equal(t, "200", output)

	// do not trust the terraform output and test it additionally with http_helper
	url := terraform.Output(t, terraformOptions, "url")
	tlsConfig := tls.Config{}
	maxRetries := 3
	timeBetweenRetries := 5 * time.Second

	http_helper.HttpGetWithRetryWithCustomValidation(t, url, &tlsConfig, maxRetries, timeBetweenRetries, validate)

	// we could also handle the same test without CustomValidation however we would need to handle the body as well
	// http_helper.HttpGetWithRetry(t, url, nil, 200, "BODY", maxRetries, timeBetweenRetries)
}

func validate(status int, body string) bool {
	return status == 200
}
