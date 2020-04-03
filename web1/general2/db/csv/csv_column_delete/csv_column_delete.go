package csv_column_delete

import (

//	    "fmt"
	    "net/http"

	    "github.com/sawaq7/go12_ver1/general/process3"
	    "strconv"
	    "html/template"
	    "github.com/sawaq7/go12_ver1/general/datastore5/trans3"
	    "github.com/sawaq7/go12_ver1/general/html5"
	    "github.com/sawaq7/go12_ver1/general/strings2"

                                    )

func Csv_column_delete(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "sky/csv_column_delete start \n" )  // γEγE―

    var err error

///
/// ε₯εγγΌγΏγGET γ
///

    string_data := r.FormValue("delete_column")  // ει€γγεγγ²γE

    strings := strings2.String_no_get( w , r , string_data  )

    delete_no := make( []int, len(strings) )

    for pos, stringsw := range strings {

      delete_no[pos] ,err = strconv.Atoi(stringsw)  // ζ΄ζ°εE	  if err != nil {
	   http.Error(w,err.Error(), http.StatusInternalServerError)

		return
	  }

	}

///
///         εγει€γγγ
///

	for _ , delete_now := range delete_no {

      process3.Csv_column_delete ( w , r ,delete_now  )

	}

///
/// γγγwebδΈγ«γcsvζE ±γθ‘¨η€Ίγγγ
///

    csv_inf := trans3.Csv_inf ( w ,r )  //     csvζE ±γγ²γE

    monitor := template.Must( template.New("html").Parse( html5.Csv_show )) // γE³γγ¬γΌγγEγγγγΌγGET

     err = monitor.Execute ( w, csv_inf )
	 if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	 }

}
