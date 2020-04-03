package csv_column_join

import (

//	    "fmt"
	    "net/http"

	    "github.com/sawaq7/go12_ver1/general/process3"
	    "github.com/sawaq7/go12_ver1/general/datastore5/trans3"
        "html/template"
        "github.com/sawaq7/go12_ver1/general/html5"
        "github.com/sawaq7/go12_ver1/general/strings2"

	    "strconv"

                                                  )

func Csv_column_join(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "sky/csv_column_join start \n" )  // γEγE―

    var err error

///
///        ε₯εγγΌγΏγGET γ
///

    string_data := r.FormValue("join_column")  // θΏ½ε γγεγγ²γE
    strings := strings2.String_no_get( w , r , string_data  )

    join_no := make( []int, len(strings) )

    for pos, stringsw := range strings {

      join_no[pos] ,err = strconv.Atoi(stringsw)  // ζ΄ζ°εE	  if err != nil {
	   http.Error(w,err.Error(), http.StatusInternalServerError)

		return
	  }

	}

	filename := r.FormValue("join_file")  // θΏ½ε γγγγ‘γ€γ«εγγ²γE

///
/// εεδ½γEγEEγΏγε γγ γ
///

    for _ , join_now := range join_no {

      process3.Csv_column_join ( w , r ,filename ,join_now )

	}

///
/// γγγwebδΈγ«γcsvζE ±γθ‘¨η€Ίγγγ
///

    csv_inf := trans3.Csv_inf ( w ,r )  //     csvζE ±γγ²γE

//    fmt.Fprintf( w, "sky/csv_column_join : csv_in %v\n", csv_inf )  // γEγE―

    monitor := template.Must( template.New("html").Parse( html5.Csv_show )) // γE³γγ¬γΌγγEγγγγΌγGET

    err = monitor.Execute ( w, csv_inf )   // webδΈγ«γcsvζE ±γθ‘¨η€Ίγγγ
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
