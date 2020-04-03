package struct_set

import (

	    "net/http"
//	    "fmt"
	    "strings"
	    "strconv"

	    "github.com/sawaq7/go12_ver1/client/tokura/suiri/type4"
                                                )

///                           　　　　　　　　　　　
///   ストリング型�EラインチE�Eタを構造体にセチE���E�Eater2_Temp)�E�水路ファイル・ストレチE���E�E///                          　　　　　　　　　　　

func Water2_temp( w http.ResponseWriter, water2_temp_string string )  ( type4.Water2_Temp ) {

//     IN         w              : レスポンスライター
//     IN   water2_temp_string   : 水路ファイルのラインチE�Eタ�E�Etring型！E//     OUT  water2_temp_struct   : 水路ファイルのラインチE�Eタ�E�Etruct型�EWater2_Temp�E�E
//    fmt.Fprintf( w, "struct_set.water2_temp start \n" )  // チE��チE��
//    fmt.Fprintf(w, "struct_set.water2_temp : water2_temp_string %s\n", water2_temp_string )  // チE��チE��

    var water2_temp_struct type4.Water2_Temp                 //   Water2_Temp　のワークエリアを確俁E
    str := strings.Fields(water2_temp_string)  // ストリング型�EラインチE�Eタを�E割

    for ii := 0 ; ii < len(str) ; ii++ {

      switch ii {

        case 0 :     //    Idの場吁E
          str2 := strings.Split(str[ii], "{"  )     // 記号　"{" をトリム

//          fmt.Printf("struct_set: str2 %v\n", str2 )
//          fmt.Printf("struct_set: len(str2) %v\n", len(str2) )

          water2_temp_struct.Id ,_ =strconv.ParseInt(str2[1],10 ,64)   // 斁E���Eをint64に変換

//          fmt.Printf("struct_set: water2_temp_struct.Id %v\n", water2_temp_struct.Id )

        break;

        case 1 :      //　水路名�E場吁E
          water2_temp_struct.Name = str[ii]

//          fmt.Printf("struct_set: water2_temp_struct.Name %v\n", water2_temp_struct.Name )

        break;

        case 2 :      //　水路高�E場吁E
          water2_temp_struct.High,_ =strconv.ParseFloat(str[ii],64)

//          fmt.Printf("struct_set: water2_temp_struct.High %v\n", water2_temp_struct.High )

        break;

        case 3 :      //　粗粒係数の場吁E
          str3 := strings.Split(str[ii], "}"  )         // 記号　"}" をトリム
          water2_temp_struct.Roughness_Factor,_ =strconv.ParseFloat(str3[0],64)

//          fmt.Printf("struct_set: water2_temp_struct.Roughness_Factor %v\n", water2_temp_struct.Roughness_Factor )

        break;

     }

   }

   return water2_temp_struct

}

