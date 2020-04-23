package process

import (

	    "net/http"
//	    "fmt"
	    "html/template"
	    "github.com/sawaq7/go12_ver1/client/sgh/html2"
	    "github.com/sawaq7/go12_ver1/client/sgh/type2"
	    "github.com/sawaq7/go12_ver1/client/sgh/datastore2"

                                                )
///
/// 　　get car inf. which is selected by district no
///

func Car_show(w http.ResponseWriter, r *http.Request ,district_no int64) {

//     IN    w      　　: レスポンスライター
//     IN    r      　　: リクエストパラメータ
//     IN  district_no  :  district no

//    fmt.Fprintf( w, "car_show start \n" )

//  set template

     monitor := template.Must(template.New("html").Parse(html2.Car_show))

//    get car inf. in d.s.

     car_view := datastore2.Datastore_sgh( "Car" ,"trans"  ,district_no , w , r  )

     value, _ := car_view.([]type2.Car)   // get value from interface data


///
///     show car inf. on web
///

	err := monitor.Execute(w, value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

