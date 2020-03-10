package storage_object_show

import (

	    "net/http"
//	    "fmt"
//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"

//	    "storage2/get"
	    "storage2"
	    "strconv"

	    "general/type5"
	    "os"
        "cloud.google.com/go/datastore"
        "context"

                                                  )

///
/// オブジェクトの内容を表示する。
///

func Storage_object_show(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "storage_object_show start \n" )  // デバック

    var bucket ,filename string

    projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if projectID == "" {

      projectID = "sample-7777"

	}

    line_no := r.FormValue("line_no")

//    fmt.Fprintf( w, "storage_object_show : line_no %v\n", line_no )  // デバック

	select_id ,_ := strconv.Atoi(line_no)

//    fmt.Fprintf( w, "storage_object_show : select_id %v\n", select_id )  // デバック


///
///   バケット名をゲット
///

//    c := appengine.NewContext(r)
    ctx := context.Background()

    client, err := datastore.NewClient(ctx, projectID)

	query := datastore.NewQuery("Storage_B_O_Temp")

//	count, err := q.Count(c)
	count, err := client.Count(ctx, query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	storage_b_o_temp      := make([]type5.Storage_B_O_Temp, 0, count)

    if _, err = client.GetAll(ctx, query , &storage_b_o_temp) ; err != nil {
//	if _, err := q.GetAll(c, &storage_b_o_temp);  err != nil {

	  http.Error(w, err.Error(), http.StatusInternalServerError)
	  return

	} else{

      for pos, storage_b_o_tempw := range storage_b_o_temp {

        if pos == 0 {
           bucket    = storage_b_o_tempw.Bucket_Name    // バケット名をゲット

        }
	  }
	}

///
///   ファイル名をゲット
///

	objects :=  storage2.Object_List ( w  ,r , bucket )

//    fmt.Fprintf( w, "storage_object_show : objects: %v\n", objects )

    for pos2 , objectsw := range objects {

      if pos2+1 == select_id {

        filename = objectsw

      }

    }


/// モニター　表示 ///

    storage2.Storage_basic( "show1" ,bucket ,filename , w , r  )

//    get.Image_file_show( w ,r ,bucket ,filename )

//	fmt.Fprintf( w, "storage_object_show : normal end \n" )  // デバック

}

