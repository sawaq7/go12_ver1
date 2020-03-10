package pipe_line_st_show

import (

	    "net/http"
	    "client/tokura/suiri/process2"
//	    "client/tokura/storage3/put1"
	    "client/tokura/storage3"
	    "client/tokura/suiri/type4"
	    "strconv"
	    "storage2"
//	    "fmt"

                                                  )

///
///     sky 水路データをファイルに追加 (ストレッジ）
///

func Pipe_line_st_show(w http.ResponseWriter, r *http.Request) {

    var water2 type4.Water2

    var idmy int64

    new_flag := 1

    bucket := "sample-7777"

///
/// key-in データをGET
///

	water2.Name = r.FormValue("water_name")  // 水路名をゲット

	water_high := r.FormValue("water_high")      // 水路高をゲット
	water2.High,_ =strconv.ParseFloat(water_high,64)  //　float64　に変換

	r_facter := r.FormValue("r_facter")      // 粗粒係数をゲット
	water2.Roughness_Factor,_ =strconv.ParseFloat(r_facter,64)  //　float64　に変換

//	fmt.Fprintf( w, "pipe_line_ds_keyin : water2.Name %v\n", water2.Name )  // デバック
//	fmt.Fprintf( w, "pipe_line_ds_keyin : water2.High %v\n", water2.High )  // デバック

///
///             Water2　ファイルがあるかチェック
///

    objects_minor , _ := storage2.Storage_basic( "list2" ,bucket ,idmy, w , r  )

    objects, _ := objects_minor.([]string)  // インターフェイス型を型変換

//    objects :=  storage2.Object_List ( w  ,r , bucket )  // バケット内のオブジェクトをゲットする

    for _ , objectsw := range objects {

      if objectsw == "Water2.txt" {

	    new_flag = 0

	  }

    }

//    fmt.Fprintf(w, "process2.pipe_line_ds_keyin : new_flag %v\n", new_flag )  // デバック

///
///         ストレッジにデータをセット
///

    if new_flag == 0 {

      _ , _ = storage3.Storage_tokura( "Water2" ,"put" ,water2 , idmy , w , r  )

//      put1.Water2 ( w , r ,water2 )   //  ファイルに追加

    } else {                          //  ファイルを新規作成

      _ , _ = storage3.Storage_tokura( "Water2" ,"put2" ,water2 , idmy , w , r  )

//	  put1.Water2_new ( w , r ,water2 )

	}

///
///           モニター表示
///

   process2.Pipe_line_st_show(w , r )

}

