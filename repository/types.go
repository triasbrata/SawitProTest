// This file contains types that are used in the repository layer.
package repository

type DoCreateEstateRequest struct {
	Width  int64 `db:"width"`
	Length int64 `db:"length"`
}
type DoCreateTreeRequest struct {
	EstateID string `db:"id"`
	X        int64  `db:"x"`
	Y        int64  `db:"y"`
	Height   int64  `db:"height"`
}
type Tree struct {
	ID       string `db:"id"`
	EstateID string `db:"estate_id"`
	X        int64  `db:"x"`
	Y        int64  `db:"y"`
	Height   int64  `db:"height"`
}
type GetEstateTreeRequest struct {
	EstateID string `db:"estate_id"`
}
type DoCreateEstateResponse struct {
	ID string
}
type DoCreateTreeResponse struct {
	ID string
}
type GetEstateTreeResponse struct {
	Data []Tree
}
type GetEstateRequest struct {
	ID string `db:"string"`
}
type GetEstateResponse struct {
	Data []Estate
}
type Estate struct {
	Width  int64  `db:"width"`
	Length int64  `db:"length"`
	ID     string `db:"id"`
}
