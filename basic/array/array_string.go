package array

import (

//	    "fmt"
//	    "html/template"

//	    "time"
                                                )

///
///     éåEãåç¸®ããããE///

func Array_string(funct int64 ,column_no int64 ,strings []string  )  ([]string ) {

//     IN  funct  ããã: ãã¡ã³ã¯ã·ã§ã³
//     ãããããEï¼E åé¤
//     ãããããEï¼E æ¿å¥
//     IN  column_no  ãã: åºæºãEéåEçªå·
//     IN  strings  : æE­åEã®ã¹ã©ã¤ã¹

//     OUT strings  : ãã©ã¼ãããå¤æ´å¾ãEæE­åEã®ã¹ã©ã¤ã¹

    strings_new := make([]string ,len(strings)+1 )  // åºåã¨ãªã¢ãç¢ºä¿E
    index:= 0

	for pos, stringsw := range strings {

//	  fmt.Fprintf( w, "array_string stringsw %v\n" ,stringsw)  // ãEãE¯

      if funct  == 0 {     // ãããåé¤ã®å ´åE
        if int64(pos+1)  == column_no {


	    } else {

          strings_new[index] = stringsw

          index ++

        }

      } else {           // ãããæ¿å¥ã®å ´åE
        if int64(pos+1)  == column_no {


          index ++           // ç©ºã®ãEEã¿ãã»ãE

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

