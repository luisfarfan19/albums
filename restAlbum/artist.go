package main

import "errors"

type Artist struct {
	Id     int
	Name   string
	Year   int
	Origin string
}

func CreateArtist(a Artist) error {

	query := `INSERT INTO artist (id, name, year, origin) VALUES ($1, $2, $3, $4)`

	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	r, err := stmt.Exec(a.Id, a.Name, a.Year, a.Origin)
	if err != nil {
		return err
	}

	i, err := r.RowsAffected()
	if i != 1 {
		return errors.New("Unable to INSERT Artist")
	}

	return nil
}

func RetrieveArtists() (albums []Artist, err error) {
	query := `SELECT * FROM artist`

	db := getConnection()
	defer db.Close()

	rows, err := db.Query(query)
	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		a := Artist{}
		err = rows.Scan(&a.Id, &a.Name, &a.Year, &a.Origin)
		if err != nil {
			return
		}
		albums = append(albums, a)
	}

	return albums, nil
}
