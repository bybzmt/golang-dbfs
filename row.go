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

type RowFile struct {
	sid int64
	path string
}

type RowStorage struct {
	Id uint32
	Host string
	Port uint32
	Path string
	Status uint8
}
