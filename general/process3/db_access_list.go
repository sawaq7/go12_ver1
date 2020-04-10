package process3

import (

	    "net/http"
//	    "fmt"
	    "html/template"
	    "github.com/sawaq7/go12_ver1/general/html5"

	    "github.com/sawaq7/go12_ver1/general/datastore5/trans3"
                                                )

func Db_access_list(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "process3.db_access_list start \n" )  // チE��チE��

///
///    set template
///

     monitor := template.Must(template.New("html").Parse(html5.Db_access_list))

///
///     get d.b. access list
///

     db_access_list := trans3.Db_access_list ( w ,r )

///
///      set web
///

	err := monitor.Execute(w, db_access_list)
	if err != nil {
	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
	}
}

