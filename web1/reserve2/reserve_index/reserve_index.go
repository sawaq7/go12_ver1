package reserve_index

import (

	    "net/http"
	    "github.com/sawaq7/go12_ver1/client/reserve/html6"
	    "html/template"
                                                  )

///
/// show the menu of medical' applications
///

func Reserve_index(w http.ResponseWriter, r *http.Request) {

   var cdmy string

/// ãE³ãã¬ã¼ããEãããã¼ãGET

    monitor := template.Must(template.New("html").Parse(html6.Reserve_index))

// ã¢ãã¿ã¼ã«è¡¨ç¤º

	err := monitor.Execute(w, cdmy)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

