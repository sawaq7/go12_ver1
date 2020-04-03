package process2

import (

	    "net/http"
//	    "fmt"
	    "html/template"
	    "github.com/sawaq7/go12_ver1/client/tokura/html4"

	    "github.com/sawaq7/go12_ver1/client/tokura/datastore4"
	    "github.com/sawaq7/go12_ver1/client/tokura/suiri/type4"
                                                )

///
///       導水勾配線リストを表示する、E///

func Water_slope_show(w http.ResponseWriter, r *http.Request) {

//     IN     w         : レスポンスライター
//     IN     r         : リクエストパラメーター

//    fmt.Fprintf( w, "process2.water_slope_show start \n" )  // チE��チE��

    var idmy int64

///
///     チE��プレート�EヘッダーをGET
///

     monitor := template.Must(template.New("html").Parse(html4.Water_slope_show))

///
///      チE�Eタストアーから、表示用チE�EタをGET
///

//     water_slope_view := trans2.Water_slope ( w ,r )

     water_slope_view := datastore4.Datastore_tokura( "Water_Slope"  ,"trans"  ,idmy , w , r  )

     value, _ := water_slope_view.([]type4.Water_Slope)    // 空インターフェイス変数よりバリュー値をゲチE��

///
///       モニターに表示
///

	err := monitor.Execute( w, value )
	if err != nil {
	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
	}
}

