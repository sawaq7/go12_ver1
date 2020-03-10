package process2

import (

	    "net/http"
//	    "fmt"
	    "html/template"
	    "client/tokura/html4"

	    "client/tokura/datastore4"
	    "client/tokura/suiri/type4"
//	    "time"
                                                )
func Pipe_line_ds_show(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "process2.pipe_line_ds_show start \n" )  // デバック

    var idmy int64

///
///      テンプレートのヘッダーをGET
///

     monitor := template.Must(template.New("html").Parse(html4.Pipe_line_ds_keyin))

///
///   データストアーから、表示用データをGET
///


//     water2_view := trans2.Water2 ( w ,r )

     water2_view := datastore4.Datastore_tokura( "Water2"  ,"trans"  ,idmy , w , r  )

     value, _ := water2_view.([]type4.Water2)    // 空インターフェイス変数よりバリュー値をゲット


///
///             モニターに表示
///

	err := monitor.Execute(w, value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

