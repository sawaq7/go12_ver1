package check3

import (

	    "net/http"
	    "fmt"
	    "errors"

                     )

///                          ///
/// チE�Eタストアをコピ�Eする  ///
///                         ///

func Name( w http.ResponseWriter, basic_name string  ) (err error ){

//     IN    w        : レスポンスライター
//     IN  basic_Name : 基本のチE�Eタストア吁E

    fmt.Fprintf( w, "check3.name start \n" )  // チE��チE��
    fmt.Fprintf( w, "check3.name basic_name %v\n" ,basic_name)  // チE��チE��

///
///  エラーメチE��ージ　セチE��
///

var (

	  Err1 = errors.New("can't find this datastore's file (check3.name)")

	                                                                        )
///
/// チE�Eタストア名　チェチE��
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

    fmt.Fprintf( w, "check3.name ok_flag %v\n" ,ok_flag)  // チE��チE��
	fmt.Fprintf( w, "check3.name normal end \n" )  // チE��チE��

    return nil
}

