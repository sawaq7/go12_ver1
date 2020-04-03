///                                           ///
/// スライス(float) を一行にしてファイルに書ぁE///
///                                          ///

package main

import (
//	    "fmt"

	    "encoding/csv"
	    "log"


	    "os"
    	     )
func main() {
	records := [][]string{
		{"first_name", "last_name", "username"},
		{"Rob", "Pike", "rob"},
		{"Ken", "Thompson", "ken"},
		{"Robert", "Griesemer", "gri"},
	}

	fname := "test.csv"

    writer , _ := os.OpenFile(fname, os.O_CREATE|os.O_RDWR, 0666)

//	w := csv.NewWriter(os.Stdout)

	w := csv.NewWriter( writer )

	for _, record := range records {
		if err := w.Write(record); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
	}

	// Write any buffered data to the underlying writer (standard output).
	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}
