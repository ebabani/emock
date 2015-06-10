package xmock
import (
	"github.com/onsi/gomega/types"
	"github.com/onsi/gomega/matchers"
)

func MatchArgs(expected ...interface{}) types.GomegaMatcher {
	return &matchers.EqualMatcher{Expected: expected}
}