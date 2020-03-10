package cal2

import (

	    "client/sgh/datastore2"
	    "net/http"
//	    "fmt"
	    "basic/date1"
	    "client/sgh/type2"

                                                )
///
///  コースNOごとにAI式を計算して荷物の数を予想する
///

func Ai_sgh_analysis(w http.ResponseWriter, r *http.Request, course_no int64 , deliver_date string )(expected_num float64 ) {

//     IN    w      　  : レスポンスライター
//     IN    r      　  : リクエストパラメータを乗車するドライバー
//     IN  course_no    : コースNo
//     IN deliver_date  : 配達日
//    OUT  expected_num : 予想荷物数

//    fmt.Fprintf( w, "cal2.ai_sgh_analysis start \n" )  // デバック
//    fmt.Fprintf( w, "cal2.ai_sgh_analysis course_no \n" ,course_no)  // デバック

    var expression string

    var factor float64

    var ii int64

    expected_num = 0.

///
/// AI 条件式を計算する
///

    sgh_ai := datastore2.Datastore_sgh( "Sgh_Ai","trans" ,course_no , w , r  )  // AI 条件式をゲット

// 空インターフェイス変数よりバリュー値をゲット

    sgh_ai_value, _ := sgh_ai.([]type2.Sgh_Ai)



	for _ , sgh_ai_valuew := range sgh_ai_value {



	  deliver_date_real := date1.Date_realdata_get( w  ,deliver_date )   // タイムデータ作成

      date_sub := deliver_date_real.Sub(sgh_ai_valuew.Date_Basic_Real)  // 基準日からの経過日数を計算（X座標）

      xx := float64(date_sub/(3600000000000*24))  //　日付に変換

	  for ii = 0 ; ii < sgh_ai_valuew.Item_Num ; ii++ {

	    if ii == 0 {                       // 計算式・係数をセット

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
//        fmt.Fprintf( w, "cal2.ai_sgh_analysis expression \n" ,expression)  // デバック
	    switch expression {


          case "*" :         // 乗算の場合

             expected_num = expected_num + xx * factor

             break;

          case "+" :         // 加算の場合

             expected_num = expected_num + factor

             break;

        }


	  }

	}

//	fmt.Fprintf( w, "cal2.ai_sgh_analysis expected_num \n" ,expected_num)  // デバック
//	fmt.Fprintf( w, "cal2.ai_sgh_analysis normal end \n" )  // デバック

    return	expected_num

}