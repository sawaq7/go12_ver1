package process2

import (

	    "net/http"
//	    "fmt"
	    "html/template"
	    "client/tokura/html4"

	    "client/tokura/datastore4"
	    "client/tokura/suiri/type4"
                                                )

///
///       導水勾配線リストを表示する。
///

func Water_slope_show(w http.ResponseWriter, r *http.Request) {

//     IN     w         : レスポンスライター
//     IN     r         : リクエストパラメーター

//    fmt.Fprintf( w, "process2.water_slope_show start \n" )  // デバック

    var idmy int64

///
///     テンプレートのヘッダーをGET
///

     monitor := template.Must(template.New("html").Parse(html4.Water_slope_show))

///
///      データストアーから、表示用データをGET
///

//     water_slope_view := trans2.Water_slope ( w ,r )

     water_slope_view := datastore4.Datastore_tokura( "Water_Slope"  ,"trans"  ,idmy , w , r  )

     value, _ := water_slope_view.([]type4.Water_Slope)    // 空インターフェイス変数よりバリュー値をゲット

///
///       モニターに表示
///

	err := monitor.Execute( w, value )
	if err != nil {
	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
	}
}

