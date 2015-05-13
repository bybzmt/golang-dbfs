package dbfs

import (
	"path"
	"time"
)

func find_dir_id(name string) (uint32, error) {
	if name == "/" {
		return 0, nil
	}

	dir, file := path.Split(name)

	did, err := find_dir_id(dir)
	if err != nil {
		return 0, err
	}

	return getDB().get_dir_id(did, file)
}

func find_node(name string) (*RowNode, error) {
	if name == "/" {
		r := new(RowNode)
		r.Id = 0
		r.Name = "/"
		r.Type = 1
		return r, nil
	}

	dir, file := path.Split(name)

	did, err := find_dir_id(dir)
	if err != nil {
		return nil, err
	}

	return getDB().GetNode(did, file)
}

func do_mkdirAll(name) error {
	dir = strings.Trim(name, "/")
	ds := strings.Split(dir, "/")

	db := getDB()

	var pid uint32
	for _, na := range ds {
		id, err := db.GetDirId(pid, na)
		if err != nil {
			//文件夹不存在时创建
			if err == ErrNoRows {
				id, err = db.DoMkdir(pid, na)
				if err != nil {
					return err
				}
			} else {
				return err
			}
		}
		pid = id
	}

	return nil
}

func do_remove(name) error {
	db := getDB()

	node := find_node(name)
	if node.Type == 1 {
		yes, err := db.HasChild(node.Id)
		if err != nil {
			return err
		}
		if yes {
			return ErrDirNotEmpty
		}
	}

	files := db.GetFile(node.Id)
	for _, f := range files {
		err := storage_remove(f)
		if err != nil {
			getLog().Warning(fmt.Sprintf("storage_remove file: %v, err:%s", f, err))
		}
	}

	err := db.DoRemoveAllFile(node.Id)
	if err != nil {
		getLog().Warning(fmt.Sprintf("db_removeAllFile id: %d, err:%s", node.Id, err))
	}

	return db.DoRemove(node.Id)
}
