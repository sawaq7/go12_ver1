package process3

import (
//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "net/http"
//	    "fmt"
//	    "html/template"
	    "general/datastore5/reformat"

//	    "general/html5"
	    "general/datastore5/trans3"

	    "cloud.google.com/go/datastore"
        "context"
        "os"
                                                )


func Csv_column_delete(w http.ResponseWriter, r *http.Request ,column_no int ) {

//     IN    w      　　: レスポンスライター
//     IN    r      　　: リクエストパラメータ
//     IN  column_no　  : 削除する行NO

//    fmt.Fprintf( w, "csv_column_delete start \n" )  // デバック

///
///   プロジェクト名をゲット
///

    project_name := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if project_name == "" {
//      fmt.Fprintf( w, "storage_bucket_list :  projectID unset \n"  )  // デバック

      project_name = "sample-7777"

	}

//    c := appengine.NewContext(r)
    ctx := context.Background()

    client, err := datastore.NewClient(ctx, project_name)
    if err != nil {
       http.Error(w,err.Error(), http.StatusInternalServerError)
		return
    }

///
/// 　　　csv情報を修正する　
///

	csv_inf := trans3.Csv_inf ( w ,r )  ///      csv情報をゲット

    csv_inf_new := reformat.Csv_inf ( 0, int64(column_no) ,csv_inf ,w ,r )      /// csv情報を修正する

///
/// 　　　データストアに、csv情報を再セットする　
///
    for _, csv_inf_neww := range csv_inf_new {

      key := datastore.IDKey("Csv_Inf", csv_inf_neww.Id, nil)
//      key := datastore.NewKey(c, "Csv_Inf", "", csv_inf_neww.Id, nil)  //　アクセスキーゲット

      if _, err := client.Put(ctx, key, &csv_inf_neww ); err != nil {
//      if _, err := datastore.Put(c, key, &csv_inf_neww); err != nil {  // データストアに再セット
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	  }

    }


//	fmt.Fprintf( w, "csv_column_delete normal end \n" )  // デバック


}

