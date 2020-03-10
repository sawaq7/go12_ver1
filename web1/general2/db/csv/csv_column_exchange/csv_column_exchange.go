package csv_column_exchange

import (
//	     "fmt"
	     "net/http"

	     "general/process3"
	     "general/html5"
	     "strconv"
	     "general/datastore5/trans3"
	     "html/template"
                                      )

func Csv_column_exchange(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "sky/csv_column_exchange start \n" )  // デバック

///
/// 入力データをGET 　
///

    exchange_column1_minor := r.FormValue("column1")  // 削除する列をゲット

	exchange_column1 ,err := strconv.Atoi(exchange_column1_minor)  // 整数化
	if err != nil {
	   http.Error(w,err.Error(), http.StatusInternalServerError)

		return
	}

    exchange_column2_minor := r.FormValue("column2")  // 削除する列をゲット

	exchange_column2 ,err := strconv.Atoi(exchange_column2_minor)  // 整数化
	if err != nil {
	   http.Error(w,err.Error(), http.StatusInternalServerError)

		return
	}

///
///    指定した1行を追加する
///

    if exchange_column1 < exchange_column2 {

	   exchange_column2 ++

	}
	   process3.Csv_column_join2 ( w , r ,exchange_column1 ,exchange_column2 )

///
///
///    指定した1行を削除する
///

    if exchange_column1 >= exchange_column2 {

	   exchange_column1 ++

	}

	process3.Csv_column_delete ( w , r ,exchange_column1  )

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
