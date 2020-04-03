package process

import (

	    "net/http"
//	    "fmt"
	    "html/template"
	    "github.com/sawaq7/go12_ver1/client/sgh/html2"
	    "github.com/sawaq7/go12_ver1/client/sgh/type2"
	    "github.com/sawaq7/go12_ver1/general/type5"
//	    "github.com/sawaq7/go12_ver1/client/sgh/datastore2/trans"

        "github.com/sawaq7/go12_ver1/client/sgh/datastore2"
//	    "time"
                                                )


func Deliver_showall2( course_no int64 ,w http.ResponseWriter, r *http.Request ) {

//     IN  course_no  : コースNo
//     IN    w      　: レスポンスライター
//     IN    r      　: リクエストパラメータ

//   fmt.Fprintf( w, "process.deliver_show_all2 start \n" )  // チE��チE��

///
///     チE��プレート�EヘッダーをGET
///

     monitor := template.Must(template.New("html").Parse(html2.Deliver_showall1))

///
///    チE�Eタストアーから、表示用チE�EタをGET
///


     general_work := make([]type5.General_Work, 2)
     general_work[0].Int64_Work = 0          // 地区惁E��
     general_work[1].Int64_Work = course_no  //　コースNO

//     deliver_view := datastore2.D_store( "Deliver"  ,"trans"  ,general_work , w , r  )
     deliver_view := datastore2.Datastore_sgh( "Deliver"  ,"trans"  ,general_work , w , r  )

     // 空インターフェイス変数よりバリュー値をゲチE��

     value, _ := deliver_view.([]type2.Deliver)

///
///      モニターに表示
///

    err := monitor.Execute(w, value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

