package ai

import (

	    "net/http"
//	    "fmt"

	    "github.com/sawaq7/go12_ver1/client/sgh/ai/analysis1"

                                                )


func Ai_sgh( course_no int64 ,w http.ResponseWriter, r *http.Request ) {

//     IN  course_no  : コースNo
//     IN    w      　: レスポンスライター
//     IN    r      　: リクエストパラメータを乗車するドライバ�E

//    fmt.Fprintf( w, "ai.ai_sgh start \n" )  // チE��チE��

///
/// 荷物の成長玁E��調査
///

//   analysis1.Deliver_growth_rate( course_no  ,w , r  )
   analysis1.Deliver_growth_rate2( course_no  ,w , r  )    // 最終二乗法で計箁E
//   fmt.Fprintf( w, "ai.ai_sgh normal end \n" )  // チE��チE��

}

