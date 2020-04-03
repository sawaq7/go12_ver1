package csv_match_exp

import (

	    "net/http"
//	    "fmt"

	    "html/template"

	    "github.com/sawaq7/go12_ver1/general/html5"
	    "github.com/sawaq7/go12_ver1/general/datastore5/trans3"

                                                  )

func Csv_match_exp(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "csv_match_exp start \n" )  // チE��チE��

///
///　　web にcsv惁E��を表示
///

     csv_inf_view := trans3.Csv_inf ( w ,r )  ///      csv惁E��をゲチE��

     monitor := template.Must( template.New("html").Parse( html5.Csv_match_exp )) // チE��プレート�EヘッダーをGET

     err := monitor.Execute ( w, csv_inf_view )
	 if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	 }

//	fmt.Fprintf( w, "csv_match_exp : normal end \n" )  // チE��チE��

}

