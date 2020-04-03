///
/// スライス(string) を一行にしてファイルに書ぁE///

package rw

import (
	    "fmt"
	    "strings"
	    "os"
    	     )

func Wrline2(  writer *os.File ,ldata []string  ) {

//     IN  writer : ファイルポインター
//     IN  ldeta  : スライス�E�Etring�E�データ

   fmt.Println ("func wrline2 start　" )
   fmt.Println ("func wrline2 ldata　",ldata )

   ldata2 := strings.Join( ldata, " " )   //　スライスチE�Eタを　1行データに変換
   writer.WriteString(ldata2)            // チE�Eタをファイルに書き込む
   writer.WriteString("\n")             // 改行すめE
   return
}
