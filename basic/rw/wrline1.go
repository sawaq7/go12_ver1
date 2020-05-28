///
///    float-datas of slice changes line-data and write in file
///

package rw

import (
	    "fmt"
	    "os"
    	     )

func Wrline1(  writer *os.File ,ldata []float64  ) {

//     IN  writer : file-writer
//     IN  ldeta  : float-datas of slice

fmt.Println ("func wrline1 ライター　",writer )
fmt.Println ("func wrline1 ldata　",ldata )

//     write float-data in file

   for  i := 0 ; i < len(ldata) ; i++ {

        fmt.Fprintf(writer ,"%f " ,ldata[i] )

   }

//  line-feed
   fmt.Fprintf(writer ,"\n" )

   return
}
