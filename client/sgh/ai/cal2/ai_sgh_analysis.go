package cal2

import (

	    "github.com/sawaq7/go12_ver1/client/sgh/datastore2"
	    "net/http"
//	    "fmt"
	    "github.com/sawaq7/go12_ver1/basic/date1"
	    "github.com/sawaq7/go12_ver1/client/sgh/type2"

                                                )
///
///    calcurate ai expression and expect delivery number  for each cource no
///

func Ai_sgh_analysis(w http.ResponseWriter, r *http.Request, course_no int64 , deliver_date string )(expected_num float64 ) {

//     IN    w      　  : レスポンスライター
//     IN    r      　  : リクエストパラメータを乗車するドライバ�E
//     IN  course_no    :
//     IN deliver_date  :
//    OUT  expected_num :

//    fmt.Fprintf( w, "cal2.ai_sgh_analysis start \n" )
//    fmt.Fprintf( w, "cal2.ai_sgh_analysis course_no \n" ,course_no)

    var expression string

    var factor float64

    var ii int64

    expected_num = 0.

///
///     get ai expression
///

    sgh_ai := datastore2.Datastore_sgh( "Sgh_Ai","trans" ,course_no , w , r  )

//    get value from interface data

    sgh_ai_value, _ := sgh_ai.([]type2.Sgh_Ai)

	for _ , sgh_ai_valuew := range sgh_ai_value {



	  deliver_date_real := date1.Date_realdata_get( w  ,deliver_date )   //    make time data

      date_sub := deliver_date_real.Sub(sgh_ai_valuew.Date_Basic_Real)
      xx := float64(date_sub/(3600000000000*24))

	  for ii = 0 ; ii < sgh_ai_valuew.Item_Num ; ii++ {

	    if ii == 0 {

	      expression = sgh_ai_valuew.Item1_Name
	      factor     = sgh_ai_valuew.Item1_Factor

	    }else if ii == 1 {

	      expression = sgh_ai_valuew.Item2_Name
	      factor     = sgh_ai_valuew.Item2_Factor


	    }else if ii == 2 {

	      expression = sgh_ai_valuew.Item3_Name
	      factor     = sgh_ai_valuew.Item3_Factor

	    }else if ii == 3 {

	      expression = sgh_ai_valuew.Item4_Name
	      factor     = sgh_ai_valuew.Item4_Factor

	    }else if ii == 4 {

	      expression = sgh_ai_valuew.Item5_Name
	      factor     = sgh_ai_valuew.Item5_Factor

	    }
//        fmt.Fprintf( w, "cal2.ai_sgh_analysis expression \n" ,expression)
	    switch expression {


          case "*" :
             expected_num = expected_num + xx * factor

             break;

          case "+" :
             expected_num = expected_num + factor

             break;

        }


	  }

	}

//	fmt.Fprintf( w, "cal2.ai_sgh_analysis expected_num \n" ,expected_num)
//	fmt.Fprintf( w, "cal2.ai_sgh_analysis normal end \n" )

    return	expected_num

}
