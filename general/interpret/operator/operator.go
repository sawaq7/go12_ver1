package operator

import (

        "net/http"
//	    "fmt"

	    "general/type5"
	    "general/datastore5/trans3"

                                    )

///
///     比較演算子のメインルーチン
///


func Operator( csv_inf []type5.Csv_Inf ,function string ,match_word string ,column_no int ,w http.ResponseWriter, r *http.Request )  ( csv_inf2 []type5.Csv_Inf ) {

//     IN 　csv_inf      : カレントのcsv情報
//     IN    function    : ファンクション　
//        　　　　　　　　　＊　eq ne ge gt le lt
//     IN   match_word   : マッチングワード　
//     IN   column_no　  : マッチング対象の行NO
//     IN    w      　　 : レスポンスライター
//     IN    r      　　 : リクエストパラメータ

//     OUT  csv_inf2     : マッチング後のcsv情報

//    fmt.Fprintf( w, "operator start \n" )  // デバック
//    fmt.Fprintf( w, "operator function %v\n" ,function )  // デバック
 //   fmt.Fprintf( w, "operator match_word %v\n" ,match_word )  // デバック

///
///     マッチングキーをゲット
///

    match_key := trans3.Csv_inf_column ( w ,r ,column_no )

///
///  ファンクションにより各種処理を分岐して行う。
///

	switch function {

      case "eq" :

        csv_inf2 = Operator_eq ( w , r ,csv_inf  ,match_word  ,match_key )

      break;

      case "ne" :

        csv_inf2 = Operator_ne ( w , r ,csv_inf  ,match_word  ,match_key )

      break;

      case "gt" :

        csv_inf2 = Operator_gt ( w , r ,csv_inf  ,match_word  ,match_key )

      break;

      case "ge" :

        csv_inf2 = Operator_ge ( w , r ,csv_inf  ,match_word  ,match_key )

      break;

      case "lt" :

        csv_inf2 = Operator_lt ( w , r ,csv_inf  ,match_word  ,match_key )

      break;

      case "le" :

        csv_inf2 = Operator_le ( w , r ,csv_inf  ,match_word  ,match_key )

      break;

    }

	return csv_inf2

}

///
/// 　　　　	メニュー　：operator/eq
///

func Operator_eq ( w http.ResponseWriter , r *http.Request ,csv_inf []type5.Csv_Inf ,match_word string ,match_key []string  )  (csv_inf2 []type5.Csv_Inf ) {

//     IN    w      　　: レスポンスライター
//     IN    r      　　: リクエストパラメータ
//     IN 　csv_inf     : csv情報
//     IN   match_word   : マッチングワード
//     IN 　match_key   : マッチングキー

//     OUT  csv_inf2    : ソート後のcsv情報

    var line_counter int64

    count := len(csv_inf)    // レコード数ゲット

//    fmt.Fprintf( w, "operator_eq start \n" )  // デバック

///
/// 　　　マッチングする　
///
	csv_inf2 = make([]type5.Csv_Inf, 0)  /// ソートテーブルを確保

    line_counter = 0

	for  loop_1 := 0 ; loop_1 < count ; loop_1++  {

	  if match_word == match_key[loop_1]  {  ///     マッチング

	    line_counter ++

//        csv_inf[loop_1].Line_No = line_counter

        csv_inf2 = append ( csv_inf2,  csv_inf[loop_1] )

	  }

	}

    return	csv_inf2
}

///
/// 　　　　	メニュー　：operator/ne
///

func Operator_ne ( w http.ResponseWriter , r *http.Request ,csv_inf []type5.Csv_Inf ,match_word string ,match_key []string  )  (csv_inf2 []type5.Csv_Inf ) {

//     IN    w      　　: レスポンスライター
//     IN    r      　　: リクエストパラメータ
//     IN 　csv_inf     : csv情報
//     IN   match_word   : マッチングワード
//     IN 　match_key   : マッチングキー

//     OUT  csv_inf2    : マッチング後のcsv情報

    var line_counter int64

    count := len(csv_inf)    // レコード数ゲット

//    fmt.Fprintf( w, "operator_ne start \n" )  // デバック

///
/// 　　　マッチングする　
///

	csv_inf2 = make([]type5.Csv_Inf, 0)  /// ソートテーブルを確保

    line_counter = 0

	for  loop_1 := 0 ; loop_1 < count ; loop_1++  {

	  if match_word != match_key[loop_1]  {  ///     マッチング

	    line_counter ++

//        csv_inf[loop_1].Line_No = line_counter

        csv_inf2 = append ( csv_inf2,  csv_inf[loop_1] )

	  }

	}

    return	csv_inf2
}

///
/// 　　　　	メニュー　：operator/gt
///

