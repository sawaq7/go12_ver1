package strings2

import (

//	    "fmt"
	    "strings"
	    "net/http"
	    "regexp"
//	    "sort"

                      )

///
///     æE­åEãããæ°å­ãæ½åºãã
///

func String_no_get ( w http.ResponseWriter, r *http.Request , string_data string  )  ([]string ) {

//     IN    w      ãã: ã¬ã¹ãã³ã¹ã©ã¤ã¿ã¼
//     IN    r      ãã: ãªã¯ã¨ã¹ããã©ã¡ã¼ã¿
//     IN  string_data  : æE­åE

//     OUT strings  : åE²å¾ãEæE­åEç¾¤

//    fmt.Fprintf( w, "string_no_get start \n" )  // ãEãE¯

    string_data2 := strings.Split ( string_data, " "  )

    if len(string_data2) == 1 {

	   string_data2 = strings.Split ( string_data, ","  )

	}


//	str_test := "1-3"
//    rep := regexp.MustCompile(`[0-9]`)
//    str := rep.ReplaceAllString(str_test, "C")

//    fmt.Fprintf( w, "string_no_get str %c\n" ,str )  // ãEãE¯


//    string_work := string_data2[0]

//    for _, string_workw := range string_work {

//      fmt.Fprintf( w, "string_no_get string_workw %c\n" ,string_workw )  // ãEãE¯

//	}

    for _, string_data2w := range string_data2 {

      judge := String_check ( w , r  , `[0-9]`, string_data2w )

      if judge == false {

//	     fmt.Fprintf( w, "string_no_get string isn't numeric \n"  )  // ãEãE¯

	     return string_data2

	  }
    }

//    fmt.Fprintf( w, "string_no_get normal end \n" )  // ãEãE¯

    return	string_data2
}

///
///     æ­£è¦è¡¨ç¾ãããæå­åEãå¤å®ããE///

func String_check ( w http.ResponseWriter, r *http.Request , reg_exp string, string_data string  )  ( bool ) {

//     IN    w      ãã: ã¬ã¹ãã³ã¹ã©ã¤ã¿ã¼
//     IN    r      ãã: ãªã¯ã¨ã¹ããã©ã¡ã¼ã¿
//     IN  reg_exp      : æ­£è¦è¡¨ç¾
//     IN  string_data  : æE­åE

//     OUT judge        : true
//                      : false

//    fmt.Fprintf( w, "string_check start \n" )  // ãEãE¯

    judge := regexp.MustCompile( reg_exp ).Match( []byte(string_data) )

//    fmt.Fprintf( w, "string_check normal end \n" )  // ãEãE¯

    return	judge
}
