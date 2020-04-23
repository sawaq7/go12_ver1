package process

import (

	    "net/http"
//	    "fmt"
	    "html/template"
	    "github.com/sawaq7/go12_ver1/client/sgh/html2"
	    "github.com/sawaq7/go12_ver1/client/sgh/datastore2/trans"

                                                )


func Ai_sgh_ex_show( course_no int64 ,w http.ResponseWriter, r *http.Request ) {

//     IN  course_no  : コースNo
//     IN    w      　: レスポンスライター
//     IN    r      　: リクエストパラメータ

//    fmt.Fprintf( w, "process.ai_sgh_ex_show start \n" )

     //  set template
     monitor := template.Must(template.New("html").Parse(html2.Ai_sgh_ex_show))


     // get ai data
     sgh_ai_view := trans.Sgh_ai ( course_no ,w ,r )

    //      show area inf. on web
    err := monitor.Execute(w, sgh_ai_view)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}