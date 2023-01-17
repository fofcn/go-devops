package mysql

import (
	"database/sql"
	"strings"
	"taskmanager/zlog"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type MySQLConfig struct {
	IP         string
	Port       string
	Username   string
	Password   string
	DBName     string
	PoolConfig MySQLPoolConfig
}

type MySQLPoolConfig struct {
	MaxConn         int
	MaxIdleConn     int
	ConnMaxIdleTime int
	Ping            string
}

type mysqlclient struct {
	config MySQLConfig
	db     *sql.DB
}

type MySQLClient interface {
	Init(config MySQLConfig) error
	Exec(query string, args ...interface{}) (int64, error)
	Add(query string, args ...interface{}) (int64, error)
	Row(query string, args ...interface{}) (map[string]string, error)
	All(query string, args ...interface{}) ([]map[string]string, error)
	Shutdown()
}

func NewClient() MySQLClient {
	return &mysqlclient{
		config: MySQLConfig{
			IP:       "127.0.0.1",
			Port:     "3306",
			Username: "mysql",
			Password: "mysqlpass",
			DBName:   "mysql",
			PoolConfig: MySQLPoolConfig{
				MaxConn:         2,
				MaxIdleConn:     1,
				ConnMaxIdleTime: 600000,
				Ping:            "SELECT 1 from dual",
			},
		},
		db: nil,
	}
}

// Init implements MySQLClient
func (client *mysqlclient) Init(config MySQLConfig) error {
	client.config = config
	path := strings.Join([]string{config.Username, ":",
		config.Password, "@tcp(", config.IP, ":", config.Port, ")/",
		config.DBName, "?charset=utf8"}, "")
	zlog.Logger.Debugf("gdbc url: %v", path)
	DB, err := sql.Open("mysql", path)
	if err != nil {
		return err
	}
	DB.SetMaxOpenConns(config.PoolConfig.MaxConn)
	DB.SetMaxIdleConns(config.PoolConfig.MaxIdleConn)
	DB.SetConnMaxIdleTime(time.Duration(config.PoolConfig.ConnMaxIdleTime))

	if err := DB.Ping(); err != nil {
		zlog.Logger.Errorf("Open database faled, %v", err)
		return err
	}

	client.db = DB
	return nil
}

// Exec implements MySQLClient
func (client *mysqlclient) Exec(query string, args ...interface{}) (int64, error) {
	stmt, err := client.db.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(args...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (client *mysqlclient) Add(query string, args ...interface{}) (int64, error) {
	stmt, err := client.db.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(args...)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}
func (client *mysqlclient) Row(query string, args ...interface{}) (map[string]string, error) {
	if !strings.Contains(strings.ToUpper(query), "LIMIT") {
		query += " LIMIT 1"
	}
	stmt, err := client.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	ret := make(map[string]string, len(scanArgs))

	for i := range values {
		scanArgs[i] = &values[i]
	}
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			break
		}
		var value string

		for i, col := range values {
			if col == nil {
				value = "" //or NULL
			} else {
				value = string(col)
			}
			ret[columns[i]] = value
		}
		break
	}
	return ret, err
}

func (client *mysqlclient) All(query string, args ...interface{}) ([]map[string]string, error) {
	stmt, err := client.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))

	ret := make([]map[string]string, 0)
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			break
		}
		var value string
		vmap := make(map[string]string, len(scanArgs))
		for i, col := range values {
			if col == nil {
				value = "" // or NULL
			} else {
				value = string(col)
			}
			vmap[columns[i]] = value
		}
		ret = append(ret, vmap)
	}
	return ret, err
}

// Shutdown implements MySQLClient
func (client *mysqlclient) Shutdown() {
	client.db.Close()
}
