package process4

import (

	    "net/http"
//	    "fmt"
	    "html/template"
	    "github.com/sawaq7/go12_ver1/client/reserve/html6"
//	    "github.com/sawaq7/go12_ver1/client/reserve/type6"
	    "github.com/sawaq7/go12_ver1/client/reserve/datastore6/trans5"
                                                )


func Reserve_register(w http.ResponseWriter, r *http.Request ,guest_no int64) {

//     IN    w      　　: レスポンスライター
//     IN    r      　　: リクエストパラメータ
//     IN  guest_no     : guest no

//    fmt.Fprintf( w, "reserve_register start \n" )

///
///     set template
///

    monitor := template.Must(template.New("html").Parse(html6.Reserve_register))

///
///      get Guest_reserve_minor in d.s.
///

//    d_area_view := datastore2.Datastore_sgh( "D_Area","trans" ,district_no , w , r  )

    guest_reserve_minor_slice := trans5.Guest_reserve_minor( guest_no , w , r  )

//    value, _ := d_area_view.([]type2.D_Area)

///
///     show reserve inf. on web
///

//   fmt.Fprintf( w, "reserve_register d_area_view %v\n" ,d_area_view)

//	err := monitor.Execute(w, value)

	err := monitor.Execute(w, guest_reserve_minor_slice)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

