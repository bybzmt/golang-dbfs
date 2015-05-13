package dbfs

import (
	"path"
)

var (
    ErrInvalid    = os.ErrInvalid
    ErrPermission = os.ErrPermission
    ErrExist      = os.ErrExist
    ErrNotExist   = os.ErrNotExist
)

var ErrDirNotEmpty = errors.New("Dir Not Empty")

type filesystem struct {
}

func (s *filesystem) Chomd(name string, mode os.FileMode) error {
	return nil
}

func (s *filesystem) Chtimes(name string, atime time.Time, mtime time.Time) error {
	return nil
}

func (s *filesystem) Mkdir(name string, perm os.FileMode) error {
	dir, file := path.Split(name)

	node, err := find_node(dir)
	if err != nil {
		return err
	}

	if node.IsDir() {
		_, err := table_node.insert(node.Id, name)
		return err
	}

	return ErrNotDir
}

func (s *filesystem) MkdirAll(name string, perm os.FileMode) error {
	dir, file := path.Split(name)

	node, err := find_node(dir)
	if err != nil {
		if err == ErrNoRows {
			e2 := s.MkdirAll(dir)
			if e2 != nil {
				return e2
			} else {
				node, err = find_node(dir)
				if err != nil {
					return err
				}
			}
		} else {
			return err
		}
	}

	if node.IsDir() {
		_, err := table_node.insert(node.Id, name)
		return err
	}

	return ErrNotDir
}

func (s *filesystem) Remove(name string) error {
	node, err := find_node(name)
	if err != nil {
		return err
	}

	if node.IsDir() {
		if find_has_child(node.Id) {
			return ErrDirNotEmpty
		}
	}

	//文件操作

	return table_node.remove(node.Id)
}

func (s *filesystem) RemoveAll(name string) error {
	node, err := find_node(name)
	if err != nil {
		return err
	}

	if node.IsDir() {
		childs := find_childs(node.Id)
		for _, child := range childs {
			err := s.RemoveAll(child)
		}
	}

	//文件操作

	return table_node.remove(node.Id)
}

func (s *filesystem) Rename(oldpath, newpath string) error {
	dir, file := path.Split(newpath)

	node, err := find_node(oldpath)
	if err != nil {
		return err
	}

	node2, err2 := find_node(dir)
	if err != nil {
		return err
	}

	_, err := table_node.update(node.Id, node2.Pid, file)
	return err
}

func (s *filesystem) Truncate(name string, size int64) error {
	node, err := find_node(oldpath)
	if err != nil {
		return err
	}

	//文件操作
}

func (s *filesystem) Create(name string) (file File, err error) {
	return s.OpenFile(name, O_RDWR|O_CREATE|O_TRUNC, 0666)
}
func (s *filesystem) Open(name string) (file File, err error) {
	return s.OpenFile(name, O_RDONLY, 0)
}

func (s *filesystem) OpenFile(name string, flag int, perm os.FileMode) (file File, err error) {
}

func (s *filesystem) Stat(name string) (fi os.FileInfo, err error) {
	node, err := find_node(name)
	if err != nil {
		if err == ErrNoRows {
			return nil, ErrNotExist
		}
		return nil, err
	}

	return node, nil
}

func (s *filesystem) Lstat(name string) (fi os.FileInfo, err error) {
	return s.Stat(name)
}
