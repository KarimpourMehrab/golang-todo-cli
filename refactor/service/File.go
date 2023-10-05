package service

import (
	"fmt"
	"os"
)

type File struct {
	path string
}

func (f *File) SetPath(path string) {
	f.path = "./files" + path
}
func (f *File) GetPath() string {
	return f.path
}

func (f *File) Get() []byte {
	file, err := os.OpenFile(f.path, os.O_RDONLY, os.ModeAppend)
	if err != nil {
		fmt.Println(err)
	}
	pData := make([]byte, 280)
	_, err = file.Read(pData)
	if err != nil {
		fmt.Println(err)
	}
	return pData
}

func (f *File) Store(data []byte) {
	var catFile *os.File
	var openError error
	var createError error

	catFile, openError = os.OpenFile(f.GetPath(), os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if openError != nil {
		catFile, createError = makeIfNotExists(openError, f.GetPath())
		if createError != nil {
			fmt.Println(createError)
		}
	}
	data = append(data, []byte("\n")...)
	_, wErr := catFile.Write(data)
	if wErr != nil {
		fmt.Println(wErr)
	}

	defer func() {
		cError := catFile.Close()
		if cError != nil {
			fmt.Println(cError)
		}
	}()
}

func makeIfNotExists(err error, path string) (*os.File, error) {
	if os.IsNotExist(err) {
		catFile, createError := os.Create(path)
		if createError != nil {
			fmt.Println(createError)
		} else {
			return catFile, nil
		}
	} else {
		fmt.Println(err)
	}
	return nil, err
}
