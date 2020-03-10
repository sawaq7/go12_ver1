package sort2

import (

	    "net/http"
//	    "fmt"
	    "general/datastore5/trans3"

	    "general/type5"

                                                )

///
/// 　　　　	csv情報をソートする
///


func Csv_inf(w http.ResponseWriter , r *http.Request ,csv_inf []type5.Csv_Inf ,sort_key_no []int  )  (csv_inf2 []type5.Csv_Inf ) {

//     IN    w      　　: レスポンスライター
//     IN    r      　　: リクエストパラメータ
//     IN 　csv_inf     : csv情報
//     IN 　sort_key_no : ソートキーNO

//     OUT  csv_inf2    : ソート後のcsv情報

//    fmt.Fprintf( w, "sort.csv_inf start \n" )  // デバック

///
///      ループ階層の判定
///

    loop_action := 0

    switch len(sort_key_no) {

          case 1 :

            loop_action = 1

          break;

          case 2 :

            loop_action = 2

          break;

          case 3 :

            loop_action = 3

          break;

      }

//      fmt.Fprintf( w, "sort2.csv_sort : loop_action %v\n", loop_action )  // デバック

///
///      　階層別にソートを行う
///

      switch loop_action {

          case 1 :

            csv_inf2 = Csv_inf_single( w ,r ,csv_inf ,sort_key_no   )   /// 1重ソートする

          break;

          case 2 :

            csv_inf2 = Csv_inf_double( w ,r ,csv_inf ,sort_key_no   )   /// 2重ソートする

          break;

          case 3 :

            csv_inf2 = Csv_inf_triple( w ,r ,csv_inf ,sort_key_no   )   /// 1重ソートする

          break;

      }




    return	csv_inf2
}

///
/// 　　　　	csv情報を1重ソートする
///

func Csv_inf_single( w http.ResponseWriter , r *http.Request ,csv_inf []type5.Csv_Inf ,sort_key_no []int  )  (csv_inf2 []type5.Csv_Inf ) {

//     IN    w      　　: レスポンスライター
//     IN    r      　　: リクエストパラメータ
//     IN 　csv_inf     : csv情報
//     IN 　sort_key_no : ソートキーNO

//     OUT  csv_inf2    : ソート後のcsv情報

    var loop_2_flag ,loop_2_min int

    var line_counter int64

    var key_1_save  string

//    fmt.Fprintf( w, "sort.csv_inf_single start \n" )  // デバック

    count := len(csv_inf)    // レコード数ゲット

//     fmt.Fprintf( w, "sort2.sort.csv_inf_single : count %v\n", count )  // デバック

     /// ソートテーブルを確保

     sort_key1 := make( []string, count )
     sort_key2 := make( []string, count )

///
/// 　　　ソートキーをセットする　
///

    for pos , sort_key_now := range sort_key_no {

      string_wk := trans3.Csv_inf_column ( w ,r ,sort_key_now )   /// 　　ソートキーをセット

      for pos2 , string_wkw := range string_wk {

        if pos  == 0  {

          sort_key1[pos2] = string_wkw

        } else {

          sort_key2[pos2] = string_wkw

        }
//        fmt.Fprintf( w, "sort2.csv_sort : string_wkw %v\n", string_wkw )  // デバック
      }
    }

///
/// 　　　ソートする　
///
	csv_inf2 = make([]type5.Csv_Inf, 0)  /// ソートテーブルを確保

	skip_check := make([]int ,count)        /// スキップの判定フラグテーブルを確保

    line_counter = 0

	for  loop_1 := 0 ; loop_1 < count ; loop_1++  {

	  loop_2_flag = -1

	  for  loop_2 := 0 ; loop_2 < count ; loop_2 ++ {

	    if skip_check[loop_2] != 1  {  /// ソート済みか否かのチェック

	      if loop_2_flag  == -1  {     ///  初期値のセット

	        loop_2_min = loop_2
	        key_1_save = sort_key1[loop_2]

	        loop_2_flag = 0

	      /// 最小値、再セット

	      }  else if key_1_save >  sort_key1[loop_2]  {

	        loop_2_min = loop_2
	        key_1_save = sort_key1[loop_2]

//            fmt.Fprintf( w, "sort2.csv_sort_single : key_1_save_single %v\n", key_1_save )  // デバック

	      }

	    }

      }

///
///   min値をテーブルにセット
///
      line_counter ++
      csv_inf[loop_2_min].Line_No = line_counter

      csv_inf2 = append ( csv_inf2,  csv_inf[loop_2_min] )

///
///  ソート済フラグをセット
///

      skip_check[loop_2_min] = 1

	}

    return	csv_inf2
}

