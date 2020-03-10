package copy3

import (
//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "net/http"
//	    "fmt"

        "client/sgh/type2"

        "os"
        "cloud.google.com/go/datastore"
        "context"

                                                )

///
/// データストアをコピーする  (　ベーシック名　：　Deliver　）
///

func Deliver( w http.ResponseWriter, r *http.Request ,basic_name string ,copy_file string ,new_file string ) {

//     IN    w        : レスポンスライター
//     IN    r        : リクエストパラメータ
//     IN  basic_name : 基本のデータストア名
//     IN  copy_file  : コピー元のデータストア名
//     IN  new_file   : ニューデータストア名
//    OUT  err        : エラーメッセージ

//    fmt.Fprintf( w, "copy3.deliver start \n" )  // デバック
//    fmt.Fprintf( w, "copy3.deliver basic_name %v\n" ,basic_name)  // デバック

///
///  プロジェクトID　ゲット
///
    projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if projectID == "" {

      projectID = "sample-7777"

	}

//	c := appengine.NewContext(r)
    ctx := context.Background()

    client, err := datastore.NewClient(ctx, projectID)

	query := datastore.NewQuery(copy_file)
//	q := datastore.NewQuery(copy_file)

//	count, err := q.Count(c)
    count, err := client.Count(ctx, query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

//    fmt.Fprintf( w, "copy3.deliver count %v\n" ,count)  // デバック

///
/// クローン情報をSETするワークエリアを確保
///

    ds_data := make([]type2.Deliver, 0, count)

    if _, err := client.GetAll(ctx, query , &ds_data);  err != nil {
//	if _, err := q.GetAll(c, &ds_data);  err != nil {         // クローン情報をGET

	  http.Error(w, err.Error(), http.StatusInternalServerError)
	  return

	} else{
      for _, ds_dataw := range ds_data {                       //　ニューファイルにクローン情報をセット

//	    if _, err := datastore.Put(c, datastore.NewIncompleteKey(c, new_file, nil), &ds_dataw); err != nil {
        new_key := datastore.IncompleteKey( new_file , nil)

        if _, err = client.Put(ctx, new_key, &ds_dataw ); err != nil {

		  http.Error(w,err.Error(), http.StatusInternalServerError)
		  return

	    }

	  }
	}

//	fmt.Fprintf( w, "copy3.deliver normal end \n" )  // デバック

    return
}

