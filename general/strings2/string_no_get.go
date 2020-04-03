package strings2

import (

//	    "fmt"
	    "strings"
	    "net/http"
	    "regexp"
//	    "sort"

                      )

///
///     斁E���Eより、数字を抽出する
///

func String_no_get ( w http.ResponseWriter, r *http.Request , string_data string  )  ([]string ) {

//     IN    w      　　: レスポンスライター
//     IN    r      　　: リクエストパラメータ
//     IN  string_data  : 斁E���E

//     OUT strings  : 刁E��後�E斁E���E群

//    fmt.Fprintf( w, "string_no_get start \n" )  // チE��チE��

    string_data2 := strings.Split ( string_data, " "  )

    if len(string_data2) == 1 {

	   string_data2 = strings.Split ( string_data, ","  )

	}


//	str_test := "1-3"
//    rep := regexp.MustCompile(`[0-9]`)
//    str := rep.ReplaceAllString(str_test, "C")

//    fmt.Fprintf( w, "string_no_get str %c\n" ,str )  // チE��チE��


//    string_work := string_data2[0]

//    for _, string_workw := range string_work {

//      fmt.Fprintf( w, "string_no_get string_workw %c\n" ,string_workw )  // チE��チE��

//	}

    for _, string_data2w := range string_data2 {

      judge := String_check ( w , r  , `[0-9]`, string_data2w )

      if judge == false {

//	     fmt.Fprintf( w, "string_no_get string isn't numeric \n"  )  // チE��チE��

	     return string_data2

	  }
    }

//    fmt.Fprintf( w, "string_no_get normal end \n" )  // チE��チE��

    return	string_data2
}

///
///     正規表現より、文字�Eを判定すめE///

func String_check ( w http.ResponseWriter, r *http.Request , reg_exp string, string_data string  )  ( bool ) {

//     IN    w      　　: レスポンスライター
//     IN    r      　　: リクエストパラメータ
//     IN  reg_exp      : 正規表現
//     IN  string_data  : 斁E���E

//     OUT judge        : true
//                      : false

//    fmt.Fprintf( w, "string_check start \n" )  // チE��チE��

    judge := regexp.MustCompile( reg_exp ).Match( []byte(string_data) )

//    fmt.Fprintf( w, "string_check normal end \n" )  // チE��チE��

    return	judge
}
