package test

import (
	"fmt"
	"io/fs"
	"path"
	"path/filepath"
	"testing"
)

func TestBase(t *testing.T) {
	path := "./tt/f2/"

	// ok.txt
	fmt.Println(filepath.Base(path))

	// tt/f2/ok.txt
	fmt.Println(filepath.Clean(path))

	// ./tt/f2/ ok.txt
	fmt.Println(filepath.Split(path))

	// /Users/zhouzhenyong/project/go-private/go-learn/src/test/tt/f2/ok.txt <nil>
	fmt.Println(filepath.Abs(path))

	// tt/f2
	fmt.Println(filepath.Dir(path))

	// tt/f2/ok.txt <nil>
	fmt.Println(filepath.EvalSymlinks(path))

	// .txt
	fmt.Println(filepath.Ext(path))

	// ./tt/f2/ok.txt
	fmt.Println(filepath.FromSlash(path))

	// tt/f2/ok.txt/===
	fmt.Println(filepath.Join(path, "==="))

	// ../../../test.go <nil>
	fmt.Println(filepath.Rel(path, "test.go"))

	// [./tt/f2/ok.txt]
	fmt.Println(filepath.SplitList(path))

	// ./tt/f2/ok.txt
	fmt.Println(filepath.ToSlash(path))
	// 空
	fmt.Println(filepath.VolumeName(path))

	// 可以看到对应目录下面的所有的子文件夹
	filepath.Walk(path, func(path string, info fs.FileInfo, err error) error {
		fmt.Println(path)
		return err
	})

	// 可以看到对应目录下面的所有的子文件夹
	filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
		fmt.Println(path)
		return err
	})
}

func TestPathJson(t *testing.T) {
	// /test/rf.state
	fmt.Println(path.Join("/test", "rf.state"))
	// /test/rf.state
	fmt.Println(path.Join("/test/", "rf.state"))
	// /test/rf.state
	fmt.Println(path.Join("/test//", "rf.state"))
	// /test/rf.state
	fmt.Println(path.Join("/test///", "rf.state"))
}
