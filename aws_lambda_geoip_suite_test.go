package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestAwsLambdaGeoip(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "AwsLambdaGeoip Suite")
}
