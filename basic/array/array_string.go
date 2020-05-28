package array

import (

//	    "fmt"
//	    "html/template"

//	    "time"
                                                )

///
///     shrink array
///

func Array_string(funct int64 ,column_no int64 ,strings []string  )  ([]string ) {

//     IN  funct  　　　: function
//     　　　　　delete
//     　　　　　insert
//     IN  column_no  　　: column-number
//     IN  strings  :  string-array

//     OUT strings  : string-array after processing

    strings_new := make([]string ,len(strings)+1 )  //  allocate out-put area
    index:= 0

	for pos, stringsw := range strings {

//	  fmt.Fprintf( w, "array_string stringsw %v\n" ,stringsw)

      if funct  == 0 {     // 　　　when column delete
        if int64(pos+1)  == column_no {


	    } else {

          strings_new[index] = stringsw

          index ++

        }

      } else {           // 　　　when column insert
        if int64(pos+1)  == column_no {


          index ++           //    set space-data

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

