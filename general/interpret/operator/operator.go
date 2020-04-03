package operator

import (

        "net/http"
//	    "fmt"

	    "github.com/sawaq7/go12_ver1/general/type5"
	    "github.com/sawaq7/go12_ver1/general/datastore5/trans3"

                                    )

///
///     æ¯è¼E¼ç®å­ãEã¡ã¤ã³ã«ã¼ãã³
///


func Operator( csv_inf []type5.Csv_Inf ,function string ,match_word string ,column_no int ,w http.ResponseWriter, r *http.Request )  ( csv_inf2 []type5.Csv_Inf ) {

//     IN ãcsv_inf      : ã«ã¬ã³ããEcsvæE ±
//     IN    function    : ãã¡ã³ã¯ã·ã§ã³ã
//        ãããããããããEãeq ne ge gt le lt
//     IN   match_word   : ãããã³ã°ã¯ã¼ãã
//     IN   column_noã  : ãããã³ã°å¯¾è±¡ã®è¡NO
//     IN    w      ãã : ã¬ã¹ãã³ã¹ã©ã¤ã¿ã¼
//     IN    r      ãã : ãªã¯ã¨ã¹ããã©ã¡ã¼ã¿

//     OUT  csv_inf2     : ãããã³ã°å¾ãEcsvæE ±

//    fmt.Fprintf( w, "operator start \n" )  // ãEãE¯
//    fmt.Fprintf( w, "operator function %v\n" ,function )  // ãEãE¯
 //   fmt.Fprintf( w, "operator match_word %v\n" ,match_word )  // ãEãE¯

///
///     ãããã³ã°ã­ã¼ãã²ãE
///

    match_key := trans3.Csv_inf_column ( w ,r ,column_no )

///
///  ãã¡ã³ã¯ã·ã§ã³ã«ããåE¨®å¦çEåE²ãã¦è¡ããE///

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
/// ãããã	ã¡ãã¥ã¼ãEoperator/eq
///

func Operator_eq ( w http.ResponseWriter , r *http.Request ,csv_inf []type5.Csv_Inf ,match_word string ,match_key []string  )  (csv_inf2 []type5.Csv_Inf ) {

//     IN    w      ãã: ã¬ã¹ãã³ã¹ã©ã¤ã¿ã¼
//     IN    r      ãã: ãªã¯ã¨ã¹ããã©ã¡ã¼ã¿
//     IN ãcsv_inf     : csvæE ±
//     IN   match_word   : ãããã³ã°ã¯ã¼ãE//     IN ãmatch_key   : ãããã³ã°ã­ã¼

//     OUT  csv_inf2    : ã½ã¼ãå¾ãEcsvæE ±

    var line_counter int64

    count := len(csv_inf)    // ã¬ã³ã¼ãæ°ã²ãE

//    fmt.Fprintf( w, "operator_eq start \n" )  // ãEãE¯

///
/// ããããããã³ã°ããã
///
	csv_inf2 = make([]type5.Csv_Inf, 0)  /// ã½ã¼ããã¼ãã«ãç¢ºä¿E
    line_counter = 0

	for  loop_1 := 0 ; loop_1 < count ; loop_1++  {

	  if match_word == match_key[loop_1]  {  ///     ãããã³ã°

	    line_counter ++

//        csv_inf[loop_1].Line_No = line_counter

        csv_inf2 = append ( csv_inf2,  csv_inf[loop_1] )

	  }

	}

    return	csv_inf2
}

///
/// ãããã	ã¡ãã¥ã¼ãEoperator/ne
///

func Operator_ne ( w http.ResponseWriter , r *http.Request ,csv_inf []type5.Csv_Inf ,match_word string ,match_key []string  )  (csv_inf2 []type5.Csv_Inf ) {

//     IN    w      ãã: ã¬ã¹ãã³ã¹ã©ã¤ã¿ã¼
//     IN    r      ãã: ãªã¯ã¨ã¹ããã©ã¡ã¼ã¿
//     IN ãcsv_inf     : csvæE ±
//     IN   match_word   : ãããã³ã°ã¯ã¼ãE//     IN ãmatch_key   : ãããã³ã°ã­ã¼

//     OUT  csv_inf2    : ãããã³ã°å¾ãEcsvæE ±

    var line_counter int64

    count := len(csv_inf)    // ã¬ã³ã¼ãæ°ã²ãE

//    fmt.Fprintf( w, "operator_ne start \n" )  // ãEãE¯

///
/// ããããããã³ã°ããã
///

	csv_inf2 = make([]type5.Csv_Inf, 0)  /// ã½ã¼ããã¼ãã«ãç¢ºä¿E
    line_counter = 0

	for  loop_1 := 0 ; loop_1 < count ; loop_1++  {

	  if match_word != match_key[loop_1]  {  ///     ãããã³ã°

	    line_counter ++

//        csv_inf[loop_1].Line_No = line_counter

        csv_inf2 = append ( csv_inf2,  csv_inf[loop_1] )

	  }

	}

    return	csv_inf2
}

///
/// ãããã	ã¡ãã¥ã¼ãEoperator/gt
///

func Operator_gt ( w http.ResponseWriter , r *http.Request ,csv_inf []type5.Csv_Inf ,match_word string ,match_key []string  )  (csv_inf2 []type5.Csv_Inf ) {

//     IN    w      ãã: ã¬ã¹ãã³ã¹ã©ã¤ã¿ã¼
//     IN    r      ãã: ãªã¯ã¨ã¹ããã©ã¡ã¼ã¿
//     IN ãcsv_inf     : csvæE ±
//     IN   match_word   : ãããã³ã°ã¯ã¼ãE//     IN ãmatch_key   : ãããã³ã°ã­ã¼

//     OUT  csv_inf2    : ãããã³ã°å¾ãEcsvæE ±

    var line_counter int64

    count := len(csv_inf)    // ã¬ã³ã¼ãæ°ã²ãE

//    fmt.Fprintf( w, "operator_ne start \n" )  // ãEãE¯

///
/// ããããããã³ã°ããã
///

	csv_inf2 = make([]type5.Csv_Inf, 0)  /// ã½ã¼ããã¼ãã«ãç¢ºä¿E
    line_counter = 0

	for  loop_1 := 0 ; loop_1 < count ; loop_1++  {

	  if  match_key[loop_1] > match_word  {  ///     ãããã³ã°

	    line_counter ++

//        csv_inf[loop_1].Line_No = line_counter

        csv_inf2 = append ( csv_inf2,  csv_inf[loop_1] )

	  }

	}

    return	csv_inf2

}

///
/// ãããã	ã¡ãã¥ã¼ãEoperator/ge
///

func Operator_ge ( w http.ResponseWriter , r *http.Request ,csv_inf []type5.Csv_Inf ,match_word string ,match_key []string  )  (csv_inf2 []type5.Csv_Inf ) {

//     IN    w      ãã: ã¬ã¹ãã³ã¹ã©ã¤ã¿ã¼
//     IN    r      ãã: ãªã¯ã¨ã¹ããã©ã¡ã¼ã¿
//     IN ãcsv_inf     : csvæE ±
//     IN   match_word   : ãããã³ã°ã¯ã¼ãE//     IN ãmatch_key   : ãããã³ã°ã­ã¼

//     OUT  csv_inf2    : ãããã³ã°å¾ãEcsvæE ±

    var line_counter int64

    count := len(csv_inf)    // ã¬ã³ã¼ãæ°ã²ãE

//    fmt.Fprintf( w, "operator_ne start \n" )  // ãEãE¯

///
/// ããããããã³ã°ããã
///

	csv_inf2 = make([]type5.Csv_Inf, 0)  /// ã½ã¼ããã¼ãã«ãç¢ºä¿E
    line_counter = 0

	for  loop_1 := 0 ; loop_1 < count ; loop_1++  {

	  if match_key[loop_1] >= match_word {  ///     ãããã³ã°

	    line_counter ++

//        csv_inf[loop_1].Line_No = line_counter

        csv_inf2 = append ( csv_inf2,  csv_inf[loop_1] )

	  }

	}

    return	csv_inf2

}

///
/// ãããã	ã¡ãã¥ã¼ãEoperator/lt
///

func Operator_lt ( w http.ResponseWriter , r *http.Request ,csv_inf []type5.Csv_Inf ,match_word string ,match_key []string  )  (csv_inf2 []type5.Csv_Inf ) {

//     IN    w      ãã: ã¬ã¹ãã³ã¹ã©ã¤ã¿ã¼
//     IN    r      ãã: ãªã¯ã¨ã¹ããã©ã¡ã¼ã¿
//     IN ãcsv_inf     : csvæE ±
//     IN   match_word   : ãããã³ã°ã¯ã¼ãE//     IN ãmatch_key   : ãããã³ã°ã­ã¼

//     OUT  csv_inf2    : ãããã³ã°å¾ãEcsvæE ±

    var line_counter int64

    count := len(csv_inf)    // ã¬ã³ã¼ãæ°ã²ãE

//    fmt.Fprintf( w, "operator_ne start \n" )  // ãEãE¯

///
/// ããããããã³ã°ããã
///

	csv_inf2 = make([]type5.Csv_Inf, 0)  /// ã½ã¼ããã¼ãã«ãç¢ºä¿E
    line_counter = 0

	for  loop_1 := 0 ; loop_1 < count ; loop_1++  {

	  if  match_key[loop_1]  < match_word  {  ///     ãããã³ã°

	    line_counter ++

//        csv_inf[loop_1].Line_No = line_counter

        csv_inf2 = append ( csv_inf2,  csv_inf[loop_1] )

	  }

	}

    return	csv_inf2

}

///
/// ãããã	ã¡ãã¥ã¼ãEoperator/le
///

func Operator_le ( w http.ResponseWriter , r *http.Request ,csv_inf []type5.Csv_Inf ,match_word string ,match_key []string  )  (csv_inf2 []type5.Csv_Inf ) {

//     IN    w      ãã: ã¬ã¹ãã³ã¹ã©ã¤ã¿ã¼
//     IN    r      ãã: ãªã¯ã¨ã¹ããã©ã¡ã¼ã¿
//     IN ãcsv_inf     : csvæE ±
//     IN   match_word   : ãããã³ã°ã¯ã¼ãE//     IN ãmatch_key   : ãããã³ã°ã­ã¼

//     OUT  csv_inf2    : ãããã³ã°å¾ãEcsvæE ±

    var line_counter int64

    count := len(csv_inf)    // ã¬ã³ã¼ãæ°ã²ãE

//    fmt.Fprintf( w, "operator_ne start \n" )  // ãEãE¯

///
/// ããããããã³ã°ããã
///

	csv_inf2 = make([]type5.Csv_Inf, 0)  /// ã½ã¼ããã¼ãã«ãç¢ºä¿E
    line_counter = 0

	for  loop_1 := 0 ; loop_1 < count ; loop_1++  {

	  if match_key[loop_1] <= match_word  {  ///     ãããã³ã°

	    line_counter ++

//        csv_inf[loop_1].Line_No = line_counter

        csv_inf2 = append ( csv_inf2,  csv_inf[loop_1] )

	  }

	}

    return	csv_inf2

}

///
///     è«çæ¼ç®å­ãEã¡ã¤ã³ã«ã¼ãã³
///


func Operator2( csv_records type5.Csv_Records ,function string ,w http.ResponseWriter, r *http.Request )  ( csv_inf []type5.Csv_Inf ) {

//     IN ãcsv_records  : csvæE ±ç¾¤
//     IN    function    : ãã¡ã³ã¯ã·ã§ã³ã
//        ãããããããããEãand , or
//     IN    w      ãã : ã¬ã¹ãã³ã¹ã©ã¤ã¿ã¼
//     IN    r      ãã : ãªã¯ã¨ã¹ããã©ã¡ã¼ã¿

//     OUT  csv_inf      : æ½åºå¾ãEcsvæE ±

//    fmt.Fprintf( w, "operator start \n" )  // ãEãE¯

///
///  ãã¡ã³ã¯ã·ã§ã³ã«ããåE¨®å¦çEåE²ãã¦è¡ããE///

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
/// ãããã	ã¡ãã¥ã¼ãEoperator/or
///

func Operator_or ( w http.ResponseWriter , r *http.Request ,csv_records type5.Csv_Records   )  (csv_inf []type5.Csv_Inf ) {

//     IN    w      ãã: ã¬ã¹ãã³ã¹ã©ã¤ã¿ã¼
//     IN    r      ãã: ãªã¯ã¨ã¹ããã©ã¡ã¼ã¿
//     IN ãcsv_records : csvæE ±ç¾¤

//     OUT  csv_inf    : æ½åºå¾ãEcsvæE ±

//    fmt.Fprintf( w, "operator_or start \n" )  // ãEãE¯

    csv_inf_wk := make([]type5.Csv_Inf, 0)  /// æ¤ç´¢ç¨ã®ã¬ã³ã¼ããã¼ãã«ãç¢ºä¿E
    for  index := 0 ; index < int(csv_records.Records_Num) ; index++  {

      csv_inf_wk = append ( csv_inf_wk, csv_records.Records[index]... )

    }

    count := len(csv_inf_wk)    // ã¬ã³ã¼ãæ°ã²ãE

///
/// ãããã¬ã³ã¼ããæ½åºããã
///

	csv_inf = make([]type5.Csv_Inf, 0)  /// ã¬ã³ã¼ãç¨ã®ãEEãã«ãç¢ºä¿E
	skip_check := make([]int ,count)        /// ã¹ã­ãEEã®å¤å®ãã©ã°ãEEãã«ãç¢ºä¿E
	for  loop_1 := 0 ; loop_1 < count ; loop_1++  {

	  if skip_check[loop_1] != 1  {  /// ãã§ãE¯æ¸ãã®ãã§ãE¯

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
/// ãããã	ã¡ãã¥ã¼ãEoperator/and
///

func Operator_and ( w http.ResponseWriter , r *http.Request ,csv_records type5.Csv_Records   )  (csv_inf []type5.Csv_Inf ) {

//     IN    w      ãã: ã¬ã¹ãã³ã¹ã©ã¤ã¿ã¼
//     IN    r      ãã: ãªã¯ã¨ã¹ããã©ã¡ã¼ã¿
//     IN ãcsv_records : csvæE ±ç¾¤

//     OUT  csv_inf    : æ½åºå¾ãEcsvæE ±

//    fmt.Fprintf( w, "operator_and start \n" )  // ãEãE¯

    var  same_count int64

    csv_inf_wk := make([]type5.Csv_Inf, 0)  /// æ¤ç´¢ç¨ã®ã¬ã³ã¼ããã¼ãã«ãç¢ºä¿E
    for  index := 0 ; index < int(csv_records.Records_Num) ; index++  {

      csv_inf_wk = append ( csv_inf_wk, csv_records.Records[index]... )

    }

//    fmt.Fprintf( w, "operator_and csv_inf_wk %v\n" ,csv_inf_wk )  // ãEãE¯

    count := len(csv_inf_wk)    // ã¬ã³ã¼ãæ°ã²ãE

///
/// ãããã¬ã³ã¼ããæ½åºããã
///

	csv_inf = make([]type5.Csv_Inf, 0)  /// ã¬ã³ã¼ãç¨ã®ãEEãã«ãç¢ºä¿E
	skip_check := make([]int ,count)        /// ã¹ã­ãEEã®å¤å®ãã©ã°ãEEãã«ãç¢ºä¿E
	for  loop_1 := 0 ; loop_1 < count ; loop_1++  {

	  if skip_check[loop_1] != 1  {  /// ãã§ãE¯æ¸ãã®ãã§ãE¯

	    same_count = 0

	    for  loop_2 := 0 ; loop_2 < count ; loop_2 ++ {

	      if csv_inf_wk[loop_1].Line_No ==  csv_inf_wk[loop_2].Line_No  {

	        skip_check[loop_2] = 1

	        same_count ++

	      }

	    }

//        fmt.Fprintf( w, "operator2/and same_count %v\n" ,same_count )  // ãEãE¯

	    if same_count == csv_records.Records_Num  {

	      csv_inf = append ( csv_inf,  csv_inf_wk[loop_1] )

	    }

      }

	}

    return	csv_inf

}
