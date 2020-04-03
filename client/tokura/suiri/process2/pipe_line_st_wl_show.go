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

//     IN  wname : 豌ｴ霍ｯ蜷・縲縲縲縲
//     IN     w  : 繝ｬ繧ｹ繝昴Φ繧ｹ繝ｩ繧､繧ｿ繝ｼ
//     IN     r  : 繝ｪ繧ｯ繧ｨ繧ｹ繝医ヱ繝ｩ繝｡繝ｼ繧ｿ繝ｼ


//    fmt.Fprintf( w, "process2.pipe_line_st_wl_show start \n" )  // 繝・ヰ繝・け

//    fmt.Fprintf( w, "process2.pipe_line_st_wl_show wname %v   \n" , wname  )  // 繝・ヰ繝・け

    var idmy int64

///
///     繝・Φ繝励Ξ繝ｼ繝医・繝倥ャ繝繝ｼ繧竪ET
///

     monitor := template.Must(template.New("html").Parse(html4.Pipe_line_st_wl_keyin))

///
///     繝・・繧ｿ繧ｹ繝医い繝ｼ縺九ｉ縲∬｡ｨ遉ｺ逕ｨ繝・・繧ｿ繧竪ET
///

       water_line_view_minor , _ := storage3.Storage_tokura( "Water_Line" ,"trans" ,wname , idmy , w , r  )

//     water_line_view := trans4.Water_line ( wname , w ,r )

       water_line_view, _ := water_line_view_minor.([]type4.Water_Line)  // 繧､繝ｳ繧ｿ繝ｼ繝輔ぉ繧､繧ｹ蝙九ｒ蝙句､画鋤

//     fmt.Fprintf( w, "process2.pipe_line_st_wl_show : len(water_line_view) %v\n", len(water_line_view) )  // 繝・ヰ繝・け

///
///     繝｢繝九ち繝ｼ縺ｫ陦ｨ遉ｺ
///

	err := monitor.Execute(w, water_line_view)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

