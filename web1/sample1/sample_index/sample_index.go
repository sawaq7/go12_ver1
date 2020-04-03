package sample_index

import (

	    "net/http"
	    "github.com/sawaq7/go12_ver1/temp/html1000"
	    "html/template"
                                                  )

///
/// show the menu of samples' applications
///

func Sample_index(w http.ResponseWriter, r *http.Request) {

   var cdmy string

/// チE��プレート�EヘッダーをGET

    monitor := template.Must(template.New("html").Parse(html1000.Sample_index))

// モニターに表示

	err := monitor.Execute(w, cdmy)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

