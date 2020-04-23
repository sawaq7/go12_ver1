package sort

import (

	    "net/http"
//	    "fmt"

	    "github.com/sawaq7/go12_ver1/client/sgh/type2"

                                                )

///
///       sort by double
///           key1 : Date  , key2 : Car_No


func Deliver(w http.ResponseWriter ,deliver []type2.Deliver  )  (deliver2 []type2.Deliver ) {

//     IN    w      　　: レスポンスライター
//     IN    r      　　: リクエストパラメータ
//     IN 　deliver     : slice of struct ( Deliver )

//     OUT  deliver2    : it's the slice of struct ( Deliver ) which is  sorted

    var j_min , j_flag   int
    var carno_save ,line_counter int64
    var date_save  string

//    fmt.Fprintf( w, "sort.deliver start \n" )

	count := len(deliver)

	deliver2 = make([]type2.Deliver, 0)
	skip_check := make([]int ,count)

	for  i := 0 ; i < count ; i++  {

	  j_flag = -1

	  for  j := 0 ; j < count ; j++ {



	    if skip_check[j] != 1  {  ///  whether or not process skips

	      if j_flag  == -1  {

	        j_min = j
	        date_save  = deliver[j].Date
	        carno_save = deliver[j].Car_No
	        j_flag = 0

	        ///  set min  value

	      }  else if ( date_save >  deliver[j].Date )                                   ||
	                 ( date_save == deliver[j].Date && carno_save > deliver[j].Car_No )     {

	        j_min = j
	        date_save  = deliver[j].Date
	        carno_save = deliver[j].Car_No

	      }
	    }
      }

                            ///  set table min value
      line_counter ++
      deliver[j_min].Line_No = line_counter

      deliver2 = append ( deliver2,  deliver[j_min] )

/// ソート済フラグをセチE��

      skip_check[j_min] = 1

	}

    return	deliver2
}

