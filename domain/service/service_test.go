package service

import (
	"reflect"
	"testing"
)

func Test_packService_Calculate(t *testing.T) {
	type args struct {
		orderItems int
	}
	tests := []struct {
		name string
		args args
		want map[int]int
	}{
		{
			name: "Should return 1 250-pack for 1 order item",
			args: args{
				orderItems: 1,
			},
			want: map[int]int{
				250: 1,
			},
		},
		{
			name: "Should return 1 250-pack for 250 order items",
			args: args{
				orderItems: 250,
			},
			want: map[int]int{
				250: 1,
			},
		},
		{
			name: "Should return 1 500-pack and 1 250-pack for 501 order items",
			args: args{
				orderItems: 501,
			},
			want: map[int]int{
				250: 1,
				500: 1,
			},
		},
		{
			name: "Should return 2 5000-pack, 1 2000-pack and 1 250-pack for 12001 order item",
			args: args{
				orderItems: 12001,
			},
			want: map[int]int{
				250:  1,
				2000: 1,
				5000: 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &packService{}
			got := p.Calculate(tt.args.orderItems)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Calculate() got = %v, want %v", got, tt.want)
			}
		})
	}
}
