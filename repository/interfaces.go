// This file contains the interfaces for the repository layer.
// The repository layer is responsible for interacting with the database.
// For testing purpose we will generate mock implementations of these
// interfaces using mockgen. See the Makefile for more information.
package repository

import "context"

type RepositoryInterface interface {
	DoCreateEstate(ctx context.Context, input DoCreateEstateRequest) (output DoCreateEstateResponse, err error)
	DoCreateTree(ctx context.Context, input DoCreateTreeRequest) (output DoCreateTreeResponse, err error)
	GetEstateTree(ctx context.Context, input GetEstateTreeRequest) (output GetEstateTreeResponse, err error)
	GetEstate(ctx context.Context, input GetEstateRequest) (output GetEstateResponse, err error)
}
