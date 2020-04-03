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

//	fmt.Fprintf( w, "sky_pipe_line_st_wl_update start %v\n" )  // ãEãE¯

	var water_line  type4.Water_Line

	var idmy1 ,idmy2 int64

///
///       ã«ã¬ã³ããEæ°´è·¯æE ±ãã²ãE
///
      water2_temp_minor , _ := storage3.Storage_tokura( "Water2_Temp" ,"trans" ,idmy1 , idmy2 , w , r  )

      water2_temp, _ := water2_temp_minor.([]type4.Water2_Temp)  // ã¤ã³ã¿ã¼ãã§ã¤ã¹åãåå¤æ

//    water2_temp := trans4.Water2_temp( w , r  )

///
///    æE®ããåç¨®ãEEã¿ãGET
///

    updidw , err := strconv.Atoi(r.FormValue("id"))

	if err  != nil {

//	   fmt.Fprintf( w, "sky_pipe_line_st_wl_update :error updidw %v\n", updidw )  // ãEãE¯

	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
	}

    updid := int64(updidw)

    water_line.Id = updid                          // ã¬ã³ã¼ãNOãã»ãE

    water_line.Name = water2_temp[0].Name        // æ°´è·¯åãã»ãE

	water_line.Section = r.FormValue("section")  // åºéåãã²ãE

	f_facter := r.FormValue("f_facter")                   // æ©æ¦ä¿æ°ãã²ãE
	water_line.Friction_Factor,_ =strconv.ParseFloat(f_facter,64)  //ãfloat64ãã«å¤æ

	velocity := r.FormValue("velocity")                   // éåº¦ãã²ãE
	water_line.Velocity,_ =strconv.ParseFloat(velocity,64)         //ãfloat64ãã«å¤æ

	p_diameter := r.FormValue("p_diameter")      // æ©æ¦ä¿æ°ãã²ãE
	water_line.Pipe_Diameter,_ =strconv.ParseFloat(p_diameter,64)  //ãfloat64ãã«å¤æ

	p_length := r.FormValue("p_length")      // æ©æ¦ä¿æ°ãã²ãE
	water_line.Pipe_Length,_ =strconv.ParseFloat(p_length,64)  //ãfloat64ãã«å¤æ

//	fmt.Fprintf( w, "pipe_line_st_wl_update : water_line.Name %v\n", water_line.Name )  // ãEãE¯
//	fmt.Fprintf( w, "pipe_line_st_wl_update : water_line.Section %v\n", water_line.Section )  // ãEãE¯
//	fmt.Fprintf( w, "pipe_line_st_wl_update : water_line.Friction_Factor %v\n", water_line.Friction_Factor )  // ãEãE¯
//	fmt.Fprintf( w, "pipe_line_st_wl_update : water_line.Velocity %v\n", water_line.Velocity )  // ãEãE¯
//	fmt.Fprintf( w, "pipe_line_st_wl_update : water_line.Pipe_Diameter %v\n", water_line.Pipe_Diameter )  // ãEãE¯
//	fmt.Fprintf( w, "pipe_line_st_wl_update : water_line.Pipe_Length %v\n", water_line.Pipe_Length )  // ãEãE¯

///
///         ééæE ±ã®å¤æ´
///

///
///          ãEEã¿ã¹ãã¢ã¼ãããè¡¨ç¤ºç¨ãEEã¿ãGET
///

     general_work := make([]type5.General_Work, 1)
     general_work[0].Int64_Work  = updid                       // ã¢ãEEãEEãEd
     general_work[0].String_Work = water2_temp[0].Name         //ãæ°´è·¯åE
/// test ç¨ãã®ã³ã¼ããstart

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

     // ç©ºã¤ã³ã¿ã¼ãã§ã¤ã¹å¤æ°ããããªã¥ã¼å¤ãã²ãE

//     flexible_out1_value, _ := flexible_out1.([]type5.General_Work)
//     flexible_out2_value, _ := flexible_out2.(type4.Struct_Colle)

//     fmt.Fprintf( w, "storage_tokura : flexible_out1_value %v\n", flexible_out1_value )
//     fmt.Fprintf( w, "storage_tokura : flexible_out2_value.Water_Line_Slice %v\n", flexible_out2_value.Water_Line_Slice )
//     fmt.Fprintf( w, "storage_tokura : flexible_out2_value.Water2_Slice %v\n", flexible_out2_value.Water2_Slice )

/// test ç¨ãã®ã³ã¼ããend

     _ , _ = storage3.Storage_tokura( "Water_Line" ,"put2" ,general_work , water_line , w , r  )

//	put1.Water_line_update ( w , r ,updid ,water2_temp[0].Name ,water_line )

///
///        ã¢ãã¿ã¼ãåè¡¨ç¤º
///

	process2.Pipe_line_st_wl_show ( water2_temp[0].Name ,w , r )

}
