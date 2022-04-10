package kata

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestKata(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Kata Suite")
}

func doChooseBestSumTest(t, k int, ls []int, exp int) {
	Expect(ChooseBestSum(t, k, ls)).To(Equal(exp))
}

var _ = Describe("Example test cases from CodeWars", func() {
	var ls []int
	When("Case 1", func() {
		It("Returns the greatest distance possible", func() {
			doChooseBestSumTest(163, 3, []int{50, 55, 56, 57, 58}, 163)
		})
	})
	When("Case 2", func() {
		It("Returns -1 - not enough towns", func() {
			doChooseBestSumTest(163, 3, []int{50}, -1)
		})
	})
	When("Case 3", func() {
		ls = []int{91, 74, 73, 85, 73, 81, 87}
		It("Returns the greatest distance possible", func() {
			doChooseBestSumTest(230, 3, ls, 228)
			doChooseBestSumTest(331, 2, ls, 178)
			doChooseBestSumTest(331, 4, ls, 331)
			doChooseBestSumTest(331, 5, ls, -1)
		})
	})
})

var _ = Describe("Array mask", func() {
	var list []int
	var mask []int8
	var expe []int
	When("Case 1", func() {
		list = []int{1, 2, 3, 4, 5}
		mask = []int8{0, 1, 0, 0, 1}
		expe = []int{0, 2, 0, 0, 5}
		listAddr := &list
		It("Masks proper values", func() {
			Expect(applyArrayMask(list, mask)).To(Equal(expe))
		})
		It("Does not modify the given list", func() {
			result := applyArrayMask(list, mask)
			Expect(list).To(Equal([]int{1, 2, 3, 4, 5}))
			Expect(listAddr).NotTo(Equal(result))
		})
	})
})
