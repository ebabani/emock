package xmock_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestXmock(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Xmock Suite")
}
