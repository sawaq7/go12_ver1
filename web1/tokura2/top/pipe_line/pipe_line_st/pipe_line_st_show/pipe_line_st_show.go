package pipe_line_st_show

import (

	    "net/http"
	    "github.com/sawaq7/go12_ver1/client/tokura/suiri/process2"
//	    "github.com/sawaq7/go12_ver1/client/tokura/storage3/put1"
	    "github.com/sawaq7/go12_ver1/client/tokura/storage3"
	    "github.com/sawaq7/go12_ver1/client/tokura/suiri/type4"
	    "strconv"
	    "storage2"
//	    "fmt"

                                                  )

///
///     sky 豌ｴ霍ｯ繝・・繧ｿ繧偵ヵ繧｡繧､繝ｫ縺ｫ霑ｽ蜉 (繧ｹ繝医Ξ繝・ず・・///

func Pipe_line_st_show(w http.ResponseWriter, r *http.Request) {

    var water2 type4.Water2

    var idmy int64

    new_flag := 1

    bucket := "sample-7777"

///
/// key-in 繝・・繧ｿ繧竪ET
///

	water2.Name = r.FormValue("water_name")  // 豌ｴ霍ｯ蜷阪ｒ繧ｲ繝・ヨ

	water_high := r.FormValue("water_high")      // 豌ｴ霍ｯ鬮倥ｒ繧ｲ繝・ヨ
	water2.High,_ =strconv.ParseFloat(water_high,64)  //縲float64縲縺ｫ螟画鋤

	r_facter := r.FormValue("r_facter")      // 邊礼ｲ剃ｿよ焚繧偵ご繝・ヨ
	water2.Roughness_Factor,_ =strconv.ParseFloat(r_facter,64)  //縲float64縲縺ｫ螟画鋤

//	fmt.Fprintf( w, "pipe_line_ds_keyin : water2.Name %v\n", water2.Name )  // 繝・ヰ繝・け
//	fmt.Fprintf( w, "pipe_line_ds_keyin : water2.High %v\n", water2.High )  // 繝・ヰ繝・け

///
///             Water2縲繝輔ぃ繧､繝ｫ縺後≠繧九°繝√ぉ繝・け
///

    objects_minor , _ := storage2.Storage_basic( "list2" ,bucket ,idmy, w , r  )

    objects, _ := objects_minor.([]string)  // 繧､繝ｳ繧ｿ繝ｼ繝輔ぉ繧､繧ｹ蝙九ｒ蝙句､画鋤

//    objects :=  storage2.Object_List ( w  ,r , bucket )  // 繝舌こ繝・ヨ蜀・・繧ｪ繝悶ず繧ｧ繧ｯ繝医ｒ繧ｲ繝・ヨ縺吶ｋ

    for _ , objectsw := range objects {

      if objectsw == "Water2.txt" {

	    new_flag = 0

	  }

    }

//    fmt.Fprintf(w, "process2.pipe_line_ds_keyin : new_flag %v\n", new_flag )  // 繝・ヰ繝・け

///
///         繧ｹ繝医Ξ繝・ず縺ｫ繝・・繧ｿ繧偵そ繝・ヨ
///

    if new_flag == 0 {

      _ , _ = storage3.Storage_tokura( "Water2" ,"put" ,water2 , idmy , w , r  )

//      put1.Water2 ( w , r ,water2 )   //  繝輔ぃ繧､繝ｫ縺ｫ霑ｽ蜉

    } else {                          //  繝輔ぃ繧､繝ｫ繧呈眠隕丈ｽ懈・

      _ , _ = storage3.Storage_tokura( "Water2" ,"put2" ,water2 , idmy , w , r  )

//	  put1.Water2_new ( w , r ,water2 )

	}

///
///           繝｢繝九ち繝ｼ陦ｨ遉ｺ
///

   process2.Pipe_line_st_show(w , r )

}

