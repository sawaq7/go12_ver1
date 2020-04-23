package process

import (

	    "net/http"
	    "fmt"
	    "html/template"
	    "github.com/sawaq7/go12_ver1/client/sgh/html2"
	    "github.com/sawaq7/go12_ver1/client/sgh/type2"
	    "github.com/sawaq7/go12_ver1/general/type5"
	    "github.com/sawaq7/go12_ver1/client/sgh/datastore2"
//	    "time"
                                                )

///
///     register  district inf. in d.s.
///

func D_district_keyin_all(w http.ResponseWriter, r *http.Request) {

//     IN    w      　　: レスポンスライター
//     IN    r      　　: リクエストパラメータ

    fmt.Fprintf( w, "d_district_keyin_all start \n" )

//    set template

     monitor := template.Must(template.New("html").Parse(html2.D_district_keyin_all))

//       get district inf in d.s.

//     d_district_view := trans.D_district ( 0 ,0 ,w ,r )

     general_work := make([]type5.General_Work, 2)
     general_work[0].Int64_Work = 0
     general_work[1].Int64_Work = 0

     deliver_view := datastore2.Datastore_sgh( "D_District" ,"trans" ,general_work , w , r  )

//    get value from interface data

     value, _ := deliver_view.([]type2.D_District)


//      show district inf. on web

    err := monitor.Execute(w, value)
//	err := monitor.Execute(w, d_district_view)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

