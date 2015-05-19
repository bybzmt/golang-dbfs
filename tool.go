package dbfs

import (
	"path"
	"time"
)

func fs_find(db *sql.DB, fi *RowNode, name string) (*RowNode, error) {
	//读cache

	//读db
}

func fs_create(fs *filesystem, fi *RowNode, name string, isdir bool) (*RowNode, error) {
	//db_create
	ctime := time.Now().Unix()
	size := isdir ? -1 : 0

	insertid, err := db_create(fs.Slave, fi, name, isdir, ctime)
	if err != nil {
		return nil, err
	}

	//成功清缓存
	r := &RowNode{
		id : insertid
		pid : fi.id
		szie : size
		mtime : ctime
		name : name
	}

	fs.DirNodeCache.Add(r.id, r)

	return r
}

func fs_remove(fs *filesystem, fi *RowNode) error {
	affected, err := db_remove(fs.Master, fi)
	if err != nil {
		return err
	}

	if affected < 1 {
		return ErrNotExist
	}

	//文件操作
	if !fi.isDir() {
		files, ok := fs.FileCache.Get(fi.id)
		if !ok {
			files, err = db_findFile(fs.Slave, fi)
			if err != nil {
				return err
			}
		}

		for _, file := range files {
			//删除真实文件
		}

		fs.FileCache.Remove(fi.id)
	}

	//请理缓存

	return nil
}
