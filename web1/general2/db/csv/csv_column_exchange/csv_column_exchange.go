package csv_column_exchange

import (
//	     "fmt"
	     "net/http"

	     "github.com/sawaq7/go12_ver1/general/process3"
	     "github.com/sawaq7/go12_ver1/general/html5"
	     "strconv"
	     "github.com/sawaq7/go12_ver1/general/datastore5/trans3"
	     "html/template"
                                      )

func Csv_column_exchange(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "sky/csv_column_exchange start \n" )

///
///     get input data　
///

    exchange_column1_minor := r.FormValue("column1")  //  get column no.

	exchange_column1 ,err := strconv.Atoi(exchange_column1_minor)  //   make an integer
	if err != nil {
	   http.Error(w,err.Error(), http.StatusInternalServerError)

		return
	}

    exchange_column2_minor := r.FormValue("column2")  //  get column no. for deleting

	exchange_column2 ,err := strconv.Atoi(exchange_column2_minor)  // make an integer
	if err != nil {
	   http.Error(w,err.Error(), http.StatusInternalServerError)

		return
	}

///
///     add column
///

    if exchange_column1 < exchange_column2 {

	   exchange_column2 ++

	}
	   process3.Csv_column_join2 ( w , r ,exchange_column1 ,exchange_column2 )

///
///
///     delete column
///

    if exchange_column1 >= exchange_column2 {

	   exchange_column1 ++

	}

	process3.Csv_column_delete ( w , r ,exchange_column1  )

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
