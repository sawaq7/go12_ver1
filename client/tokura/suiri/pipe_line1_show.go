package suiri

import (
	    "net/http"
//	    "fmt"
	    "html/template"
	    "client/tokura/suiri/type4"

        "client/tokura/html4"
                                      )


func Pipe_line1_show( w http.ResponseWriter , wnumber int , water []type4.Water  ) {

//     IN    w     : レスポンスライター
//     IN  wnumbwr : ポイント損失のスライス
//     IN  water   : 水路情報群のスライス



//    fmt.Fprintf( w, "suiri/Pipe_line1_show start \n" )  // デバック
//    fmt.Fprintf( w, "suiri/Pipe_line1_show wnumber  %v \n" , wnumber)  // デバック

    water2 := make([]type4.Water, wnumber)

    for pos, waterwk := range water {

			water2[pos].No  = waterwk.No
			water2[pos].Name = waterwk.Name
			water2[pos].High = waterwk.High
            water2[pos].Roughness_factor = waterwk.Roughness_factor
//			fmt.Fprintf( w, "suiri/Pipe_line1_show pos ,NO  ,Name %v %v %v \n" , pos ,water2[pos].No ,water2[pos].Name)  // デバック
//			fmt.Fprintf( w, "suiri/Pipe_line1_show name High ,Roughness_factor %v %v  \n" , water2[pos].High ,water2[pos].Roughness_factor)  // デバック
	}
// テンプレートのヘッダーをGET

	monitor := template.Must(template.New("html").Parse( html4.Pipe_line1_show))
//	monitor := template.Must(template.New("html").Parse( html4.Pipe_line1_show2))



// 各種入力データを表示

	err := monitor.Execute(w, water2)

//    var str_dmy string
//	err := monitor.Execute(w, water)

    if err != nil {
	   http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

