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

//   fmt.Fprintf( w, "sky/pipe_line_st_cal start \n"  )  // 繝・ヰ繝・け

   var water type4.Water2

   var idmy ,idmy2 int64

///
///       繧ｫ繝ｬ繝ｳ繝医・豌ｴ霍ｯ諠・ｱ繧偵ご繝・ヨ
///

   water2_temp_minor , _ := storage3.Storage_tokura( "Water2_Temp" ,"trans" ,idmy , idmy2 , w , r  )

   //    water2_temp := trans4.Water2_temp( w , r  )

   water2_temp, _ := water2_temp_minor.([]type4.Water2_Temp)  // 繧､繝ｳ繧ｿ繝ｼ繝輔ぉ繧､繧ｹ蝙九ｒ蝙句､画鋤



    water.Name = water2_temp[0].Name
	water.High = water2_temp[0].High
	water.Roughness_Factor = water2_temp[0].Roughness_Factor

///
///    豌ｴ霍ｯ繝ｩ繧､繝ｳ縺ｮ繝・・繧ｿ繧偵ご繝・ヨ
///
      water_line_minor , _ := storage3.Storage_tokura( "Water_Line" ,"trans" ,water.Name , idmy , w , r  )
//    water_line := trans4.Water_line ( water.Name , w ,r )

      water_line, _ := water_line_minor.([]type4.Water_Line)  // 繧､繝ｳ繧ｿ繝ｼ繝輔ぉ繧､繧ｹ蝙九ｒ蝙句､画鋤

///
///       蜍墓ｰｴ蜍ｾ驟咲ｷ壹・險育ｮ・///

    p_number ,ad_eneup ,ad_enedown ,ad_glineup ,ad_glinedown := cal.Pipe_line1( water  ,water_line  )

//    fmt.Fprintf( w, "sky/pipe_line_st_cal : p_number %v\n", p_number )  // 繝・ヰ繝・け
//    fmt.Fprintf( w, "sky/pipe_line_st_cal : ad_eneup %v\n", ad_eneup )  // 繝・ヰ繝・け
//    fmt.Fprintf( w, "sky/pipe_line_st_cal : ad_enedown %v\n", ad_enedown )  // 繝・ヰ繝・け
//   fmt.Fprintf( w, "sky/pipe_line_st_cal : ad_glineup %v\n", ad_glineup )  // 繝・ヰ繝・け
//    fmt.Fprintf( w, "sky/pipe_line_st_cal : ad_glinedown %v\n", ad_glinedown )  // 繝・ヰ繝・け
//    fmt.Fprintf( w, "sky/pipe_line_st_cal : ad_eneup len %v\n", len(ad_eneup) )  // 繝・ヰ繝・け

///
///           繧ｰ繝ｩ繝輔・菴懈・
///

    f_name := cal.Pipe_line1_make_graf( w ,r ,p_number ,ad_eneup ,ad_enedown ,ad_glineup ,ad_glinedown  )

//    fmt.Fprintf( w, "sky/pipe_line_st_cal : f_name %v\n", f_name )  // 繝・ヰ繝・け

///
///           繧ｰ繝ｩ繝輔・陦ｨ遉ｺ
///

    cal.Pipe_line1_show_graf( w ,r ,f_name )

}

