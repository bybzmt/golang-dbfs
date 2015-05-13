package netfs

import (
	"time"
	"os"
	"log"
	"github.com/bybzmt/golang-filelog"
)

const VERSION uint32 = 1

const (
	TYTE_INIT uint8 = iota + 1
	TYPE_REQUEST
	TYPE_RESPONSE
)

const (
	//连接管理
	LINK_CLOSE uint8 = iota + 1

	//ping
	LINK_PING
)

const (
	//文件系统操作码
	FS_CHMOD uint8 = iota + 20
	FS_CHTIMES
	FS_MKDIR
	FS_MKDIRALL
	FS_REMOVE
	FS_REMOVEALL
	FS_RENAME
	FS_TRUNCATE
	FS_CREATE
	FS_OPEN
	FS_OPENFILE
	FS_LSTAT
	FS_STAT
)

const (
	//文件对象操作码
	FILE_CHMOD uint8 = iota + 60
	FILE_CLOSE
	FILE_READ
	FILE_READAT
	FILE_READDIR
	FILE_READDIRNAMES
	FILE_SEEK
	FILE_STAT
	FILE_SYNC
	FILE_TRUNCATE
	FILE_WRITE
	FILE_WRITEAT
)

const (
	//特殊状态码
	ERROR_MAX uint16 = 0xFF00
	ERROR_NIL uint16 = 0xFF01
	ERROR_EOF uint16 = 0xFF02
)

type IO_Error string
func (e IO_Error) Error() string {
	return string(e)
}

type Data_Error string
func (e Data_Error) Error() string {
	return string(e)
}

type FileSystem interface {
	Chmod(name string, mode os.FileMode) error
	Chtimes(name string, atime time.Time, mtime time.Time) error
	Mkdir(name string, perm os.FileMode) error
	MkdirAll(path string, perm os.FileMode) error
	Remove(name string) error
	RemoveAll(path string) error
	Rename(oldpath, newpath string) error
	Truncate(name string, size int64) error
	Create(name string) (file File, err error)
	Open(name string) (file File, err error)
	OpenFile(name string, flag int, perm os.FileMode) (file File, err error)
	Stat(name string) (fi os.FileInfo, err error)
	Lstat(name string) (fi os.FileInfo, err error)
}

type File interface {
	Chmod(mode os.FileMode) error
	Close() error
	Name() string
	Read(b []byte) (n int, err error)
	ReadAt(b []byte, off int64) (n int, err error)
	Readdir(n int) (fi []os.FileInfo, err error)
	Readdirnames(n int) (names []string, err error)
	Seek(offset int64, whence int) (ret int64, err error)
	Stat() (fi os.FileInfo, err error)
	Sync() (err error)
	Truncate(size int64) error
	Write(b []byte) (n int, err error)
	WriteAt(b []byte, off int64) (n int, err error)
	WriteString(s string) (ret int, err error) //非交互
}

var Logger flog.Writer

func getLog() flog.Writer {
	if Logger == nil {
		var err error
		Logger, err = flog.New("", "LOCAL0:NOTICE", "netfs")
		if err != nil {
			log.Panicln(err)
		}
	}
	return Logger
}

func onPanic(err *error) {
	if x := recover(); x != nil {
		switch v := x.(type) {
		case IO_Error :
			*err = v
		default:
			panic(x)
		}
	}
}

