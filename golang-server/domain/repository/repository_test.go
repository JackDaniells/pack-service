package repository

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func Test_packRepository_Create(t *testing.T) {
	type fields struct {
		packs []int
	}
	type args struct {
		pack int
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantList []int
	}{
		{
			name: "Should add pack if it's not exists",
			fields: fields{
				packs: []int{1, 2, 3},
			},
			args: args{
				pack: 4,
			},
			wantList: []int{1, 2, 3, 4},
		},
		{
			name: "Should ignore if pack already exists",
			fields: fields{
				packs: []int{1, 2, 3},
			},
			args: args{
				pack: 3,
			},
			wantList: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &packRepository{
				packs: tt.fields.packs,
			}
			p.Create(tt.args.pack)
			assert.Equal(t, p.packs, tt.wantList)
		})
	}
}

func Test_packRepository_GetAll(t *testing.T) {
	type fields struct {
		packs []int
	}
	tests := []struct {
		name   string
		fields fields
		want   []int
	}{
		{
			name: "Should return all packs in descending order",
			fields: fields{
				packs: []int{1, 3, 2, 5, 4},
			},
			want: []int{5, 4, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &packRepository{
				packs: tt.fields.packs,
			}
			if got := p.GetAll(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_packRepository_Remove(t *testing.T) {
	type fields struct {
		packs []int
	}
	type args struct {
		packToRemove int
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantList []int
	}{
		{
			name: "Should remove if element exists in pack list",
			fields: fields{
				packs: []int{1, 3, 2, 5, 4},
			},
			args: args{
				packToRemove: 1,
			},
			wantList: []int{3, 2, 5, 4},
		},
		{
			name: "Should did nothing if element doesnt exists in pack list",
			fields: fields{
				packs: []int{1, 3, 2, 5, 4},
			},
			args: args{
				packToRemove: 6,
			},
			wantList: []int{1, 3, 2, 5, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &packRepository{
				packs: tt.fields.packs,
			}
			p.Remove(tt.args.packToRemove)
			assert.Equal(t, p.packs, tt.wantList)

		})
	}
}

func Test_packRepository_AddList(t *testing.T) {
	type fields struct {
		packs []int
	}
	type args struct {
		packsToAdd []int
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantList []int
	}{
		{
			name: "Should add all non-duplicate items to the pack list",
			fields: fields{
				packs: []int{1, 2, 3, 4},
			},
			args: args{
				packsToAdd: []int{3, 4, 5, 6, 7, 8},
			},
			wantList: []int{1, 2, 3, 4, 5, 6, 7, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &packRepository{
				packs: tt.fields.packs,
			}
			p.AddList(tt.args.packsToAdd)

			assert.Equal(t, tt.wantList, p.packs)
		})
	}
}

func Test_packRepository_RemoveList(t *testing.T) {
	type fields struct {
		packs []int
	}
	type args struct {
		packsToRemove []int
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantList []int
	}{
		{
			name: "Should remove all sent items to the pack list",
			fields: fields{
				packs: []int{1, 2, 3, 4, 5, 6},
			},
			args: args{
				packsToRemove: []int{3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
			},
			wantList: []int{1, 2},
		},
		{
			name: "Should not remove items if no item from the sent list is found in the pack array",
			fields: fields{
				packs: []int{1, 2, 3, 4},
			},
			args: args{
				packsToRemove: []int{5, 6, 7, 8, 9, 10},
			},
			wantList: []int{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &packRepository{
				packs: tt.fields.packs,
			}
			p.RemoveList(tt.args.packsToRemove)

			assert.Equal(t, tt.wantList, p.packs)
		})
	}
}
