package handler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/repository"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestServer_GetEstateIdDronePlan(t *testing.T) {
	ctrl := gomock.NewController(t)
	maxDistance := 100
	mockRepo := repository.NewMockRepositoryInterface(ctrl)
	defer ctrl.Finish()

	type args struct {
		ctx    echo.Context
		id     string
		params generated.GetEstateIdDronePlanParams
	}
	tests := []struct {
		name     string
		mock     func()
		ctx      func() (echo.Context, *httptest.ResponseRecorder)
		args     args
		want     string
		wantCode int
	}{
		{
			name: "success",
			mock: func() {
				mockRepo.EXPECT().GetEstate(gomock.All(), gomock.Eq(repository.GetEstateRequest{
					ID: "a",
				})).Return(repository.GetEstateResponse{
					Data: []repository.Estate{
						{
							Width:  2,
							Length: 5,
						},
					},
				}, nil)
				mockRepo.EXPECT().GetEstateTree(gomock.All(), gomock.Eq(repository.GetEstateTreeRequest{
					EstateID: "a",
				})).Return(repository.GetEstateTreeResponse{
					Data: []repository.Tree{
						{
							EstateID: "a",
							ID:       "b",
							X:        3,
							Y:        3,
							Height:   2,
						},
					},
				}, nil)
			},
			ctx: func() (echo.Context, *httptest.ResponseRecorder) {
				req := httptest.NewRequest(http.MethodGet, "/", nil)
				res := httptest.NewRecorder()
				return echo.New().NewContext(req, res), res
			},
			args: args{
				id:     "a",
				params: generated.GetEstateIdDronePlanParams{},
			},
			wantCode: 200,
			want:     `{"distance":146}`,
		},
		{
			name: "error when request tree data",
			mock: func() {
				mockRepo.EXPECT().GetEstate(gomock.All(), gomock.Eq(repository.GetEstateRequest{
					ID: "a",
				})).Return(repository.GetEstateResponse{
					Data: []repository.Estate{
						{
							Width:  2,
							Length: 5,
						},
					},
				}, nil)
				mockRepo.EXPECT().GetEstateTree(gomock.All(), gomock.Eq(repository.GetEstateTreeRequest{
					EstateID: "a",
				})).Return(repository.GetEstateTreeResponse{
					Data: []repository.Tree{
						{
							EstateID: "a",
							ID:       "b",
							X:        3,
							Y:        3,
							Height:   2,
						},
					},
				}, fmt.Errorf("boom"))
			},
			ctx: func() (echo.Context, *httptest.ResponseRecorder) {
				req := httptest.NewRequest(http.MethodGet, "/", nil)
				res := httptest.NewRecorder()
				return echo.New().NewContext(req, res), res
			},
			args: args{
				id:     "a",
				params: generated.GetEstateIdDronePlanParams{},
			},
			wantCode: 500,
			want:     `{"statusCode":500,"success":false,"message":"boom"}`,
		},
		{
			name: "error when request estate data",
			mock: func() {
				mockRepo.EXPECT().GetEstate(gomock.All(), gomock.Eq(repository.GetEstateRequest{
					ID: "a",
				})).Return(repository.GetEstateResponse{
					Data: []repository.Estate{
						{
							Width:  2,
							Length: 5,
						},
					},
				}, fmt.Errorf("boom"))
			},
			ctx: func() (echo.Context, *httptest.ResponseRecorder) {
				req := httptest.NewRequest(http.MethodGet, "/", nil)
				res := httptest.NewRecorder()
				return echo.New().NewContext(req, res), res
			},
			args: args{
				id:     "a",
				params: generated.GetEstateIdDronePlanParams{},
			},
			wantCode: 500,
			want:     `{"statusCode":500,"success":false,"message":"boom"}`,
		},
		{
			name: "error estate not found",
			mock: func() {
				mockRepo.EXPECT().GetEstate(gomock.All(), gomock.Eq(repository.GetEstateRequest{
					ID: "a",
				})).Return(repository.GetEstateResponse{
					Data: []repository.Estate{},
				}, nil)
			},
			ctx: func() (echo.Context, *httptest.ResponseRecorder) {
				req := httptest.NewRequest(http.MethodGet, "/", nil)
				res := httptest.NewRecorder()
				return echo.New().NewContext(req, res), res
			},
			args: args{
				id:     "a",
				params: generated.GetEstateIdDronePlanParams{},
			},
			wantCode: 404,
			want:     `{"statusCode":404,"success":false,"message":"estate with id a is not found"}`,
		},

		{
			name: "success with max distance",
			mock: func() {
				mockRepo.EXPECT().GetEstate(gomock.All(), gomock.Eq(repository.GetEstateRequest{
					ID: "a",
				})).Return(repository.GetEstateResponse{
					Data: []repository.Estate{
						{
							Width:  2,
							Length: 5,
						},
					},
				}, nil)
				mockRepo.EXPECT().GetEstateTree(gomock.All(), gomock.Eq(repository.GetEstateTreeRequest{
					EstateID: "a",
				})).Return(repository.GetEstateTreeResponse{
					Data: []repository.Tree{
						{
							EstateID: "a",
							ID:       "b",
							X:        3,
							Y:        3,
							Height:   2,
						},
					},
				}, nil)
			},
			ctx: func() (echo.Context, *httptest.ResponseRecorder) {
				req := httptest.NewRequest(http.MethodGet, "/", nil)
				res := httptest.NewRecorder()
				return echo.New().NewContext(req, res), res
			},
			args: args{
				id: "a",
				params: generated.GetEstateIdDronePlanParams{
					MaxDistance: &maxDistance,
				},
			},
			wantCode: 200,
			want:     `{"distance":100,"rest":{"x":1,"y":3}}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			s := &Server{
				Repository: mockRepo,
			}
			ctx, res := tt.ctx()
			s.GetEstateIdDronePlan(ctx, tt.args.id, tt.args.params)
			assert.Equal(t, tt.want, strings.Trim(res.Body.String(), "\n"))
			assert.Equal(t, tt.wantCode, res.Code)
		})
	}
}

func TestServer_GetEstateIdStats(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := repository.NewMockRepositoryInterface(ctrl)
	defer ctrl.Finish()

	type args struct {
		ctx echo.Context
		id  string
	}
	tests := []struct {
		name     string
		mock     func()
		ctx      func() (echo.Context, *httptest.ResponseRecorder)
		args     args
		want     string
		wantCode int
	}{
		{
			name: "success",
			mock: func() {
				mockRepo.EXPECT().GetEstate(gomock.All(), gomock.Eq(repository.GetEstateRequest{
					ID: "a",
				})).Return(repository.GetEstateResponse{
					Data: []repository.Estate{
						{
							Width:  2,
							Length: 5,
						},
					},
				}, nil)
				mockRepo.EXPECT().GetEstateTree(gomock.All(), gomock.Eq(repository.GetEstateTreeRequest{
					EstateID: "a",
				})).Return(repository.GetEstateTreeResponse{
					Data: []repository.Tree{
						{
							EstateID: "a",
							ID:       "b",
							X:        3,
							Y:        3,
							Height:   2,
						},
					},
				}, nil)
			},
			ctx: func() (echo.Context, *httptest.ResponseRecorder) {
				req := httptest.NewRequest(http.MethodGet, "/", nil)
				res := httptest.NewRecorder()
				return echo.New().NewContext(req, res), res
			},
			args: args{
				id: "a",
			},
			wantCode: 200,
			want:     `{"count":1,"max":2,"min":2,"median":2}`,
		},
		{
			name: "error when request tree data",
			mock: func() {
				mockRepo.EXPECT().GetEstate(gomock.All(), gomock.Eq(repository.GetEstateRequest{
					ID: "a",
				})).Return(repository.GetEstateResponse{
					Data: []repository.Estate{
						{
							Width:  2,
							Length: 5,
						},
					},
				}, nil)
				mockRepo.EXPECT().GetEstateTree(gomock.All(), gomock.Eq(repository.GetEstateTreeRequest{
					EstateID: "a",
				})).Return(repository.GetEstateTreeResponse{
					Data: []repository.Tree{
						{
							EstateID: "a",
							ID:       "b",
							X:        3,
							Y:        3,
							Height:   2,
						},
					},
				}, fmt.Errorf("boom"))
			},
			ctx: func() (echo.Context, *httptest.ResponseRecorder) {
				req := httptest.NewRequest(http.MethodGet, "/", nil)
				res := httptest.NewRecorder()
				return echo.New().NewContext(req, res), res
			},
			args: args{
				id: "a",
			},
			wantCode: 500,
			want:     `{"statusCode":500,"success":false,"message":"boom"}`,
		},
		{
			name: "error when request estate data",
			mock: func() {
				mockRepo.EXPECT().GetEstate(gomock.All(), gomock.Eq(repository.GetEstateRequest{
					ID: "a",
				})).Return(repository.GetEstateResponse{
					Data: []repository.Estate{
						{
							Width:  2,
							Length: 5,
						},
					},
				}, fmt.Errorf("boom"))
			},
			ctx: func() (echo.Context, *httptest.ResponseRecorder) {
				req := httptest.NewRequest(http.MethodGet, "/", nil)
				res := httptest.NewRecorder()
				return echo.New().NewContext(req, res), res
			},
			args: args{
				id: "a",
			},
			wantCode: 500,
			want:     `{"statusCode":500,"success":false,"message":"boom"}`,
		},
		{
			name: "error estate not found",
			mock: func() {
				mockRepo.EXPECT().GetEstate(gomock.All(), gomock.Eq(repository.GetEstateRequest{
					ID: "a",
				})).Return(repository.GetEstateResponse{
					Data: []repository.Estate{},
				}, nil)
			},
			ctx: func() (echo.Context, *httptest.ResponseRecorder) {
				req := httptest.NewRequest(http.MethodGet, "/", nil)
				res := httptest.NewRecorder()
				return echo.New().NewContext(req, res), res
			},
			args: args{
				id: "a",
			},
			wantCode: 404,
			want:     `{"statusCode":404,"success":false,"message":"estate with id a is not found"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			s := &Server{
				Repository: mockRepo,
			}
			ctx, res := tt.ctx()
			s.GetEstateIdStats(ctx, tt.args.id)
			assert.Equal(t, tt.want, strings.Trim(res.Body.String(), "\n"))
			assert.Equal(t, tt.wantCode, res.Code)
		})
	}
}

