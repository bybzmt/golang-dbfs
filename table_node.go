package dbfs

import (
	"database/sql"
	"time"
)

type RowNode struct {
	id int64
	pid int64
	size int64
	mtime int64
	name string
}

func (fi *RowNode) Name() string {
	return fi.name
}

func (fi *RowNode) Size() int64 {
	return fi.size
}

func (fi *RowNode) Mode() os.FileMode {
	if fi.IsDir() {
		return os.ModeDir | 0777
	}
	return 0777
}

func (fi *RowNode) ModTime() time.Time {
	return time.Unix(fi.mtime, 0)
}

func (fi *RowNode) IsDir() bool {
	return fi.size == -1
}

func (fi *RowNode) Sys() interface{} {
	return nil
}


var sql_find_node = "SELECT id, pid, name, type, mtime, size FROM node WHERE pid = ? and name = ?"

var sql_exist_child = "SELECT 1 FROM node WHERE pid = ? LIMIT 1"

type table_node struct {
	Master *sql.DB
	Salve *sql.DB
}

func (t *table_node) Find(pid uint32, name string) (*RowNode, error) {
	row, err := db.QueryRow(sql_find_node, pid, file)
	if err != nil {
		return nil, err
	}

	var id, pid, mtime, size *int64
	var name *string

	err := row.Scan(&id, &pid, &name, &ty, &ctime, &mtime, &size)
	if err != nil {
		return nil, err
	}

	r := &RowNode{
		id : *id
		pid : *pid
		szie : *size
		mtime : *mtime
		name : *name
	}

	return r, nil
}

func (t *table_node) HasChild(pid uint32, name string, ty uint8) (uint32, error) {
	row, err := db.QueryRow(sql_test_has_child, pid)
	if err != nil {
		return false, err
	}

	err := row.Scan(&id)
	if err != nil {
		if err == ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
func (t *table_node) Insert(pid uint32, name string, ty uint8) (uint32, error) {
}

func (t *table_node) Update(id, pid uint32, name string) error {
}

func (t *table_node) UpdateTime(id, mtime uint32) error {
}

func (t *table_node) UpdateSize(id uint32, size int64) error {
}
