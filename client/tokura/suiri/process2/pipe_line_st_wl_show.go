package process2

import (

	    "net/http"
//	    "fmt"
	    "html/template"
	    "client/tokura/html4"

//	    "client/tokura/storage3/trans4"
	    "client/tokura/storage3"
	    "client/tokura/suiri/type4"
//	    "time"
                                                )

func Pipe_line_st_wl_show( wname string ,w http.ResponseWriter, r *http.Request) {

//     IN  wname : 水路名 　　　　
//     IN     w  : レスポンスライター
//     IN     r  : リクエストパラメーター


//    fmt.Fprintf( w, "process2.pipe_line_st_wl_show start \n" )  // デバック

//    fmt.Fprintf( w, "process2.pipe_line_st_wl_show wname %v   \n" , wname  )  // デバック

    var idmy int64

///
///     テンプレートのヘッダーをGET
///

     monitor := template.Must(template.New("html").Parse(html4.Pipe_line_st_wl_keyin))

///
///     データストアーから、表示用データをGET
///

       water_line_view_minor , _ := storage3.Storage_tokura( "Water_Line" ,"trans" ,wname , idmy , w , r  )

//     water_line_view := trans4.Water_line ( wname , w ,r )

       water_line_view, _ := water_line_view_minor.([]type4.Water_Line)  // インターフェイス型を型変換

//     fmt.Fprintf( w, "process2.pipe_line_st_wl_show : len(water_line_view) %v\n", len(water_line_view) )  // デバック

///
///     モニターに表示
///

	err := monitor.Execute(w, water_line_view)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

