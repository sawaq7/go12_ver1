package tokura_index_pl

import (

	    "net/http"
	    "github.com/sawaq7/go12_ver1/client/tokura/html4"
	    "html/template"
                                                  )

///
/// show the menu of tokuras' applications
///


func Tokura_index_pl(w http.ResponseWriter, r *http.Request) {

   var cdmy string

/// ãE³ãã¬ã¼ããEãããã¼ãGET

    monitor := template.Must(template.New("html").Parse(html4.Tokura_index_pl))

// ã¢ãã¿ã¼ã«è¡¨ç¤º

	err := monitor.Execute(w, cdmy)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

