package test
import (
  "testing"
  "github.com/stretchr/testify/assert"
  "github.com/gruntwork-io/terratest/modules/docker"

)
func TestDocker(t *testing.T) {
	tag := "myapp"
  	buildOptions := &docker.BuildOptions{
    	Tags: []string{tag},
  	}

	docker.Build(t, ".", buildOptions)
	opts := &docker.RunOptions{Command: []string{"cat", "/helloworld.txt"}}
	output := docker.Run(t, tag, opts)
  	assert.Equal(t, "Hello world", output)

}