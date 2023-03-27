package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/docker"
	"github.com/stretchr/testify/assert"
)

func TestDockerHelloWorldExample(t *testing.T) {
	tag := "dorukakinci/terraform:1.4.2"
	buildOptions := &docker.BuildOptions{
		// name of the Docker image
		Tags: []string{tag},
	}

	// Build the Docker image from the Dockerfile terraformcli.Dockerfile
	docker.Build(t, "../docker", buildOptions)

	// website::tag::3:: Run the Docker image, read the text file from it, and make sure it contains the expected output.
	opts := &docker.RunOptions{Command: []string{"terraform", "--version"}}
	output := docker.Run(t, tag, opts)
	assert.Contains(t, output, "Terraform v1.4.2")
}
