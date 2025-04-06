package alg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTwoSum(t *testing.T) {
	result := onePass([]int{2, 7, 11, 15}, 9)

	assert.Equal(t, result, []int{0, 1})
}

var (
	nums   = makeRange(0, 1_000_000)
	target = 1_999_998
)

func makeRange(min, max int) []int {
	r := make([]int, max-min+1)
	for i := range r {
		r[i] = min + i
	}
	return r
}

func BenchmarkBruteForce(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bruteForce(nums, target)
	}
}

func BenchmarkOnePass(b *testing.B) {
	for i := 0; i < b.N; i++ {
		onePass(nums, target)
	}
}

func BenchmarkTwoPass(b *testing.B) {
	for i := 0; i < b.N; i++ {
		twoPass(nums, target)
	}
}

func onePass(nums []int, target int) []int {
	seen := make(map[int]int) // num -> index

	for i, num := range nums {
		complement := target - num
		if j, found := seen[complement]; found {
			return []int{j, i}
		}
		seen[num] = i
	}
	return nil
}

func twoPass(nums []int, target int) []int {
	seen := make(map[int]int)

	for i, num := range nums {
		seen[num] = i
	}

	for i, num := range nums {
		if j, found := seen[target-num]; found && i != j {
			return []int{i, j}
		}
	}

	return nil
}

func bruteForce(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
