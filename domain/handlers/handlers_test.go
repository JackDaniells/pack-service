package handlers

import (
	"github.com/JackDaniells/pack-service/api"
	"github.com/JackDaniells/pack-service/domain/contracts"
	"github.com/JackDaniells/pack-service/domain/contracts/mocks"
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
			name: "should return 200 when calculate return ok",
			fields: fields{
				packService: func() contracts.PackService {
					packService := &mocks.PackService{}
					packService.On("Calculate", 123).Return(nil)
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
			name: "should return 201 when calculate return ok",
			fields: fields{
				packService: func() contracts.PackService {
					packService := &mocks.PackService{}
					packService.On("Create", 1).Return()
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
			req, _ := http.NewRequest(http.MethodPost, "/pack", strings.NewReader(tt.args.request))
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
			name: "should return 200 when calculate return ok",
			fields: fields{
				packService: func() contracts.PackService {
					packService := &mocks.PackService{}
					packService.On("Remove", 123).Return()
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
			req, _ := http.NewRequest(http.MethodDelete, "/pack/"+tt.args.pack, nil)
			res := httptest.NewRecorder()
			router.ServeHTTP(res, req)

			assert.Equal(t, tt.expectedCode, res.Code)
		})
	}
}
