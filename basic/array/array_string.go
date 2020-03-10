package array

import (

//	    "fmt"
//	    "html/template"

//	    "time"
                                                )

///
///     配列を収縮させる。
///

func Array_string(funct int64 ,column_no int64 ,strings []string  )  ([]string ) {

//     IN  funct  　　　: ファンクション
//     　　　　　：０  削除
//     　　　　　：１  挿入
//     IN  column_no  　　: 基準の配列番号
//     IN  strings  : 文字列のスライス

//     OUT strings  : フォーマット変更後の文字列のスライス

    strings_new := make([]string ,len(strings)+1 )  // 出力エリアを確保

    index:= 0

	for pos, stringsw := range strings {

//	  fmt.Fprintf( w, "array_string stringsw %v\n" ,stringsw)  // デバック

      if funct  == 0 {     // 　　　削除の場合

        if int64(pos+1)  == column_no {


	    } else {

          strings_new[index] = stringsw

          index ++

        }

      } else {           // 　　　挿入の場合

        if int64(pos+1)  == column_no {


          index ++           // 空のデータをセット

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

