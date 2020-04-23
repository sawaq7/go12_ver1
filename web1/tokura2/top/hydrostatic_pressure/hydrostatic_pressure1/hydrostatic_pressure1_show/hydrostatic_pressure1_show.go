package hydrostatic_pressure1_show

import (
	     "net/http"
	     "html/template"
	     "github.com/sawaq7/go12_ver1/client/tokura/suiri/type4"

         "github.com/sawaq7/go12_ver1/client/tokura/html4"

	                                  )

func Hydrostatic_pressure1_show(w http.ResponseWriter, r *http.Request) {

   var seisui type4.Seisui

///      set template

	monitor := template.Must( template.New("html").Parse( html4.Hydrostatic_pressure1_show) )

///    kei-in item of seisui on web

	err := monitor.Execute(w, seisui)

    if err != nil {
	   http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
