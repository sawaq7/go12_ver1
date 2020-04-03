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

//	fmt.Fprintf( w, "sky_pipe_line_st_wl_update start %v\n" )  // チE��チE��

	var water_line  type4.Water_Line

	var idmy1 ,idmy2 int64

///
///       カレント�E水路惁E��をゲチE��
///
      water2_temp_minor , _ := storage3.Storage_tokura( "Water2_Temp" ,"trans" ,idmy1 , idmy2 , w , r  )

      water2_temp, _ := water2_temp_minor.([]type4.Water2_Temp)  // インターフェイス型を型変換

//    water2_temp := trans4.Water2_temp( w , r  )

///
///    持E��した各種チE�EタをGET
///

    updidw , err := strconv.Atoi(r.FormValue("id"))

	if err  != nil {

//	   fmt.Fprintf( w, "sky_pipe_line_st_wl_update :error updidw %v\n", updidw )  // チE��チE��

	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
	}

    updid := int64(updidw)

    water_line.Id = updid                          // レコードNOをセチE��

    water_line.Name = water2_temp[0].Name        // 水路名をセチE��

	water_line.Section = r.FormValue("section")  // 区間名をゲチE��

	f_facter := r.FormValue("f_facter")                   // 摩擦係数をゲチE��
	water_line.Friction_Factor,_ =strconv.ParseFloat(f_facter,64)  //　float64　に変換

	velocity := r.FormValue("velocity")                   // 速度をゲチE��
	water_line.Velocity,_ =strconv.ParseFloat(velocity,64)         //　float64　に変換

	p_diameter := r.FormValue("p_diameter")      // 摩擦係数をゲチE��
	water_line.Pipe_Diameter,_ =strconv.ParseFloat(p_diameter,64)  //　float64　に変換

	p_length := r.FormValue("p_length")      // 摩擦係数をゲチE��
	water_line.Pipe_Length,_ =strconv.ParseFloat(p_length,64)  //　float64　に変換

//	fmt.Fprintf( w, "pipe_line_st_wl_update : water_line.Name %v\n", water_line.Name )  // チE��チE��
//	fmt.Fprintf( w, "pipe_line_st_wl_update : water_line.Section %v\n", water_line.Section )  // チE��チE��
//	fmt.Fprintf( w, "pipe_line_st_wl_update : water_line.Friction_Factor %v\n", water_line.Friction_Factor )  // チE��チE��
//	fmt.Fprintf( w, "pipe_line_st_wl_update : water_line.Velocity %v\n", water_line.Velocity )  // チE��チE��
//	fmt.Fprintf( w, "pipe_line_st_wl_update : water_line.Pipe_Diameter %v\n", water_line.Pipe_Diameter )  // チE��チE��
//	fmt.Fprintf( w, "pipe_line_st_wl_update : water_line.Pipe_Length %v\n", water_line.Pipe_Length )  // チE��チE��

///
///         配達惁E��の変更
///

///
///          チE�Eタストアーから、表示用チE�EタをGET
///

     general_work := make([]type5.General_Work, 1)
     general_work[0].Int64_Work  = updid                       // アチE�EチE�EチEd
     general_work[0].String_Work = water2_temp[0].Name         //　水路吁E
/// test 用　のコード　start

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

     // 空インターフェイス変数よりバリュー値をゲチE��

//     flexible_out1_value, _ := flexible_out1.([]type5.General_Work)
//     flexible_out2_value, _ := flexible_out2.(type4.Struct_Colle)

//     fmt.Fprintf( w, "storage_tokura : flexible_out1_value %v\n", flexible_out1_value )
//     fmt.Fprintf( w, "storage_tokura : flexible_out2_value.Water_Line_Slice %v\n", flexible_out2_value.Water_Line_Slice )
//     fmt.Fprintf( w, "storage_tokura : flexible_out2_value.Water2_Slice %v\n", flexible_out2_value.Water2_Slice )

/// test 用　のコード　end

     _ , _ = storage3.Storage_tokura( "Water_Line" ,"put2" ,general_work , water_line , w , r  )

//	put1.Water_line_update ( w , r ,updid ,water2_temp[0].Name ,water_line )

///
///        モニター　再表示
///

	process2.Pipe_line_st_wl_show ( water2_temp[0].Name ,w , r )

}
