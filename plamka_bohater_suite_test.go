package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestPlamkaBohater(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "PlamkaBohater Suite")
}
