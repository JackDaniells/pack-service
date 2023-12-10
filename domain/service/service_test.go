package service

import (
	"errors"
	"github.com/JackDaniells/pack-service/domain/contracts"
	"github.com/JackDaniells/pack-service/domain/contracts/mocks"
	"github.com/JackDaniells/pack-service/domain/entity"
	"reflect"
	"testing"
)

func Test_packService_Calculate(t *testing.T) {
	type args struct {
		items int
	}

	type fields struct {
		packRepository contracts.PackRepository
	}

	tests := []struct {
		name    string
		args    args
		fields  fields
		want    []entity.Pack
		wantErr error
	}{
		{
			name: "Should return error if items is invalid",
			args: args{
				items: 0,
			},
			wantErr: errors.New("invalid items size"),
		},
		{
			name: "Should return 1 250-pack for 1 item",
			args: args{
				items: 1,
			},
			fields: fields{
				packRepository: func() contracts.PackRepository {
					packService := &mocks.PackRepository{}
					packService.On("GetAll").Return([]int{5000, 2000, 1000, 500, 250})
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
			name: "Should return 1 250-pack for 250 items",
			args: args{
				items: 250,
			},
			fields: fields{
				packRepository: func() contracts.PackRepository {
					packService := &mocks.PackRepository{}
					packService.On("GetAll").Return([]int{5000, 2000, 1000, 500, 250})
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
			name: "Should return 1 500-pack for 251 items",
			args: args{
				items: 251,
			},
			fields: fields{
				packRepository: func() contracts.PackRepository {
					packService := &mocks.PackRepository{}
					packService.On("GetAll").Return([]int{5000, 2000, 1000, 500, 250})
					return packService
				}(),
			},
			want: []entity.Pack{
				{
					Size:     500,
					Quantity: 1,
				},
			},
		},
		{
			name: "Should return 1 500-pack for 499 items",
			args: args{
				items: 499,
			},
			fields: fields{
				packRepository: func() contracts.PackRepository {
					packService := &mocks.PackRepository{}
					packService.On("GetAll").Return([]int{5000, 2000, 1000, 500, 250})
					return packService
				}(),
			},
			want: []entity.Pack{
				{
					Size:     500,
					Quantity: 1,
				},
			},
		},
		{
			name: "Should return 1 500-pack and 1 250-pack for 501 items",
			args: args{
				items: 501,
			},
			fields: fields{
				packRepository: func() contracts.PackRepository {
					packService := &mocks.PackRepository{}
					packService.On("GetAll").Return([]int{5000, 2000, 1000, 500, 250})
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
			name: "Should return 1 1000-pack for 751 items",
			args: args{
				items: 751,
			},
			fields: fields{
				packRepository: func() contracts.PackRepository {
					packService := &mocks.PackRepository{}
					packService.On("GetAll").Return([]int{5000, 2000, 1000, 500, 250})
					return packService
				}(),
			},
			want: []entity.Pack{
				{
					Size:     1000,
					Quantity: 1,
				},
			},
		},
		{
			name: "Should return 2 5000-pack, 1 2000-pack and 1 250-pack for 12001 item",
			args: args{
				items: 12001,
			},
			fields: fields{
				packRepository: func() contracts.PackRepository {
					packService := &mocks.PackRepository{}
					packService.On("GetAll").Return([]int{5000, 2000, 1000, 500, 250})
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
		{
			name: "Should return 1 1-pack and 1 250-pack for 251 items adding 1 in default pack sizes",
			args: args{
				items: 251,
			},
			fields: fields{
				packRepository: func() contracts.PackRepository {
					packService := &mocks.PackRepository{}
					packService.On("GetAll").Return([]int{5000, 2000, 1000, 500, 250, 1})
					return packService
				}(),
			},
			want: []entity.Pack{
				{
					Size:     250,
					Quantity: 1,
				},
				{
					Size:     1,
					Quantity: 1,
				},
			},
		},
		{
			name: "Should return 83 3-pack and 1-250-pack for 499 items adding 3 in default pack sizes",
			args: args{
				items: 499,
			},
			fields: fields{
				packRepository: func() contracts.PackRepository {
					packService := &mocks.PackRepository{}
					packService.On("GetAll").Return([]int{5000, 2000, 1000, 500, 250, 3})
					return packService
				}(),
			},
			want: []entity.Pack{
				{
					Size:     250,
					Quantity: 1,
				},
				{
					Size:     3,
					Quantity: 83,
				},
			},
		},
		{
			name: "Should return 1 5000-pack for 4001 items",
			args: args{
				items: 4001,
			},
			fields: fields{
				packRepository: func() contracts.PackRepository {
					packService := &mocks.PackRepository{}
					packService.On("GetAll").Return([]int{5000, 2000})
					return packService
				}(),
			},
			want: []entity.Pack{
				{
					Size:     5000,
					Quantity: 1,
				},
			},
		},
		{
			name: "Should return 2 5000-pack for 9100 items",
			args: args{
				items: 9100,
			},
			fields: fields{
				packRepository: func() contracts.PackRepository {
					packService := &mocks.PackRepository{}
					packService.On("GetAll").Return([]int{5000, 2000})
					return packService
				}(),
			},
			want: []entity.Pack{
				{
					Size:     5000,
					Quantity: 2,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &packService{
				repository: tt.fields.packRepository,
			}
			got, gotErr := p.Calculate(tt.args.items)
			if gotErr != nil {
				if !reflect.DeepEqual(gotErr, tt.wantErr) {
					t.Errorf("Calculate() gotErr = %v, wantErr %v", gotErr, tt.wantErr)
				}
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Calculate() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_packService_Create(t *testing.T) {
	type fields struct {
		repository contracts.PackRepository
	}
	type args struct {
		pack int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Should return error if pack size is invalid",
			args: args{
				pack: 0,
			},
			wantErr: true,
		},
		{
			name: "Should return ok if pack was created with success",
			args: args{
				pack: 3,
			},
			fields: fields{
				repository: func() contracts.PackRepository {
					packService := &mocks.PackRepository{}
					packService.On("Create", 3).Return()
					return packService
				}(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &packService{
				repository: tt.fields.repository,
			}
			if err := p.Create(tt.args.pack); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_packService_Remove(t *testing.T) {
	type fields struct {
		repository contracts.PackRepository
	}
	type args struct {
		pack int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Should return ok if pack was created with success",
			args: args{
				pack: 3,
			},
			fields: fields{
				repository: func() contracts.PackRepository {
					packService := &mocks.PackRepository{}
					packService.On("Remove", 3).Return()
					return packService
				}(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &packService{
				repository: tt.fields.repository,
			}
			if err := p.Remove(tt.args.pack); (err != nil) != tt.wantErr {
				t.Errorf("Remove() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
