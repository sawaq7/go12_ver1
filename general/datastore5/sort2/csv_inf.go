package sort2

import (

	    "net/http"
//	    "fmt"
	    "github.com/sawaq7/go12_ver1/general/datastore5/trans3"

	    "github.com/sawaq7/go12_ver1/general/type5"

                                                )

///
/// 　　　　	 sort csv_inf


func Csv_inf(w http.ResponseWriter , r *http.Request ,csv_inf []type5.Csv_Inf ,sort_key_no []int  )  (csv_inf2 []type5.Csv_Inf ) {

//     IN    w      　　: レスポンスライター
//     IN    r      　　: リクエストパラメータ
//     IN 　csv_inf     :
//     IN 　sort_key_no :

//     OUT  csv_inf2    :   it's the csv_inf which is  sorted

//    fmt.Fprintf( w, "sort.csv_inf start \n" )

///
///      judge loop levels

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

//      fmt.Fprintf( w, "sort2.csv_sort : loop_action %v\n", loop_action )

///
///      　階層別にソートを行う
///

      switch loop_action {

          case 1 :

            csv_inf2 = Csv_inf_single( w ,r ,csv_inf ,sort_key_no   )   ///  sort of one level
          break;

          case 2 :

            csv_inf2 = Csv_inf_double( w ,r ,csv_inf ,sort_key_no   )   ///    sort of two levels
          break;

          case 3 :

            csv_inf2 = Csv_inf_triple( w ,r ,csv_inf ,sort_key_no   )   ///   sort of three levels
          break;

      }




    return	csv_inf2
}

///
/// 　　　　	sort of one level

func Csv_inf_single( w http.ResponseWriter , r *http.Request ,csv_inf []type5.Csv_Inf ,sort_key_no []int  )  (csv_inf2 []type5.Csv_Inf ) {

//     IN    w      　　: レスポンスライター
//     IN    r      　　: リクエストパラメータ
//     IN 　csv_inf     :
//     IN 　sort_key_no :

//     OUT  csv_inf2    : it's the csv_inf which is  sorted

    var loop_2_flag ,loop_2_min int

    var line_counter int64

    var key_1_save  string

//    fmt.Fprintf( w, "sort.csv_inf_single start \n" )

    count := len(csv_inf)    //   get record number

//     fmt.Fprintf( w, "sort2.sort.csv_inf_single : count %v\n", count )  // チE��チE��

     /// ソートテーブルを確俁E
     sort_key1 := make( []string, count )
     sort_key2 := make( []string, count )

///
/// 　　　sort of two levels　
///

    for pos , sort_key_now := range sort_key_no {

      string_wk := trans3.Csv_inf_column ( w ,r ,sort_key_now )   /// 　　ソートキーをセチE��

      for pos2 , string_wkw := range string_wk {

        if pos  == 0  {

          sort_key1[pos2] = string_wkw

        } else {

          sort_key2[pos2] = string_wkw

        }
//        fmt.Fprintf( w, "sort2.csv_sort : string_wkw %v\n", string_wkw )  // チE��チE��
      }
    }

///
/// 　　　ソートする　
///
	csv_inf2 = make([]type5.Csv_Inf, 0)  /// ソートテーブルを確俁E
	skip_check := make([]int ,count)        /// スキチE�Eの判定フラグチE�Eブルを確俁E
    line_counter = 0

	for  loop_1 := 0 ; loop_1 < count ; loop_1++  {

	  loop_2_flag = -1

	  for  loop_2 := 0 ; loop_2 < count ; loop_2 ++ {

	    if skip_check[loop_2] != 1  {  /// ソート済みか否か�EチェチE��

	      if loop_2_flag  == -1  {     ///  初期値のセチE��

	        loop_2_min = loop_2
	        key_1_save = sort_key1[loop_2]

	        loop_2_flag = 0

	      /// 最小値、�EセチE��

	      }  else if key_1_save >  sort_key1[loop_2]  {

	        loop_2_min = loop_2
	        key_1_save = sort_key1[loop_2]

//            fmt.Fprintf( w, "sort2.csv_sort_single : key_1_save_single %v\n", key_1_save )  // チE��チE��

	      }

	    }

      }

///
///     set min value in table
///
      line_counter ++
      csv_inf[loop_2_min].Line_No = line_counter

      csv_inf2 = append ( csv_inf2,  csv_inf[loop_2_min] )

///
///  ソート済フラグをセチE��
///

      skip_check[loop_2_min] = 1

	}

    return	csv_inf2
}

///
/// 　　　　	sort of two levels
///

