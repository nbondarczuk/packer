package packer

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMerge(t *testing.T) {
	var tests = []struct {
		input    []int
		buckets  []int
		expected []int
	}{
		{
			[]int{250, 250},
			[]int{5000, 2000, 1000, 500, 250},
			[]int{500},
		},
		{
			[]int{250},
			[]int{5000, 2000, 1000, 500, 250},
			[]int{250},
		},
		{
			[]int{1000, 1000, 250, 250},
			[]int{5000, 2000, 1000, 500, 250},
			[]int{2000, 500},
		},

	}

	for i := range tests {
		t.Run(fmt.Sprintf("%d-%v->%v", tests[i].input, tests[i].buckets, tests[i].expected), func(t *testing.T) {
			result := merge(tests[i].input, tests[i].buckets)
			assert.Equal(t, tests[i].expected, result)
		})
	}
}

func TestPack(t *testing.T) {
	var tests = []struct {
		input    int
		buckets  []int
		expected []int
	}{
		{
			100,
			[]int{5000, 2000, 1000, 500, 250},
			[]int{250},
		},
		{
			300,
			[]int{5000, 2000, 1000, 500, 250},
			[]int{500},
		},
		{
			500,
			[]int{5000, 2000, 1000, 500, 250},
			[]int{500},
		},
		{
			1,
			[]int{5000, 2000, 1000, 500, 250},
			[]int{250},
		},
		{
			250,
			[]int{5000, 2000, 1000, 500, 250},
			[]int{250},
		},
		{
			251,
			[]int{5000, 2000, 1000, 500, 250},
			[]int{500},
		},
		{
			501,
			[]int{5000, 2000, 1000, 500, 250},
			[]int{500, 250},
		},
		{
			12001,
			[]int{5000, 2000, 1000, 500, 250},
			[]int{5000, 5000, 2000, 250},
		},
	}
	for i := range tests {
		t.Run(fmt.Sprintf("%d-%v->%v", tests[i].input, tests[i].buckets, tests[i].expected), func(t *testing.T) {
			result := Pack(tests[i].input, tests[i].buckets)
			assert.Equal(t, tests[i].expected, result)
		})
	}
}
