package operator

import (

        "net/http"
//	    "fmt"

	    "github.com/sawaq7/go12_ver1/general/type5"
	    "github.com/sawaq7/go12_ver1/general/datastore5/trans3"

                                    )

///
///     main for comparison operator
///


func Operator( csv_inf []type5.Csv_Inf ,function string ,match_word string ,column_no int ,w http.ResponseWriter, r *http.Request )  ( csv_inf2 []type5.Csv_Inf ) {

//     IN 　csv_inf      :
//     IN    function    : 　
//        　　　　　　　　　　eq ne ge gt le lt
//     IN   match_word   : 　
//     IN   column_no　  : the column no which is maked matching
//     IN    w      　　 : レスポンスライター
//     IN    r      　　 : リクエストパラメータ

//     OUT  csv_inf2     :  csv inf. which is after matching

//    fmt.Fprintf( w, "operator start \n" )
//    fmt.Fprintf( w, "operator function %v\n" ,function )
 //   fmt.Fprintf( w, "operator match_word %v\n" ,match_word )

///
///    get matching key
///

    match_key := trans3.Csv_inf_column ( w ,r ,column_no )

///
///  jump some action in function

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
/// 　　　　	 operator (eq)
///

func Operator_eq ( w http.ResponseWriter , r *http.Request ,csv_inf []type5.Csv_Inf ,match_word string ,match_key []string  )  (csv_inf2 []type5.Csv_Inf ) {

//     IN    w      　　: レスポンスライター
//     IN    r      　　: リクエストパラメータ
//     IN 　csv_inf     :
//     IN   match_word   :
//     IN 　match_key   :

//     OUT  csv_inf2    : csv inf. which is after sorting

    var line_counter int64

    count := len(csv_inf)    //    set record number

//    fmt.Fprintf( w, "operator_eq start \n" )

///
/// 　　　make matching　
///

	csv_inf2 = make([]type5.Csv_Inf, 0)  ///   allocate work area for sort table
    line_counter = 0

	for  loop_1 := 0 ; loop_1 < count ; loop_1++  {

	  if match_word == match_key[loop_1]  {  ///     matching

	    line_counter ++

//        csv_inf[loop_1].Line_No = line_counter

        csv_inf2 = append ( csv_inf2,  csv_inf[loop_1] )

	  }

	}

    return	csv_inf2
}

///
/// 　     operator (ne)
///

func Operator_ne ( w http.ResponseWriter , r *http.Request ,csv_inf []type5.Csv_Inf ,match_word string ,match_key []string  )  (csv_inf2 []type5.Csv_Inf ) {

//     IN    w      　　: レスポンスライター
//     IN    r      　　: リクエストパラメータ
//     IN 　csv_inf     :  csv inf.
//     IN   match_word   :
//     IN 　match_key   :

//     OUT  csv_inf2    : csv inf. which is after matching

    var line_counter int64

    count := len(csv_inf)    //  get record number

//    fmt.Fprintf( w, "operator_ne start \n" )  //

///
/// 　　　make matching　
///

	csv_inf2 = make([]type5.Csv_Inf, 0)  /// allocate work area for sorting
    line_counter = 0

	for  loop_1 := 0 ; loop_1 < count ; loop_1++  {

	  if match_word != match_key[loop_1]  {  ///     make matching

	    line_counter ++

//        csv_inf[loop_1].Line_No = line_counter

        csv_inf2 = append ( csv_inf2,  csv_inf[loop_1] )

	  }

	}

    return	csv_inf2
}

///
/// 　　　　	operator (gt)
///

func Operator_gt ( w http.ResponseWriter , r *http.Request ,csv_inf []type5.Csv_Inf ,match_word string ,match_key []string  )  (csv_inf2 []type5.Csv_Inf ) {

//     IN    w      　　: レスポンスライター
//     IN    r      　　: リクエストパラメータ
//     IN 　csv_inf     :
//     IN   match_word   :
//     IN 　match_key   :

//     OUT  csv_inf2    : csv inf. which is after matching

    var line_counter int64

    count := len(csv_inf)    // レコード数ゲチE��

//    fmt.Fprintf( w, "operator_ne start \n" )

///
/// 　　make matching
///

	csv_inf2 = make([]type5.Csv_Inf, 0)  /// allocate work area for sorting
    line_counter = 0

	for  loop_1 := 0 ; loop_1 < count ; loop_1++  {

	  if  match_key[loop_1] > match_word  {         ///   make matching

	    line_counter ++

//        csv_inf[loop_1].Line_No = line_counter

        csv_inf2 = append ( csv_inf2,  csv_inf[loop_1] )

	  }

	}

    return	csv_inf2

}

///
/// 　　　　  operator (ge)
///

func Operator_ge ( w http.ResponseWriter , r *http.Request ,csv_inf []type5.Csv_Inf ,match_word string ,match_key []string  )  (csv_inf2 []type5.Csv_Inf ) {

//     IN    w      　　: レスポンスライター
//     IN    r      　　: リクエストパラメータ
//     IN 　csv_inf     :
//     IN   match_word   :
//     IN 　match_key   :

//     OUT  csv_inf2    : csv inf. which is after matching

    var line_counter int64

    count := len(csv_inf)

//    fmt.Fprintf( w, "operator_ne start \n" )

///
/// 　　  make matching　
///

	csv_inf2 = make([]type5.Csv_Inf, 0)  /// allocate work area for sorting
    line_counter = 0

	for  loop_1 := 0 ; loop_1 < count ; loop_1++  {

	  if match_key[loop_1] >= match_word {  ///   make matching

	    line_counter ++

//        csv_inf[loop_1].Line_No = line_counter

        csv_inf2 = append ( csv_inf2,  csv_inf[loop_1] )

	  }

	}

    return	csv_inf2

}

///
/// 　　　　  operator (lt)
///

func Operator_lt ( w http.ResponseWriter , r *http.Request ,csv_inf []type5.Csv_Inf ,match_word string ,match_key []string  )  (csv_inf2 []type5.Csv_Inf ) {

//     IN    w      　　: レスポンスライター
//     IN    r      　　: リクエストパラメータ
//     IN 　csv_inf     :
//     IN   match_word   :
//     IN 　match_key   :

//     OUT  csv_inf2    : csv inf. which is after matching

    var line_counter int64

    count := len(csv_inf)    //  get some records

//    fmt.Fprintf( w, "operator_ne start \n" )

///
/// 　　make matching　
///

	csv_inf2 = make([]type5.Csv_Inf, 0)  /// allocate work area for matching
    line_counter = 0

	for  loop_1 := 0 ; loop_1 < count ; loop_1++  {

	  if  match_key[loop_1]  < match_word  {  ///    make matching　

	    line_counter ++

//        csv_inf[loop_1].Line_No = line_counter

        csv_inf2 = append ( csv_inf2,  csv_inf[loop_1] )

	  }

	}

    return	csv_inf2

}

///
/// 　　　　  operator (le)
///

func Operator_le ( w http.ResponseWriter , r *http.Request ,csv_inf []type5.Csv_Inf ,match_word string ,match_key []string  )  (csv_inf2 []type5.Csv_Inf ) {

//     IN    w      　　: レスポンスライター
//     IN    r      　　: リクエストパラメータ
//     IN 　csv_inf     :
//     IN   match_word   :
//     IN 　match_key   :

//     OUT  csv_inf2    : csv inf. which is after matching

    var line_counter int64

    count := len(csv_inf)    //   set record number

//    fmt.Fprintf( w, "operator_ne start \n" )

///
/// 　　　make matching　
///

	csv_inf2 = make([]type5.Csv_Inf, 0)  /// allocate work area for sorting
    line_counter = 0

	for  loop_1 := 0 ; loop_1 < count ; loop_1++  {

	  if match_key[loop_1] <= match_word  {  ///   make matching

	    line_counter ++

//        csv_inf[loop_1].Line_No = line_counter

        csv_inf2 = append ( csv_inf2,  csv_inf[loop_1] )

	  }

	}

    return	csv_inf2

}

///
///     main (logical operator)
///


func Operator2( csv_records type5.Csv_Records ,function string ,w http.ResponseWriter, r *http.Request )  ( csv_inf []type5.Csv_Inf ) {

//     IN 　csv_records  : some csv records
//     IN    function    : 　
//        　　　　　　　　　　and , or
//     IN    w      　　 : レスポンスライター
//     IN    r      　　 : リクエストパラメータ

//     OUT  csv_inf      : csv inf. which is after extracting

//    fmt.Fprintf( w, "operator start \n" )

///
///  jump some action in function

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
/// 　　　　  operator (or)
///

func Operator_or ( w http.ResponseWriter , r *http.Request ,csv_records type5.Csv_Records   )  (csv_inf []type5.Csv_Inf ) {

//     IN    w      　　: レスポンスライター
//     IN    r      　　: リクエストパラメータ
//     IN 　csv_records :

//     OUT  csv_inf    : csv inf. which is after extracting

//    fmt.Fprintf( w, "operator_or start \n" )

    csv_inf_wk := make([]type5.Csv_Inf, 0)  // allocate work area for records

    for  index := 0 ; index < int(csv_records.Records_Num) ; index++  {

      csv_inf_wk = append ( csv_inf_wk, csv_records.Records[index]... )

    }

    count := len(csv_inf_wk)    // レコード数ゲチE��

///
/// 　　　レコードを抽出する　
///

	csv_inf = make([]type5.Csv_Inf, 0)  /// レコード用のチE�Eブルを確俁E
	skip_check := make([]int ,count)        /// スキチE�Eの判定フラグチE�Eブルを確俁E
	for  loop_1 := 0 ; loop_1 < count ; loop_1++  {

	  if skip_check[loop_1] != 1  {  /// チェチE��済かのチェチE��

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
/// 　　　　	operator (and)
///

func Operator_and ( w http.ResponseWriter , r *http.Request ,csv_records type5.Csv_Records   )  (csv_inf []type5.Csv_Inf ) {

//     IN    w      　　: レスポンスライター
//     IN    r      　　: リクエストパラメータ
//     IN 　csv_records :

//     OUT  csv_inf    : csv inf. which is after extracting

//    fmt.Fprintf( w, "operator_and start \n" )


    var  same_count int64

    csv_inf_wk := make([]type5.Csv_Inf, 0)  // allocate work area for records

    for  index := 0 ; index < int(csv_records.Records_Num) ; index++  {

      csv_inf_wk = append ( csv_inf_wk, csv_records.Records[index]... )

    }

//    fmt.Fprintf( w, "operator_and csv_inf_wk %v\n" ,csv_inf_wk )

    count := len(csv_inf_wk)    // get record number

///
/// 　　　extract some records　
///

	csv_inf = make([]type5.Csv_Inf, 0)  ///   allocate work area for records
	skip_check := make([]int ,count)        /// allocate work area for skip flag
	for  loop_1 := 0 ; loop_1 < count ; loop_1++  {

	  if skip_check[loop_1] != 1  {    // judge skip

	    same_count = 0

	    for  loop_2 := 0 ; loop_2 < count ; loop_2 ++ {

	      if csv_inf_wk[loop_1].Line_No ==  csv_inf_wk[loop_2].Line_No  {

	        skip_check[loop_2] = 1

	        same_count ++

	      }

	    }

//        fmt.Fprintf( w, "operator2/and same_count %v\n" ,same_count )

	    if same_count == csv_records.Records_Num  {

	      csv_inf = append ( csv_inf,  csv_inf_wk[loop_1] )

	    }

      }

	}

    return	csv_inf

}
