package csv_match_wording2

import (

	    "net/http"
	    "fmt"

	    "strings"
                                                  )

func Csv_match_wording2(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "csv_match_wording2 start \n" )

    ex_data := r.FormValue("text2")

	fmt.Fprintf( w, "csv_match_wording2 ex_data:  %v\n", ex_data )

	ex_data2 := strings.Split ( ex_data, " "  )

    index:= 0

    for ii := 0 ; ii < len(ex_data2) ; ii++ {

      index = len(ex_data2) - ii - 1

      fmt.Fprintf( w, "csv_match_wording2 ex_data2:  %v\n", ex_data2[index] )

    }

	fmt.Fprintf( w, "csv_match_wording2 ex_data num:  %v\n", len(ex_data2) )

//	fmt.Fprintf( w, "csv_match_wording2 : normal end \n" )

}
