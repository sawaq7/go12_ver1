package copy3

import (
	    "google.golang.org/appengine"
	    "google.golang.org/appengine/datastore"
	    "net/http"
	    "fmt"

//	    "github.com/sawaq7/go12_ver1/general/datastore5/check3"
        "github.com/sawaq7/go12_ver1/client/sgh/type2"

                                                )

///                          ///
/// チE�Eタストアをコピ�Eする  ///
///                         ///

// func All( w http.ResponseWriter, r *http.Request ,basic_name string ,copy_file string ,new_file string ) {
 func All( w http.ResponseWriter, r *http.Request ,basic_name string ,copy_file string ,new_file string )( err error) {

//     IN    w        : レスポンスライター
//     IN    r        : リクエストパラメータ
//     IN  basic_name : 基本のチE�Eタストア吁E//     IN  copy_file  : コピ�E允E�EチE�Eタストア吁E//     IN  new_file   : ニューチE�Eタストア吁E//    OUT  err        : エラーメチE��ージ

    fmt.Fprintf( w, "copy3.all start \n" )  // チE��チE��
    fmt.Fprintf( w, "copy3.all basic_name %v\n" ,basic_name)  // チE��チE��

///
///  エラーメチE��ージ　セチE��
///

    c := appengine.NewContext(r)       ///　コンチE��ストを取得すめE
	q := datastore.NewQuery(copy_file) /// クローンのリーダーをGET

	count, err := q.Count(c)           ///　レコード数をGET
	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
		return  err

	}
    fmt.Fprintf( w, "copy3.all count %v\n" ,count)  // チE��チE��


//   if err :=check3.Name( w , basic_name   ) ;  err != nil{
//       http.Error(w, err.Error(), http.StatusInternalServerError)

//	  return err

//	}

///
/// クローン惁E��をSETするワークエリアを確俁E///

    ds_data := make([]type2.Deliver, 0, count)



	if _, err := q.GetAll(c, &ds_data);  err != nil {         // クローン惁E��をGET

//	  http.Error(w, err.Error(), http.StatusInternalServerError)
	  return  err

	} else{
      for _, ds_dataw := range ds_data {                       //　ニューファイルにクローン惁E��をセチE��

	    if _, err := datastore.Put(c, datastore.NewIncompleteKey(c, new_file, nil), &ds_dataw); err != nil {

//		  http.Error(w,err.Error(), http.StatusInternalServerError)
		  return  err

	    }

	  }
	}

	fmt.Fprintf( w, "copy3.all normal end \n" )  // チE��チE��

    return nil
}

