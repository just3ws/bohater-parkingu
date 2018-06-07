package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("PunktBohater", func() {
	It("should greet the world", func() {
		message := "Hello, World!"
		Expect(Say()).To(Equal(message))
	})
})
