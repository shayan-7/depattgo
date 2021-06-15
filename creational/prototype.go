package creational

type AccessLevel int

const (
	Root AccessLevel = iota
	Guest
)

type inode interface {
	Clone() inode
	GetReadOnly() bool
	SetReadOnly(bool)
	GetOwner() AccessLevel
	SetOwner(AccessLevel)
}

type File struct {
	name     string
	readonly bool
	owner    AccessLevel
}

func NewFile(name string, readonly bool, owner AccessLevel) *File {
	return &File{name, readonly, owner}
}

func (f *File) GetReadOnly() bool {
	return f.readonly
}

func (f *File) SetReadOnly(v bool) {
	f.readonly = v
}

func (f *File) GetOwner() AccessLevel {
	return f.owner
}

func (f *File) SetOwner(owner AccessLevel) {
	f.owner = owner
}

func (f *File) Clone() inode {
	return &File{
		name:     "clone_" + f.name,
		readonly: f.readonly,
		owner:    f.owner,
	}
}

type Directory struct {
	name     string
	readonly bool
	owner    AccessLevel
	children []inode
}

func NewDirectory(n string, r bool, o AccessLevel, c ...inode) *Directory {
	return &Directory{name: n, readonly: r, owner: o, children: c}
}

func (d *Directory) Clone() inode {
	cloned := &Directory{
		name:     "clone_" + d.name,
		readonly: d.readonly,
		owner:    d.owner,
	}
	tempChildren := make([]inode, 0)
	for _, i := range d.children {
		copy := i.Clone()
		tempChildren = append(tempChildren, copy)
	}
	cloned.children = tempChildren
	return cloned
}

func (d *Directory) GetReadOnly() bool {
	return d.readonly
}

func (d *Directory) SetReadOnly(v bool) {
	d.readonly = v
}

func (d *Directory) GetOwner() AccessLevel {
	return d.owner
}

func (d *Directory) SetOwner(owner AccessLevel) {
	d.owner = owner
}
