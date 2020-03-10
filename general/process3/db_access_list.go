package process3

import (

	    "net/http"
//	    "fmt"
	    "html/template"
	    "general/html5"

	    "general/datastore5/trans3"
                                                )

func Db_access_list(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "process3.db_access_list start \n" )  // デバック

///
///    テンプレートのヘッダーをGET
///

     monitor := template.Must(template.New("html").Parse(html5.Db_access_list))

///
///     データストアーから、表示用データをGET
///

     db_access_list := trans3.Db_access_list ( w ,r )

///
///       モニターに表示
///

	err := monitor.Execute(w, db_access_list)
	if err != nil {
	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
	}
}

