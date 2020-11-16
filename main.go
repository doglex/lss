package main

import (
	"fmt"
	"github.com/dustin/go-humanize"
	"github.com/jedib0t/go-pretty/table"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"
)

var Folder string = ""
var T = table.NewWriter()
var mutex sync.Mutex
var totalNum int64 = 0
var totalSize int64 = 0

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

type FileSizeStruct struct {
	File     os.FileInfo
	FileType string
	Size     int64
	Num      int64
	StrPath  string
}

func GetSize(idx int, file os.FileInfo, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	fileType := "file"
	if file.IsDir() {
		fileType = "directory"
	}
	strAbsPath, errPath := filepath.Abs(Folder + "/" + file.Name())
	if errPath != nil {
		fmt.Println(errPath)
		return
	}
	size, num, err := DirSize(strAbsPath)
	if err != nil {
		fmt.Println("Warn", " cannot compute ", strAbsPath)
		return
	}
	f := FileSizeStruct{file, fileType, size, num, strAbsPath}
	mutex.Lock()
	totalNum += f.Num
	totalSize += f.Size

	T.AppendRow([]interface{}{idx, f.FileType, f.File.Name(), humanize.Bytes(uint64(f.Size)), f.Size, f.Num})
	mutex.Unlock()
}

func ListDir(folder string) {
	Folder = folder
	T.SetOutputMirror(os.Stdout)
	T.AppendHeader(table.Row{"#", "Type", "Name", "Size(human)", "Size(B)", "Items"})

	files, errDir := ioutil.ReadDir(Folder)
	if errDir != nil {
		log.Fatal(errDir)
	}
	var wg sync.WaitGroup
	for idx, file := range files {
		wg.Add(1)
		go GetSize(idx, file, &wg)
	}
	wg.Wait()
	T.SortBy([]table.SortBy{{Name: "Size(B)", Mode: table.DscNumeric}})
	T.AppendFooter(table.Row{"$", "", "Total", humanize.Bytes(uint64(totalSize)), totalSize, totalNum})
	T.SetStyle(table.StyleColoredBlackOnMagentaWhite)
	T.Render()
}

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalln("Error", "cannot get cwd")
	}
	fmt.Println("CWD:", cwd)
	ListDir(cwd)
}
