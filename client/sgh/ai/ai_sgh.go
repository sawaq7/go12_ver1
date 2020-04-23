package ai

import (

	    "net/http"
//	    "fmt"

	    "github.com/sawaq7/go12_ver1/client/sgh/ai/analysis1"

                                                )


func Ai_sgh( course_no int64 ,w http.ResponseWriter, r *http.Request ) {

//     IN  course_no  : コースNo
//     IN    w      　: レスポンスライター
//     IN    r      　: リクエストパラメータを乗車するドライバ

//    fmt.Fprintf( w, "ai.ai_sgh start \n" )

///
///    expect delivery number by least squares method
///

//   analysis1.Deliver_growth_rate( course_no  ,w , r  )
   analysis1.Deliver_growth_rate2( course_no  ,w , r  )
//   fmt.Fprintf( w, "ai.ai_sgh normal end \n" )

}

