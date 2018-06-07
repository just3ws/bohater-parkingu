package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestPunktBohater(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "PunktBohater Suite")
}
