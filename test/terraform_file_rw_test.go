package test

import (
	"testing"

	//"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	//test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
)

// TEST READ WRITE FILE
func TestTerraformFileRW(t *testing.T) {
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// Set the path to the Terraform code that will be tested.
		TerraformDir: "../terraform/example-file-rw",
	})

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	//read test.txt file and compare with assert
	file, err := ioutil.ReadFile("../terraform/example-file-rw/test.txt")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "hello world\n", string(file))
}

// TEST READ WRITE FILE
func TestTerraformFileRWwithInput(t *testing.T) {
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// Set the path to the Terraform code that will be tested.
		TerraformDir: "../terraform/example-file-rw",
	})

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	//read test.txt file and compare with assert
	file, err := ioutil.ReadFile("../terraform/example-file-rw/test.txt")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "hello world\n", string(file))
}
