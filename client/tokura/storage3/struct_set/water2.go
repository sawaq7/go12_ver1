package struct_set

import (

	    "net/http"
//	    "fmt"
	    "strings"
	    "strconv"

	    "client/tokura/suiri/type4"
                                                )

///                           　　　　　　　　　　　
///   ストリング型のラインデータを構造体にセット（Water2)（水路ファイル・ストレッジ）
///                          　　　　　　　　　　　

func Water2( w http.ResponseWriter, water2_string string )  ( type4.Water2 ) {

//     IN         w         : レスポンスライター
//     IN   water2_string   : 水路ファイルのラインデータ（string型）
//     OUT  water2_struct     : 水路ファイルのラインデータ（struct型・Water2）

//    fmt.Fprintf( w, "struct_set.water2 start \n" )  // デバック
//    fmt.Fprintf(w, "struct_set.water2 : water2_string %s\n", water2_string )  // デバック

    var water2_struct type4.Water2                 //   Water2　のワークエリアを確保

    str := strings.Fields(water2_string)  // ストリング型のラインデータを分割

    for ii := 0 ; ii < len(str) ; ii++ {

      switch ii {

        case 0 :     //    Idの場合

          str2 := strings.Split(str[ii], "{"  )     // 記号　"{" をトリム

//          fmt.Printf("struct_set: str2 %v\n", str2 )
//         fmt.Printf("struct_set: len(str2) %v\n", len(str2) )

          water2_struct.Id ,_ =strconv.ParseInt(str2[1],10 ,64)   // 文字列をint64に変換

//          fmt.Printf("struct_set: water2_struct.Id %v\n", water2_struct.Id )

        break;

        case 1 :      //　水路名の場合

          water2_struct.Name = str[ii]

//          fmt.Printf("struct_set: water2_struct.Name %v\n", water2_struct.Name )

        break;

        case 2 :      //　水路高の場合

          water2_struct.High,_ =strconv.ParseFloat(str[ii],64)

//          fmt.Printf("struct_set: water2_struct.High %v\n", water2_struct.High )

        break;

        case 3 :      //　粗粒係数の場合

          str3 := strings.Split(str[ii], "}"  )         // 記号　"}" をトリム
          water2_struct.Roughness_Factor,_ =strconv.ParseFloat(str3[0],64)

//          fmt.Printf("struct_set: water2_struct.Roughness_Factor %v\n", water2_struct.Roughness_Factor )

        break;

     }

   }

   return water2_struct

}

