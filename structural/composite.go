package structural

type Node interface {
	Size() int
	Search(string) string
}

type File struct {
	name string
	size int
}

func NewFile(n string, s int) *File {
	return &File{n, s}
}

func (f *File) Size() int {
	return f.size
}

func (f *File) Search(keyword string) string {
	if f.name == keyword {
		return keyword
	}
	return ""
}

type Directory struct {
	name     string
	children []Node
}

func NewDirectory(name string, children ...Node) *Directory {
	return &Directory{name, children}
}

func (d *Directory) Size() int {
	var sum int
	for _, ch := range d.children {
		sum += ch.Size()
	}
	return sum + 10 // Assume the directory size is 10 KB
}

func (d *Directory) Search(keyword string) string {
	for _, ch := range d.children {
		v := ch.Search(keyword)
		if v != "" {
			return d.name + "/" + v
		}
	}
	return ""
}
