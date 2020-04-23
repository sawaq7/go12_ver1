package process

import (

	    "net/http"
	    "fmt"
	    "html/template"
	    "github.com/sawaq7/go12_ver1/client/sgh/html2"
	    "github.com/sawaq7/go12_ver1/client/sgh/datastore2"

	    "github.com/sawaq7/go12_ver1/client/sgh/type2"
	    "github.com/sawaq7/go12_ver1/general/type5"

//	    "time"
                                                )

///
///      show deliver inf. for each private in d.s.
///

func Deliver_showprivate(w http.ResponseWriter, r *http.Request ,private_no int64) {


//     IN    w      　　: レスポンスライター
//     IN    r      　　: リクエストパラメータ
//     IN  private_no   : private no

    fmt.Fprintf( w, "deliver_showprivate start \n" )
    fmt.Fprintf( w, "deliver_showprivate : private_no %v\n", private_no )

//     set template

     monitor := template.Must(template.New("html").Parse(html2.Deliver_showprivate))

//     get deliver inf in d.s.

     general_work := make([]type5.General_Work, 2)
     general_work[0].Int64_Work = 2          //
     general_work[1].Int64_Work = private_no

     deliver_view := datastore2.Datastore_sgh( "Deliver" ,"trans" ,general_work , w , r  )

     //  get value from interface data

     value, _ := deliver_view.([]type2.Deliver)

///
///    show deliver inf. on web
///

	err := monitor.Execute(w, value)
//	err := monitor.Execute(w, deliver_view)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

