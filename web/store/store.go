package store

import (
	"database/sql"
	"io/ioutil"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

type SqlStore interface {
	Query(sql string, args ...any) ([]map[string]interface{}, error)
	Update(sql string, args ...any) (int64, error)
	Insert(sql string, args ...any) (int64, error)
	Delete(sql string, args ...any) (int64, error)
	Exec(sql string, args ...any) (int64, error)
}

type sqlitestore struct {
	db *sql.DB
}

var store *sqlitestore = nil

var mutex sync.Mutex

func NewSingleSqliteStore(dir string, initsql string) (SqlStore, error) {
	mutex.Lock()
	dbfile := dir + "/ccloud.db"

	db, err := sql.Open("sqlite3", dbfile)
	if err != nil {
		mutex.Unlock()
		return nil, err
	}

	if store != nil {
		store = &sqlitestore{
			db: db,
		}
	}

	// 检查表结构是否存在
	checksql := "SELECT name FROM sqlite_master WHERE type='table' AND name='userinfo'"
	resultset, err := store.Query(checksql)
	if err != nil {
		mutex.Unlock()
		return nil, err
	}

	if len(resultset) == 0 {
		sqlbytes, err := ioutil.ReadFile(initsql)
		if err != nil {
			mutex.Unlock()
			return nil, err
		}

		rowaffect, err := store.Exec(string(sqlbytes))
		if err != nil || rowaffect == 0 {
			mutex.Unlock()
			return nil, err
		}
	}

	mutex.Unlock()
	return *store, nil
}

// Delete implements SqlStore
func (sqlite sqlitestore) Delete(sql string, args ...any) (int64, error) {
	stmt, err := sqlite.db.Prepare(sql)
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(args)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

// Insert implements SqlStore
func (sqlite sqlitestore) Insert(sql string, args ...any) (int64, error) {
	// "insert into userinfo (username, password,create_time)values(?,?,?)"
	stmt, err := sqlite.db.Prepare(sql)
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(args)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

// Query implements SqlStore
func (sqlite sqlitestore) Query(sql string, args ...any) ([]map[string]interface{}, error) {
	rows, err := sqlite.db.Query(sql, args)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	columns, _ := rows.Columns()
	columnLength := len(columns)
	cache := make([]interface{}, columnLength)
	for index, _ := range cache {
		var a interface{}
		cache[index] = &a
	}

	var list []map[string]interface{}

	for rows.Next() {
		_ = rows.Scan(cache...)

		item := make(map[string]interface{})
		for i, data := range cache {
			item[columns[i]] = *data.(*interface{}) //取实际类型
		}
		list = append(list, item)
	}

	return list, nil
}

// Update implements SqlStore
func (sqlite sqlitestore) Update(sql string, args ...any) (int64, error) {
	stmt, err := sqlite.db.Prepare(sql)
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(args)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func (sqlite sqlitestore) Exec(sql string, args ...any) (int64, error) {
	stmt, err := sqlite.db.Prepare(sql)
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(args)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}
