package csv_match_exp

import (

	    "net/http"
//	    "fmt"

	    "html/template"

	    "github.com/sawaq7/go12_ver1/general/html5"
	    "github.com/sawaq7/go12_ver1/general/datastore5/trans3"

                                                  )

func Csv_match_exp(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "csv_match_exp start \n" )  // γEγE―

///
///γγweb γ«csvζE ±γθ‘¨η€Ί
///

     csv_inf_view := trans3.Csv_inf ( w ,r )  ///      csvζE ±γγ²γE

     monitor := template.Must( template.New("html").Parse( html5.Csv_match_exp )) // γE³γγ¬γΌγγEγγγγΌγGET

     err := monitor.Execute ( w, csv_inf_view )
	 if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	 }

//	fmt.Fprintf( w, "csv_match_exp : normal end \n" )  // γEγE―

}