func Operator_gt ( w http.ResponseWriter , r *http.Request ,csv_inf []type5.Csv_Inf ,match_word string ,match_key []string  )  (csv_inf2 []type5.Csv_Inf ) {

//     IN    w      　　: レスポンスライター
//     IN    r      　　: リクエストパラメータ
//     IN 　csv_inf     : csv情報
//     IN   match_word   : マッチングワード
//     IN 　match_key   : マッチングキー

//     OUT  csv_inf2    : マッチング後のcsv情報

    var line_counter int64

    count := len(csv_inf)    // レコード数ゲット

//    fmt.Fprintf( w, "operator_ne start \n" )  // デバック

///
/// 　　　マッチングする　
///

	csv_inf2 = make([]type5.Csv_Inf, 0)  /// ソートテーブルを確保

    line_counter = 0

	for  loop_1 := 0 ; loop_1 < count ; loop_1++  {

	  if  match_key[loop_1] > match_word  {  ///     マッチング

	    line_counter ++

//        csv_inf[loop_1].Line_No = line_counter

        csv_inf2 = append ( csv_inf2,  csv_inf[loop_1] )

	  }

	}

    return	csv_inf2

}

///
/// 　　　　	メニュー　：operator/ge
///

func Operator_ge ( w http.ResponseWriter , r *http.Request ,csv_inf []type5.Csv_Inf ,match_word string ,match_key []string  )  (csv_inf2 []type5.Csv_Inf ) {

//     IN    w      　　: レスポンスライター
//     IN    r      　　: リクエストパラメータ
//     IN 　csv_inf     : csv情報
//     IN   match_word   : マッチングワード
//     IN 　match_key   : マッチングキー

//     OUT  csv_inf2    : マッチング後のcsv情報

    var line_counter int64

    count := len(csv_inf)    // レコード数ゲット

//    fmt.Fprintf( w, "operator_ne start \n" )  // デバック

///
/// 　　　マッチングする　
///

	csv_inf2 = make([]type5.Csv_Inf, 0)  /// ソートテーブルを確保

    line_counter = 0

	for  loop_1 := 0 ; loop_1 < count ; loop_1++  {

	  if match_key[loop_1] >= match_word {  ///     マッチング

	    line_counter ++

//        csv_inf[loop_1].Line_No = line_counter

        csv_inf2 = append ( csv_inf2,  csv_inf[loop_1] )

	  }

	}

    return	csv_inf2

}

///
/// 　　　　	メニュー　：operator/lt
///

func Operator_lt ( w http.ResponseWriter , r *http.Request ,csv_inf []type5.Csv_Inf ,match_word string ,match_key []string  )  (csv_inf2 []type5.Csv_Inf ) {

//     IN    w      　　: レスポンスライター
//     IN    r      　　: リクエストパラメータ
//     IN 　csv_inf     : csv情報
//     IN   match_word   : マッチングワード
//     IN 　match_key   : マッチングキー

//     OUT  csv_inf2    : マッチング後のcsv情報

    var line_counter int64

    count := len(csv_inf)    // レコード数ゲット

//    fmt.Fprintf( w, "operator_ne start \n" )  // デバック

///
/// 　　　マッチングする　
///

	csv_inf2 = make([]type5.Csv_Inf, 0)  /// ソートテーブルを確保

    line_counter = 0

	for  loop_1 := 0 ; loop_1 < count ; loop_1++  {

	  if  match_key[loop_1]  < match_word  {  ///     マッチング

	    line_counter ++

//        csv_inf[loop_1].Line_No = line_counter

        csv_inf2 = append ( csv_inf2,  csv_inf[loop_1] )

	  }

	}

    return	csv_inf2

}

///
/// 　　　　	メニュー　：operator/le
///

func Operator_le ( w http.ResponseWriter , r *http.Request ,csv_inf []type5.Csv_Inf ,match_word string ,match_key []string  )  (csv_inf2 []type5.Csv_Inf ) {

//     IN    w      　　: レスポンスライター
//     IN    r      　　: リクエストパラメータ
//     IN 　csv_inf     : csv情報
//     IN   match_word   : マッチングワード
//     IN 　match_key   : マッチングキー

//     OUT  csv_inf2    : マッチング後のcsv情報

    var line_counter int64

    count := len(csv_inf)    // レコード数ゲット

//    fmt.Fprintf( w, "operator_ne start \n" )  // デバック

///
/// 　　　マッチングする　
///

	csv_inf2 = make([]type5.Csv_Inf, 0)  /// ソートテーブルを確保

    line_counter = 0

	for  loop_1 := 0 ; loop_1 < count ; loop_1++  {

	  if match_key[loop_1] <= match_word  {  ///     マッチング

	    line_counter ++

//        csv_inf[loop_1].Line_No = line_counter

        csv_inf2 = append ( csv_inf2,  csv_inf[loop_1] )

	  }

	}

    return	csv_inf2

}

