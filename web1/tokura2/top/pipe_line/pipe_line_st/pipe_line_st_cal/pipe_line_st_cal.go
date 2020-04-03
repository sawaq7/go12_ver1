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

//   fmt.Fprintf( w, "sky/pipe_line_st_cal start \n"  )  // ćEćEÆ

   var water type4.Water2

   var idmy ,idmy2 int64

///
///       ć«ć¬ć³ććEę°“č·ÆęE ±ćć²ćE
///

   water2_temp_minor , _ := storage3.Storage_tokura( "Water2_Temp" ,"trans" ,idmy , idmy2 , w , r  )

   //    water2_temp := trans4.Water2_temp( w , r  )

   water2_temp, _ := water2_temp_minor.([]type4.Water2_Temp)  // ć¤ć³ćæć¼ćć§ć¤ć¹åćåå¤ę



    water.Name = water2_temp[0].Name
	water.High = water2_temp[0].High
	water.Roughness_Factor = water2_temp[0].Roughness_Factor

///
///    ę°“č·Æć©ć¤ć³ć®ćEEćæćć²ćE
///
      water_line_minor , _ := storage3.Storage_tokura( "Water_Line" ,"trans" ,water.Name , idmy , w , r  )
//    water_line := trans4.Water_line ( water.Name , w ,r )

      water_line, _ := water_line_minor.([]type4.Water_Line)  // ć¤ć³ćæć¼ćć§ć¤ć¹åćåå¤ę

///
///       åę°“å¾éē·ćEčØē®E///

    p_number ,ad_eneup ,ad_enedown ,ad_glineup ,ad_glinedown := cal.Pipe_line1( water  ,water_line  )

//    fmt.Fprintf( w, "sky/pipe_line_st_cal : p_number %v\n", p_number )  // ćEćEÆ
//    fmt.Fprintf( w, "sky/pipe_line_st_cal : ad_eneup %v\n", ad_eneup )  // ćEćEÆ
//    fmt.Fprintf( w, "sky/pipe_line_st_cal : ad_enedown %v\n", ad_enedown )  // ćEćEÆ
//   fmt.Fprintf( w, "sky/pipe_line_st_cal : ad_glineup %v\n", ad_glineup )  // ćEćEÆ
//    fmt.Fprintf( w, "sky/pipe_line_st_cal : ad_glinedown %v\n", ad_glinedown )  // ćEćEÆ
//    fmt.Fprintf( w, "sky/pipe_line_st_cal : ad_eneup len %v\n", len(ad_eneup) )  // ćEćEÆ

///
///           ć°ć©ććEä½ęE
///

    f_name := cal.Pipe_line1_make_graf( w ,r ,p_number ,ad_eneup ,ad_enedown ,ad_glineup ,ad_glinedown  )

//    fmt.Fprintf( w, "sky/pipe_line_st_cal : f_name %v\n", f_name )  // ćEćEÆ

///
///           ć°ć©ććEč”Øē¤ŗ
///

    cal.Pipe_line1_show_graf( w ,r ,f_name )

}

