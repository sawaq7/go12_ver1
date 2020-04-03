package pipe_line_st_wl_update

import (

	    "strconv"
	    "net/http"
//	    "fmt"
	    "github.com/sawaq7/go12_ver1/client/tokura/suiri/process2"
//	    "github.com/sawaq7/go12_ver1/client/tokura/storage3/trans4"
	    "github.com/sawaq7/go12_ver1/client/tokura/storage3"
//	    "github.com/sawaq7/go12_ver1/client/tokura/storage3/put1"

	    "github.com/sawaq7/go12_ver1/client/tokura/suiri/type4"
	    "github.com/sawaq7/go12_ver1/general/type5"

                                                   )

func Pipe_line_st_wl_update(w http.ResponseWriter, r *http.Request) {

//	fmt.Fprintf( w, "sky_pipe_line_st_wl_update start %v\n" )  // 繝・ヰ繝・け

	var water_line  type4.Water_Line

	var idmy1 ,idmy2 int64

///
///       繧ｫ繝ｬ繝ｳ繝医・豌ｴ霍ｯ諠・ｱ繧偵ご繝・ヨ
///
      water2_temp_minor , _ := storage3.Storage_tokura( "Water2_Temp" ,"trans" ,idmy1 , idmy2 , w , r  )

      water2_temp, _ := water2_temp_minor.([]type4.Water2_Temp)  // 繧､繝ｳ繧ｿ繝ｼ繝輔ぉ繧､繧ｹ蝙九ｒ蝙句､画鋤

//    water2_temp := trans4.Water2_temp( w , r  )

///
///    謖・ｮ壹＠縺溷推遞ｮ繝・・繧ｿ繧竪ET
///

    updidw , err := strconv.Atoi(r.FormValue("id"))

	if err  != nil {

//	   fmt.Fprintf( w, "sky_pipe_line_st_wl_update :error updidw %v\n", updidw )  // 繝・ヰ繝・け

	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
	}

    updid := int64(updidw)

    water_line.Id = updid                          // 繝ｬ繧ｳ繝ｼ繝丑O繧偵そ繝・ヨ

    water_line.Name = water2_temp[0].Name        // 豌ｴ霍ｯ蜷阪ｒ繧ｻ繝・ヨ

	water_line.Section = r.FormValue("section")  // 蛹ｺ髢灘錐繧偵ご繝・ヨ

	f_facter := r.FormValue("f_facter")                   // 鞫ｩ謫ｦ菫よ焚繧偵ご繝・ヨ
	water_line.Friction_Factor,_ =strconv.ParseFloat(f_facter,64)  //縲float64縲縺ｫ螟画鋤

	velocity := r.FormValue("velocity")                   // 騾溷ｺｦ繧偵ご繝・ヨ
	water_line.Velocity,_ =strconv.ParseFloat(velocity,64)         //縲float64縲縺ｫ螟画鋤

	p_diameter := r.FormValue("p_diameter")      // 鞫ｩ謫ｦ菫よ焚繧偵ご繝・ヨ
	water_line.Pipe_Diameter,_ =strconv.ParseFloat(p_diameter,64)  //縲float64縲縺ｫ螟画鋤

	p_length := r.FormValue("p_length")      // 鞫ｩ謫ｦ菫よ焚繧偵ご繝・ヨ
	water_line.Pipe_Length,_ =strconv.ParseFloat(p_length,64)  //縲float64縲縺ｫ螟画鋤

//	fmt.Fprintf( w, "pipe_line_st_wl_update : water_line.Name %v\n", water_line.Name )  // 繝・ヰ繝・け
//	fmt.Fprintf( w, "pipe_line_st_wl_update : water_line.Section %v\n", water_line.Section )  // 繝・ヰ繝・け
//	fmt.Fprintf( w, "pipe_line_st_wl_update : water_line.Friction_Factor %v\n", water_line.Friction_Factor )  // 繝・ヰ繝・け
//	fmt.Fprintf( w, "pipe_line_st_wl_update : water_line.Velocity %v\n", water_line.Velocity )  // 繝・ヰ繝・け
//	fmt.Fprintf( w, "pipe_line_st_wl_update : water_line.Pipe_Diameter %v\n", water_line.Pipe_Diameter )  // 繝・ヰ繝・け
//	fmt.Fprintf( w, "pipe_line_st_wl_update : water_line.Pipe_Length %v\n", water_line.Pipe_Length )  // 繝・ヰ繝・け

///
///         驟埼＃諠・ｱ縺ｮ螟画峩
///

///
///          繝・・繧ｿ繧ｹ繝医い繝ｼ縺九ｉ縲∬｡ｨ遉ｺ逕ｨ繝・・繧ｿ繧竪ET
///

     general_work := make([]type5.General_Work, 1)
     general_work[0].Int64_Work  = updid                       // 繧｢繝・・繝・・繝・d
     general_work[0].String_Work = water2_temp[0].Name         //縲豌ｴ霍ｯ蜷・
/// test 逕ｨ縲縺ｮ繧ｳ繝ｼ繝峨start

//    var struct_colle    type4.Struct_Colle

//	water_line2 := make([]type4.Water_Line, 2)

//	water_line2[0].Name             = "water1"
//	water_line2[0].Friction_Factor  = 1.1
//	water_line2[1].Name             = "water2"
//	water_line2[1].Friction_Factor  = 2.2

//	water2 := make([]type4.Water2, 2)

//	water2[0].Name             = "water1"
//	water2[0].High             = 100.1
//	water2[1].Name             = "water2"
//	water2[1].High             = 200.2

//    struct_colle.Water_Line_Slice = water_line2
//    struct_colle.Water2_Slice = water2

//     flexible_out1 , flexible_out2 := storage3.Storage_tokura( "Water_Line" ,"put_test" ,general_work , struct_colle , w , r  )

     // 遨ｺ繧､繝ｳ繧ｿ繝ｼ繝輔ぉ繧､繧ｹ螟画焚繧医ｊ繝舌Μ繝･繝ｼ蛟､繧偵ご繝・ヨ

//     flexible_out1_value, _ := flexible_out1.([]type5.General_Work)
//     flexible_out2_value, _ := flexible_out2.(type4.Struct_Colle)

//     fmt.Fprintf( w, "storage_tokura : flexible_out1_value %v\n", flexible_out1_value )
//     fmt.Fprintf( w, "storage_tokura : flexible_out2_value.Water_Line_Slice %v\n", flexible_out2_value.Water_Line_Slice )
//     fmt.Fprintf( w, "storage_tokura : flexible_out2_value.Water2_Slice %v\n", flexible_out2_value.Water2_Slice )

/// test 逕ｨ縲縺ｮ繧ｳ繝ｼ繝峨end

     _ , _ = storage3.Storage_tokura( "Water_Line" ,"put2" ,general_work , water_line , w , r  )

//	put1.Water_line_update ( w , r ,updid ,water2_temp[0].Name ,water_line )

///
///        繝｢繝九ち繝ｼ縲蜀崎｡ｨ遉ｺ
///

	process2.Pipe_line_st_wl_show ( water2_temp[0].Name ,w , r )

}
