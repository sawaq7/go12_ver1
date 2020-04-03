///                                           ///
/// スライス(float) を一行にしてファイルに書ぁE///
///                                          ///

package main

import (
	     "fmt"
	     "encoding/csv"

	     "os"
    	                 )

func main() {

    fmt.Println ("csv_write start　" )

	records := [][]string{
		{"first_name", "last_name", "username"},
		{"Rob", "Pike", "rob"},
		{"Ken", "Thompson", "ken"},
		{"Robert", "Griesemer", "gri"},
	}

	fname := "csv_write.csv"

    writer , _ := os.OpenFile(fname, os.O_CREATE|os.O_RDWR, 0666)

//	w := csv.NewWriter(os.Stdout)

	csv_writer := csv.NewWriter( writer )

	for _, record := range records {
		if err := csv_writer.Write(record); err != nil {
//			log.Fatalln("error writing record to csv:", err)
			fmt.Println ("csv_write 　err1 " , err )
		}
	}

	// Write any buffered data to the underlying writer (standard output).
	csv_writer.Flush()

	if err := csv_writer.Error(); err != nil {

	   fmt.Println ("csv_write 　err2 " , err )

	}

	fmt.Println ("csv_write normal end　" )

}
