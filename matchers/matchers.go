package emock
import (
	"github.com/onsi/gomega/types"
	"github.com/onsi/gomega/matchers"
)

// Helper matcher for emock and ginkgo
// Use for verifying a function was called with the expected args
//
// Expect(mockObj.GetArgsForCall(0)).To(MatchArgs(123, "abc"))
//
func MatchArgs(expected ...interface{}) types.GomegaMatcher {
	return &matchers.EqualMatcher{Expected: expected}
}