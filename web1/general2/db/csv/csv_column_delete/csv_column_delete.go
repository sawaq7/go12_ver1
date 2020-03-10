package csv_column_delete

import (

//	    "fmt"
	    "net/http"

	    "general/process3"
	    "strconv"
	    "html/template"
	    "general/datastore5/trans3"
	    "general/html5"
	    "general/strings2"

                                    )

func Csv_column_delete(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "sky/csv_column_delete start \n" )  // デバック

    var err error

///
/// 入力データをGET 　
///

    string_data := r.FormValue("delete_column")  // 削除する列をゲット

    strings := strings2.String_no_get( w , r , string_data  )

    delete_no := make( []int, len(strings) )

    for pos, stringsw := range strings {

      delete_no[pos] ,err = strconv.Atoi(stringsw)  // 整数化
	  if err != nil {
	   http.Error(w,err.Error(), http.StatusInternalServerError)

		return
	  }

	}

///
///         列を削除する　
///

	for _ , delete_now := range delete_no {

      process3.Csv_column_delete ( w , r ,delete_now  )

	}

///
/// 　　　web上に、csv情報を表示する　
///

    csv_inf := trans3.Csv_inf ( w ,r )  //     csv情報をゲット

    monitor := template.Must( template.New("html").Parse( html5.Csv_show )) // テンプレートのヘッダーをGET

     err = monitor.Execute ( w, csv_inf )
	 if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	 }

}
