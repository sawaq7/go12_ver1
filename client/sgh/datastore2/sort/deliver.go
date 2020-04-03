package sort

import (

	    "net/http"
//	    "fmt"

	    "github.com/sawaq7/go12_ver1/client/sgh/type2"

                                                )

///
/// 配達惁E��を、E重sortする
///           key1 : Date  , key2 : Car_No


func Deliver(w http.ResponseWriter ,deliver []type2.Deliver  )  (deliver2 []type2.Deliver ) {

//     IN    w      　　: レスポンスライター
//     IN    r      　　: リクエストパラメータ
//     IN 　deliver     : 構造体　”�E達情報”�Eスライス

//     OUT  deliver2    : ソート後�E構造体　”�E達情報”�Eスライス

    var j_min , j_flag   int
    var carno_save ,line_counter int64
    var date_save  string

//    fmt.Fprintf( w, "sort.deliver start \n" )  // チE��チE��

	count := len(deliver)

	deliver2 = make([]type2.Deliver, 0)  /// ソートテーブルを確俁E
	skip_check := make([]int ,count)        /// スキチE�Eの判定フラグチE�Eブルを確俁E
    line_counter = 0

	for  i := 0 ; i < count ; i++  {

	  j_flag = -1

	  for  j := 0 ; j < count ; j++ {



	    if skip_check[j] != 1  {  /// ソート済みか否か�EチェチE��

	      if j_flag  == -1  {     ///  初期値のセチE��

	        j_min = j
	        date_save  = deliver[j].Date
	        carno_save = deliver[j].Car_No
	        j_flag = 0

	        /// 最小値、�EセチE��

	      }  else if ( date_save >  deliver[j].Date )                                   ||
	                 ( date_save == deliver[j].Date && carno_save > deliver[j].Car_No )     {

	        j_min = j
	        date_save  = deliver[j].Date
	        carno_save = deliver[j].Car_No

	      }
	    }
      }

/// min値をテーブルにセチE��
      line_counter ++
      deliver[j_min].Line_No = line_counter

      deliver2 = append ( deliver2,  deliver[j_min] )

/// ソート済フラグをセチE��

      skip_check[j_min] = 1

	}

    return	deliver2
}

