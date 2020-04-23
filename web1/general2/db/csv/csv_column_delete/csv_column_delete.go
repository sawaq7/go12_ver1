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

//    fmt.Fprintf( w, "sky/csv_column_delete start \n" )

    var err error

///
/// 入力データをGET 　
///

    string_data := r.FormValue("delete_column")

    strings := strings2.String_no_get( w , r , string_data  )

    delete_no := make( []int, len(strings) )

    for pos, stringsw := range strings {

      delete_no[pos] ,err = strconv.Atoi(stringsw)  // make an integer
      if err != nil {
	   http.Error(w,err.Error(), http.StatusInternalServerError)

		return
	  }

	}

///
///         delete column　
///

	for _ , delete_now := range delete_no {

      process3.Csv_column_delete ( w , r ,delete_now  )

	}

///
/// 　　　csv inf. show web　
///

    csv_inf := trans3.Csv_inf ( w ,r )  //     get csv inf.

    monitor := template.Must( template.New("html").Parse( html5.Csv_show )) // set template

     err = monitor.Execute ( w, csv_inf )
	 if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	 }

}
