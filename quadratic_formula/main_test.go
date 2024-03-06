package main

import (
	"testing"
)

func Test_quadraticFormula(t *testing.T) {
	type args struct {
		a float64
		b float64
		c float64
	}
	tests := []struct {
		name    string
		args    args
		want1   float64
		want2   float64
		wantErr bool
	}{
		//  The test case should check to make sure the roots are correct when a = 1, b = -3 and c = -4
		{
			name: "a = 1, b = -3 c = -4",
			args: args{
				a: 1.0,
				b: -3.0,
				c: -4.0,
			},
			want1: 4.0,
			want2: -1.0,
		},
		//  The test case should check to make sure the roots are correct when a = 2, b = 1 and c = 1 this test case should fail
		{
			name: "a = 2, b = 1 c = 1",
			args: args{
				a: 2.0,
				b: 1.0,
				c: 1.0,
			},
			wantErr: true,
		},
		//  The test case should check to make sure the roots are correct when a = 0, b = 1 and c = 3
		{
			name: "a = 0, b = 1 c = 3",
			args: args{
				a: 0.0,
				b: 1.0,
				c: 3.0,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := quadraticFormula(tt.args.a, tt.args.b, tt.args.c)

			if (err != nil) != tt.wantErr {
				t.Errorf("quadraticFormula() got = %v, want %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {

				if got != tt.want1 {
					t.Errorf("quadraticFormula() got = %v, want %v", got, tt.want1)
				}
				if got1 != tt.want2 {
					t.Errorf("quadraticFormula() got1 = %v, want %v", got1, tt.want2)
				}
			}
		})
	}
}
