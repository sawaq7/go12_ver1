package sort

import (

	    "net/http"
//	    "fmt"

	    "client/sgh/type2"

                                                )

///
/// 配達情報を、2重sortする
///           key1 : Date  , key2 : Car_No


func Deliver(w http.ResponseWriter ,deliver []type2.Deliver  )  (deliver2 []type2.Deliver ) {

//     IN    w      　　: レスポンスライター
//     IN    r      　　: リクエストパラメータ
//     IN 　deliver     : 構造体　”配達情報”のスライス

//     OUT  deliver2    : ソート後の構造体　”配達情報”のスライス

    var j_min , j_flag   int
    var carno_save ,line_counter int64
    var date_save  string

//    fmt.Fprintf( w, "sort.deliver start \n" )  // デバック

	count := len(deliver)

	deliver2 = make([]type2.Deliver, 0)  /// ソートテーブルを確保

	skip_check := make([]int ,count)        /// スキップの判定フラグテーブルを確保

    line_counter = 0

	for  i := 0 ; i < count ; i++  {

	  j_flag = -1

	  for  j := 0 ; j < count ; j++ {



	    if skip_check[j] != 1  {  /// ソート済みか否かのチェック

	      if j_flag  == -1  {     ///  初期値のセット

	        j_min = j
	        date_save  = deliver[j].Date
	        carno_save = deliver[j].Car_No
	        j_flag = 0

	        /// 最小値、再セット

	      }  else if ( date_save >  deliver[j].Date )                                   ||
	                 ( date_save == deliver[j].Date && carno_save > deliver[j].Car_No )     {

	        j_min = j
	        date_save  = deliver[j].Date
	        carno_save = deliver[j].Car_No

	      }
	    }
      }

/// min値をテーブルにセット
      line_counter ++
      deliver[j_min].Line_No = line_counter

      deliver2 = append ( deliver2,  deliver[j_min] )

/// ソート済フラグをセット

      skip_check[j_min] = 1

	}

    return	deliver2
}

