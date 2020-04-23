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

//    fmt.Fprintf( w, "sky/csv_column_join start \n" )

    var err error

///
///        get input data　
///

    string_data := r.FormValue("join_column")
    strings := strings2.String_no_get( w , r , string_data  )  // get some numbers

    join_no := make( []int, len(strings) )

    for pos, stringsw := range strings {

      join_no[pos] ,err = strconv.Atoi(stringsw)  //  make an interer
      if err != nil {
	   http.Error(w,err.Error(), http.StatusInternalServerError)

		return
	  }

	}

	filename := r.FormValue("join_file")  // get file name which is added

///
///       add some columns 　
///

    for _ , join_now := range join_no {

      process3.Csv_column_join ( w , r ,filename ,join_now )

	}

///
/// 　　　csv inf. show web　
///

    csv_inf := trans3.Csv_inf ( w ,r )  //     get csv inf.

//    fmt.Fprintf( w, "sky/csv_column_join : csv_in %v\n", csv_inf )

    monitor := template.Must( template.New("html").Parse( html5.Csv_show )) // set template

    err = monitor.Execute ( w, csv_inf )   //   show csv inf. on web　
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