func Csv_inf_double( w http.ResponseWriter , r *http.Request ,csv_inf []type5.Csv_Inf ,sort_key_no []int  )  (csv_inf2 []type5.Csv_Inf ) {

//     IN    w      　　: レスポンスライター
//     IN    r      　　: リクエストパラメータ
//     IN 　csv_inf     : csv惁E��
//     IN 　sort_key_no : ソートキーNO

//     OUT  csv_inf2    : it's the csv_inf which is  sorted

    var loop_2_flag ,loop_2_min int

    var line_counter int64

    var key_1_save ,key_2_save string

//    fmt.Fprintf( w, "sort.csv_inf_double start \n" )

    count := len(csv_inf)    // get record number

//     fmt.Fprintf( w, "sort2.sort.csv_inf_double : count %v\n", count )

     /// ソートテーブルを確俁E
     sort_key1 := make( []string, count )
     sort_key2 := make( []string, count )

///
/// 　　　set sort key　
///

    for pos , sort_key_now := range sort_key_no {

      string_wk := trans3.Csv_inf_column ( w ,r ,sort_key_now )   /// 　　set sort key

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
	csv_inf2 = make([]type5.Csv_Inf, 0)  //   allocate work area for sorting
	skip_check := make([]int ,count)     //   allocate work area for skiping
    line_counter = 0

	for  loop_1 := 0 ; loop_1 < count ; loop_1++  {

	  loop_2_flag = -1

	  for  loop_2 := 0 ; loop_2 < count ; loop_2 ++ {

	    if skip_check[loop_2] != 1  {  ///  whether or not loop2 skips

	      if loop_2_flag  == -1  {     ///  loop2 flag ini.

	        loop_2_min = loop_2
	        key_1_save = sort_key1[loop_2]
	        key_2_save = sort_key2[loop_2]

	        loop_2_flag = 0

	      ///    set min value

	      }  else if ( key_1_save >  sort_key1[loop_2] )                                    ||
	                 ( key_1_save == sort_key1[loop_2] && key_2_save  > sort_key2[loop_2] )     {

	        loop_2_min = loop_2
	        key_1_save = sort_key1[loop_2]
	        key_2_save = sort_key2[loop_2]

//            fmt.Fprintf( w, "sort2.csv_sort_double : key_1_save %v\n", key_1_save )
//            fmt.Fprintf( w, "sort2.csv_sort_double : key_2_save %v\n", key_2_save )

	      }

	    }

      }

///
///   set min value in table
///
      line_counter ++
      csv_inf[loop_2_min].Line_No = line_counter

      csv_inf2 = append ( csv_inf2,  csv_inf[loop_2_min] )

      skip_check[loop_2_min] = 1    //  set skip flag

	}

    return	csv_inf2
}

///
/// 　　　　	sort of three levels
///

func Csv_inf_triple( w http.ResponseWriter , r *http.Request ,csv_inf []type5.Csv_Inf ,sort_key_no []int  )  (csv_inf2 []type5.Csv_Inf ) {

//     IN    w      　　: レスポンスライター
//     IN    r      　　: リクエストパラメータ
//     IN 　csv_inf     :
//     IN 　sort_key_no :

//     OUT  csv_inf2    : it's the csv_inf which is  sorted

    var loop_2_flag ,loop_2_min int

    var line_counter int64

    var key_1_save ,key_2_save ,key_3_save string

//    fmt.Fprintf( w, "sort.csv_inf_triple start \n" )

    count := len(csv_inf)    // get record number

//     fmt.Fprintf( w, "sort2.sort.csv_inf_triple : count %v\n", count )

///
///    allocate work area for records
///

     sort_key1 := make( []string, count )
     sort_key2 := make( []string, count )
     sort_key3 := make( []string, count )

///
/// 　　　ソートキーをセチE��する　
///

    for pos , sort_key_now := range sort_key_no {

      string_wk := trans3.Csv_inf_column ( w ,r ,sort_key_now ) /// 　　set sort key

      for pos2 , string_wkw := range string_wk {

        if pos  == 0  {

          sort_key1[pos2] = string_wkw

        } else if pos == 1 {

          sort_key2[pos2] = string_wkw

        } else {

          sort_key3[pos2] = string_wkw

        }
//        fmt.Fprintf( w, "sort2.csv_sort : string_wkw %v\n", string_wkw )
      }
    }

///
/// 　　　ソートする　
///
	csv_inf2 = make([]type5.Csv_Inf, 0)  /// allocate work area for sorting
	skip_check := make([]int ,count)     ///  allocate work area for juding wheter or not skip
    line_counter = 0

	for  loop_1 := 0 ; loop_1 < count ; loop_1++  {

	  loop_2_flag = -1

	  for  loop_2 := 0 ; loop_2 < count ; loop_2 ++ {

	    if skip_check[loop_2] != 1  {  //whether or not loop2 skips

	      if loop_2_flag  == -1  {     //  loop2 flag ini.

	        loop_2_min = loop_2
	        key_1_save = sort_key1[loop_2]
	        key_2_save = sort_key2[loop_2]
	        key_3_save = sort_key3[loop_2]

	        loop_2_flag = 0

	      /// 最小値、�EセチE��

	      }  else if ( key_1_save >  sort_key1[loop_2] )                                      ||

	                 ( key_1_save == sort_key1[loop_2] && key_2_save  > sort_key2[loop_2] )   ||

	                 ( key_1_save == sort_key1[loop_2] && key_2_save == sort_key2[loop_2] && key_3_save  > sort_key3[loop_2] ) {

	        loop_2_min = loop_2
	        key_1_save = sort_key1[loop_2]
	        key_2_save = sort_key2[loop_2]
	        key_3_save = sort_key3[loop_2]

//            fmt.Fprintf( w, "sort2.csv_sort : key_1_save %v\n", key_1_save )
//            fmt.Fprintf( w, "sort2.csv_sort : key_2_save %v\n", key_2_save )
//            fmt.Fprintf( w, "sort2.csv_sort : key_3_save %v\n", key_3_save )

	      }

	    }

      }

///
///   set min value in table
///
      line_counter ++

      csv_inf[loop_2_min].Line_No = line_counter

      csv_inf2 = append ( csv_inf2,  csv_inf[loop_2_min] )

      skip_check[loop_2_min] = 1     ///      set skip flag

	}

    return	csv_inf2
}
