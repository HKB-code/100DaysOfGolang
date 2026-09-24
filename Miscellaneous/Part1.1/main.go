// Manual Loop Wrap Example (os.ReadDir like)

package main 

type DisEntry interface {
	Name()string
	IsDir()bool
}

type disentry struct{
	name string
	isDir bool
}

func (d *disentry)Name()string{
	return  d.name
}

func (d *disentry)IsDir()bool{
	return  d.isDir
}

func ReadMyDir(dirName string) []DisEntry{

	rawFiles := []disentry{
		{name: "main.go", isDir: false},
		{name: "utils", isDir: true},
		{name: "README.md", isDir: false},
		{name: "cmd", isDir: true},
	}

	var result []DisEntry
	for _,f := range rawFiles{
		entry := DisEntry(&f)
		result = append(result, entry)

	}
	return  result
}

