///                                           ///
/// スライス(float) を一行にしてファイルに書ぁE///
///                                          ///

package main

import (
	     "fmt"
	     "encoding/csv"
         "strings"
	     "os"
    	                 )

func main() {

    fmt.Println ("csv_read start　" )

	fname := "csv_write.csv"

    reader , _ := os.OpenFile(fname, os.O_CREATE|os.O_RDWR, 0666)

//	w := csv.NewWriter(os.Stdout)

	csv_reader := csv.NewReader ( strings.NewReader(reader) )

	for {

	    record, err := csv_reader.Read()
	    if err == io.EOF {

			break
		}
		if err != nil {

		   fmt.Println ("csv_read err1 " , err )

		}

	    fmt.Println ("csv_read record " record )

	}

	fmt.Println ("csv_read normal end　" )

}
