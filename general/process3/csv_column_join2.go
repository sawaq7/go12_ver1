package process3

import (

	    "net/http"
//	    "fmt"

	    "github.com/sawaq7/go12_ver1/general/datastore5/reformat"

	    "github.com/sawaq7/go12_ver1/general/datastore5/trans3"
	    "github.com/sawaq7/go12_ver1/general/datastore5/set1"

	    "cloud.google.com/go/datastore"
        "context"
        "os"
                                                )

///
///     add one column
///

func Csv_column_join2(w http.ResponseWriter, r *http.Request , column_no1 int ,column_no2 int ) {

//     IN    w      　　: レスポンスライター
//     IN    r      　　: リクエストパラメータ
//     IN  column_no1　 : delete column no.
//     IN  column_no2　 : add column no.

//    fmt.Fprintf( w, "process3.csv_column_join2 start \n" )

///
///   get project name
///

    project_name := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if project_name == "" {
//      fmt.Fprintf( w, "storage_bucket_list :  projectID unset \n"  )

      project_name = "sample-7777"

	}

    ctx := context.Background()

    client, err := datastore.NewClient(ctx, project_name)
    if err != nil {
       http.Error(w,err.Error(), http.StatusInternalServerError)
		return
    }

///
/// 　　　expand format　
///

	csv_inf := trans3.Csv_inf ( w ,r )  ///      get csv inf.

    csv_inf2 := reformat.Csv_inf ( 1 ,int64(column_no2) ,csv_inf ,w ,r )
                                                           /// expand one low
///
///      add csv column
///

    csv_inf_join := trans3.Csv_inf_column ( w ,r ,column_no1 )         /// add csv data

    csv_inf_new := set1.Csv_inf (  csv_inf2 ,csv_inf_join ,column_no2 , w ,r )
                                                        /// set csv data

///
/// 　　　put csv inf.　
///

    for _, csv_inf_neww := range csv_inf_new {

//   	  fmt.Fprintf( w, "process3.csv_column_join2 csv_inf_neww %v\n", csv_inf_neww )

      key := datastore.IDKey("Csv_Inf", csv_inf_neww.Id, nil)

      if _, err := client.Put(ctx, key, &csv_inf_neww ); err != nil {

		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	  }

    }

}
