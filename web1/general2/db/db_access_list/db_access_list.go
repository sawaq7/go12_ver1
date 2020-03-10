package db_access_list

import (

	    "net/http"
	    "general/process3"
                                                  )

///
///    main　データーベースのアクセスリストを表示する
///

func Db_access_list(w http.ResponseWriter, r *http.Request) {

///
/// モニター表示
///

    process3.Db_access_list(w , r )

}