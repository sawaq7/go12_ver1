package check3

import (

	    "net/http"
	    "fmt"
	    "errors"

                     )

///                          ///
/// データストアをコピーする  ///
///                         ///

func Name( w http.ResponseWriter, basic_name string  ) (err error ){

//     IN    w        : レスポンスライター
//     IN  basic_Name : 基本のデータストア名


    fmt.Fprintf( w, "check3.name start \n" )  // デバック
    fmt.Fprintf( w, "check3.name basic_name %v\n" ,basic_name)  // デバック

///
///  エラーメッセージ　セット
///

var (

	  Err1 = errors.New("can't find this datastore's file (check3.name)")

	                                                                        )
///
/// データストア名　チェック
///
    ok_flag := 0

    switch basic_name {

      case "Deliver" :

        ok_flag = 1


      break;

    }

    if ok_flag == 0 {

		return Err1
	}

    fmt.Fprintf( w, "check3.name ok_flag %v\n" ,ok_flag)  // デバック
	fmt.Fprintf( w, "check3.name normal end \n" )  // デバック

    return nil
}

