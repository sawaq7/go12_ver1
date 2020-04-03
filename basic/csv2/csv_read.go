///                                           ///
/// スライス(float) を一行にしてファイルに書ぁE///
///                                          ///

package main

import (
	     "fmt"
//	     "encoding/csv"
         "bufio"
	     "os"
	     "io"
    	                 )

func main() {

    fmt.Println ("csv_read start　" )

	fname := "csv_write.csv"

    reader , _ := os.OpenFile(fname, os.O_CREATE|os.O_RDWR, 0666)

    csv_reader := bufio.NewReaderSize(reader, 4096)

	for {

	    record ,err  := csv_reader.ReadString('\n')

	    if err == io.EOF {

			break
		}
		if err != nil {

		   fmt.Println ("csv_read err1 " , err )

		}

	    fmt.Println ("csv_read record " , record )

	}

	fmt.Println ("csv_read normal end　" )

}
