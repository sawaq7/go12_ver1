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
///     show deliver inf. for each car no in d.s.
///

func Deliver_showcar(w http.ResponseWriter, r *http.Request ,car_no int64) {

//     IN    w      　　: レスポンスライター
//     IN    r      　　: リクエストパラメータ
//     IN  car_no　　   : カーNo

    fmt.Fprintf( w, "deliver_showcar start \n" )
    fmt.Fprintf( w, "deliver_showcar : car_no %v\n", car_no )

//      set template

     monitor := template.Must(template.New("html").Parse(html2.Deliver_showcar))

//    get deliver inf in d.s.

     general_work := make([]type5.General_Work, 2)
     general_work[0].Int64_Work = 1          //   district no
     general_work[1].Int64_Work = car_no  //　  car no

     deliver_view := datastore2.Datastore_sgh( "Deliver" ,"trans"  ,general_work , w , r  )

    //  get value from interface data

     value, _ := deliver_view.([]type2.Deliver)

///     show deliver inf. on web

	err := monitor.Execute(w, value)
//	err := monitor.Execute(w, deliver_view)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

