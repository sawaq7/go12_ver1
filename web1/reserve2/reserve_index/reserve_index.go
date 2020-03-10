package reserve_index

import (

	    "net/http"
	    "client/reserve/html6"
	    "html/template"
                                                  )

///
/// show the menu of medical' applications
///

func Reserve_index(w http.ResponseWriter, r *http.Request) {

   var cdmy string

/// テンプレートのヘッダーをGET

    monitor := template.Must(template.New("html").Parse(html6.Reserve_index))

// モニターに表示

	err := monitor.Execute(w, cdmy)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

