package process

import (
//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "net/http"
	    "fmt"
	    "html/template"
	    "github.com/sawaq7/go12_ver1/client/sgh/html2"
	    "github.com/sawaq7/go12_ver1/client/sgh/datastore2"

	    "github.com/sawaq7/go12_ver1/client/sgh/type2"
	    "github.com/sawaq7/go12_ver1/general/type5"

//	    "time"
                                                )

/// sub  car no ごとの　配達チE�Eタを表示する　 ///

func Deliver_showprivate(w http.ResponseWriter, r *http.Request ,private_no int64) {


//     IN    w      　　: レスポンスライター
//     IN    r      　　: リクエストパラメータ
//     IN  private_no   : プライベ�EチEo

    fmt.Fprintf( w, "deliver_showprivate start \n" )  // チE��チE��
    fmt.Fprintf( w, "deliver_showprivate : private_no %v\n", private_no )  // チE��チE��

// チE��プレート�EヘッダーをGET

     monitor := template.Must(template.New("html").Parse(html2.Deliver_showprivate))

// チE�Eタストアーから、表示用チE�EタをGET

     general_work := make([]type5.General_Work, 2)
     general_work[0].Int64_Work = 2          // 個人惁E��
     general_work[1].Int64_Work = private_no

//     deliver_view := datastore2.D_store( "Deliver" ,"trans" ,general_work , w , r  )
     deliver_view := datastore2.Datastore_sgh( "Deliver" ,"trans" ,general_work , w , r  )

     // 空インターフェイス変数よりバリュー値をゲチE��

     value, _ := deliver_view.([]type2.Deliver)

// モニターに表示

	err := monitor.Execute(w, value)

//	err := monitor.Execute(w, deliver_view)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

