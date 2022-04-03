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

	_, err = db.Exec("CREATE TABLE Users (id varchar(37) PRIMARY KEY, urls varchar NOT NULL);")
	if err != nil {
		log.Println(err)
	}
	_, err = db.Exec("CREATE TABLE Urls (id varchar(37) PRIMARY KEY, url varchar unique NOT NULL);")
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
	err = rows.Scan(&url)
	if err != nil {
		return "", err
	}

	return url, nil
}

func (d *Database) Set(val, pth string) (string, error) {
	id := uuid.New()

	stmt := `insert into Urls(id,url) values($1,$2) ON CONFLICT DO NOTHING`
	res, err := d.Conn.Query(stmt, id.String(), val)
	if err.(*pq.Error).Code == pgerrcode.UniqueViolation {
		stmt = `select id from Urls where url=$1`
		rows, err := d.Conn.Query(stmt, val)
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
		log.Println(err, res.Err())
		return "", err
	}

	return id.String(), nil
}

func (d *Database) UsersSet(id, url string) error {
	urls, err := d.UsersGet(id)
	if err == nil {
		urls = append(urls, url)
		stmt := `update "Users" set "urls"=$1 where "id"=$2`
		_, err = d.Conn.Exec(stmt, strings.Join(urls, ","), id)
		if err != nil {
			return err
		}
	} else {
		stmt := `insert into "Users"("id","urls") values($1,$2)`
		_, err = d.Conn.Exec(stmt, id, url)
		if err != nil {
			return err
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
		res = append(res, u)
	}

	return res, nil
}

func (d *Database) InsertMany(m []CustomIDSet) ([]CustomIDSet, error) {
	var res []CustomIDSet

	for _, el := range m {
		stmt := `insert into "Urls"("id","url") values($1,$2) ON CONFLICT DO NOTHING`
		_, err := d.Conn.Exec(stmt, el.CorrelationID, el.OriginalURL)
		if err.(*pq.Error).Code == pgerrcode.UniqueViolation {
			stmt = `select "id" from "Urls" where url=$1`
			rows, err := d.Conn.Query(stmt, el.OriginalURL)
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
