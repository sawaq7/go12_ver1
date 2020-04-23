package sgh_index

import (

	    "net/http"
	    "github.com/sawaq7/go12_ver1/client/sgh/html2"
	    "html/template"
                                                  )

///
/// show the menu of tokuras' applications
///

func Sgh_index(w http.ResponseWriter, r *http.Request) {

   var cdmy string

///     set template

    monitor := template.Must(template.New("html").Parse(html2.Sgh_index))

///   show sgh menu on web

	err := monitor.Execute(w, cdmy)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

