package pipe_line_st_wl_show

import (

	    "net/http"
	    "github.com/sawaq7/go12_ver1/client/tokura/suiri/process2"
	    "github.com/sawaq7/go12_ver1/client/tokura/suiri/type4"
	    "strconv"
//	    "fmt"

//	    "github.com/sawaq7/go12_ver1/client/tokura/storage3/trans4"

//	    "github.com/sawaq7/go12_ver1/client/tokura/storage3/check4"
	    "github.com/sawaq7/go12_ver1/client/tokura/storage3"
                                                  )

///
///         入力した水路ラインのチE�Eタを表示する
///

func Pipe_line_st_wl_show(w http.ResponseWriter, r *http.Request) {

//   fmt.Fprintf( w, "sky/pipe_line_st_wl_show start \n"  )  // チE��チE��

/// key-in チE�EタをGET ///

   var water_line type4.Water_Line

   var idmy ,idmy2 int64

///
///         チE��ポラリーファイルより、水路名をゲチE��
///
     water2_temp_minor , _ := storage3.Storage_tokura( "Water2_Temp" ,"trans" ,idmy , idmy2 , w , r  )

     water2_temp, _ := water2_temp_minor.([]type4.Water2_Temp)  // インターフェイス型を型変換

//   water2_temp := trans4.Water2_temp ( w ,r )

///
///         持E��した水路の既存�E水路ラインの数をゲチE��
///

    record_number_temp , _ := storage3.Storage_tokura( "Water_Line" ,"check" ,water2_temp[0].Name , idmy , w , r  )

    record_number, _ := record_number_temp.(int64)

//    record_number := check4.Water_line_re_num( water2_temp[0].Name  ,w , r  )

	for _, water2_tempw := range water2_temp {

       water_line.Name = water2_tempw.Name          /// 水路名�EセチE��
       water_line.Id   = record_number + int64( 1 ) /// idのセチE��

    }

	water_line.Section = r.FormValue("section")  // 区間名をゲチE��

	f_facter := r.FormValue("f_facter")                   // 摩擦係数をゲチE��
	water_line.Friction_Factor,_ =strconv.ParseFloat(f_facter,64)  //　float64　に変換

	velocity := r.FormValue("velocity")                   // 速度をゲチE��
	water_line.Velocity,_ =strconv.ParseFloat(velocity,64)         //　float64　に変換

	p_diameter := r.FormValue("p_diameter")      // 摩擦係数をゲチE��
	water_line.Pipe_Diameter,_ =strconv.ParseFloat(p_diameter,64)  //　float64　に変換

	p_length := r.FormValue("p_length")      // 摩擦係数をゲチE��
	water_line.Pipe_Length,_ =strconv.ParseFloat(p_length,64)  //　float64　に変換

//	fmt.Fprintf( w, "sky/pipe_line_st_wl_show : water_line.Name %v\n", water_line.Name )  // チE��チE��
//	fmt.Fprintf( w, "sky/pipe_line_st_wl_show : water_line.Section %v\n", water_line.Section )  // チE��チE��
//	fmt.Fprintf( w, "sky/pipe_line_st_wl_show : water_line.Friction_Factor %v\n", water_line.Friction_Factor )  // チE��チE��
//	fmt.Fprintf( w, "sky/pipe_line_st_wl_show : water_line.Velocity %v\n", water_line.Velocity )  // チE��チE��
//	fmt.Fprintf( w, "sky/pipe_line_st_wl_show : water_line.Pipe_Diameter %v\n", water_line.Pipe_Diameter )  // チE��チE��
//	fmt.Fprintf( w, "sky/pipe_line_st_wl_show : water_line.Pipe_Length %v\n", water_line.Pipe_Length )  // チE��チE��

///                           　　　　　　　　　　　
///   ストレチE��ファイルに水路ラインファイル惁E��を書ぁE///                          　　　　　　　　　　　

    _ , _ = storage3.Storage_tokura( "Water_Line" ,"put" ,water_line , idmy , w , r  )

//   put1.Water_line( w , r ,water_line )

///
///             モニター表示
///

   process2.Pipe_line_st_wl_show ( water_line.Name ,w , r )


}

