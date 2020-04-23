package process

import (

	    "net/http"
	    "fmt"
	    "html/template"
	    "github.com/sawaq7/go12_ver1/client/sgh/datastore2"
	    "github.com/sawaq7/go12_ver1/client/sgh/html2"
	    "github.com/sawaq7/go12_ver1/client/sgh/type2"
	    "github.com/sawaq7/go12_ver1/general/type5"
//	    "time"
                                                )


func Ai_district_showall(w http.ResponseWriter, r *http.Request) {

//     IN    w      　　: レスポンスライター
//     IN    r      　　: リクエストパラメータ

    fmt.Fprintf( w, "ai_district_show_all start \n" )


     //  set template
     monitor := template.Must(template.New("html").Parse(html2.Ai_district_showall))

     general_work := make([]type5.General_Work, 2)
     general_work[0].Int64_Work = 0
     general_work[1].Int64_Work = 0

     deliver_view := datastore2.Datastore_sgh( "D_District" ,"trans" ,general_work , w , r  )

     value, _ := deliver_view.([]type2.D_District)

///
///       show area inf. on web
///

    err := monitor.Execute(w, value)
//	err := monitor.Execute(w, d_district_view)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

