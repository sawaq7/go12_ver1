package db_access_list

import (

	    "net/http"
	    "github.com/sawaq7/go12_ver1/general/process3"
                                                  )

///
///    main　チE�Eターベ�Eスのアクセスリストを表示する
///

func Db_access_list(w http.ResponseWriter, r *http.Request) {

///
/// モニター表示
///

    process3.Db_access_list(w , r )

}
