package csv_column_exchange

import (
//	     "fmt"
	     "net/http"

	     "github.com/sawaq7/go12_ver1/general/process3"
	     "github.com/sawaq7/go12_ver1/general/html5"
	     "strconv"
	     "github.com/sawaq7/go12_ver1/general/datastore5/trans3"
	     "html/template"
                                      )

func Csv_column_exchange(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "sky/csv_column_exchange start \n" )  // ใEใEฏ

///
/// ๅฅๅใใผใฟใGET ใ
///

    exchange_column1_minor := r.FormValue("column1")  // ๅ้คใใๅใใฒใE

	exchange_column1 ,err := strconv.Atoi(exchange_column1_minor)  // ๆดๆฐๅE	if err != nil {
	   http.Error(w,err.Error(), http.StatusInternalServerError)

		return
	}

    exchange_column2_minor := r.FormValue("column2")  // ๅ้คใใๅใใฒใE

	exchange_column2 ,err := strconv.Atoi(exchange_column2_minor)  // ๆดๆฐๅE	if err != nil {
	   http.Error(w,err.Error(), http.StatusInternalServerError)

		return
	}

///
///    ๆEฎใใE่กใ่ฟฝๅ ใใ
///

    if exchange_column1 < exchange_column2 {

	   exchange_column2 ++

	}
	   process3.Csv_column_join2 ( w , r ,exchange_column1 ,exchange_column2 )

///
///
///    ๆEฎใใE่กใๅ้คใใ
///

    if exchange_column1 >= exchange_column2 {

	   exchange_column1 ++

	}

	process3.Csv_column_delete ( w , r ,exchange_column1  )

///
/// ใใใwebไธใซใcsvๆE ฑใ่กจ็คบใใใ
///

    csv_inf := trans3.Csv_inf ( w ,r )  //     csvๆE ฑใใฒใE

    monitor := template.Must( template.New("html").Parse( html5.Csv_show )) // ใEณใใฌใผใใEใใใใผใGET

     err = monitor.Execute ( w, csv_inf )
	 if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	 }

}
