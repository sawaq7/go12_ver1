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
///    配達ファイルより、該当するコースNOの生産性を算�Eする
///

func Deliver_car_no( car_no int64 ,w http.ResponseWriter, r *http.Request  ) (int64 ,float64 ,float64 ) {

//     IN  car_no    　 : 号車No
//     IN    w      　  : レスポンスライター
//     IN    r      　  : リクエストパラメータ
//    OUT number_total  : 配達個数のΣ
//    OUT time_total    : 配達時間のΣ (h)
//    OUT productivity  : 配達生産性　(　配達個数のΣ　/ 配達時間のΣ

//    fmt.Fprintf( w, "cal3.deliver_car_no start \n" )
//    fmt.Fprintf( w, "cal3.deliver_car_no : car_no %v\n"  ,  car_no )

    var number_total int64

    var time_total ,productivity float64

///
///       get car-no in d.s.
///

//     deliver_view := trans.Deliver ( 1 ,car_no ,w ,r )

//     deliver_view2 := sort.Deliver ( w ,deliver_view )
     general_work := make([]type5.General_Work, 2)
     general_work[0].Int64_Work = 1
     general_work[1].Int64_Work = car_no

     deliver_view := datastore2.Datastore_sgh( "Deliver" ,"trans"  ,general_work , w , r )

     value, _ := deliver_view.([]type2.Deliver)

     number_total = 0
     time_total   = 0.

///
///        make data for productivity
///　

     for _, deliverw := range value {

       if deliverw.Car_No ==  car_no  {        //   check  car-no which fit these conditions

          number_total = number_total  + deliverw.Number
          time_total   = time_total  + 10.

       }
	 }

	 if time_total   == 0. ||
	    number_total == 0     {

        productivity = 0.

     } else  {

        productivity = float64(number_total) / time_total

     }


//     fmt.Fprintf( w, "cal3.deliver_car_no : number_total %v\n"  ,  number_total )
//     fmt.Fprintf( w, "cal3.deliver_car_no : time_total %f\n"  ,  time_total )
//     fmt.Fprintf( w, "cal3.deliver_car_no : productivity %f\n"  , productivity )

//     fmt.Fprintf( w, "cal3.deliver_car_no normal end \n" )

     return number_total ,time_total , productivity
}
