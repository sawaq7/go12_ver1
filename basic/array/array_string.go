package array

import (

//	    "fmt"
//	    "html/template"

//	    "time"
                                                )

///
///     配�Eを収縮させる、E///

func Array_string(funct int64 ,column_no int64 ,strings []string  )  ([]string ) {

//     IN  funct  　　　: ファンクション
//     　　　　　�E�！E 削除
//     　　　　　�E�！E 挿入
//     IN  column_no  　　: 基準�E配�E番号
//     IN  strings  : 斁E���Eのスライス

//     OUT strings  : フォーマット変更後�E斁E���Eのスライス

    strings_new := make([]string ,len(strings)+1 )  // 出力エリアを確俁E
    index:= 0

	for pos, stringsw := range strings {

//	  fmt.Fprintf( w, "array_string stringsw %v\n" ,stringsw)  // チE��チE��

      if funct  == 0 {     // 　　　削除の場吁E
        if int64(pos+1)  == column_no {


	    } else {

          strings_new[index] = stringsw

          index ++

        }

      } else {           // 　　　挿入の場吁E
        if int64(pos+1)  == column_no {


          index ++           // 空のチE�EタをセチE��

          strings_new[index] = stringsw

          index ++

	    } else {

          strings_new[index] = stringsw

          index ++

        }
      }
    }

    return	strings_new
}

