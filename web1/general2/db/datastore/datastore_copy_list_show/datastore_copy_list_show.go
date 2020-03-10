package datastore_copy_list_show

import (

//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "net/http"
//	    "fmt"
	    "general/process3"

	    "general/type5"
//	    "strconv"
//	    "time"

        "cloud.google.com/go/datastore"
        "context"
        "os"
                                                  )

func Datastore_copy_list_show(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "datastore_copy_list_show start \n" )  // デバック

    var g type5.Ds_Copy_List

    project_name := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if project_name == "" {
//      fmt.Fprintf( w, "storage_bucket_list :  projectID unset \n"  )  // デバック

      project_name = "sample-7777"

	}
    ctx := context.Background()
//	c := appengine.NewContext(r)

    client, err := datastore.NewClient(ctx, project_name)
    if err != nil {
       http.Error(w, err.Error(), http.StatusInternalServerError)
       return
    }

	g.Basic_Name = r.FormValue("basic_name")  // 基本のデータストア名をゲット
	g.Copy_Name  = r.FormValue("copy_file")   // コピー元のデータストア名をゲット
	g.New_Name   = r.FormValue("new_file")    // ニューデータストア名をゲット

/// データストアーにデータをセット ///
    new_key := datastore.IncompleteKey("Ds_Copy_List", nil)

//    fmt.Fprintf(w, "datastore_copy_list_show: new_key %v\n", new_key )  // デバック

    if   _, err = client.Put(ctx, new_key, &g ); err != nil {
//	if _, err := datastore.Put(c, datastore.NewIncompleteKey(c, "Ds_Copy_List", nil), &g); err != nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}

//   	fmt.Fprintf( w, "datastore_copy_list_show : g.Basic_Name %v\n", g.Basic_Name )  // デバック
//    fmt.Fprintf( w, "datastore_copy_list_show : g.Copy_Name %v\n", g.Copy_Name )  // デバック
//    fmt.Fprintf( w, "datastore_copy_list_show : g.New_Name %v\n", g.New_Name )  // デバック

/// モニター　再表示 ///

    process3.Datastore_copy_list_keyin(w , r )

//	fmt.Fprintf( w, "datastore_copy_list_show : normal end \n" )  // デバック




}
