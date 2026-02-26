package transactionhistory

import (
	"errors"
	"fmt"
	"os"
)

const HistoryFile = "Transaction.txt"

func WriteTOFile(data string) error {
	f,err:=os.OpenFile(HistoryFile,os.O_APPEND|os.O_CREATE|os.O_WRONLY,0600)
	if  err!= nil {
		return  errors.New("error in writing transaction to a file.")
	}
	defer f.Close()
	_,err = f.WriteString(data + "\n")
	return  err
}


func ReadFile()error{
	data,err:=os.ReadFile(HistoryFile)
	if  err!= nil {
		return  errors.New("error occured in reading transaction file")
	}
	b:= string(data)
	fmt.Println(b)
	return  nil
}