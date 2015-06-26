package emock_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestEmock(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Emock Suite")
}
