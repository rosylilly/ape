package ape_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestApe(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Ape Suite")
}
