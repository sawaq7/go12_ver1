package cal3

import (

	    "net/http"
//	    "fmt"
	    "github.com/sawaq7/go12_ver1/client/sgh/datastore2"
//	    "github.com/sawaq7/go12_ver1/client/sgh/datastore2/sort"
        "github.com/sawaq7/go12_ver1/client/sgh/type2"
	    "github.com/sawaq7/go12_ver1/general/type5"

                                                )
///
///    calculate for each cource-no in d.s.
///

func Deliver_course_no( course_no int64 ,w http.ResponseWriter, r *http.Request  ) (int64 ,float64 ,float64 ) {

//     IN  course_no    : コースNo
//     IN    w      　  : レスポンスライター
//     IN    r      　  : リクエストパラメータ
//    OUT number_total  : 配達個数のΣ
//    OUT time_total    : 配達時間のΣ (h)
//    OUT productivity  : 配達生産性　(　配達個数のΣ　/ 配達時間のΣ

//   fmt.Fprintf( w, "cal3.deliver_course_no start \n" )

    var number_total int64

    var time_total ,productivity float64

//     get course-no in d.s

//     deliver_view := trans.Deliver ( 0 ,course_no ,w ,r )

//     deliver_view2 := sort.Deliver ( w ,deliver_view )
     general_work := make([]type5.General_Work, 2)
     general_work[0].Int64_Work = 0
     general_work[1].Int64_Work = course_no

     deliver_view := datastore2.Datastore_sgh( "Deliver" ,"trans"  ,general_work , w , r  )

     value, _ := deliver_view.([]type2.Deliver)


     number_total = 0
     time_total   = 0.

///
/// 生産性用のチE�Eタを作�Eする make data for productivity
///　

     for _, deliverw := range value {

       if deliverw.Course_No ==  course_no  {     // check  course-no which fit these conditions
          number_total = number_total  + deliverw.Number
          time_total   = time_total  + 10.

       }
	 }
     productivity = float64(number_total) / time_total

//     fmt.Fprintf( w, "cal3.deliver_course_no : number_total %v\n"  ,  number_total )
//     fmt.Fprintf( w, "cal3.deliver_course_no : time_total %f\n"  ,  time_total )
//     fmt.Fprintf( w, "cal3.deliver_course_no : productivity %f\n"  , productivity )

//    fmt.Fprintf( w, "cal3.deliver_course_no normal end \n" )

     return number_total ,time_total , productivity
}
