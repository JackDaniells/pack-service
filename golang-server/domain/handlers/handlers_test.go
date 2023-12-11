package handlers

import (
	"errors"
	"github.com/JackDaniells/pack-service/api"
	"github.com/JackDaniells/pack-service/domain/contracts"
	"github.com/JackDaniells/pack-service/domain/contracts/mocks"
	"github.com/JackDaniells/pack-service/domain/entity"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_packHandler_Calculate(t *testing.T) {
	type fields struct {
		packService contracts.PackService
	}
	type args struct {
		items string
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		expectedCode int
	}{
		{
			name: "should return 400 when parse items fails",
			args: args{
				items: "abc",
			},
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "should return 400 when calculate service return error",
			fields: fields{
				packService: func() contracts.PackService {
					packService := &mocks.PackService{}
					packService.On("Calculate", 123).
						Return(nil, errors.New("mock error"))
					return packService
				}(),
			},
			args: args{
				items: "123",
			},
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "should return 200 when calculate return ok",
			fields: fields{
				packService: func() contracts.PackService {
					packService := &mocks.PackService{}
					packService.On("Calculate", 123).Return(nil, nil)
					return packService
				}(),
			},
			args: args{
				items: "123",
			},
			expectedCode: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := &packHandler{
				packService: tt.fields.packService,
			}

			router := api.NewMuxRouter(handler)
			req, _ := http.NewRequest(http.MethodGet, "/calculate?items="+tt.args.items, nil)
			res := httptest.NewRecorder()
			router.ServeHTTP(res, req)

			assert.Equal(t, tt.expectedCode, res.Code)
		})
	}
}

func Test_packHandler_GetAll(t *testing.T) {
	type fields struct {
		packService contracts.PackService
	}
	tests := []struct {
		name         string
		fields       fields
		expectedCode int
	}{
		{
			name: "should return 400 when getAll service return error",
			fields: fields{
				packService: func() contracts.PackService {
					packService := &mocks.PackService{}
					packService.On("GetAll").
						Return(nil, errors.New("mock error"))
					return packService
				}(),
			},
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "should return 200 when calculate return ok",
			fields: fields{
				packService: func() contracts.PackService {
					packService := &mocks.PackService{}
					packService.On("GetAll").Return([]entity.Pack{}, nil)
					return packService
				}(),
			},
			expectedCode: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := &packHandler{
				packService: tt.fields.packService,
			}

			router := api.NewMuxRouter(handler)
			req, _ := http.NewRequest(http.MethodGet, "/packs", nil)
			res := httptest.NewRecorder()
			router.ServeHTTP(res, req)

			assert.Equal(t, tt.expectedCode, res.Code)
		})
	}
}

func Test_packHandler_Create(t *testing.T) {
	type fields struct {
		packService contracts.PackService
	}
	type args struct {
		request string
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		expectedCode int
	}{
		{
			name: "should return 400 when parse request fails",
			args: args{
				request: "mock error",
			},
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "should return 201 when create service return error",
			fields: fields{
				packService: func() contracts.PackService {
					packService := &mocks.PackService{}
					packService.On("Create", 1).Return(errors.New("mock error"))
					return packService
				}(),
			},
			args: args{
				request: `
					{
						"size": 1
					}
				`,
			},
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "should return 201 when create return ok",
			fields: fields{
				packService: func() contracts.PackService {
					packService := &mocks.PackService{}
					packService.On("Create", 1).Return(nil)
					return packService
				}(),
			},
			args: args{
				request: `
					{
						"size": 1
					}
				`,
			},
			expectedCode: http.StatusCreated,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := &packHandler{
				packService: tt.fields.packService,
			}

			router := api.NewMuxRouter(handler)
			req, _ := http.NewRequest(http.MethodPost, "/packs", strings.NewReader(tt.args.request))
			res := httptest.NewRecorder()
			router.ServeHTTP(res, req)

			assert.Equal(t, tt.expectedCode, res.Code)
		})
	}
}

func Test_packHandler_Remove(t *testing.T) {
	type fields struct {
		packService contracts.PackService
	}
	type args struct {
		pack string
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		expectedCode int
	}{
		{
			name: "should return 400 when parse pack to int fails",
			args: args{
				pack: "abc",
			},
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "should return 400 when remove service return error",
			fields: fields{
				packService: func() contracts.PackService {
					packService := &mocks.PackService{}
					packService.On("Remove", 123).Return(errors.New("mock error"))
					return packService
				}(),
			},
			args: args{
				pack: "123",
			},
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "should return 200 when remove pack return ok",
			fields: fields{
				packService: func() contracts.PackService {
					packService := &mocks.PackService{}
					packService.On("Remove", 123).Return(nil)
					return packService
				}(),
			},
			args: args{
				pack: "123",
			},
			expectedCode: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := &packHandler{
				packService: tt.fields.packService,
			}

			router := api.NewMuxRouter(handler)
			req, _ := http.NewRequest(http.MethodDelete, "/packs/"+tt.args.pack, nil)
			res := httptest.NewRecorder()
			router.ServeHTTP(res, req)

			assert.Equal(t, tt.expectedCode, res.Code)
		})
	}
}
