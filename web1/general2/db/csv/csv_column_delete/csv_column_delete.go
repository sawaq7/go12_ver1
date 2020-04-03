package csv_column_delete

import (

//	    "fmt"
	    "net/http"

	    "github.com/sawaq7/go12_ver1/general/process3"
	    "strconv"
	    "html/template"
	    "github.com/sawaq7/go12_ver1/general/datastore5/trans3"
	    "github.com/sawaq7/go12_ver1/general/html5"
	    "github.com/sawaq7/go12_ver1/general/strings2"

                                    )

func Csv_column_delete(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "sky/csv_column_delete start \n" )  // チE��チE��

    var err error

///
/// 入力データをGET 　
///

    string_data := r.FormValue("delete_column")  // 削除する列をゲチE��

    strings := strings2.String_no_get( w , r , string_data  )

    delete_no := make( []int, len(strings) )

    for pos, stringsw := range strings {

      delete_no[pos] ,err = strconv.Atoi(stringsw)  // 整数匁E	  if err != nil {
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
/// 　　　web上に、csv惁E��を表示する　
///

    csv_inf := trans3.Csv_inf ( w ,r )  //     csv惁E��をゲチE��

    monitor := template.Must( template.New("html").Parse( html5.Csv_show )) // チE��プレート�EヘッダーをGET

     err = monitor.Execute ( w, csv_inf )
	 if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	 }

}
