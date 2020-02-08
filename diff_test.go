package main

import (
	"math"
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

type Sex int

const (
	_ Sex = iota
	Male
	Female
	NoAnswer
)

type User struct {
	Name  string
	Age   int
	IsVIP bool
	Sex   Sex
}

func TestDiff(t *testing.T) {
	tests := []struct {
		name string
		got  *User
		want *User
	}{
		{"Equal", &User{"tom", 25, false, Male}, &User{"tom", 25, false, Male}},
		{"NotEqual", &User{"tom", 25, false, Male}, &User{"tom", 25, true, Male}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if diff := cmp.Diff(tt.want, tt.got); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestEqual(t *testing.T) {
	tests := []struct {
		name string
		got  *User
		want *User
	}{
		{"Equal", &User{"tom", 25, false, Male}, &User{"tom", 25, false, Male}},
		{"NotEqual", &User{"tom", 25, false, Male}, &User{"tom", 25, true, Male}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !cmp.Equal(tt.want, tt.got) {
				t.Errorf("got %v, but want %v", tt.got, tt.want)
			}
		})
	}
}

func TestEqualWithOptionComparer(t *testing.T) {
	opt := cmp.Comparer(func(x, y float64) bool {
		return math.Abs(x-y) < 0.01
	})

	tests := []struct {
		name string
		got  float64
		want float64
	}{
		{"Equal", 1, 1},
		{"NearlyEqual", 1, 1.001},
		{"NotEqual", 1, 1.01},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !cmp.Equal(tt.want, tt.got, opt) {
				t.Errorf("got %v, but want %v", tt.got, tt.want)
			}
		})
	}
}

func TestEqualWithOptionFilterValues(t *testing.T) {
	alwaysEqual := cmp.Comparer(func(_, _ interface{}) bool { return true })
	opt := cmp.Options{
		cmp.FilterValues(func(x, y *User) bool {
			return x.Name == y.Name
		}, alwaysEqual),
		cmp.FilterValues(func(x, y *User) bool {
			return x.Age == y.Age && x.Sex == y.Sex
		}, alwaysEqual),
	}

	tests := []struct {
		name string
		got  *User
		want *User
	}{
		{"EqualName", &User{"tom", 25, false, Male}, &User{"tom", 32, false, NoAnswer}},
		{"EqualAgeAndSex", &User{"bob", 25, true, Male}, &User{"tom", 25, false, Male}},
		{"NotEqual", &User{"bob", 32, false, Male}, &User{"tom", 32, false, NoAnswer}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !cmp.Equal(tt.want, tt.got, opt) {
				t.Errorf("got %v, but want %v", tt.got, tt.want)
			}
		})
	}
}

func TestEqualWithOptionTransformer(t *testing.T) {
	trans := cmp.Transformer("Sort", func(in []int) []int {
		out := append([]int(nil), in...)
		sort.Ints(out)
		return out
	})

	tests := []struct {
		name string
		got  []int
		want []int
	}{
		{"Equal", []int{1, 2, 3}, []int{3, 2, 1}},
		{"NotEqual", []int{1, 2, 3}, []int{1, 1, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !cmp.Equal(tt.want, tt.got, trans) {
				t.Errorf("got %v, but want %v", tt.got, tt.want)
			}
		})
	}
}

type Score struct {
	Pass  bool
	score int
}

func TestEqualWithOptionIgnoreUnexported(t *testing.T) {
	tests := []struct {
		name string
		got  *Score
		want *Score
	}{
		{"Equal", &Score{true, 90}, &Score{true, 85}},
		{"NotEqual", &Score{false, 90}, &Score{true, 85}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !cmp.Equal(tt.want, tt.got, cmpopts.IgnoreUnexported(*tt.want)) {
				t.Errorf("got %v, but want %v", tt.got, tt.want)
			}
		})
	}
}
