package process3

import (

//        "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
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
///     1列�EチE�Eタを追加する�E�Eolumn/move)
///

func Csv_column_join2(w http.ResponseWriter, r *http.Request , column_no1 int ,column_no2 int ) {

//     IN    w      　　: レスポンスライター
//     IN    r      　　: リクエストパラメータ
//     IN  column_no1　 : 削除予定�E行NO
//     IN  column_no2　 : 追加する行NO

//    fmt.Fprintf( w, "process3.csv_column_join2 start \n" )  // チE��チE��

///
///   プロジェクト名をゲチE��
///

    project_name := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if project_name == "" {
//      fmt.Fprintf( w, "storage_bucket_list :  projectID unset \n"  )  // チE��チE��

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
/// 　　　csv惁E��のフォーマットを拡張する　
///

	csv_inf := trans3.Csv_inf ( w ,r )  ///      csv惁E��をゲチE��

    csv_inf2 := reformat.Csv_inf ( 1 ,int64(column_no2) ,csv_inf ,w ,r )
                                                           /// フォーマットを1列拡張する
///
///      csv惁E��を追加する
///

    csv_inf_join := trans3.Csv_inf_column ( w ,r ,column_no1 )         /// 追加するcsv惁E��をゲチE��

    csv_inf_new := set1.Csv_inf (  csv_inf2 ,csv_inf_join ,column_no2 , w ,r )
                                                        /// 追加するチE�Eタ1列をセチE��

///
/// 　　　チE�Eタストアに、csv惁E��を�EセチE��する　
///

    for _, csv_inf_neww := range csv_inf_new {

//   	  fmt.Fprintf( w, "process3.csv_column_join2 csv_inf_neww %v\n", csv_inf_neww )  // チE��チE��

//      key := datastore.NewKey(c, "Csv_Inf", "", csv_inf_neww.Id, nil)  //　アクセスキーゲチE��
      key := datastore.IDKey("Csv_Inf", csv_inf_neww.Id, nil)

      if _, err := client.Put(ctx, key, &csv_inf_neww ); err != nil {
//      if _, err := datastore.Put(c, key, &csv_inf_neww); err != nil {  // チE�Eタストアに再セチE��
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	  }

    }

///
/// 　　　web上に、csv惁E��を表示する　
///

//    monitor := template.Must( template.New("html").Parse( html5.Csv_show )) // チE��プレート�EヘッダーをGET

//     err = monitor.Execute ( w, csv_inf_new )
//	 if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//	 }
//	fmt.Fprintf( w, "process3.csv_column_join2 normal end \n" )  // チE��チE��


}

