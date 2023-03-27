package test

import (
	"testing"

	//"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	//test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
)

func TestTerraformCurl(t *testing.T) {
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// Set the path to the Terraform code that will be tested.
		TerraformDir: "../terraform/example-curl-status-code",
	})

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	//output := terraform.Output(t, terraformOptions, "hello_world")
	//	assert.Equal(t, "Hello, World!", output)

	//read test.txt file and compare with assert
}

func TestTerraformFileRW(t *testing.T) {
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// Set the path to the Terraform code that will be tested.
		TerraformDir: "../terraform/example-file-rw",
	})

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	//output := terraform.Output(t, terraformOptions, "hello_world")
	//	assert.Equal(t, "Hello, World!", output)

	//read test.txt file and compare with assert
	file, err := ioutil.ReadFile("../terraform/example-file-rw/test.txt")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "hello world\n", string(file))
}
