package process3

import (

	    "net/http"
//	    "fmt"
	    "html/template"
	    "general/html5"
	    "storage2"
	    "general/type5"
//	    "time"
                                                )

///
/// 指示したバケット内のオブジェクトを、ウエブ上に表示する。
///

func Storage_object_show ( w http.ResponseWriter, r *http.Request ,project string ,bucket string ) {

//     IN    w      : レスポンスライター
//     IN    r      : リクエストパラメータ
//     IN  project  : プロジェクト名
//     IN  bucket   : バケット名

//    fmt.Fprintf( w, "process3.storage_object_show start \n" )  // デバック

//    var t_dmy   time.Time

//    var idmy int64

    storage_b_o_view := make([]type5.Storage_B_O_View, 0) // バケットリストの表示用エリアを確保

//    objects_minor , _ := storage2.Storage_basic( "list3" ,bucket ,idmy, w , r  )

//    objects, _ := objects_minor.([]type5.Storage_B_O_View)  // インターフェイス型を型変換

    objects :=  storage2.Object_List_Detail ( w  ,r , bucket )  // 指示したバケット内のオブジェクト情報をゲットする

    for pos , objectsw := range objects {

//      fmt.Fprintf( w, "storage_bucket_list : pos: %v\n", pos )

      storage_b_o_view = append(storage_b_o_view, type5.Storage_B_O_View { int64(pos+1)            ,
                                                                              project              ,
                                                                              bucket               ,
                                                                              objectsw.Object_Name ,
                                                                              objectsw.Created       })

    }

    monitor := template.Must(template.New("html").Parse(html5.Storage_object_list))

	err := monitor.Execute( w, storage_b_o_view )
	if err != nil {
	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
	}

//	fmt.Fprintf( w, "process3.storage_object_show normal end \n" )  // デバック
}

