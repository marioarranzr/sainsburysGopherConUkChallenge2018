package main

import "testing"

func Test_calculateHouses(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Long string",
			args: args{s: "<><<>>^v^^<"},
			want: 6,
		},
		{
			name: "Empty string",
			args: args{s: ""},
			want: 1,
		},
		{
			name: "One element",
			args: args{s: "<"},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateHouses(tt.args.s); got != tt.want {
				t.Errorf("count() = %v, want %v", got, tt.want)
			}
		})
	}
}