///
/// 　　　　	csv情報を2重ソートする
///

func Csv_inf_double( w http.ResponseWriter , r *http.Request ,csv_inf []type5.Csv_Inf ,sort_key_no []int  )  (csv_inf2 []type5.Csv_Inf ) {

//     IN    w      　　: レスポンスライター
//     IN    r      　　: リクエストパラメータ
//     IN 　csv_inf     : csv情報
//     IN 　sort_key_no : ソートキーNO

//     OUT  csv_inf2    : ソート後のcsv情報

    var loop_2_flag ,loop_2_min int

    var line_counter int64

    var key_1_save ,key_2_save string

//    fmt.Fprintf( w, "sort.csv_inf_double start \n" )  // デバック

    count := len(csv_inf)    // レコード数ゲット

//     fmt.Fprintf( w, "sort2.sort.csv_inf_double : count %v\n", count )  // デバック

     /// ソートテーブルを確保

     sort_key1 := make( []string, count )
     sort_key2 := make( []string, count )

///
/// 　　　ソートキーをセットする　
///

    for pos , sort_key_now := range sort_key_no {

      string_wk := trans3.Csv_inf_column ( w ,r ,sort_key_now )   /// 　　ソートキーをセット

      for pos2 , string_wkw := range string_wk {

        if pos  == 0  {

          sort_key1[pos2] = string_wkw

        } else  {

          sort_key2[pos2] = string_wkw

        }

      }
    }

///
/// 　　　ソートする　
///
	csv_inf2 = make([]type5.Csv_Inf, 0)  /// ソートテーブルを確保

	skip_check := make([]int ,count)        /// スキップの判定フラグテーブルを確保

    line_counter = 0

	for  loop_1 := 0 ; loop_1 < count ; loop_1++  {

	  loop_2_flag = -1

	  for  loop_2 := 0 ; loop_2 < count ; loop_2 ++ {

	    if skip_check[loop_2] != 1  {  /// ソート済みか否かのチェック

	      if loop_2_flag  == -1  {     ///  初期値のセット

	        loop_2_min = loop_2
	        key_1_save = sort_key1[loop_2]
	        key_2_save = sort_key2[loop_2]

	        loop_2_flag = 0

	      /// 最小値、再セット

	      }  else if ( key_1_save >  sort_key1[loop_2] )                                    ||
	                 ( key_1_save == sort_key1[loop_2] && key_2_save  > sort_key2[loop_2] )     {

	        loop_2_min = loop_2
	        key_1_save = sort_key1[loop_2]
	        key_2_save = sort_key2[loop_2]

//            fmt.Fprintf( w, "sort2.csv_sort_double : key_1_save %v\n", key_1_save )  // デバック
//            fmt.Fprintf( w, "sort2.csv_sort_double : key_2_save %v\n", key_2_save )  // デバック

	      }

	    }

      }

///
///   min値をテーブルにセット
///
      line_counter ++
      csv_inf[loop_2_min].Line_No = line_counter

      csv_inf2 = append ( csv_inf2,  csv_inf[loop_2_min] )

///
///  ソート済フラグをセット
///

      skip_check[loop_2_min] = 1

	}

    return	csv_inf2
}

