package sort

import (

	    "net/http"
//	    "fmt"

	    "github.com/sawaq7/go12_ver1/client/sgh/type2"

                                                )

///
/// ééæE ±ããEésortãã
///           key1 : Date  , key2 : Car_No


func Deliver(w http.ResponseWriter ,deliver []type2.Deliver  )  (deliver2 []type2.Deliver ) {

//     IN    w      ãã: ã¬ã¹ãã³ã¹ã©ã¤ã¿ã¼
//     IN    r      ãã: ãªã¯ã¨ã¹ããã©ã¡ã¼ã¿
//     IN ãdeliver     : æ§é ä½ãâéEéæå ±âãEã¹ã©ã¤ã¹

//     OUT  deliver2    : ã½ã¼ãå¾ãEæ§é ä½ãâéEéæå ±âãEã¹ã©ã¤ã¹

    var j_min , j_flag   int
    var carno_save ,line_counter int64
    var date_save  string

//    fmt.Fprintf( w, "sort.deliver start \n" )  // ãEãE¯

	count := len(deliver)

	deliver2 = make([]type2.Deliver, 0)  /// ã½ã¼ããã¼ãã«ãç¢ºä¿E
	skip_check := make([]int ,count)        /// ã¹ã­ãEEã®å¤å®ãã©ã°ãEEãã«ãç¢ºä¿E
    line_counter = 0

	for  i := 0 ; i < count ; i++  {

	  j_flag = -1

	  for  j := 0 ; j < count ; j++ {



	    if skip_check[j] != 1  {  /// ã½ã¼ãæ¸ã¿ãå¦ããEãã§ãE¯

	      if j_flag  == -1  {     ///  åæå¤ã®ã»ãE

	        j_min = j
	        date_save  = deliver[j].Date
	        carno_save = deliver[j].Car_No
	        j_flag = 0

	        /// æå°å¤ãåEã»ãE

	      }  else if ( date_save >  deliver[j].Date )                                   ||
	                 ( date_save == deliver[j].Date && carno_save > deliver[j].Car_No )     {

	        j_min = j
	        date_save  = deliver[j].Date
	        carno_save = deliver[j].Car_No

	      }
	    }
      }

/// minå¤ããã¼ãã«ã«ã»ãE
      line_counter ++
      deliver[j_min].Line_No = line_counter

      deliver2 = append ( deliver2,  deliver[j_min] )

/// ã½ã¼ãæ¸ãã©ã°ãã»ãE

      skip_check[j_min] = 1

	}

    return	deliver2
}

