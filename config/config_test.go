package config

import (
    "fmt"
    "testing" // import the testing package to use the testing functions
)

func TestConfig(t *testing.T) { // define a test function named TestConfig
    // TODO: write tests for the Config struct
	// Init()
	fmt.Println(EnvConfig.Jwt.Expires)
}