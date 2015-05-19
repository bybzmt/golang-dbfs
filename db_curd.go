package dbfs

import (
	"database/sql"
	"time"
)

var sql_find_node = "SELECT id, pid, name, type, mtime, size FROM node WHERE pid = ? and name = ?"

var sql_exist_child = "SELECT 1 FROM node WHERE pid = ? LIMIT 1"

var sql_remove_node = "DELETE FROM node WHERE id = ?"

var sql_mkdir = "INSERT INTO node (pid, name, size, ctime, mtime) VALUES(?, ?, 1, ?, ?)"

var sql_move = "UPDATE node SET pid = ?,  name = ? WHERE id = ?"

var sql_resize = "UPDATE node SET size = ? WHERE id = ?"



func db_find(db *sql.DB, fi *RowNode, name string) (*RowNode, error) {
	row, err := db.QueryRow(sql_find_node, fi.pid, name)
	if err != nil {
		return nil, err
	}

	var id, pid, mtime, size *int64
	var name *string

	err := row.Scan(&id, &pid, &name, &mtime, &size)
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

func db_findFile(db *sql.DB, fi *RowNode) ([]RowFile, error) {
}

func db_hasChild(db *sql.DB, fi *RowNode) (bool, error) {
	row, err := db.QueryRow(sql_test_has_child, fi.id)
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

func db_create(db *sql.DB, fi *RowNode, name string, size int64, ctime time.Time) (insertid int64, error) {
	res, err := db.Exec(sql_mkdir, fi.id, name, size, ctime, ctime)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

func db_move(db *sql.DB, fi *RowNode, pid int64, name string) (affected int64, error) {
	res, err := db.Exec(sql_move, pid, name, fi.id)
	if err != nil {
		return err
	}

	return res.RowsAffected()
}

func db_reszie(db *sql.DB, fi *RowNode, size int64) (affected int64, error) {
	res, err := db.Exec(sql_resize, size, fi.id)
	if err != nil {
		return err
	}

	return res.RowsAffected()
}

func db_remove(db *sql.DB, fi *RowNode) (affected int64, error) {
	res, err := db.Exec(sql_remove_node, fi.id)
	if err != nil {
		return err
	}

	return res.RowsAffected()
}



