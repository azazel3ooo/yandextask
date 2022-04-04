package models

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"log"
	"strings"
)

func (d *Database) Init(cfg Config) {
	//var connectionInfo string
	name := "postgres"
	//host := strings.Split(cfg.DatabaseDsn, ":")
	//if len(host) > 1 {
	//	connectionInfo = fmt.Sprintf("host=%s port=%s user=postgres password=%s dbname=myDB sslmode=disable", host[0], host[1], "Ne8GowT4_")
	//} else {
	//	connectionInfo = fmt.Sprintf("host=localhost port=%s user=postgres password=%s dbname=myDB sslmode=disable", host[0], "Ne8GowT4_")
	//}

	db, err := sql.Open(name, cfg.DatabaseDsn)
	if err != nil {
		log.Println(err)
	}

	_, err = db.Exec("CREATE TABLE Users (id varchar PRIMARY KEY, urls varchar NOT NULL);")
	if err != nil {
		log.Println(err)
	}
	_, err = db.Exec("CREATE TABLE Urls (id varchar PRIMARY KEY, url varchar unique NOT NULL);")
	if err != nil {
		log.Println(err)
	}

	err = db.Ping()
	if err != nil {
		log.Println(err, "from db.Init()")
	} else {
		log.Println("Connected")
	}

	d.Conn = db
}

func (d *Database) Ping() error {
	return d.Conn.Ping()
}

func (d *Database) Get(key string) (string, error) {
	stmt := `select url from Urls where id=$1`
	rows, err := d.Conn.Query(stmt, key)
	if err != nil {
		return "", err
	}
	if rows.Err() != nil {
		return "", rows.Err()
	}
	defer rows.Close()

	var url string
	for rows.Next() {
		err = rows.Scan(&url)
		if err != nil {
			return "", err
		}
	}

	return url, nil
}

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

	stmt = `insert into Urls(id,url) values($1,$2)`
	rows, err = d.Conn.Query(stmt, id.String(), val)
	if err != nil {
		return "", err
	}
	if rows.Err() != nil {
		log.Println(rows.Err())
	}
	return id.String(), nil
}

func (d *Database) UsersSet(id, url string) error {
	urls, err := d.UsersGet(id)
	if err == nil {
		urls = append(urls, url)
		stmt := `update Users set urls=$1 where id=$2`
		rows, err := d.Conn.Query(stmt, strings.Join(urls, ","), id)
		if err != nil {
			return err
		}
		if rows.Err() != nil {
			log.Println(rows.Err())
		}
	} else {
		stmt := `insert into Users(id,urls) values($1,$2)`
		rows, err := d.Conn.Query(stmt, id, url)
		if err != nil {
			return err
		}
		if rows.Err() != nil {
			log.Println(rows.Err())
		}
	}
	return nil
}

func (d *Database) UsersGet(id string) ([]string, error) {
	stmt := `select "urls" from "Users" where id=$1`
	row, err := d.Conn.Query(stmt, id)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	if row.Err() != nil {
		log.Println(row.Err())
	}

	var s string
	row.Scan(&s)
	return strings.Split(s, ","), nil
}

func (d *Database) GetUrlsForUser(ids []string) ([]UserResponse, error) {
	for idx, id := range ids {
		ids[idx] = fmt.Sprintf("'%s'", id)
	}

	stmt := `select * from "Urls" where id IN ($1)`
	rows, err := d.Conn.Query(stmt, strings.Join(ids, ","))
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
		if rows.Err() != nil {
			log.Println(rows.Err())
		}
		res = append(res, u)
	}

	return res, nil
}

func (d *Database) InsertMany(m []CustomIDSet) ([]CustomIDSet, error) {
	var res []CustomIDSet

	for _, el := range m {
		stmt := `select id from Urls where url=$1`
		rows, _ := d.Conn.Query(stmt, el.OriginalURL)
		if rows.Err() != nil {
			log.Println(rows.Err())
		}

		var i string
		err := rows.Scan(&i)
		if err != nil {
			stmt := `insert into Urls(id,url) values($1,$2)`
			rows, err := d.Conn.Query(stmt, el.CorrelationID, el.OriginalURL)
			if err != nil {
				continue
			}
			if rows.Err() != nil {
				log.Println(rows.Err())
			}
			res = append(res, CustomIDSet{CorrelationID: el.CorrelationID, ShortURL: el.CorrelationID})
			continue
		}
		defer rows.Close()

		res = append(res, CustomIDSet{CorrelationID: el.CorrelationID, ShortURL: i})
	}
	return res, nil
}
