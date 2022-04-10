package kata_test

import (
	"lc/besttravel-codewars/kata"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestKata(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Kata Suite")
}

func doTest(t, k int, ls []int, exp int) {
	Expect(kata.ChooseBestSum(t, k, ls)).To(Equal(exp))
}

var _ = Describe("Example test cases from CodeWars", func() {
	var ls []int
	When("Case 1", func() {
		It("Returns the greatest distance possible", func() {
			doTest(163, 3, []int{50, 55, 56, 57, 58}, 163)
		})
	})
	When("Case 2", func() {
		It("Returns -1 - not enough towns", func() {
			doTest(163, 3, []int{50}, -1)
		})
	})
	When("Case 3", func() {
		ls = []int{91, 74, 73, 85, 73, 81, 87}
		It("Returns the greatest distance possible", func() {
			doTest(230, 3, ls, 228)
			doTest(331, 2, ls, 178)
			doTest(331, 4, ls, 331)
			doTest(331, 5, ls, -1)
		})
	})
})
