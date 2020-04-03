package process2

import (

	    "net/http"
//	    "fmt"
	    "html/template"
	    "github.com/sawaq7/go12_ver1/client/tokura/html4"

//	    "github.com/sawaq7/go12_ver1/client/tokura/storage3/trans4"
	    "github.com/sawaq7/go12_ver1/client/tokura/storage3"
	    "github.com/sawaq7/go12_ver1/client/tokura/suiri/type4"
//	    "time"
                                                )

func Pipe_line_st_wl_show( wname string ,w http.ResponseWriter, r *http.Request) {

//     IN  wname : 水路吁E　　　　
//     IN     w  : レスポンスライター
//     IN     r  : リクエストパラメーター


//    fmt.Fprintf( w, "process2.pipe_line_st_wl_show start \n" )  // チE��チE��

//    fmt.Fprintf( w, "process2.pipe_line_st_wl_show wname %v   \n" , wname  )  // チE��チE��

    var idmy int64

///
///     チE��プレート�EヘッダーをGET
///

     monitor := template.Must(template.New("html").Parse(html4.Pipe_line_st_wl_keyin))

///
///     チE�Eタストアーから、表示用チE�EタをGET
///

       water_line_view_minor , _ := storage3.Storage_tokura( "Water_Line" ,"trans" ,wname , idmy , w , r  )

//     water_line_view := trans4.Water_line ( wname , w ,r )

       water_line_view, _ := water_line_view_minor.([]type4.Water_Line)  // インターフェイス型を型変換

//     fmt.Fprintf( w, "process2.pipe_line_st_wl_show : len(water_line_view) %v\n", len(water_line_view) )  // チE��チE��

///
///     モニターに表示
///

	err := monitor.Execute(w, water_line_view)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

