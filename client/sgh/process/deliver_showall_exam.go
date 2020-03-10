package process

import (
//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "net/http"
	    "fmt"
	    "html/template"
	    "client/sgh/html2"
//	    "client/sgh/types"
	    "client/sgh/datastore2/trans"
//	    "time"
                                                )


func Deliver_showall_exam(w http.ResponseWriter, r *http.Request) {

//     IN    w      　: レスポンスライター
//     IN    r      　: リクエストパラメータ

    fmt.Fprintf( w, "process.deliver_show_all_exam start \n" )  // デバック

// テンプレートのヘッダーをGET

     monitor := template.Must(template.New("html").Parse(html2.Deliver_showall1))

// データストアーから、表示用データをGET

     deliver_view := trans.Deliver2 (w ,r )

// モニターに表示

	err := monitor.Execute(w, deliver_view)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

