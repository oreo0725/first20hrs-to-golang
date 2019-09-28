package math

import "testing"

func TestNstar(t *testing.T) {
	// type args struct {
	// 	n int
	// }
	tests := []struct {
		name string
		num  int
		want int
	}{
		// TODO: Add test cases.
		{"1 should get 1", 1, 1},
		{"0 should get 0", 0, 0},
		{"2 should get 2", 2, 2},
		{"3 should get 6", 3, 6},
		{"5 should get 120", 5, 120},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NStar(tt.num); got != tt.want {
				t.Errorf("nstar() = %v, want %v", got, tt.want)
			}
		})
	}
}
