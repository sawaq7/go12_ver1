package pipe_line_st_cal

import (

	    "net/http"
	    "github.com/sawaq7/go12_ver1/client/tokura/suiri/cal"
	    "github.com/sawaq7/go12_ver1/client/tokura/suiri/type4"
//	    "github.com/sawaq7/go12_ver1/client/tokura/storage3/trans4"
	    "github.com/sawaq7/go12_ver1/client/tokura/storage3"
//	    "strconv"
//	    "fmt"

                                                  )

func Pipe_line_st_cal(w http.ResponseWriter, r *http.Request) {

//   fmt.Fprintf( w, "sky/pipe_line_st_cal start \n"  )  // チE��チE��

   var water type4.Water2

   var idmy ,idmy2 int64

///
///       カレント�E水路惁E��をゲチE��
///

   water2_temp_minor , _ := storage3.Storage_tokura( "Water2_Temp" ,"trans" ,idmy , idmy2 , w , r  )

   //    water2_temp := trans4.Water2_temp( w , r  )

   water2_temp, _ := water2_temp_minor.([]type4.Water2_Temp)  // インターフェイス型を型変換



    water.Name = water2_temp[0].Name
	water.High = water2_temp[0].High
	water.Roughness_Factor = water2_temp[0].Roughness_Factor

///
///    水路ラインのチE�EタをゲチE��
///
      water_line_minor , _ := storage3.Storage_tokura( "Water_Line" ,"trans" ,water.Name , idmy , w , r  )
//    water_line := trans4.Water_line ( water.Name , w ,r )

      water_line, _ := water_line_minor.([]type4.Water_Line)  // インターフェイス型を型変換

///
///       動水勾配線�E計箁E///

    p_number ,ad_eneup ,ad_enedown ,ad_glineup ,ad_glinedown := cal.Pipe_line1( water  ,water_line  )

//    fmt.Fprintf( w, "sky/pipe_line_st_cal : p_number %v\n", p_number )  // チE��チE��
//    fmt.Fprintf( w, "sky/pipe_line_st_cal : ad_eneup %v\n", ad_eneup )  // チE��チE��
//    fmt.Fprintf( w, "sky/pipe_line_st_cal : ad_enedown %v\n", ad_enedown )  // チE��チE��
//   fmt.Fprintf( w, "sky/pipe_line_st_cal : ad_glineup %v\n", ad_glineup )  // チE��チE��
//    fmt.Fprintf( w, "sky/pipe_line_st_cal : ad_glinedown %v\n", ad_glinedown )  // チE��チE��
//    fmt.Fprintf( w, "sky/pipe_line_st_cal : ad_eneup len %v\n", len(ad_eneup) )  // チE��チE��

///
///           グラフ�E作�E
///

    f_name := cal.Pipe_line1_make_graf( w ,r ,p_number ,ad_eneup ,ad_enedown ,ad_glineup ,ad_glinedown  )

//    fmt.Fprintf( w, "sky/pipe_line_st_cal : f_name %v\n", f_name )  // チE��チE��

///
///           グラフ�E表示
///

    cal.Pipe_line1_show_graf( w ,r ,f_name )

}

