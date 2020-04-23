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
///     show water-line from water-name
///                          　　　　　　　　　　　

func Pipe_line_ds_wl_show(funct int64 ,wname string ,w http.ResponseWriter, r *http.Request) {

//     IN  funct : ファンクション　0:   show all water-line
//               　　　　　　　　　1:   show water-line which was selected
//     IN  wname :  water-name
//     IN    w   : レスポンスライター
//     IN    r   : リクエストパラメータ
//     OUT  one  : water-line slice
//     IN    w      　: レスポンスライター
//     IN    r      　: リクエストパラメータ

//    fmt.Fprintf( w, "process2.pipe_line_ds_wl_show start \n" )
//    fmt.Fprintf( w, "process2.pipe_line_ds_wl_show funct %v   \n" , funct  )
//    fmt.Fprintf( w, "process2.pipe_line_ds_wl_show wname %v   \n" , wname  )

///
///          set template
///

     monitor := template.Must(template.New("html").Parse(html4.Pipe_line_ds_wl_keyin))

///
///           get water-line-inf.
///

//     water_line_view := trans2.Water_line (funct  ,wname , w ,r )

      water_line_view := datastore4.Datastore_tokura( "Water_Line"  ,"trans"  ,wname , w , r  )



     value, _ := water_line_view.([]type4.Water_Line)

//     fmt.Fprintf( w, "process2.pipe_line_ds_wl_show : len(water_line_view) %v\n", len(water_line_view) )  // チE��チE��


///    show water-line on web

	err := monitor.Execute ( w, value )
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

