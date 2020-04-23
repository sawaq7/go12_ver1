package process4

import (

	    "net/http"
//	    "fmt"
	    "html/template"
	    "github.com/sawaq7/go12_ver1/client/reserve/html6"
//	    "github.com/sawaq7/go12_ver1/client/reserve/type6"
	    "github.com/sawaq7/go12_ver1/client/reserve/datastore6/trans5"
//	    "strconv"
                                                )


func Medical_xray_show(w http.ResponseWriter, r *http.Request ,guest_no int64) {

//     IN    w      　　: レスポンスライター
//     IN    r      　　: リクエストパラメータ
//     IN  guest_no     : guest no

//    fmt.Fprintf( w, "medical_xray_show start \n" )

///
///     set template
///

    monitor := template.Must(template.New("html").Parse(html6.Medical_xray_show))

///
///      get Reserve in d.s.
///

//    d_area_view := datastore2.D_Area( "D_Area","trans" ,district_no , w , r  )

    guest_medical_xray_slice := trans5.Guest_medical_xray( guest_no , w , r  )

//    get value from interface data

//    value, _ := d_area_view.([]type2.D_Area) //    get value from interface data

///
///     show reserve inf. on web
///

//   fmt.Fprintf( w, "medical_xray_show d_area_view %v\n" ,d_area_view)

	err := monitor.Execute(w, guest_medical_xray_slice)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

