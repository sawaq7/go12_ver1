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
/// ããæE®ããå°åNO.ã®å·è»æå ±ãè¡¨ç¤ºãã
///

func Car_show(w http.ResponseWriter, r *http.Request ,district_no int64) {

//     IN    w      ãã: ã¬ã¹ãã³ã¹ã©ã¤ã¿ã¼
//     IN    r      ãã: ãªã¯ã¨ã¹ããã©ã¡ã¼ã¿
//     IN  district_no  : å°åNo

//    fmt.Fprintf( w, "car_show start \n" )  // ãEãE¯}

// ãE³ãã¬ã¼ããEãããã¼ãGET

     monitor := template.Must(template.New("html").Parse(html2.Car_show))

// ãEEã¿ã¹ãã¢ã¼ãããè¡¨ç¤ºç¨ãEEã¿ãGET

//     car_view := datastore2.D_store( "Car" ,"trans"  ,district_no , w , r  )
     car_view := datastore2.Datastore_sgh( "Car" ,"trans"  ,district_no , w , r  )

     // ç©ºã¤ã³ã¿ã¼ãã§ã¤ã¹å¤æ°ããããªã¥ã¼å¤ãã²ãE

     value, _ := car_view.([]type2.Car)


// ã¢ãã¿ã¼ã«è¡¨ç¤º

	err := monitor.Execute(w, value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

