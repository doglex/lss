package main

import (
	"fmt"
	"github.com/dustin/go-humanize"
	"github.com/jedib0t/go-pretty/table"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func DirSize(path string) (size int64, num int64, err error) {
	err = filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
			num++
		}
		return err
	})
	return
}

func ListDir(folder string, noPath bool) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Type", "Name", "Size(human)", "Size(B)", "Items", "Path"})

	files, errDir := ioutil.ReadDir(folder)
	if errDir != nil {
		log.Fatal(errDir)
	}
	var totalNum int64 = 0
	var totalSize int64 = 0
	for idx, file := range files {
		fileType := "file"
		if file.IsDir() {
			fileType = "directory"
		}
		strAbsPath, errPath := filepath.Abs(folder + "/" + file.Name())
		if errPath != nil {
			fmt.Println(errPath)
		}
		size, num, err := DirSize(strAbsPath)
		if err != nil {
			fmt.Println("Warn", " cannot compute ", strAbsPath)
			continue
		}
		totalNum += num
		totalSize += size
		//fmt.Println(size, num, err)
		//fmt.Println(file.Name(),file.Size(), strAbsPath,)
		if noPath {
			t.AppendRow([]interface{}{idx, fileType, file.Name(), humanize.Bytes(uint64(size)), size, num})
		} else {
			t.AppendRow([]interface{}{idx, fileType, file.Name(), humanize.Bytes(uint64(size)), size, num, strAbsPath})
		}
	}
	t.SortBy([]table.SortBy{{Name: "Size(B)", Mode: table.DscNumeric}})
	t.AppendFooter(table.Row{"$", "", "Total", humanize.Bytes(uint64(totalSize)), totalSize, totalNum})
	t.SetStyle(table.StyleColoredBlackOnMagentaWhite)
	t.Render()
}

func main() {
	noPath := false
	if len(os.Args) > 1 && os.Args[1] == "-s" { // provide noPath
		fmt.Println(len(os.Args))
		noPath = true
	}
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalln("Error", "cannot get cwd")
	}
	fmt.Println("CWD:", cwd)
	ListDir(cwd, noPath)
}
