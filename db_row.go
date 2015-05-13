package dbfs


type RowFile struct {
	Nid uint32
	Sid uint16
	File string
	Mtime uint32
}

type RowStorage struct {
	Id uint32
	Host string
	Port uint32
	Path string
	Status uint8
}
