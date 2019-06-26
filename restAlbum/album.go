package main

import "errors"

type Album struct {
	Id       int
	ArtistId int
	Name     string
	Year     int
	Price    int
}

func CreateAlbum(a Album) error {

	query := `INSERT INTO album (id, id_artist, name, YEAR, PRICE) VALUES ($1, $2, $3, $4, $5)`

	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	r, err := stmt.Exec(a.Id, a.ArtistId, a.Name, a.Year, a.Price)
	if err != nil {
		return err
	}

	i, err := r.RowsAffected()
	if i != 1 {
		return errors.New("Unable to INSERT Album")
	}

	return nil
}

func RetrieveAlbums() (albums []Album, err error) {
	query := `SELECT * FROM album`

	db := getConnection()
	defer db.Close()

	rows, err := db.Query(query)
	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		a := Album{}
		err = rows.Scan(&a.Id, &a.ArtistId, &a.Name, &a.Year, &a.Price)
		if err != nil {
			return
		}
		albums = append(albums, a)
	}

	return albums, nil
}

func RetrieveAlbumsByArtist(id string) (albums []Album, err error) {

	query := `SELECT * from ALBUM where id_Artist = $1`

	db := getConnection()
	defer db.Close()

	rows, err := db.Query(query, id)
	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		a := Album{}
		err = rows.Scan(&a.Id, &a.ArtistId, &a.Name, &a.Year, &a.Price)
		if err != nil {
			return
		}
		albums = append(albums, a)
	}

	return albums, nil
}
