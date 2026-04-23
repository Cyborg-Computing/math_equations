package main

import (
	"testing"
)

func Test_division(t *testing.T) {
	type args struct {
		a float64
		b float64
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		//  The test case should check to make sure the division is correct when a = 3 and b = 4
		{
			name: "a = 3, b = 4",
			args: args{
				a: 3.0,
				b: 4.0,
			},
			want: 0.75,
		},
		//  The test case should check to make sure the division is correct when a = 9 and b = 0
		{
			name: "a = 9, b = 0",
			args: args{
				a: 9.0,
				b: 0.0,
			},
			wantErr: true, // this test case should fail because b cannot be 0
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := division(tt.args.a, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("division() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("division() = %v, want %v", got, tt.want)
			}
		})
	}
}
