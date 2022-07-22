package db

import (
	"errors"
	"github.com/azazel3ooo/yandextask/internal/models"
	"github.com/google/uuid"
	"log"
)

// Get - возвращает оригинальную ссылку по ее id
func (d *Database) Get(key string) (string, error) {
	stmt := `select url, deleted from Urls where id=$1`
	rows, err := d.Conn.Query(stmt, key)
	if err != nil {
		return "", err
	}
	if rows.Err() != nil {
		return "", rows.Err()
	}
	defer rows.Close()

	var (
		url string
		del bool
	)
	for rows.Next() {
		err = rows.Scan(&url, &del)
		if err != nil {
			return "", err
		}
	}
	if del {
		return "deleted", nil
	}

	return url, nil
}

// Set - возвращает id для нвовой короткой ссылки
func (d *Database) Set(val, pth string) (string, error) {
	id := uuid.New()

	stmt := `select id from Urls where url=$1`
	rows, err := d.Conn.Query(stmt, val)
	if err != nil {
		return "", err
	}
	if rows.Err() != nil {
		log.Println(rows.Err())
	}

	var i string

	if rows.Next() {
		err = rows.Scan(&i)
		if err != nil {
			return "", err
		}
		if rows.Err() != nil {
			log.Println(rows.Err())
		}
		return i, errors.New("conflict")
	}
	defer rows.Close()

	stmt = `insert into Urls(id,url,deleted) values($1,$2,$3)`
	rows, err = d.Conn.Query(stmt, id.String(), val, false)
	if err != nil {
		return "", err
	}
	if rows.Err() != nil {
		log.Println(rows.Err())
	}
	return id.String(), nil
}

// InsertMany - добавляет несколько новых ссылок с кастомными id
func (d *Database) InsertMany(m []models.CustomIDSet) ([]models.CustomIDSet, error) {
	var res []models.CustomIDSet

	for _, el := range m {
		stmt := `select id from Urls where url=$1`
		rows, _ := d.Conn.Query(stmt, el.OriginalURL)
		if rows.Err() != nil {
			log.Println(rows.Err())
		}

		var i string
		err := rows.Scan(&i)
		if err != nil {
			stmt := `insert into Urls(id,url,deleted) values($1,$2,$3)`
			row, err := d.Conn.Query(stmt, el.CorrelationID, el.OriginalURL, false)
			if err != nil {
				continue
			}
			if row.Err() != nil {
				log.Println(row.Err())
			}
			res = append(res, models.CustomIDSet{CorrelationID: el.CorrelationID, ShortURL: el.CorrelationID})
			continue
		}
		defer rows.Close()

		res = append(res, models.CustomIDSet{CorrelationID: el.CorrelationID, ShortURL: i})
	}
	return res, nil
}

// Delete - "удаляет" ссылку, добавляя пометку об удалении
func (d *Database) Delete(ids []string) error {
	stmt := `update Urls SET deleted=true WHERE id=$1`
	for _, id := range ids {
		_, err := d.Conn.Exec(stmt, id)
		if err != nil {
			log.Println(err)
		}
	}

	return nil
}

func (d *Database) GetUrlsStat() (int, error) {
	stmt := "SELECT COUNT(*) FROM Urls"
	row, err := d.Conn.Query(stmt)
	if err != nil {
		return 0, err
	}
	defer row.Close()

	var size int
	if row.Err() == nil && row.Next() {
		err = row.Scan(&size)
		if err != nil {
			return 0, err
		}
	}

	return size, row.Err()
}
