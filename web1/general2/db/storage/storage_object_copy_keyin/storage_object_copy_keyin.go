package storage_object_copy_keyin

import (

	    "net/http"
	    "github.com/sawaq7/go12_ver1/general/process3"
                                                  )

///
///    main　チE�Eターベ�Eスのアクセスリストを表示する
///



func Storage_object_copy_keyin(w http.ResponseWriter, r *http.Request) {

///
/// モニター表示
///

//    process3.Db_access_list(w , r )
    process3.Storage_object_copy_keyin(w , r )


}

