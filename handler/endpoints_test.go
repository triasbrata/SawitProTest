package handler

import (
	"testing"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/repository"
	"github.com/labstack/echo/v4"
)

func TestServer_GetEstateIdDronePlan(t *testing.T) {
	type fields struct {
		Repository repository.RepositoryInterface
	}
	type args struct {
		ctx    echo.Context
		id     string
		params generated.GetEstateIdDronePlanParams
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				Repository: tt.fields.Repository,
			}
			if err := s.GetEstateIdDronePlan(tt.args.ctx, tt.args.id, tt.args.params); (err != nil) != tt.wantErr {
				t.Errorf("Server.GetEstateIdDronePlan() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestServer_GetEstateIdStats(t *testing.T) {
	type fields struct {
		Repository repository.RepositoryInterface
	}
	type args struct {
		ctx echo.Context
		id  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				Repository: tt.fields.Repository,
			}
			if err := s.GetEstateIdStats(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Server.GetEstateIdStats() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestServer_PostEstate(t *testing.T) {
	type fields struct {
		Repository repository.RepositoryInterface
	}
	type args struct {
		ctx echo.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				Repository: tt.fields.Repository,
			}
			if err := s.PostEstate(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Server.PostEstate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestServer_PostEstateIdTree(t *testing.T) {
	type fields struct {
		Repository repository.RepositoryInterface
	}
	type args struct {
		ctx echo.Context
		id  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				Repository: tt.fields.Repository,
			}
			if err := s.PostEstateIdTree(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Server.PostEstateIdTree() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
