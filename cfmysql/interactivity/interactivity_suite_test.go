package interactivity_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestInteractivity(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Interactivity Suite")
}
