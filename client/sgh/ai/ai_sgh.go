package ai

import (

	    "net/http"
//	    "fmt"

	    "client/sgh/ai/analysis1"

                                                )


func Ai_sgh( course_no int64 ,w http.ResponseWriter, r *http.Request ) {

//     IN  course_no  : コースNo
//     IN    w      　: レスポンスライター
//     IN    r      　: リクエストパラメータを乗車するドライバー

//    fmt.Fprintf( w, "ai.ai_sgh start \n" )  // デバック

///
/// 荷物の成長率を調査
///

//   analysis1.Deliver_growth_rate( course_no  ,w , r  )
   analysis1.Deliver_growth_rate2( course_no  ,w , r  )    // 最終二乗法で計算

//   fmt.Fprintf( w, "ai.ai_sgh normal end \n" )  // デバック

}

