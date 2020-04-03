package pipe_line_st_update

import (

	    "net/http"
	    "github.com/sawaq7/go12_ver1/client/tokura/suiri/process2"
//	    "github.com/sawaq7/go12_ver1/client/tokura/storage3/put1"
	    "github.com/sawaq7/go12_ver1/client/tokura/storage3"
	    "github.com/sawaq7/go12_ver1/client/tokura/suiri/type4"
	    "strconv"
//	    "storage2"
//	    "fmt"

                                                  )

func Pipe_line_st_update(w http.ResponseWriter, r *http.Request) {

    var water2 type4.Water2

///
///          key-in 繝・・繧ｿ繧竪ET
///

    updidw , err := strconv.Atoi(r.FormValue("id"))     // id繧偵ご繝・ヨ
	if err  != nil {

//	   fmt.Fprintf( w, "sky_pipe_line_ds_update :error updidw %v\n", updidw )  // 繝・ヰ繝・け

	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
	}
    updid := int64(updidw)

    water2.Id = updid                      // 繝ｬ繧ｳ繝ｼ繝永d繧ｻ繝・ヨ

	water2.Name = r.FormValue("water_name")  // 豌ｴ霍ｯ蜷阪ｒ繧ｲ繝・ヨ

	water_high := r.FormValue("water_high")      // 豌ｴ霍ｯ鬮倥ｒ繧ｲ繝・ヨ
	water2.High,_ =strconv.ParseFloat(water_high,64)  //縲float64縲縺ｫ螟画鋤

	r_facter := r.FormValue("r_facter")      // 邊礼ｲ剃ｿよ焚繧偵ご繝・ヨ
	water2.Roughness_Factor,_ =strconv.ParseFloat(r_facter,64)  //縲float64縲縺ｫ螟画鋤

//	fmt.Fprintf( w, "sky.pipe_line_st_update : water2.Name %v\n", water2.Name )  // 繝・ヰ繝・け
//	fmt.Fprintf( w, "sky.pipe_line_st_update : water2.High %v\n", water2.High )  // 繝・
///
///         繧ｹ繝医Ξ繝・ず縺ｫ繧｢繝・・繝・・繝医ョ繝ｼ繧ｿ繧偵そ繝・ヨ
///

    _ , _ = storage3.Storage_tokura( "Water2" ,"put3" ,updid , water2 , w , r  )

//	put1.Water2_update ( w , r ,updid ,water2 )

///
///           繝｢繝九ち繝ｼ陦ｨ遉ｺ
///

   process2.Pipe_line_st_show(w , r )

}

