package csv_column_join

import (

//	    "fmt"
	    "net/http"

	    "github.com/sawaq7/go12_ver1/general/process3"
	    "github.com/sawaq7/go12_ver1/general/datastore5/trans3"
        "html/template"
        "github.com/sawaq7/go12_ver1/general/html5"
        "github.com/sawaq7/go12_ver1/general/strings2"

	    "strconv"

                                                  )

func Csv_column_join(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "sky/csv_column_join start \n" )  // チE��チE��

    var err error

///
///        入力データをGET 　
///

    string_data := r.FormValue("join_column")  // 追加する列をゲチE��
    strings := strings2.String_no_get( w , r , string_data  )

    join_no := make( []int, len(strings) )

    for pos, stringsw := range strings {

      join_no[pos] ,err = strconv.Atoi(stringsw)  // 整数匁E	  if err != nil {
	   http.Error(w,err.Error(), http.StatusInternalServerError)

		return
	  }

	}

	filename := r.FormValue("join_file")  // 追加するファイル名をゲチE��

///
/// 列単位�EチE�Eタを加える 　
///

    for _ , join_now := range join_no {

      process3.Csv_column_join ( w , r ,filename ,join_now )

	}

///
/// 　　　web上に、csv惁E��を表示する　
///

    csv_inf := trans3.Csv_inf ( w ,r )  //     csv惁E��をゲチE��

//    fmt.Fprintf( w, "sky/csv_column_join : csv_in %v\n", csv_inf )  // チE��チE��

    monitor := template.Must( template.New("html").Parse( html5.Csv_show )) // チE��プレート�EヘッダーをGET

    err = monitor.Execute ( w, csv_inf )   // web上に、csv惁E��を表示する　
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
