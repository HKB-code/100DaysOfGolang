package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

type FileManager struct{
	InputFilePath string
	OutputFilePath string
}

// func ReadFiles(path string)([]string,error) {
// 	file, err := os.Open(path)
// 	if err != nil {
		
// 		return nil,errors.New("Failed to open file!")
// 	}
// 	scanner := bufio.NewScanner(file)
// 	var lines []string
// 	for scanner.Scan() {
// 		lines = append(lines, scanner.Text())
// 	}
// 	if err = scanner.Err(); err != nil {
	
// 		file.Close()
// 		return nil, errors.New("Reading the file content Failed")
// 	}
// 	file.Close()
// 	return lines,nil
// }

func (fm FileManager) ReadFiles()([]string,error) {
	file, err := os.Open(fm.InputFilePath)
	if err != nil {
		
		return nil,errors.New("Failed to open file!")
	}
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err = scanner.Err(); err != nil {
	
		file.Close()
		return nil, errors.New("Reading the file content Failed")
	}
	file.Close()
	return lines,nil
}
func (fm FileManager) WriteResult(data any)error{
	file,err:= os.Create(fm.OutputFilePath)
	if err != nil {
	return  errors.New("Failed to create File")
	}

	encoder:=json.NewEncoder(file)
	err=encoder.Encode(data)
	if err != nil {
		return errors.New("Failed to convert data into json")
	}
	return  nil
}

func NewFm(inputPath string, outputPath string)*FileManager{
	return &FileManager{
		InputFilePath: inputPath,
		OutputFilePath: outputPath,
	}
}