///
///     論理演算子のメインルーチン
///


func Operator2( csv_records type5.Csv_Records ,function string ,w http.ResponseWriter, r *http.Request )  ( csv_inf []type5.Csv_Inf ) {

//     IN 　csv_records  : csv情報群
//     IN    function    : ファンクション　
//        　　　　　　　　　＊　and , or
//     IN    w      　　 : レスポンスライター
//     IN    r      　　 : リクエストパラメータ

//     OUT  csv_inf      : 抽出後のcsv情報

//    fmt.Fprintf( w, "operator start \n" )  // デバック

///
///  ファンクションにより各種処理を分岐して行う。
///

	switch function {

      case "or" :

        csv_inf = Operator_or ( w , r ,csv_records  )

      break;

      case "and" :

        csv_inf = Operator_and ( w , r ,csv_records  )

      break;

    }

	return csv_inf

}

///
/// 　　　　	メニュー　：operator/or
///

func Operator_or ( w http.ResponseWriter , r *http.Request ,csv_records type5.Csv_Records   )  (csv_inf []type5.Csv_Inf ) {

//     IN    w      　　: レスポンスライター
//     IN    r      　　: リクエストパラメータ
//     IN 　csv_records : csv情報群

//     OUT  csv_inf    : 抽出後のcsv情報

//    fmt.Fprintf( w, "operator_or start \n" )  // デバック

    csv_inf_wk := make([]type5.Csv_Inf, 0)  /// 検索用のレコードテーブルを確保

    for  index := 0 ; index < int(csv_records.Records_Num) ; index++  {

      csv_inf_wk = append ( csv_inf_wk, csv_records.Records[index]... )

    }

    count := len(csv_inf_wk)    // レコード数ゲット

///
/// 　　　レコードを抽出する　
///

	csv_inf = make([]type5.Csv_Inf, 0)  /// レコード用のテーブルを確保

	skip_check := make([]int ,count)        /// スキップの判定フラグテーブルを確保

	for  loop_1 := 0 ; loop_1 < count ; loop_1++  {

	  if skip_check[loop_1] != 1  {  /// チェック済かのチェック

	    for  loop_2 := 0 ; loop_2 < count ; loop_2 ++ {

	      if csv_inf_wk[loop_1].Line_No ==  csv_inf_wk[loop_2].Line_No  {

	        skip_check[loop_2] = 1

	      }

	    }

	    csv_inf = append ( csv_inf,  csv_inf_wk[loop_1] )

      }

	}

    return	csv_inf

}

///
/// 　　　　	メニュー　：operator/and
///

func Operator_and ( w http.ResponseWriter , r *http.Request ,csv_records type5.Csv_Records   )  (csv_inf []type5.Csv_Inf ) {

//     IN    w      　　: レスポンスライター
//     IN    r      　　: リクエストパラメータ
//     IN 　csv_records : csv情報群

//     OUT  csv_inf    : 抽出後のcsv情報

//    fmt.Fprintf( w, "operator_and start \n" )  // デバック

    var  same_count int64

    csv_inf_wk := make([]type5.Csv_Inf, 0)  /// 検索用のレコードテーブルを確保

    for  index := 0 ; index < int(csv_records.Records_Num) ; index++  {

      csv_inf_wk = append ( csv_inf_wk, csv_records.Records[index]... )

    }

//    fmt.Fprintf( w, "operator_and csv_inf_wk %v\n" ,csv_inf_wk )  // デバック

    count := len(csv_inf_wk)    // レコード数ゲット

///
/// 　　　レコードを抽出する　
///

	csv_inf = make([]type5.Csv_Inf, 0)  /// レコード用のテーブルを確保

	skip_check := make([]int ,count)        /// スキップの判定フラグテーブルを確保

	for  loop_1 := 0 ; loop_1 < count ; loop_1++  {

	  if skip_check[loop_1] != 1  {  /// チェック済かのチェック

	    same_count = 0

	    for  loop_2 := 0 ; loop_2 < count ; loop_2 ++ {

	      if csv_inf_wk[loop_1].Line_No ==  csv_inf_wk[loop_2].Line_No  {

	        skip_check[loop_2] = 1

	        same_count ++

	      }

	    }

//        fmt.Fprintf( w, "operator2/and same_count %v\n" ,same_count )  // デバック

	    if same_count == csv_records.Records_Num  {

	      csv_inf = append ( csv_inf,  csv_inf_wk[loop_1] )

	    }

      }

	}

    return	csv_inf

}