func TestServer_PostEstate(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := repository.NewMockRepositoryInterface(ctrl)
	defer ctrl.Finish()

	tests := []struct {
		name     string
		mock     func()
		ctx      func() (echo.Context, *httptest.ResponseRecorder)
		want     string
		wantCode int
	}{
		{
			name: "success",
			mock: func() {
				mockRepo.EXPECT().DoCreateEstate(gomock.All(), gomock.Eq(repository.DoCreateEstateRequest{
					Width:  100,
					Length: 100,
				})).Return(repository.DoCreateEstateResponse{
					ID: "a",
				}, nil)
			},
			ctx: func() (echo.Context, *httptest.ResponseRecorder) {
				req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"width":100,"length":100}`))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				res := httptest.NewRecorder()
				return echo.New().NewContext(req, res), res
			},
			wantCode: 200,
			want:     `{"id":"a"}`,
		},
		{
			name: "error when create data",
			mock: func() {
				mockRepo.EXPECT().DoCreateEstate(gomock.All(), gomock.Eq(repository.DoCreateEstateRequest{
					Width:  100,
					Length: 100,
				})).Return(repository.DoCreateEstateResponse{
					ID: "a",
				}, fmt.Errorf("boom"))
			},
			ctx: func() (echo.Context, *httptest.ResponseRecorder) {
				req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"width":100,"length":100}`))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				res := httptest.NewRecorder()
				return echo.New().NewContext(req, res), res
			},
			wantCode: 500,
			want:     `{"statusCode":500,"success":false,"message":"boom"}`,
		},
		{
			name: "error parse json",
			mock: func() {
			},
			ctx: func() (echo.Context, *httptest.ResponseRecorder) {
				req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"width:100,"length":100}`))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				res := httptest.NewRecorder()
				return echo.New().NewContext(req, res), res
			},
			wantCode: 400,
			want:     `{"statusCode":400,"success":false,"message":"code=400, message=Syntax error: offset=14, error=invalid character 'l' after object key, internal=invalid character 'l' after object key"}`,
		},
		{
			name: "error out of bound",
			mock: func() {
			},
			ctx: func() (echo.Context, *httptest.ResponseRecorder) {
				req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"width":10000000,"length":100}`))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				res := httptest.NewRecorder()
				return echo.New().NewContext(req, res), res
			},
			wantCode: 400,
			want:     `{"statusCode":400,"success":false,"message":"Width of estate max is only 50000 and min is 1"}`,
		},
		{
			name: "error out of bound length",
			mock: func() {
			},
			ctx: func() (echo.Context, *httptest.ResponseRecorder) {
				req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"width":100,"length":10000000}`))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				res := httptest.NewRecorder()
				return echo.New().NewContext(req, res), res
			},
			wantCode: 400,
			want:     `{"statusCode":400,"success":false,"message":"Length of estate max is only 50000 and min is 1"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			s := &Server{
				Repository: mockRepo,
			}
			ctx, res := tt.ctx()
			s.PostEstate(ctx)
			assert.Equal(t, tt.want, strings.Trim(res.Body.String(), "\n"))
			assert.Equal(t, tt.wantCode, res.Code)
		})
	}
}

func TestServer_PostEstateIdTree(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := repository.NewMockRepositoryInterface(ctrl)
	defer ctrl.Finish()

	tests := []struct {
		name     string
		mock     func()
		ctx      func() (echo.Context, *httptest.ResponseRecorder)
		want     string
		wantCode int
	}{
		{
			name: "success",
			mock: func() {
				mockRepo.EXPECT().GetEstate(gomock.All(), gomock.Eq(repository.GetEstateRequest{
					ID: "a",
				})).Return(repository.GetEstateResponse{
					Data: []repository.Estate{
						{
							Width:  15,
							Length: 15,
							ID:     "a",
						},
					},
				}, nil)
				mockRepo.EXPECT().DoCreateTree(gomock.All(), gomock.Eq(repository.DoCreateTreeRequest{
					EstateID: "a",
					X:        1,
					Y:        2,
					Height:   3,
				})).Return(repository.DoCreateTreeResponse{
					ID: "b",
				}, nil)
			},
			ctx: func() (echo.Context, *httptest.ResponseRecorder) {
				req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{ "x": 1, "y": 2, "height": 3 }`))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				res := httptest.NewRecorder()
				return echo.New().NewContext(req, res), res
			},
			wantCode: 200,
			want:     `{"id":"b"}`,
		},
		{
			name: "error when create tree",
			mock: func() {
				mockRepo.EXPECT().GetEstate(gomock.All(), gomock.Eq(repository.GetEstateRequest{
					ID: "a",
				})).Return(repository.GetEstateResponse{
					Data: []repository.Estate{
						{
							Width:  15,
							Length: 15,
							ID:     "a",
						},
					},
				}, nil)
				mockRepo.EXPECT().DoCreateTree(gomock.All(), gomock.Eq(repository.DoCreateTreeRequest{
					EstateID: "a",
					X:        1,
					Y:        2,
					Height:   3,
				})).Return(repository.DoCreateTreeResponse{
					ID: "b",
				}, fmt.Errorf("boom"))
			},
			ctx: func() (echo.Context, *httptest.ResponseRecorder) {
				req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{ "x": 1, "y": 2, "height": 3 }`))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				res := httptest.NewRecorder()
				return echo.New().NewContext(req, res), res
			},
			wantCode: 500,
			want:     `{"statusCode":500,"success":false,"message":"boom"}`,
		},
		{
			name: "error when create tree",
			mock: func() {
				mockRepo.EXPECT().GetEstate(gomock.All(), gomock.Eq(repository.GetEstateRequest{
					ID: "a",
				})).Return(repository.GetEstateResponse{
					Data: []repository.Estate{
						{
							Width:  1,
							Length: 1,
							ID:     "a",
						},
					},
				}, nil)
			},
			ctx: func() (echo.Context, *httptest.ResponseRecorder) {
				req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{ "x": 1, "y": 2, "height": 3 }`))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				res := httptest.NewRecorder()
				return echo.New().NewContext(req, res), res
			},
			wantCode: 400,
			want:     `{"statusCode":400,"success":false,"message":"out of bound"}`,
		},
		{
			name: "error when fetch data estate",
			mock: func() {
				mockRepo.EXPECT().GetEstate(gomock.All(), gomock.Eq(repository.GetEstateRequest{
					ID: "a",
				})).Return(repository.GetEstateResponse{
					Data: []repository.Estate{
						{
							Width:  1,
							Length: 1,
							ID:     "a",
						},
					},
				}, fmt.Errorf("boom"))
			},
			ctx: func() (echo.Context, *httptest.ResponseRecorder) {
				req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{ "x": 1, "y": 2, "height": 3 }`))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				res := httptest.NewRecorder()
				return echo.New().NewContext(req, res), res
			},
			wantCode: 500,
			want:     `{"statusCode":500,"success":false,"message":"boom"}`,
		},
		{
			name: "error estate is empty",
			mock: func() {
				mockRepo.EXPECT().GetEstate(gomock.All(), gomock.Eq(repository.GetEstateRequest{
					ID: "a",
				})).Return(repository.GetEstateResponse{
					Data: []repository.Estate{},
				}, nil)
			},
			ctx: func() (echo.Context, *httptest.ResponseRecorder) {
				req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{ "x": 1, "y": 2, "height": 3 }`))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				res := httptest.NewRecorder()
				return echo.New().NewContext(req, res), res
			},
			wantCode: 404,
			want:     `{"statusCode":404,"success":false,"message":"estate not found"}`,
		},
		{
			name: "error out of bound",
			mock: func() {
			},
			ctx: func() (echo.Context, *httptest.ResponseRecorder) {
				req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{ "x": 0, "y": 0, "height": 3 }`))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				res := httptest.NewRecorder()
				return echo.New().NewContext(req, res), res
			},
			wantCode: 400,
			want:     `{"statusCode":400,"success":false,"message":"out of bound"}`,
		},
		{
			name: "error out of bound hight",
			mock: func() {
			},
			ctx: func() (echo.Context, *httptest.ResponseRecorder) {
				req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{ "x": 1, "y": 1, "height": 31 }`))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				res := httptest.NewRecorder()
				return echo.New().NewContext(req, res), res
			},
			wantCode: 400,
			want:     `{"statusCode":400,"success":false,"message":"height of tree max is only 30"}`,
		},
		{
			name: "error bind",
			mock: func() {
			},
			ctx: func() (echo.Context, *httptest.ResponseRecorder) {
				req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{ "x: 1, "y": 1, "height": 31 }`))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				res := httptest.NewRecorder()
				return echo.New().NewContext(req, res), res
			},
			wantCode: 400,
			want:     `{"statusCode":400,"success":false,"message":"code=400, message=Syntax error: offset=11, error=invalid character 'y' after object key, internal=invalid character 'y' after object key"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			s := &Server{
				Repository: mockRepo,
			}
			ctx, res := tt.ctx()
			s.PostEstateIdTree(ctx, "a")
			assert.Equal(t, tt.want, strings.Trim(res.Body.String(), "\n"))
			assert.Equal(t, tt.wantCode, res.Code)
		})
	}
}
