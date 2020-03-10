package struct_set

import (

	    "net/http"
//	    "fmt"
	    "strings"
	    "strconv"

	    "client/tokura/suiri/type4"
                                                )

///                           　　　　　　　　　　　
///   ストリング型のラインデータを構造体にセット（Water＿Line)（水路ラインファイル・ストレッジ）
///                          　　　　　　　　　　　

func Water_line( w http.ResponseWriter, water_line_string string )  ( type4.Water_Line ) {

//     IN         w             : レスポンスライター
//     IN   water_line_string   : 水路ラインファイルのラインデータ（string型）
//     OUT  water_line_struct   : 水路ラインファイルのラインデータ（struct型・Water_Line）

//    fmt.Fprintf( w, "struct_set.water_line start \n" )  // デバック
//    fmt.Fprintf(w, "struct_set.water_line : water_line_string %s\n", water_line_string )  // デバック

    var water_line_struct type4.Water_Line                 //   Water_Line　のワークエリアを確保

    str := strings.Fields(water_line_string)  // ストリング型のラインデータを分割

    for ii := 0 ; ii < len(str) ; ii++ {

      switch ii {

        case 0 :     //    Idの場合

          str2 := strings.Split(str[ii], "{"  )     // 記号　"{" をトリム

//          fmt.Printf("struct_set: str2 %v\n", str2 )
//          fmt.Printf("struct_set: len(str2) %v\n", len(str2) )

          water_line_struct.Id ,_ =strconv.ParseInt(str2[1],10 ,64)   // 文字列をint64に変換

//          fmt.Printf("struct_set: water_line_struct.Id %v\n", water_line_struct.Id )

        break;

        case 1 :      //　水路名の場合

          water_line_struct.Name = str[ii]

//          fmt.Printf("struct_set: water_line_struct.Name %v\n", water_line_struct.Name )

        break;

        case 2 :      //　区間名の場合

          water_line_struct.Section = str[ii]

//          fmt.Printf("struct_set: water_line_struct.Section %v\n", water_line_struct.Section )

        break;

        case 3 :      //　摩擦係数の場合

          water_line_struct.Friction_Factor,_ =strconv.ParseFloat(str[ii],64)

//          fmt.Printf("struct_set: water_line_struct.High %v\n", water_line_struct.Friction_Factor )

        break;

        case 4 :      //　速度の場合

          water_line_struct.Velocity,_ =strconv.ParseFloat(str[ii],64)

//          fmt.Printf("struct_set: water_line_struct.High %v\n", water_line_struct.Velocity )

        break;

        case 5 :      //　管径

          water_line_struct.Pipe_Diameter,_ =strconv.ParseFloat(str[ii],64)

//          fmt.Printf("struct_set: water_line_struct.High %v\n", water_line_struct.Pipe_Diameter )

        break;

        case 6 :      //　管長

          str3 := strings.Split(str[ii], "}"  )         // 記号　"}" をトリム
          water_line_struct.Pipe_Length,_ =strconv.ParseFloat(str3[0],64)

//         fmt.Printf("struct_set: water_line_struct.Roughness_Factor %v\n", water_line_struct.Pipe_Length )

        break;

     }

   }

   return water_line_struct

}

