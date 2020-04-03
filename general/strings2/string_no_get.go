package strings2

import (

//	    "fmt"
	    "strings"
	    "net/http"
	    "regexp"
//	    "sort"

                      )

///
///     譁・ｭ怜・繧医ｊ縲∵焚蟄励ｒ謚ｽ蜃ｺ縺吶ｋ
///

func String_no_get ( w http.ResponseWriter, r *http.Request , string_data string  )  ([]string ) {

//     IN    w      縲縲: 繝ｬ繧ｹ繝昴Φ繧ｹ繝ｩ繧､繧ｿ繝ｼ
//     IN    r      縲縲: 繝ｪ繧ｯ繧ｨ繧ｹ繝医ヱ繝ｩ繝｡繝ｼ繧ｿ
//     IN  string_data  : 譁・ｭ怜・

//     OUT strings  : 蛻・牡蠕後・譁・ｭ怜・鄒､

//    fmt.Fprintf( w, "string_no_get start \n" )  // 繝・ヰ繝・け

    string_data2 := strings.Split ( string_data, " "  )

    if len(string_data2) == 1 {

	   string_data2 = strings.Split ( string_data, ","  )

	}


//	str_test := "1-3"
//    rep := regexp.MustCompile(`[0-9]`)
//    str := rep.ReplaceAllString(str_test, "C")

//    fmt.Fprintf( w, "string_no_get str %c\n" ,str )  // 繝・ヰ繝・け


//    string_work := string_data2[0]

//    for _, string_workw := range string_work {

//      fmt.Fprintf( w, "string_no_get string_workw %c\n" ,string_workw )  // 繝・ヰ繝・け

//	}

    for _, string_data2w := range string_data2 {

      judge := String_check ( w , r  , `[0-9]`, string_data2w )

      if judge == false {

//	     fmt.Fprintf( w, "string_no_get string isn't numeric \n"  )  // 繝・ヰ繝・け

	     return string_data2

	  }
    }

//    fmt.Fprintf( w, "string_no_get normal end \n" )  // 繝・ヰ繝・け

    return	string_data2
}

///
///     豁｣隕剰｡ｨ迴ｾ繧医ｊ縲∵枚蟄怜・繧貞愛螳壹☆繧・///

func String_check ( w http.ResponseWriter, r *http.Request , reg_exp string, string_data string  )  ( bool ) {

//     IN    w      縲縲: 繝ｬ繧ｹ繝昴Φ繧ｹ繝ｩ繧､繧ｿ繝ｼ
//     IN    r      縲縲: 繝ｪ繧ｯ繧ｨ繧ｹ繝医ヱ繝ｩ繝｡繝ｼ繧ｿ
//     IN  reg_exp      : 豁｣隕剰｡ｨ迴ｾ
//     IN  string_data  : 譁・ｭ怜・

//     OUT judge        : true
//                      : false

//    fmt.Fprintf( w, "string_check start \n" )  // 繝・ヰ繝・け

    judge := regexp.MustCompile( reg_exp ).Match( []byte(string_data) )

//    fmt.Fprintf( w, "string_check normal end \n" )  // 繝・ヰ繝・け

    return	judge
}
