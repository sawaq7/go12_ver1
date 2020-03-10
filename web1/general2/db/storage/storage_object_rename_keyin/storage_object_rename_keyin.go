package storage_object_rename_keyin

import (

	    "net/http"
	    "general/process3"
                                                  )

///
///    main　データーベースのアクセスリストを表示する
///

func Storage_object_rename_keyin(w http.ResponseWriter, r *http.Request) {

///
/// モニター表示
///

    process3.Storage_object_rename_keyin(w , r )

}

