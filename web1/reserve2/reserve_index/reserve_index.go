package reserve_index

import (

	    "net/http"
	    "github.com/sawaq7/go12_ver1/client/reserve/html6"
	    "html/template"
                                                  )

///
/// show the menu of medical's applications
///

func Reserve_index(w http.ResponseWriter, r *http.Request) {

   var cdmy string

///      set template

    monitor := template.Must(template.New("html").Parse(html6.Reserve_index))

///           show medical's menu on web

	err := monitor.Execute(w, cdmy)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

