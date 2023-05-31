package gini_test

import (
	"fmt"
	"testing"

	"github.com/haijima/gini"
)

func TestGiniInt(t *testing.T) {
	tests := []struct {
		name    string
		values  []int
		want    float64
		wantErr bool
	}{
		{name: "empty", values: []int{}, want: 0, wantErr: false},
		{name: "single", values: []int{1}, want: 0, wantErr: false},
		{name: "equality", values: []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10}, want: 0, wantErr: false},
		{name: "monopoly", values: []int{100, 0, 0, 0, 0, 0, 0, 0, 0, 0}, want: 0.9, wantErr: false},
		{name: "unsorted", values: []int{0, 0, 0, 100, 0, 0, 0, 0, 0, 0}, want: 0.9, wantErr: false},
		{name: "negative", values: []int{0, 1, 5, 3, 2, -1, 4, 6, 3, 2}, want: 0, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := gini.Gini(tt.values)
			if (err != nil) != tt.wantErr {
				t.Errorf("Gini() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Gini() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGiniFloat64(t *testing.T) {
	got, err := gini.Gini([]float64{1.1, 1.1, 1.1})
	if err != nil {
		t.Errorf("Gini() error = %v", err)
		return
	}
	if got != 0 {
		t.Errorf("Gini() got = %v, want 0", got)
	}
}

func TestGiniCustomType(t *testing.T) {
	type MyInt int

	got, err := gini.Gini([]MyInt{1, 2, 3, 4})
	if err != nil {
		t.Errorf("Gini() error = %v", err)
		return
	}
	if got != 0.25 {
		t.Errorf("Gini() got = %v, want 0.25", got)
	}
}

func ExampleGini() {
	index, err := gini.Gini([]int{1, 2, 3, 4})
	if err != nil {
		panic(err)
	}
	fmt.Println(index)
	// Output: 0.25
}
