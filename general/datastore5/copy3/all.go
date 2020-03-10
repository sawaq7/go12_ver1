package copy3

import (
	    "google.golang.org/appengine"
	    "google.golang.org/appengine/datastore"
	    "net/http"
	    "fmt"

//	    "general/datastore5/check3"
        "client/sgh/type2"

                                                )

///                          ///
/// データストアをコピーする  ///
///                         ///

// func All( w http.ResponseWriter, r *http.Request ,basic_name string ,copy_file string ,new_file string ) {
 func All( w http.ResponseWriter, r *http.Request ,basic_name string ,copy_file string ,new_file string )( err error) {

//     IN    w        : レスポンスライター
//     IN    r        : リクエストパラメータ
//     IN  basic_name : 基本のデータストア名
//     IN  copy_file  : コピー元のデータストア名
//     IN  new_file   : ニューデータストア名
//    OUT  err        : エラーメッセージ

    fmt.Fprintf( w, "copy3.all start \n" )  // デバック
    fmt.Fprintf( w, "copy3.all basic_name %v\n" ,basic_name)  // デバック

///
///  エラーメッセージ　セット
///

    c := appengine.NewContext(r)       ///　コンテキストを取得する

	q := datastore.NewQuery(copy_file) /// クローンのリーダーをGET

	count, err := q.Count(c)           ///　レコード数をGET
	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
		return  err

	}
    fmt.Fprintf( w, "copy3.all count %v\n" ,count)  // デバック


//   if err :=check3.Name( w , basic_name   ) ;  err != nil{
//       http.Error(w, err.Error(), http.StatusInternalServerError)

//	  return err

//	}

///
/// クローン情報をSETするワークエリアを確保
///

    ds_data := make([]type2.Deliver, 0, count)



	if _, err := q.GetAll(c, &ds_data);  err != nil {         // クローン情報をGET

//	  http.Error(w, err.Error(), http.StatusInternalServerError)
	  return  err

	} else{
      for _, ds_dataw := range ds_data {                       //　ニューファイルにクローン情報をセット

	    if _, err := datastore.Put(c, datastore.NewIncompleteKey(c, new_file, nil), &ds_dataw); err != nil {

//		  http.Error(w,err.Error(), http.StatusInternalServerError)
		  return  err

	    }

	  }
	}

	fmt.Fprintf( w, "copy3.all normal end \n" )  // デバック

    return nil
}

