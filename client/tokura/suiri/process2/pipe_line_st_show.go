package process2

import (

	    "net/http"
//	    "fmt"
	    "html/template"
	    "client/tokura/html4"
        "client/tokura/suiri/type4"
	    "client/tokura/storage3"
	    "storage2"
                                                )
func Pipe_line_st_show(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "process2.pipe_line_st_show start \n" )  // デバック

    var idmy1 ,idmy2 int64

    skip_flag := 1

    bucket := "sample-7777"

///
///             Water2　ファイルがあるかチェック
///

    objects :=  storage2.Object_List ( w  ,r , bucket )  // バケット内のオブジェクトをゲットする

    for _ , objectsw := range objects {

      if objectsw == "Water2.txt" {

	    skip_flag = 0

	  }

    }

//    fmt.Fprintf(w, "process2.pipe_line_st_show : skip_flag %v\n", skip_flag )  // デバック

///
///            、表示用データ・テンプレートのヘッダーをGETして表示
///

     monitor := template.Must(template.New("html").Parse(html4.Pipe_line_st_keyin))

     if skip_flag == 0 {

     water2_view_minor , _ := storage3.Storage_tokura( "Water2" ,"trans" ,idmy1 , idmy2 , w , r  )

//       water2_view := trans4.Water2 ( w ,r )

       water2_view, _ := water2_view_minor.([]type4.Water2)  // インターフェイス型を型変換

       err := monitor.Execute(w, water2_view)

       if err != nil {

	    http.Error(w, err.Error(), http.StatusInternalServerError)

	   }

     } else {

       water2_view := make([]type4.Water2, 0)   //   Water2　の表示エリアを確保

       err := monitor.Execute(w, water2_view)

       if err != nil {

	    http.Error(w, err.Error(), http.StatusInternalServerError)

	   }

     }

}

