package main

import "testing"

func Test_pythagoreanTheorem(t *testing.T) {
	tests := []struct {
		name string
		a    float64
		b    float64
		want float64
	}{
		//  The test case should check to make sure the hypotenuse is correct when a = 3 and b = 4
		{
			name: "a = 3, b = 4",
			a:    3.0,
			b:    4.0,
			want: 5.0,
		},
		//  The test case should check to make sure the hypotenuse is correct when a = 5 and b = 12
		{
			name: "a = 5, b = 12",
			a:    5.0,
			b:    12.0,
			want: 13.0,
		},
		//  The test case should check to make sure the hypotenuse is correct when a = 8 and b = 6
		{
			name: "a = 8, b = 6",
			a:    8.0,
			b:    6.0,
			want: 10.0,
		},
		//  The test case should check to make sure the hypotenuse is correct when a = 9 and b = 12
		{
			name: "a = 9, b = 12",
			a:    9.0,
			b:    12.0,
			want: 15.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pythagoreanTheorem(tt.a, tt.b); got != tt.want {
				t.Errorf("pythagoreanTheorem() = %v, want %v", got, tt.want)
			}
		})
	}
}
