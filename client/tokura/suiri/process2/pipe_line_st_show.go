package process2

import (

	    "net/http"
//	    "fmt"
	    "html/template"
	    "github.com/sawaq7/go12_ver1/client/tokura/html4"
        "github.com/sawaq7/go12_ver1/client/tokura/suiri/type4"
	    "github.com/sawaq7/go12_ver1/client/tokura/storage3"
	    "storage2"
                                                )
func Pipe_line_st_show(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "process2.pipe_line_st_show start \n" )  // 繝・ヰ繝・け

    var idmy1 ,idmy2 int64

    skip_flag := 1

    bucket := "sample-7777"

///
///             Water2縲繝輔ぃ繧､繝ｫ縺後≠繧九°繝√ぉ繝・け
///

    objects :=  storage2.Object_List ( w  ,r , bucket )  // 繝舌こ繝・ヨ蜀・・繧ｪ繝悶ず繧ｧ繧ｯ繝医ｒ繧ｲ繝・ヨ縺吶ｋ

    for _ , objectsw := range objects {

      if objectsw == "Water2.txt" {

	    skip_flag = 0

	  }

    }

//    fmt.Fprintf(w, "process2.pipe_line_st_show : skip_flag %v\n", skip_flag )  // 繝・ヰ繝・け

///
///            縲∬｡ｨ遉ｺ逕ｨ繝・・繧ｿ繝ｻ繝・Φ繝励Ξ繝ｼ繝医・繝倥ャ繝繝ｼ繧竪ET縺励※陦ｨ遉ｺ
///

     monitor := template.Must(template.New("html").Parse(html4.Pipe_line_st_keyin))

     if skip_flag == 0 {

     water2_view_minor , _ := storage3.Storage_tokura( "Water2" ,"trans" ,idmy1 , idmy2 , w , r  )

//       water2_view := trans4.Water2 ( w ,r )

       water2_view, _ := water2_view_minor.([]type4.Water2)  // 繧､繝ｳ繧ｿ繝ｼ繝輔ぉ繧､繧ｹ蝙九ｒ蝙句､画鋤

       err := monitor.Execute(w, water2_view)

       if err != nil {

	    http.Error(w, err.Error(), http.StatusInternalServerError)

	   }

     } else {

       water2_view := make([]type4.Water2, 0)   //   Water2縲縺ｮ陦ｨ遉ｺ繧ｨ繝ｪ繧｢繧堤｢ｺ菫・
       err := monitor.Execute(w, water2_view)

       if err != nil {

	    http.Error(w, err.Error(), http.StatusInternalServerError)

	   }

     }

}

