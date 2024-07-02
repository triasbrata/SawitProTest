package repository

import (
	"context"
)

// DoCreateEstate implements RepositoryInterface.
func (r *Repository) DoCreateEstate(ctx context.Context, input DoCreateEstateRequest) (output DoCreateEstateResponse, err error) {
	output = DoCreateEstateResponse{}
	query, args, err := r.Db.BindNamed(queryCreateEstate, input)
	if err != nil {
		return DoCreateEstateResponse{}, err
	}

	err = r.Db.SelectContext(ctx, &output.ID, query, args...)
	if err != nil {
		return DoCreateEstateResponse{}, err
	}
	return

}

// GetEstate implements RepositoryInterface.
func (r *Repository) GetEstate(ctx context.Context, input GetEstateRequest) (output GetEstateResponse, err error) {
	query, args, err := r.Db.BindNamed(queryGetEstate, input)
	if err != nil {
		return GetEstateResponse{}, err
	}
	output = GetEstateResponse{
		Data: make([]Estate, 0),
	}

	err = r.Db.SelectContext(ctx, &output.Data, query, args...)
	if err != nil {
		return GetEstateResponse{}, err
	}
	return

}

// DoCreateTree implements RepositoryInterface.
func (r *Repository) DoCreateTree(ctx context.Context, input DoCreateTreeRequest) (output DoCreateTreeResponse, err error) {
	output = DoCreateTreeResponse{}
	query, args, err := r.Db.BindNamed(queryDoCreateTree, input)
	if err != nil {
		return DoCreateTreeResponse{}, err
	}

	err = r.Db.SelectContext(ctx, &output.ID, query, args...)
	if err != nil {
		return DoCreateTreeResponse{}, err
	}
	return
}

// GetEstateTree implements RepositoryInterface.
func (r *Repository) GetEstateTree(ctx context.Context, input GetEstateTreeRequest) (output GetEstateTreeResponse, err error) {
	query, args, err := r.Db.BindNamed(queryGetTree, input)
	if err != nil {
		return GetEstateTreeResponse{}, err
	}
	output = GetEstateTreeResponse{
		Data: make([]Tree, 0),
	}

	err = r.Db.SelectContext(ctx, &output.Data, query, args...)
	if err != nil {
		return GetEstateTreeResponse{}, err
	}
	return
}
