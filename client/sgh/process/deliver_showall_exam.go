package process

import (

	    "net/http"
	    "fmt"
	    "html/template"
	    "github.com/sawaq7/go12_ver1/client/sgh/html2"
//	    "github.com/sawaq7/go12_ver1/client/sgh/types"
	    "github.com/sawaq7/go12_ver1/client/sgh/datastore2/trans"
//	    "time"
                                                )

///
///     show deliver inf. in d.s.
///

func Deliver_showall_exam(w http.ResponseWriter, r *http.Request) {

//     IN    w      　: レスポンスライター
//     IN    r      　: リクエストパラメータ

    fmt.Fprintf( w, "process.deliver_show_all_exam start \n" )

//   set template

     monitor := template.Must(template.New("html").Parse(html2.Deliver_showall1))

//   get deliver inf in d.s.

     deliver_view := trans.Deliver2 (w ,r )

//   show dliver inf. on web

	err := monitor.Execute(w, deliver_view)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

