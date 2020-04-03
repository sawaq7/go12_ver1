package csv_match_exp

import (

	    "net/http"
//	    "fmt"

	    "html/template"

	    "github.com/sawaq7/go12_ver1/general/html5"
	    "github.com/sawaq7/go12_ver1/general/datastore5/trans3"

                                                  )

func Csv_match_exp(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "csv_match_exp start \n" )  // 繝・ヰ繝・け

///
///縲縲web 縺ｫcsv諠・ｱ繧定｡ｨ遉ｺ
///

     csv_inf_view := trans3.Csv_inf ( w ,r )  ///      csv諠・ｱ繧偵ご繝・ヨ

     monitor := template.Must( template.New("html").Parse( html5.Csv_match_exp )) // 繝・Φ繝励Ξ繝ｼ繝医・繝倥ャ繝繝ｼ繧竪ET

     err := monitor.Execute ( w, csv_inf_view )
	 if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	 }

//	fmt.Fprintf( w, "csv_match_exp : normal end \n" )  // 繝・ヰ繝・け

}