///
/// 　　　　	csv情報を3重ソートする
///

func Csv_inf_triple( w http.ResponseWriter , r *http.Request ,csv_inf []type5.Csv_Inf ,sort_key_no []int  )  (csv_inf2 []type5.Csv_Inf ) {

//     IN    w      　　: レスポンスライター
//     IN    r      　　: リクエストパラメータ
//     IN 　csv_inf     : csv情報
//     IN 　sort_key_no : ソートキーNO

//     OUT  csv_inf2    : ソート後のcsv情報

    var loop_2_flag ,loop_2_min int

    var line_counter int64

    var key_1_save ,key_2_save ,key_3_save string

//    fmt.Fprintf( w, "sort.csv_inf_triple start \n" )  // デバック

    count := len(csv_inf)    // レコード数ゲット

//     fmt.Fprintf( w, "sort2.sort.csv_inf_triple : count %v\n", count )  // デバック

     /// ソートテーブルを確保

     sort_key1 := make( []string, count )
     sort_key2 := make( []string, count )
     sort_key3 := make( []string, count )

///
/// 　　　ソートキーをセットする　
///

    for pos , sort_key_now := range sort_key_no {

      string_wk := trans3.Csv_inf_column ( w ,r ,sort_key_now )   /// 　　ソートキーをセット

      for pos2 , string_wkw := range string_wk {

        if pos  == 0  {

          sort_key1[pos2] = string_wkw

        } else if pos == 1 {

          sort_key2[pos2] = string_wkw

        } else {

          sort_key3[pos2] = string_wkw

        }
//        fmt.Fprintf( w, "sort2.csv_sort : string_wkw %v\n", string_wkw )  // デバック
      }
    }

///
/// 　　　ソートする　
///
	csv_inf2 = make([]type5.Csv_Inf, 0)  /// ソートテーブルを確保

	skip_check := make([]int ,count)        /// スキップの判定フラグテーブルを確保

    line_counter = 0

	for  loop_1 := 0 ; loop_1 < count ; loop_1++  {

	  loop_2_flag = -1

	  for  loop_2 := 0 ; loop_2 < count ; loop_2 ++ {

	    if skip_check[loop_2] != 1  {  /// ソート済みか否かのチェック

	      if loop_2_flag  == -1  {     ///  初期値のセット

	        loop_2_min = loop_2
	        key_1_save = sort_key1[loop_2]
	        key_2_save = sort_key2[loop_2]
	        key_3_save = sort_key3[loop_2]

	        loop_2_flag = 0

	      /// 最小値、再セット

	      }  else if ( key_1_save >  sort_key1[loop_2] )                                      ||

	                 ( key_1_save == sort_key1[loop_2] && key_2_save  > sort_key2[loop_2] )   ||

	                 ( key_1_save == sort_key1[loop_2] && key_2_save == sort_key2[loop_2] && key_3_save  > sort_key3[loop_2] ) {

	        loop_2_min = loop_2
	        key_1_save = sort_key1[loop_2]
	        key_2_save = sort_key2[loop_2]
	        key_3_save = sort_key3[loop_2]

//            fmt.Fprintf( w, "sort2.csv_sort : key_1_save %v\n", key_1_save )  // デバック
//            fmt.Fprintf( w, "sort2.csv_sort : key_2_save %v\n", key_2_save )  // デバック
//            fmt.Fprintf( w, "sort2.csv_sort : key_3_save %v\n", key_3_save )  // デバック

	      }

	    }

      }

///
///   min値をテーブルにセット
///
      line_counter ++
      csv_inf[loop_2_min].Line_No = line_counter

      csv_inf2 = append ( csv_inf2,  csv_inf[loop_2_min] )

///
///  ソート済フラグをセット
///

      skip_check[loop_2_min] = 1

	}

    return	csv_inf2
}
