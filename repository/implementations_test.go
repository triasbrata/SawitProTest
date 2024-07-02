package repository

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/triasbrata/golibs/pkg/dbx"
	"github.com/triasbrata/golibs/pkg/dbx/dbxm"
)

func TestRepository_DoCreateEstate(t *testing.T) {
	dbmock := dbxm.NewDB(t)
	type fields struct {
		Db dbx.DB
	}
	type args struct {
		ctx   context.Context
		input DoCreateEstateRequest
	}
	tests := []struct {
		name       string
		mock       func()
		args       args
		wantOutput DoCreateEstateResponse
		wantErr    bool
	}{
		{
			name: "success",
			mock: func() {
				dbmock.On("BindNamed", queryCreateEstate, DoCreateEstateRequest{
					Width:  10,
					Length: 10,
				}).Return("a", []interface{}{int64(10), int64(10)}, nil).Once()
				dbmock.On("SelectContext", context.Background(), mock.Anything, "a", int64(10), int64(10)).Return(nil).Once().Run(func(args mock.Arguments) {
					mod := args[1].(*[]Estate)
					*mod = append(*mod, Estate{
						ID: "b",
					})
				})
			},
			args: args{
				ctx: context.Background(),
				input: DoCreateEstateRequest{
					Width:  10,
					Length: 10,
				},
			},
			wantOutput: DoCreateEstateResponse{
				ID: "b",
			},
			wantErr: false,
		},
		{
			name: "err",
			mock: func() {
				dbmock.On("BindNamed", queryCreateEstate, DoCreateEstateRequest{
					Width:  10,
					Length: 10,
				}).Return("a", []interface{}{int64(10), int64(10)}, fmt.Errorf("boom")).Once()
			},
			args: args{
				ctx: context.Background(),
				input: DoCreateEstateRequest{
					Width:  10,
					Length: 10,
				},
			},
			wantOutput: DoCreateEstateResponse{},
			wantErr:    true,
		},
		{
			name: "err",
			mock: func() {
				dbmock.On("BindNamed", queryCreateEstate, DoCreateEstateRequest{
					Width:  10,
					Length: 10,
				}).Return("a", []interface{}{int64(10), int64(10)}, nil).Once()
				dbmock.On("SelectContext", context.Background(), mock.Anything, "a", int64(10), int64(10)).Return(fmt.Errorf("boom")).Once()
			},
			args: args{
				ctx: context.Background(),
				input: DoCreateEstateRequest{
					Width:  10,
					Length: 10,
				},
			},
			wantOutput: DoCreateEstateResponse{},
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			r := &Repository{
				Db: dbmock,
			}
			gotOutput, err := r.DoCreateEstate(tt.args.ctx, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Repository.DoCreateEstate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOutput, tt.wantOutput) {
				t.Errorf("Repository.DoCreateEstate() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}

func TestRepository_GetEstateTree(t *testing.T) {
	dbmock := dbxm.NewDB(t)
	type fields struct {
		Db dbx.DB
	}
	type args struct {
		ctx   context.Context
		input GetEstateTreeRequest
	}
	tests := []struct {
		name       string
		mock       func()
		args       args
		wantOutput GetEstateTreeResponse
		wantErr    bool
	}{
		{
			name: "success",
			mock: func() {
				dbmock.On("BindNamed", queryGetTree, GetEstateTreeRequest{
					EstateID: "a",
				}).Return("a", []interface{}{int64(10), int64(10)}, nil).Once()
				dbmock.On("SelectContext", context.Background(), mock.Anything, "a", int64(10), int64(10)).Return(nil).Once().Run(func(args mock.Arguments) {
					mod := args[1].(*[]Tree)
					*mod = append(*mod, Tree{
						ID:       "1",
						EstateID: "a",
						X:        1,
						Y:        2,
						Height:   3,
					})
				})
			},
			args: args{
				ctx: context.Background(),
				input: GetEstateTreeRequest{
					EstateID: "a",
				},
			},
			wantOutput: GetEstateTreeResponse{
				Data: []Tree{
					{
						ID:       "1",
						EstateID: "a",
						X:        1,
						Y:        2,
						Height:   3,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "err",
			mock: func() {
				dbmock.On("BindNamed", queryGetTree, GetEstateTreeRequest{
					EstateID: "a",
				}).Return("a", []interface{}{int64(10), int64(10)}, fmt.Errorf("boom")).Once()
			},
			args: args{
				ctx: context.Background(),
				input: GetEstateTreeRequest{
					EstateID: "a",
				},
			},
			wantOutput: GetEstateTreeResponse{},
			wantErr:    true,
		},
		{
			name: "err",
			mock: func() {
				dbmock.On("BindNamed", queryGetTree, GetEstateTreeRequest{
					EstateID: "a",
				}).Return("a", []interface{}{"a"}, nil).Once()
				dbmock.On("SelectContext", context.Background(), mock.Anything, "a", "a").Return(fmt.Errorf("boom")).Once()
			},
			args: args{
				ctx: context.Background(),
				input: GetEstateTreeRequest{
					EstateID: "a",
				},
			},
			wantOutput: GetEstateTreeResponse{},
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			r := &Repository{
				Db: dbmock,
			}
			gotOutput, err := r.GetEstateTree(tt.args.ctx, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Repository.GetEstateTree() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOutput, tt.wantOutput) {
				t.Errorf("Repository.GetEstateTree() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}
func TestRepository_GetEstate(t *testing.T) {
	dbmock := dbxm.NewDB(t)
	type fields struct {
		Db dbx.DB
	}
	type args struct {
		ctx   context.Context
		input GetEstateRequest
	}
	tests := []struct {
		name       string
		mock       func()
		args       args
		wantOutput GetEstateResponse
		wantErr    bool
	}{
		{
			name: "success",
			mock: func() {
				dbmock.On("BindNamed", queryGetEstate, GetEstateRequest{
					ID: "a",
				}).Return("a", []interface{}{int64(10), int64(10)}, nil).Once()
				dbmock.On("SelectContext", context.Background(), mock.Anything, "a", int64(10), int64(10)).Return(nil).Once().Run(func(args mock.Arguments) {
					mod := args[1].(*[]Estate)
					*mod = append(*mod, Estate{
						ID:     "a",
						Width:  1,
						Length: 2,
					})
				})
			},
			args: args{
				ctx: context.Background(),
				input: GetEstateRequest{
					ID: "a",
				},
			},
			wantOutput: GetEstateResponse{
				Data: []Estate{
					{
						ID:     "a",
						Width:  1,
						Length: 2,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "err",
			mock: func() {
				dbmock.On("BindNamed", queryGetEstate, GetEstateRequest{
					ID: "a",
				}).Return("a", []interface{}{int64(10), int64(10)}, fmt.Errorf("boom")).Once()
			},
			args: args{
				ctx: context.Background(),
				input: GetEstateRequest{
					ID: "a",
				},
			},
			wantOutput: GetEstateResponse{},
			wantErr:    true,
		},
		{
			name: "err",
			mock: func() {
				dbmock.On("BindNamed", queryGetEstate, GetEstateRequest{
					ID: "a",
				}).Return("a", []interface{}{"a"}, nil).Once()
				dbmock.On("SelectContext", context.Background(), mock.Anything, "a", "a").Return(fmt.Errorf("boom")).Once()
			},
			args: args{
				ctx: context.Background(),
				input: GetEstateRequest{
					ID: "a",
				},
			},
			wantOutput: GetEstateResponse{},
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			r := &Repository{
				Db: dbmock,
			}
			gotOutput, err := r.GetEstate(tt.args.ctx, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Repository.GetEstate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOutput, tt.wantOutput) {
				t.Errorf("Repository.GetEstate() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}
func TestRepository_DoCreateTree(t *testing.T) {
	dbmock := dbxm.NewDB(t)
	type fields struct {
		Db dbx.DB
	}
	type args struct {
		ctx   context.Context
		input DoCreateTreeRequest
	}
	tests := []struct {
		name       string
		mock       func()
		args       args
		wantOutput DoCreateTreeResponse
		wantErr    bool
	}{
		{
			name: "success",
			mock: func() {
				dbmock.On("BindNamed", queryDoCreateTree, DoCreateTreeRequest{
					EstateID: "a",
					X:        1,
					Y:        2,
					Height:   3,
				}).Return("a", []interface{}{int64(10), int64(10)}, nil).Once()
				dbmock.On("SelectContext", context.Background(), mock.Anything, "a", int64(10), int64(10)).Return(nil).Once().Run(func(args mock.Arguments) {
					mod := args[1].(*[]Tree)
					*mod = append(*mod, Tree{
						ID: "a",
					})
				})
			},
			args: args{
				ctx: context.Background(),
				input: DoCreateTreeRequest{
					EstateID: "a",
					X:        1,
					Y:        2,
					Height:   3,
				},
			},
			wantOutput: DoCreateTreeResponse{
				ID: "a",
			},
			wantErr: false,
		},
		{
			name: "err",
			mock: func() {
				dbmock.On("BindNamed", queryDoCreateTree, DoCreateTreeRequest{
					EstateID: "a",
					X:        1,
					Y:        2,
					Height:   3,
				}).Return("a", []interface{}{int64(10), int64(10)}, fmt.Errorf("boom")).Once()
			},
			args: args{
				ctx: context.Background(),
				input: DoCreateTreeRequest{
					EstateID: "a",
					X:        1,
					Y:        2,
					Height:   3,
				},
			},
			wantOutput: DoCreateTreeResponse{},
			wantErr:    true,
		},
		{
			name: "err",
			mock: func() {
				dbmock.On("BindNamed", queryDoCreateTree, DoCreateTreeRequest{
					EstateID: "a",
					X:        1,
					Y:        2,
					Height:   3,
				}).Return("a", []interface{}{"a"}, nil).Once()
				dbmock.On("SelectContext", context.Background(), mock.Anything, "a", "a").Return(fmt.Errorf("boom")).Once()
			},
			args: args{
				ctx: context.Background(),
				input: DoCreateTreeRequest{
					EstateID: "a",
					X:        1,
					Y:        2,
					Height:   3,
				},
			},
			wantOutput: DoCreateTreeResponse{},
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			r := &Repository{
				Db: dbmock,
			}
			gotOutput, err := r.DoCreateTree(tt.args.ctx, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Repository.DoCreateTree() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOutput, tt.wantOutput) {
				t.Errorf("Repository.DoCreateTree() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}
