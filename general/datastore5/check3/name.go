package check3

import (

	    "net/http"
	    "fmt"
	    "errors"

                     )

///                          ///
/// 繝・・繧ｿ繧ｹ繝医い繧偵さ繝斐・縺吶ｋ  ///
///                         ///

func Name( w http.ResponseWriter, basic_name string  ) (err error ){

//     IN    w        : 繝ｬ繧ｹ繝昴Φ繧ｹ繝ｩ繧､繧ｿ繝ｼ
//     IN  basic_Name : 蝓ｺ譛ｬ縺ｮ繝・・繧ｿ繧ｹ繝医い蜷・

    fmt.Fprintf( w, "check3.name start \n" )  // 繝・ヰ繝・け
    fmt.Fprintf( w, "check3.name basic_name %v\n" ,basic_name)  // 繝・ヰ繝・け

///
///  繧ｨ繝ｩ繝ｼ繝｡繝・そ繝ｼ繧ｸ縲繧ｻ繝・ヨ
///

var (

	  Err1 = errors.New("can't find this datastore's file (check3.name)")

	                                                                        )
///
/// 繝・・繧ｿ繧ｹ繝医い蜷阪繝√ぉ繝・け
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

    fmt.Fprintf( w, "check3.name ok_flag %v\n" ,ok_flag)  // 繝・ヰ繝・け
	fmt.Fprintf( w, "check3.name normal end \n" )  // 繝・ヰ繝・け

    return nil
}

