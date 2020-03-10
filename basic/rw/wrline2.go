///
/// スライス(string) を一行にしてファイルに書く
///

package rw

import (
	    "fmt"
	    "strings"
	    "os"
    	     )

func Wrline2(  writer *os.File ,ldata []string  ) {

//     IN  writer : ファイルポインター
//     IN  ldeta  : スライス（string）データ

   fmt.Println ("func wrline2 start　" )
   fmt.Println ("func wrline2 ldata　",ldata )

   ldata2 := strings.Join( ldata, " " )   //　スライスデータを　1行データに変換
   writer.WriteString(ldata2)            // データをファイルに書き込む
   writer.WriteString("\n")             // 改行する

   return
}
