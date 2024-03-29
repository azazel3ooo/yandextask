package db

import (
	"fmt"
	"log"
	"strings"

	"github.com/azazel3ooo/yandextask/internal/models"
)

// UsersSet - добавляет пользователю новую ссылку
func (d *Database) UsersSet(id, url string) error {
	urls, err := d.UsersGet(id)
	if err != nil {
		return err
	}

	urls = append(urls, url)

	stmt := `insert into Users(id, urls) values($1,$2) on conflict (id) do update set urls=$2`
	_, err = d.Conn.Exec(stmt, id, strings.Join(urls, ","))
	if err != nil {
		return err
	}
	return nil
}

// UsersGet - вовзращает ссылки по id пользователя
func (d *Database) UsersGet(id string) ([]string, error) {
	stmt := `select urls from Users where id=$1`
	row, err := d.Conn.Query(stmt, id)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	if row.Err() != nil {
		log.Println(row.Err())
	}

	var s string
	for row.Next() {
		err = row.Scan(&s)
		if err != nil {
			log.Println(err)
			return nil, err
		}
	}
	return strings.Split(s, ","), nil
}

// GetUrlsForUser - возвращает все ссылки по массиву id
func (d *Database) GetUrlsForUser(ids []string) ([]models.UserResponse, error) {
	for idx, id := range ids {
		ids[idx] = fmt.Sprintf("'%s'", id)
	}

	stmt := `select * from Urls where id IN ($1)`
	rows, err := d.Conn.Query(stmt, strings.Join(ids, ","))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var (
		u   models.UserResponse
		res []models.UserResponse
		del bool
	)
	for rows.Next() {
		err = rows.Scan(&u.Short, &u.Original, &del)
		if err != nil {
			log.Println(err)
			continue
		}
		if rows.Err() != nil {
			log.Println(rows.Err())
		}

		if !del {
			res = append(res, u)
		}
	}

	return res, nil
}

func (d *Database) GetUserStat() (int, error) {
	stmt := "SELECT COUNT(*) FROM Users"
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
