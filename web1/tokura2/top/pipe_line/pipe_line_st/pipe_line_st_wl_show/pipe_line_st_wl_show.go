package pipe_line_st_wl_show

import (

	    "net/http"
	    "client/tokura/suiri/process2"
	    "client/tokura/suiri/type4"
	    "strconv"
//	    "fmt"

//	    "client/tokura/storage3/trans4"

//	    "client/tokura/storage3/check4"
	    "client/tokura/storage3"
                                                  )

///
///         入力した水路ラインのデータを表示する
///

func Pipe_line_st_wl_show(w http.ResponseWriter, r *http.Request) {

//   fmt.Fprintf( w, "sky/pipe_line_st_wl_show start \n"  )  // デバック

/// key-in データをGET ///

   var water_line type4.Water_Line

   var idmy ,idmy2 int64

///
///         テンポラリーファイルより、水路名をゲット
///
     water2_temp_minor , _ := storage3.Storage_tokura( "Water2_Temp" ,"trans" ,idmy , idmy2 , w , r  )

     water2_temp, _ := water2_temp_minor.([]type4.Water2_Temp)  // インターフェイス型を型変換

//   water2_temp := trans4.Water2_temp ( w ,r )

///
///         指示した水路の既存の水路ラインの数をゲット
///

    record_number_temp , _ := storage3.Storage_tokura( "Water_Line" ,"check" ,water2_temp[0].Name , idmy , w , r  )

    record_number, _ := record_number_temp.(int64)

//    record_number := check4.Water_line_re_num( water2_temp[0].Name  ,w , r  )

	for _, water2_tempw := range water2_temp {

       water_line.Name = water2_tempw.Name          /// 水路名のセット
       water_line.Id   = record_number + int64( 1 ) /// idのセット

    }

	water_line.Section = r.FormValue("section")  // 区間名をゲット

	f_facter := r.FormValue("f_facter")                   // 摩擦係数をゲット
	water_line.Friction_Factor,_ =strconv.ParseFloat(f_facter,64)  //　float64　に変換

	velocity := r.FormValue("velocity")                   // 速度をゲット
	water_line.Velocity,_ =strconv.ParseFloat(velocity,64)         //　float64　に変換

	p_diameter := r.FormValue("p_diameter")      // 摩擦係数をゲット
	water_line.Pipe_Diameter,_ =strconv.ParseFloat(p_diameter,64)  //　float64　に変換

	p_length := r.FormValue("p_length")      // 摩擦係数をゲット
	water_line.Pipe_Length,_ =strconv.ParseFloat(p_length,64)  //　float64　に変換

//	fmt.Fprintf( w, "sky/pipe_line_st_wl_show : water_line.Name %v\n", water_line.Name )  // デバック
//	fmt.Fprintf( w, "sky/pipe_line_st_wl_show : water_line.Section %v\n", water_line.Section )  // デバック
//	fmt.Fprintf( w, "sky/pipe_line_st_wl_show : water_line.Friction_Factor %v\n", water_line.Friction_Factor )  // デバック
//	fmt.Fprintf( w, "sky/pipe_line_st_wl_show : water_line.Velocity %v\n", water_line.Velocity )  // デバック
//	fmt.Fprintf( w, "sky/pipe_line_st_wl_show : water_line.Pipe_Diameter %v\n", water_line.Pipe_Diameter )  // デバック
//	fmt.Fprintf( w, "sky/pipe_line_st_wl_show : water_line.Pipe_Length %v\n", water_line.Pipe_Length )  // デバック

///                           　　　　　　　　　　　
///   ストレッジファイルに水路ラインファイル情報を書く
///                          　　　　　　　　　　　

    _ , _ = storage3.Storage_tokura( "Water_Line" ,"put" ,water_line , idmy , w , r  )

//   put1.Water_line( w , r ,water_line )

///
///             モニター表示
///

   process2.Pipe_line_st_wl_show ( water_line.Name ,w , r )


}

