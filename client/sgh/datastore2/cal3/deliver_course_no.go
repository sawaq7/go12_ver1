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

func Deliver_course_no( course_no int64 ,w http.ResponseWriter, r *http.Request  ) (int64 ,float64 ,float64 ) {

//     IN  course_no    : コースNo
//     IN    w      　  : レスポンスライター
//     IN    r      　  : リクエストパラメータ
//    OUT number_total  : 配達個数のΣ
//    OUT time_total    : 配達時間のΣ (h)
//    OUT productivity  : 配達生産性　(　配達個数のΣ　/ 配達時間のΣ

//   fmt.Fprintf( w, "cal3.deliver_course_no start \n" )  // チE��チE��

    var number_total int64

    var time_total ,productivity float64

// チE�Eタストアーから、該当するコースNo.のチE�EタをGET

//     deliver_view := trans.Deliver ( 0 ,course_no ,w ,r ) /// セレクトデータをＧ�E��E�

//     deliver_view2 := sort.Deliver ( w ,deliver_view )       /// 2重ソーチE日付�E号車！E
     general_work := make([]type5.General_Work, 2)
     general_work[0].Int64_Work = 0          // 地区惁E��
     general_work[1].Int64_Work = course_no  //　コースNO

     deliver_view := datastore2.Datastore_sgh( "Deliver" ,"trans"  ,general_work , w , r  )

     // 空インターフェイス変数よりバリュー値をゲチE��

     value, _ := deliver_view.([]type2.Deliver)


     number_total = 0
     time_total   = 0.

///
/// 生産性用のチE�Eタを作�Eする
///　

     for _, deliverw := range value {

       if deliverw.Course_No ==  course_no  {        // 該当するコースNOかチェチE��

          number_total = number_total  + deliverw.Number
          time_total   = time_total  + 10.               // "temp" 配達ファイルのフォーマット追加後に修正

       }
	 }
     productivity = float64(number_total) / time_total

//     fmt.Fprintf( w, "cal3.deliver_course_no : number_total %v\n"  ,  number_total )  // チE��チE��
//     fmt.Fprintf( w, "cal3.deliver_course_no : time_total %f\n"  ,  time_total )  // チE��チE��
//     fmt.Fprintf( w, "cal3.deliver_course_no : productivity %f\n"  , productivity )  // チE��チE��

//    fmt.Fprintf( w, "cal3.deliver_course_no normal end \n" )  // チE��チE��

     return number_total ,time_total , productivity


}
