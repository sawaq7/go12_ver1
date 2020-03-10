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
///                           　　　　　　　　　　
///      水路名より水路ライン情報を表示する
///                          　　　　　　　　　　　

func Pipe_line_ds_wl_show(funct int64 ,wname string ,w http.ResponseWriter, r *http.Request) {

//     IN  funct : ファンクション　0:すべての水路ラインを表示
//               　　　　　　　　　1:指定した水路名の水路ラインを表示
//     IN  wname : 水路名 　　　　 * funct= 0の場合はダミー
//     IN    w      　: レスポンスライター
//     IN    r      　: リクエストパラメータ

//    fmt.Fprintf( w, "process2.pipe_line_ds_wl_show start \n" )  // デバック
//    fmt.Fprintf( w, "process2.pipe_line_ds_wl_show funct %v   \n" , funct  )  // デバック
//    fmt.Fprintf( w, "process2.pipe_line_ds_wl_show wname %v   \n" , wname  )  // デバック

///
///           テンプレートのヘッダーをGET
///

     monitor := template.Must(template.New("html").Parse(html4.Pipe_line_ds_wl_keyin))

///
///           データストアーから、表示用データをGET
///

//     water_line_view := trans2.Water_line (funct  ,wname , w ,r )

      water_line_view := datastore4.Datastore_tokura( "Water_Line"  ,"trans"  ,wname , w , r  )



     value, _ := water_line_view.([]type4.Water_Line)    // 空インターフェイス変数よりバリュー値をゲット

//     fmt.Fprintf( w, "process2.pipe_line_ds_wl_show : len(water_line_view) %v\n", len(water_line_view) )  // デバック


// モニターに表示

	err := monitor.Execute ( w, value )
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

