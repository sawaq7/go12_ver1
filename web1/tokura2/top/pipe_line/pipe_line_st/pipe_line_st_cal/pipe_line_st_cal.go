package pipe_line_st_cal

import (

	    "net/http"
	    "client/tokura/suiri/cal"
	    "client/tokura/suiri/type4"
//	    "client/tokura/storage3/trans4"
	    "client/tokura/storage3"
//	    "strconv"
//	    "fmt"

                                                  )

func Pipe_line_st_cal(w http.ResponseWriter, r *http.Request) {

//   fmt.Fprintf( w, "sky/pipe_line_st_cal start \n"  )  // デバック

   var water type4.Water2

   var idmy ,idmy2 int64

///
///       カレントの水路情報をゲット
///

   water2_temp_minor , _ := storage3.Storage_tokura( "Water2_Temp" ,"trans" ,idmy , idmy2 , w , r  )

   //    water2_temp := trans4.Water2_temp( w , r  )

   water2_temp, _ := water2_temp_minor.([]type4.Water2_Temp)  // インターフェイス型を型変換



    water.Name = water2_temp[0].Name
	water.High = water2_temp[0].High
	water.Roughness_Factor = water2_temp[0].Roughness_Factor

///
///    水路ラインのデータをゲット
///
      water_line_minor , _ := storage3.Storage_tokura( "Water_Line" ,"trans" ,water.Name , idmy , w , r  )
//    water_line := trans4.Water_line ( water.Name , w ,r )

      water_line, _ := water_line_minor.([]type4.Water_Line)  // インターフェイス型を型変換

///
///       動水勾配線の計算
///

    p_number ,ad_eneup ,ad_enedown ,ad_glineup ,ad_glinedown := cal.Pipe_line1( water  ,water_line  )

//    fmt.Fprintf( w, "sky/pipe_line_st_cal : p_number %v\n", p_number )  // デバック
//    fmt.Fprintf( w, "sky/pipe_line_st_cal : ad_eneup %v\n", ad_eneup )  // デバック
//    fmt.Fprintf( w, "sky/pipe_line_st_cal : ad_enedown %v\n", ad_enedown )  // デバック
//   fmt.Fprintf( w, "sky/pipe_line_st_cal : ad_glineup %v\n", ad_glineup )  // デバック
//    fmt.Fprintf( w, "sky/pipe_line_st_cal : ad_glinedown %v\n", ad_glinedown )  // デバック
//    fmt.Fprintf( w, "sky/pipe_line_st_cal : ad_eneup len %v\n", len(ad_eneup) )  // デバック

///
///           グラフの作成
///

    f_name := cal.Pipe_line1_make_graf( w ,r ,p_number ,ad_eneup ,ad_enedown ,ad_glineup ,ad_glinedown  )

//    fmt.Fprintf( w, "sky/pipe_line_st_cal : f_name %v\n", f_name )  // デバック

///
///           グラフの表示
///

    cal.Pipe_line1_show_graf( w ,r ,f_name )

}

