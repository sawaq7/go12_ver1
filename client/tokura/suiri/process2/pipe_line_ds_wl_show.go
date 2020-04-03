package process2

import (

	    "net/http"
//	    "fmt"
	    "html/template"
	    "github.com/sawaq7/go12_ver1/client/tokura/html4"

	    "github.com/sawaq7/go12_ver1/client/tokura/datastore4"
	    "github.com/sawaq7/go12_ver1/client/tokura/suiri/type4"
//	    "time"
                                                )
///                           　　　　　　　　　　
///      水路名より水路ライン惁E��を表示する
///                          　　　　　　　　　　　

func Pipe_line_ds_wl_show(funct int64 ,wname string ,w http.ResponseWriter, r *http.Request) {

//     IN  funct : ファンクション　0:すべての水路ラインを表示
//               　　　　　　　　　1:持E��した水路名�E水路ラインを表示
//     IN  wname : 水路吁E　　　　 * funct= 0の場合�Eダミ�E
//     IN    w      　: レスポンスライター
//     IN    r      　: リクエストパラメータ

//    fmt.Fprintf( w, "process2.pipe_line_ds_wl_show start \n" )  // チE��チE��
//    fmt.Fprintf( w, "process2.pipe_line_ds_wl_show funct %v   \n" , funct  )  // チE��チE��
//    fmt.Fprintf( w, "process2.pipe_line_ds_wl_show wname %v   \n" , wname  )  // チE��チE��

///
///           チE��プレート�EヘッダーをGET
///

     monitor := template.Must(template.New("html").Parse(html4.Pipe_line_ds_wl_keyin))

///
///           チE�Eタストアーから、表示用チE�EタをGET
///

//     water_line_view := trans2.Water_line (funct  ,wname , w ,r )

      water_line_view := datastore4.Datastore_tokura( "Water_Line"  ,"trans"  ,wname , w , r  )



     value, _ := water_line_view.([]type4.Water_Line)    // 空インターフェイス変数よりバリュー値をゲチE��

//     fmt.Fprintf( w, "process2.pipe_line_ds_wl_show : len(water_line_view) %v\n", len(water_line_view) )  // チE��チE��


// モニターに表示

	err := monitor.Execute ( w, value )
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

