package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"errors"
	"strings"
	"sync/atomic"
)

type MultiMysqlOptions struct {
	OptionsType string        `toml:"optionstype"json:"optionstype"` //readwrite,readonly,writeonly
	Options     *MysqlOptions `toml:"options"json:"options"`
}

type MysqlOptions struct {
	Username     string `toml:"username"json:"username"`
	Password     string `toml:"password"json:"password"`
	Host         string `toml:"host"json:"host"`
	Port         int64  `toml:"port"json:"port"`
	Name         string `toml:"name"json:"name"`
	MaxOpenConns int64  `toml:"maxopenconns"json:"maxopenconns"`
	MaxIdleConns int64  `toml:"maxidleconns"json:"maxidleconns"`
}

type MultiMysqlClient struct {
	readClientList  []*MysqlClient
	writeClientList []*MysqlClient
	bothClientList  []*MysqlClient
	readIndex       uint64
	writeIndex      uint64
	bothIndex       uint64
	readLen         uint64
	writeLen        uint64
	bothLen         uint64
}

type MysqlClient struct {
	Db *sql.DB
}

func NewMultiMsqlClient(optList []*MultiMysqlOptions) *MultiMysqlClient {
	bothClientList := make([]*MysqlClient, 0)
	readClientList := make([]*MysqlClient, 0)
	writeClientList := make([]*MysqlClient, 0)

	for i := range (optList) {
		client, err := NewMysqlClient(optList[i].Options)
		if err == nil {
			if optList[i].OptionsType == "readwrite" {
				bothClientList = append(bothClientList, client)
				readClientList = append(readClientList, client)
				writeClientList = append(writeClientList, client)
			} else if optList[i].OptionsType == "readonly" {
				readClientList = append(readClientList, client)
			} else if optList[i].OptionsType == "writeonly" {
				writeClientList = append(writeClientList, client)
			}
		}
	}
	readLen := uint64(len(readClientList))
	writeLen := uint64(len(writeClientList))
	bothLen := uint64(len(bothClientList))
	return &MultiMysqlClient{
		readClientList:  readClientList,
		writeClientList: writeClientList,
		bothClientList:  writeClientList,
		readIndex:       uint64(0),
		writeIndex:      uint64(0),
		bothIndex:       uint64(0),
		readLen:         readLen,
		writeLen:        writeLen,
		bothLen:         bothLen,
	}
}

func (m *MultiMysqlClient) GetBoth() *MysqlClient {
	if m.bothLen == 0 {
		return nil
	}
	atomic.AddUint64(&m.bothIndex, ^uint64(0))
	bothIndex := atomic.LoadUint64(&m.bothIndex)
	index := bothIndex % m.bothLen
	client := m.bothClientList[index]
	return client
}

func (m *MultiMysqlClient) GetRead() *MysqlClient {
	if m.readLen == 0 {
		return nil
	}
	atomic.AddUint64(&m.readIndex, ^uint64(0))
	readIndex := atomic.LoadUint64(&m.readIndex)

	index := readIndex % m.readLen
	client := m.readClientList[index]
	return client
}

func (m *MultiMysqlClient) GetWrite() *MysqlClient {
	if m.writeLen == 0 {
		return nil
	}
	atomic.AddUint64(&m.writeIndex, ^uint64(0))
	writeIndex := atomic.LoadUint64(&m.writeIndex)

	index := writeIndex % m.writeLen
	client := m.writeClientList[index]
	return client
}

func NewMysqlClient(options *MysqlOptions) (mysqlClient *MysqlClient, err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		options.Username,
		options.Password,
		options.Host,
		options.Port,
		options.Name)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return
	}

	var maxOpenConns, maxIdleConns int
	if options.MaxOpenConns == 0 {
		maxOpenConns = 20
	} else {
		maxOpenConns = int(options.MaxOpenConns)
	}

	if options.MaxIdleConns == 0 {
		maxIdleConns = 20
	} else {
		maxIdleConns = int(options.MaxIdleConns)
	}

	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConns)
	mysqlClient = &MysqlClient{
		Db: db,
	}
	return
}

func (m *MysqlClient) QueryCount(format string, a ...interface{}) (n int64, err error) {
	sql, err := MysqlFilter(format, a...)
	if err != nil {
		return
	}
	rows, err := m.Db.Query(sql)
	if err != nil {
		return
	}
	rows.Next()
	rows.Scan(&n)
	rows.Close()
	return
}

