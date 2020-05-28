///
///    string-datas of slice changes line-data and write in file
///

package rw

import (
	    "fmt"
	    "strings"
	    "os"
    	     )

func Wrline2(  writer *os.File ,ldata []string  ) {

//     IN  writer : file-writer
//     IN  ldeta  : string-datas of slice

   fmt.Println ("func wrline2 start　" )
   fmt.Println ("func wrline2 ldata　",ldata )

   ldata2 := strings.Join( ldata, " " )   //　change data-type from  string-datas of slice to line-data
   writer.WriteString(ldata2)             //  write line-data
   writer.WriteString("\n")               // line-feed
   return
}
