package storage_object_copy_keyin

import (

	    "net/http"
	    "general/process3"
                                                  )

///
///    main　データーベースのアクセスリストを表示する
///



func Storage_object_copy_keyin(w http.ResponseWriter, r *http.Request) {

///
/// モニター表示
///

//    process3.Db_access_list(w , r )
    process3.Storage_object_copy_keyin(w , r )


}