func (m *MysqlClient) QueryCallMethon(format string, a ...interface{}) (n float64, err error) {
	sql, err := MysqlFilter(format, a...)
	if err != nil {
		return
	}
	rows, err := m.Db.Query(sql)
	if err != nil {
		return
	}
	rows.Next()
	rows.Scan(&n)
	rows.Close()
	return
}

func (m *MysqlClient) QueryRows(format string, a ...interface{}) (row *sql.Rows, err error) {
	sql, err := MysqlFilter(format, a...)
	if err != nil {
		return
	}
	row, err = m.Db.Query(sql)
	return
}

func (m *MysqlClient) QueryOne(format string, a ...interface{}) (data map[string]string, err error) {
	sql, err := MysqlFilter(format, a...)
	if err != nil {
		return
	}
	rows, err := m.Db.Query(sql)
	if err != nil {
		return
	}
	data, err = MysqlGetOneByRows(rows)
	rows.Close()
	return
}

func (m *MysqlClient) QueryGetRow(format string, a ...interface{}) (rows *sql.Rows, err error) {
	sql, err := MysqlFilter(format, a...)
	if err != nil {
		return
	}
	rows, err = m.Db.Query(sql)
	if err != nil {
		return
	}
	rows.Close()
	return
}

func (m *MysqlClient) Query(format string, a ...interface{}) (data []map[string]string, err error) {
	sql, err := MysqlFilter(format, a...)
	if err != nil {
		return
	}
	rows, err := m.Db.Query(sql)
	if err != nil {
		return
	}
	data, err = MysqlGetByRows(rows)
	rows.Close()
	return
}

func (m *MysqlClient) Exec(format string, a ...interface{}) (res sql.Result, err error) {
	sql, err := MysqlFilter(format, a...)
	if err != nil {
		return
	}
	res, err = m.Db.Exec(sql)
	return
}

func (m *MysqlClient) MysqlSetOneData(dest []interface{}, sql string) (err error) {
	rows, err := m.Db.Query(sql)
	if err != nil {
		return
	}

	if rows.Next() {
		err = rows.Scan(dest...)
	} else {
		return nil
	}
	return err
}

func MysqlGetOneByRows(rows *sql.Rows) (map[string]string, error) {
	if rows == nil {
		return nil, errors.New("rows is nil")
	}
	cols, err := rows.Columns()

	rawResult := make([][]byte, len(cols))
	result := make(map[string]string, len(cols))
	dest := make([]interface{}, len(cols))
	for i, _ := range rawResult {
		dest[i] = &rawResult[i]
	}
	if rows.Next() {
		err = rows.Scan(dest...)
		for i, raw := range rawResult {
			if raw == nil {
				result[cols[i]] = ""
			} else {
				result[cols[i]] = string(raw)
			}
		}
	} else {
		return nil, err
	}
	return result, err
}

func MysqlGetByRows(rows *sql.Rows) ([]map[string]string, error) {
	if rows == nil {
		return make([]map[string]string, 0), errors.New("rows is nil")
	}
	cols, err := rows.Columns()
	if err != nil {
		return make([]map[string]string, 0), err
	}
	lennum := len(cols)
	result := make([]map[string]string, 0)

	for rows.Next() {
		dest := make([]interface{}, lennum)
		rawResult := make([][]byte, lennum)
		result2 := make(map[string]string, lennum)
		for i, _ := range rawResult {
			dest[i] = &rawResult[i]
		}
		err = rows.Scan(dest...)
		for i, raw := range rawResult {
			if raw == nil {
				result2[cols[i]] = ""
			} else {
				result2[cols[i]] = string(raw)
			}
		}
		result = append(result, result2)
	}
	return result, err
}

func MysqlFilter(format string, a ...interface{}) (str string, err error) {
	for i := range a {
		switch ai := a[i].(type) {
		default:
			err = errors.New("不能处理这种类型")
			return
		case int64:
			a[i] = int(ai)
		case int32:
			a[i] = int(ai)
		case int:
			a[i] = int(ai)
		case []byte:
			str := strings.Replace(string(a[i].([]byte)), `\"`, `"`, -1)
			str = strings.Replace(str, `\'`, `'`, -1)
			str = strings.Replace(str, `"`, `\"`, -1)
			str = strings.Replace(str, `'`, `\'`, -1)
			a[i] = str
		case string:
			str := strings.Replace(a[i].(string), `\"`, `"`, -1)
			str = strings.Replace(str, `\'`, `'`, -1)
			str = strings.Replace(str, `"`, `\"`, -1)
			str = strings.Replace(str, `'`, `\'`, -1)
			a[i] = str
		}
	}
	str = fmt.Sprintf(format, a...)
	return
}
