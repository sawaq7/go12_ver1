///                                           ///
/// スライス(float) を一行にしてファイルに書く ///
///                                          ///

package rw

import (
	    "fmt"
	    "os"
    	     )

func Wrline1(  writer *os.File ,ldata []float64  ) {

//     IN  writer : ファイルポインター
//     IN  ldeta  : スライス（float）データ

fmt.Println ("func wrline1 ライター　",writer )
fmt.Println ("func wrline1 ldata　",ldata )

// データをファイルに書き込む

   for  i := 0 ; i < len(ldata) ; i++ {

        fmt.Fprintf(writer ,"%f " ,ldata[i] )

   }

// 改行する

   fmt.Fprintf(writer ,"\n" )

   return
}
