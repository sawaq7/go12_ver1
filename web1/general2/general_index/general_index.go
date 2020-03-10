package general_index

import (

	    "net/http"
	    "general/html5"
	    "html/template"
                                                  )

///
/// show the menu of general' applications
///

func General_index(w http.ResponseWriter, r *http.Request) {

   var cdmy string

/// テンプレートのヘッダーをGET

    monitor := template.Must(template.New("html").Parse(html5.General_index))

// モニターに表示

	err := monitor.Execute(w, cdmy)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

