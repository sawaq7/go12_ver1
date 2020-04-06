package copy3

import (
//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "net/http"
//	    "fmt"

        "github.com/sawaq7/go12_ver1/client/sgh/type2"

        "os"
        "cloud.google.com/go/datastore"
        "context"

                                                )

///
/// チE�Eタストアをコピ�Eする  (　ベ�EシチE��名　�E�　Deliver　�E�E///

func Deliver( w http.ResponseWriter, r *http.Request ,basic_name string ,copy_file string ,new_file string ) {

//     IN    w        : レスポンスライター
//     IN    r        : リクエストパラメータ
//     IN  basic_name : 基本のチE�Eタストア吁E//     IN  copy_file  : コピ�E允E�EチE�Eタストア吁E//     IN  new_file   : ニューチE�Eタストア吁E//    OUT  err        : エラーメチE��ージ

//    fmt.Fprintf( w, "copy3.deliver start \n" )  // チE��チE��
//    fmt.Fprintf( w, "copy3.deliver basic_name %v\n" ,basic_name)  // チE��チE��

///
///  プロジェクチED　ゲチE��
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

//    fmt.Fprintf( w, "copy3.deliver count %v\n" ,count)  // チE��チE��

///
/// クローン惁E��をSETするワークエリアを確俁E///

    ds_data := make([]type2.Deliver, 0, count)

    if _, err := client.GetAll(ctx, query , &ds_data);  err != nil {
//	if _, err := q.GetAll(c, &ds_data);  err != nil {         // クローン惁E��をGET

	  http.Error(w, err.Error(), http.StatusInternalServerError)
	  return

	} else{
      for _, ds_dataw := range ds_data {                       //　ニューファイルにクローン惁E��をセチE��

//	    if _, err := datastore.Put(c, datastore.NewIncompleteKey(c, new_file, nil), &ds_dataw); err != nil {
        new_key := datastore.IncompleteKey( new_file , nil)

        if _, err = client.Put(ctx, new_key, &ds_dataw ); err != nil {

		  http.Error(w,err.Error(), http.StatusInternalServerError)
		  return

	    }

	  }
	}

//	fmt.Fprintf( w, "copy3.deliver normal end \n" )  // チE��チE��

    return
}
