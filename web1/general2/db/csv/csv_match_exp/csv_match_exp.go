package csv_match_exp

import (

	    "net/http"
//	    "fmt"

	    "html/template"

	    "general/html5"
	    "general/datastore5/trans3"

                                                  )

func Csv_match_exp(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "csv_match_exp start \n" )  // デバック

///
///　　web にcsv情報を表示
///

     csv_inf_view := trans3.Csv_inf ( w ,r )  ///      csv情報をゲット

     monitor := template.Must( template.New("html").Parse( html5.Csv_match_exp )) // テンプレートのヘッダーをGET

     err := monitor.Execute ( w, csv_inf_view )
	 if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	 }

//	fmt.Fprintf( w, "csv_match_exp : normal end \n" )  // デバック

}

