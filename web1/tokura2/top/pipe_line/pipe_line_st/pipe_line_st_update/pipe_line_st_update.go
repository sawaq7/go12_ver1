package pipe_line_st_update

import (

	    "net/http"
	    "client/tokura/suiri/process2"
//	    "client/tokura/storage3/put1"
	    "client/tokura/storage3"
	    "client/tokura/suiri/type4"
	    "strconv"
//	    "storage2"
//	    "fmt"

                                                  )

func Pipe_line_st_update(w http.ResponseWriter, r *http.Request) {

    var water2 type4.Water2

///
///          key-in データをGET
///

    updidw , err := strconv.Atoi(r.FormValue("id"))     // idをゲット
	if err  != nil {

//	   fmt.Fprintf( w, "sky_pipe_line_ds_update :error updidw %v\n", updidw )  // デバック

	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
	}
    updid := int64(updidw)

    water2.Id = updid                      // レコードidセット

	water2.Name = r.FormValue("water_name")  // 水路名をゲット

	water_high := r.FormValue("water_high")      // 水路高をゲット
	water2.High,_ =strconv.ParseFloat(water_high,64)  //　float64　に変換

	r_facter := r.FormValue("r_facter")      // 粗粒係数をゲット
	water2.Roughness_Factor,_ =strconv.ParseFloat(r_facter,64)  //　float64　に変換

//	fmt.Fprintf( w, "sky.pipe_line_st_update : water2.Name %v\n", water2.Name )  // デバック
//	fmt.Fprintf( w, "sky.pipe_line_st_update : water2.High %v\n", water2.High )  // デ

///
///         ストレッジにアップデートデータをセット
///

    _ , _ = storage3.Storage_tokura( "Water2" ,"put3" ,updid , water2 , w , r  )

//	put1.Water2_update ( w , r ,updid ,water2 )

///
///           モニター表示
///

   process2.Pipe_line_st_show(w , r )

}

