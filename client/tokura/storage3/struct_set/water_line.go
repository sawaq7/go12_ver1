package struct_set

import (

	    "net/http"
//	    "fmt"
	    "strings"
	    "strconv"

	    "github.com/sawaq7/go12_ver1/client/tokura/suiri/type4"
                                                )

///                           　　　　　　　　　　　
///   ストリング型�EラインチE�Eタを構造体にセチE���E�Eater�E�Line)�E�水路ラインファイル・ストレチE���E�E///                          　　　　　　　　　　　

func Water_line( w http.ResponseWriter, water_line_string string )  ( type4.Water_Line ) {

//     IN         w             : レスポンスライター
//     IN   water_line_string   : 水路ラインファイルのラインチE�Eタ�E�Etring型！E//     OUT  water_line_struct   : 水路ラインファイルのラインチE�Eタ�E�Etruct型�EWater_Line�E�E
//    fmt.Fprintf( w, "struct_set.water_line start \n" )  // チE��チE��
//    fmt.Fprintf(w, "struct_set.water_line : water_line_string %s\n", water_line_string )  // チE��チE��

    var water_line_struct type4.Water_Line                 //   Water_Line　のワークエリアを確俁E
    str := strings.Fields(water_line_string)  // ストリング型�EラインチE�Eタを�E割

    for ii := 0 ; ii < len(str) ; ii++ {

      switch ii {

        case 0 :     //    Idの場吁E
          str2 := strings.Split(str[ii], "{"  )     // 記号　"{" をトリム

//          fmt.Printf("struct_set: str2 %v\n", str2 )
//          fmt.Printf("struct_set: len(str2) %v\n", len(str2) )

          water_line_struct.Id ,_ =strconv.ParseInt(str2[1],10 ,64)   // 斁E���Eをint64に変換

//          fmt.Printf("struct_set: water_line_struct.Id %v\n", water_line_struct.Id )

        break;

        case 1 :      //　水路名�E場吁E
          water_line_struct.Name = str[ii]

//          fmt.Printf("struct_set: water_line_struct.Name %v\n", water_line_struct.Name )

        break;

        case 2 :      //　区間名の場吁E
          water_line_struct.Section = str[ii]

//          fmt.Printf("struct_set: water_line_struct.Section %v\n", water_line_struct.Section )

        break;

        case 3 :      //　摩擦係数の場吁E
          water_line_struct.Friction_Factor,_ =strconv.ParseFloat(str[ii],64)

//          fmt.Printf("struct_set: water_line_struct.High %v\n", water_line_struct.Friction_Factor )

        break;

        case 4 :      //　速度の場吁E
          water_line_struct.Velocity,_ =strconv.ParseFloat(str[ii],64)

//          fmt.Printf("struct_set: water_line_struct.High %v\n", water_line_struct.Velocity )

        break;

        case 5 :      //　管征E
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

