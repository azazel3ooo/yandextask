package models

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgerrcode"
	"github.com/lib/pq"
	"log"
	"strings"
)

func (d *Database) Init(cfg Config) {
	d.name = "postgres"
	host := strings.Split(cfg.DatabaseDsn, ":")
	if len(host) > 1 {
		d.connectionInfo = fmt.Sprintf("host=%s port=%s user=postgres password=%s dbname=myDB sslmode=disable", host[0], host[1], "Ne8GowT4_")
	} else {
		d.connectionInfo = fmt.Sprintf("host=localhost port=%s user=postgres password=%s dbname=myDB sslmode=disable", host[0], "Ne8GowT4_")
	}

	db, err := sql.Open(d.name, d.connectionInfo)
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	db.Exec("CREATE TABLE Users (id varchar(36) PRIMARY KEY, urls varchar NOT NULL);")
	db.Exec("CREATE TABLE Urls (id varchar(36) PRIMARY KEY, url varchar PRIMARY KEY);")
}

func (d *Database) Ping() error {
	db, err := sql.Open(d.name, d.connectionInfo)
	if err != nil {
		return err
	}
	defer db.Close()

	err = db.Ping()
	return err
}

func (d *Database) Get(key string) (string, error) {
	db, err := sql.Open(d.name, d.connectionInfo)
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	stmt := `select "url" from "Urls" where id=$1`
	rows, err := db.Query(stmt, key)
	if err != nil {
		return "", err
	}
	if rows.Err() != nil {
		return "", rows.Err()
	}
	defer rows.Close()

	var url string
	err = rows.Scan(&url)
	if err != nil {
		return "", err
	}

	return url, nil
}

func (d *Database) Set(val, pth string) (string, error) {
	id := uuid.New()

	db, err := sql.Open(d.name, d.connectionInfo)
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	stmt := `insert into "Urls"("id","url") values($1,$2) ON CONFLICT DO NOTHING`
	_, err = db.Exec(stmt, id.String(), val)
	if err.(*pq.Error).Code == pgerrcode.UniqueViolation {
		stmt = `select "id" from "Urls" where url=$1`
		rows, err := db.Query(stmt, val)
		if err != nil {
			return "", err
		}
		defer rows.Close()

		var i string
		err = rows.Scan(&i)
		if err != nil {
			return "", err
		}
		return i, errors.New("conflict")
	} else if err != nil {
		return "", err
	}

	return id.String(), nil
}

func (d *Database) UsersSet(id, url string) error {
	db, err := sql.Open(d.name, d.connectionInfo)
	if err != nil {
		return err
	}
	defer db.Close()

	urls, err := d.UsersGet(id)
	if err == nil {
		urls = append(urls, url)
		stmt := `update "Users" set "urls"=$1 where "id"=$2`
		_, err = db.Exec(stmt, strings.Join(urls, ","), id)
		if err != nil {
			return err
		}
	} else {
		stmt := `insert into "Users"("id","urls") values($1,$2)`
		_, err = db.Exec(stmt, id, url)
		if err != nil {
			return err
		}
	}
	return nil
}

func (d *Database) UsersGet(id string) ([]string, error) {
	db, err := sql.Open(d.name, d.connectionInfo)
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	stmt := `select "urls" from "Users" where id=$1`
	row, err := db.Query(stmt, id)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	var s string
	row.Scan(&s)
	return strings.Split(s, ","), nil
}

func (d *Database) GetUrlsForUser(ids []string) ([]UserResponse, error) {
	for idx, id := range ids {
		ids[idx] = fmt.Sprintf("'%s'", id)
	}

	db, err := sql.Open(d.name, d.connectionInfo)
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	stmt := `select * from "Urls" where id IN ($1)`
	rows, err := db.Query(stmt, strings.Join(ids, ","))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var (
		u   UserResponse
		res []UserResponse
	)
	for rows.Next() {
		err = rows.Scan(&u.Short, &u.Original)
		if err != nil {
			log.Println(err)
			continue
		}
		res = append(res, u)
	}

	return res, nil
}

func (d *Database) InsertMany(m []CustomIDSet) ([]CustomIDSet, error) {
	var res []CustomIDSet

	db, err := sql.Open(d.name, d.connectionInfo)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	for _, el := range m {
		stmt := `insert into "Urls"("id","url") values($1,$2) ON CONFLICT DO NOTHING`
		_, err = db.Exec(stmt, el.CorrelationID, el.OriginalURL)
		if err.(*pq.Error).Code == pgerrcode.UniqueViolation {
			stmt = `select "id" from "Urls" where url=$1`
			rows, err := db.Query(stmt, el.OriginalURL)
			if err != nil {
				log.Println(err)
				continue
			}
			defer rows.Close()

			var i string
			err = rows.Scan(&i)
			if err != nil {
				log.Println(err)
			}

			res = append(res, CustomIDSet{CorrelationID: el.CorrelationID, ShortURL: i})
			continue
		} else if err != nil {
			continue
		}
		res = append(res, CustomIDSet{CorrelationID: el.CorrelationID, ShortURL: el.CorrelationID})
	}
	return res, nil
}
