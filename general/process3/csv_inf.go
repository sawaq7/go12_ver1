package process3

import (

	    "net/http"
//	    "fmt"
	    "html/template"

	    "github.com/sawaq7/go12_ver1/general/html5"
	    "github.com/sawaq7/go12_ver1/general/datastore5/trans3"
//	    "time"
                                                )


func Csv_inf(w http.ResponseWriter, r *http.Request ) {

//     IN    w      　　: レスポンスライター
//     IN    r      　　: リクエストパラメータ
//     IN  district_no  : 地域No

//    fmt.Fprintf( w, "csv_inf start \n" )

//    set template

     monitor := template.Must(template.New("html").Parse(html5.Csv_show))


//    get csv inf.

//    d_area_view := datastore2.Datastore_sgh( "D_Area","trans" ,district_no , w , r  )

    csv_inf := trans3.Csv_inf ( w ,r )

//  set web
//   fmt.Fprintf( w, "csv_inf csv_inf %v\n" ,csv_inf)

	err := monitor.Execute( w, csv_inf )

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

