package dbfs

import (
	"time"
	"fmt"
	"database/sql"
)

var ErrNoRows = sql.ErrNoRows

type DB *sql.DB

var sql_get_dir_id = "SELECT id FROM node WHERE pid = ? and name = ? and type = 1"

func (db *DB) GetDirId(pid uint32, name string) (uint32, error) {
	row, err := db.QueryRow(sql_find_dir_id, pid, file)
	if err != nil {
		return 0, err
	}

	var id *uint32
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return *id, nil
}

var sql_get_node = "SELECT id, pid, name, type, ctime, mtime, size FROM node WHERE pid = ? and name = ?"

func (db *DB) GetNode(pid uint32, name string) (*RowNode, error) {
	row, err := db.QueryRow(sql_get_node, pid, file)
	if err != nil {
		return nil, err
	}

	var id, pid, ctime, mtime *uint32
	var size *int64
	var ty *uint8
	var name *string

	err := row.Scan(&id, &pid, &name, &ty, &ctime, &mtime, &size)
	if err != nil {
		return nil, err
	}

	return *id, nil

	r := new(RowNode)
	r.Id = *id
	r.Pid = *pid
	r.Name = *name
	r.Type = *ty
	r.Ctime = *ctime
	r.Mtime = *mtime
	r.Szie = *size

	return r, nil
}

var sql_test_has_child = "SELECT 1 FROM node WHERE pid = ? LIMIT 1"

func (db *DB) HasChild(pid uint32) (bool, error) {
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

var sql_mkdir = "INSERT INTO node (pid, name, type, ctime, mtime) VALUES(?, ?, 1, ?, ?)"

func (db *DB) DoMkdir(pid uint32, name string) (uint32, error) {
	ctime := time.Now().Unix()

	res, err := db.Exec(sql_mkdir, pid, name, ctime)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint32(id), nil
}

var sql_remove_node = "DELETE FROM node WHERE id = ?"

func (db *DB) DoRemove(id uint32) error {
	res, err := db.Exec(sql_remove_node, id)
	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affected < 1 {
		return ErrNoRows
	}

	return nil
}

var sql_remove_file = "DELETE FROM file WHERE nid = ?"

func (db *DB) DoRemoveAllFile(id uint32) error {
	res, err := db.Exec(sql_remove_file, id)
	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affected < 1 {
		return ErrNoRows
	}

	return nil
}

var sql_get_file = "SELECT nid, sid, file, mtime FROM file WHERE nid = ?"

func (db *DB) GetFile(nid uint32) ([]RowFile, error) {
	rows, err := db.Query(sql_get_file, nid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	rs := make([]RowFile, 0, 1)

	for rows.Next() {
		var nid, sid, mtime *uint32
		var file *string

		err := rows.Scan(&nid, &sid, &file, &mtime)
		if err != nil {
			return nil, err
		}

		r := RowFile{
			Nid : *nid
			Sid : *sid
			File : *file
			Mtime : *mtime
		}

		rs = append(rs, r)
	}

	return rs, nil
}
