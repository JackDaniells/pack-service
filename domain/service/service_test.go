package service

import (
	"github.com/JackDaniells/pack-service/domain/contracts"
	"github.com/JackDaniells/pack-service/domain/contracts/mocks"
	"github.com/JackDaniells/pack-service/domain/entity"
	"reflect"
	"testing"
)

func Test_packService_Calculate(t *testing.T) {
	type args struct {
		orderItems int
	}

	type fields struct {
		packRepository contracts.PackRepository
	}

	tests := []struct {
		name   string
		args   args
		fields fields
		want   []entity.Pack
	}{
		{
			name: "Should return 1 250-pack for 1 order item",
			args: args{
				orderItems: 1,
			},
			fields: fields{
				packRepository: func() contracts.PackRepository {
					packService := &mocks.PackRepository{}
					packService.On("GetAll").Return([]int{250, 500, 1000, 2000, 5000})
					return packService
				}(),
			},
			want: []entity.Pack{
				{
					Size:     250,
					Quantity: 1,
				},
			},
		},
		{
			name: "Should return 1 250-pack for 250 order items",
			args: args{
				orderItems: 250,
			},
			fields: fields{
				packRepository: func() contracts.PackRepository {
					packService := &mocks.PackRepository{}
					packService.On("GetAll").Return([]int{250, 500, 1000, 2000, 5000})
					return packService
				}(),
			},
			want: []entity.Pack{
				{
					Size:     250,
					Quantity: 1,
				},
			},
		},
		{
			name: "Should return 1 500-pack and 1 250-pack for 501 order items",
			args: args{
				orderItems: 501,
			},
			fields: fields{
				packRepository: func() contracts.PackRepository {
					packService := &mocks.PackRepository{}
					packService.On("GetAll").Return([]int{250, 500, 1000, 2000, 5000})
					return packService
				}(),
			},
			want: []entity.Pack{
				{
					Size:     500,
					Quantity: 1,
				},
				{
					Size:     250,
					Quantity: 1,
				},
			},
		},
		{
			name: "Should return 2 5000-pack, 1 2000-pack and 1 250-pack for 12001 order item",
			args: args{
				orderItems: 12001,
			},
			fields: fields{
				packRepository: func() contracts.PackRepository {
					packService := &mocks.PackRepository{}
					packService.On("GetAll").Return([]int{250, 500, 1000, 2000, 5000})
					return packService
				}(),
			},
			want: []entity.Pack{
				{
					Size:     5000,
					Quantity: 2,
				},
				{
					Size:     2000,
					Quantity: 1,
				},
				{
					Size:     250,
					Quantity: 1,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &packService{
				repository: tt.fields.packRepository,
			}
			got := p.Calculate(tt.args.orderItems)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Calculate() got = %v, want %v", got, tt.want)
			}
		})
	}
}
