package csv_match_wording

import (

	    "net/http"
//	    "fmt"

	    "github.com/sawaq7/go12_ver1/general/html5"
	    "html/template"
                                                  )

func Csv_match_wording(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "csv_match_wording start \n" )

     var s_dmy string

    monitor := template.Must(template.New("html").Parse(html5.Csv_match_wording))

    err := monitor.Execute(w, s_dmy )

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

//	fmt.Fprintf( w, "csv_match_wording : normal end \n" )

}
