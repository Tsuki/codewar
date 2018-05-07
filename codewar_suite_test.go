package kata_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCodewar(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Codewar Suite")
}
