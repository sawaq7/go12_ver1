package process

import (
//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "net/http"
	    "fmt"
	    "html/template"
	    "github.com/sawaq7/go12_ver1/client/sgh/html2"
//	    "github.com/sawaq7/go12_ver1/client/sgh/types"
	    "github.com/sawaq7/go12_ver1/client/sgh/datastore2/trans"
//	    "time"
                                                )


func Deliver_showall_exam(w http.ResponseWriter, r *http.Request) {

//     IN    w      　: レスポンスライター
//     IN    r      　: リクエストパラメータ

    fmt.Fprintf( w, "process.deliver_show_all_exam start \n" )  // チE��チE��

// チE��プレート�EヘッダーをGET

     monitor := template.Must(template.New("html").Parse(html2.Deliver_showall1))

// チE�Eタストアーから、表示用チE�EタをGET

     deliver_view := trans.Deliver2 (w ,r )

// モニターに表示

	err := monitor.Execute(w, deliver_view)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

