package repository

const (
	queryCreateEstate = `INSERT INTO estate (width, length) VALUES (:width, :length) RETURNING id;`
	queryGetEstate    = `SELECT * FROM estate where id = :id limit 1`
	queryDoCreateTree = `INSERT INTO tree (estate_id, x, y, height) VALUES (:estate_id, :x, :y, :height) RETURNING id; `
	queryGetTree      = `SELECT * FROM tree where estate_id = :estate_id`
)
