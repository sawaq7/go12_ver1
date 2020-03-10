package water_slope_delete

import (

//	"google.golang.org/appengine"
//	"google.golang.org/appengine/datastore"
	"net/http"
	"strconv"
//	"fmt"
	"client/tokura/suiri/process2"
	"client/tokura/suiri/type4"
	"storage2"

	"cloud.google.com/go/datastore"
    "context"
    "os"
                                            )

///
///   導水勾配線を削除する
///

func Water_slope_delete(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "water_slope_delete start \n" )  // デバック

    var g type4.Water_Slope

    project_name := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if project_name == "" {

      project_name = "sample-7777"

	}

    ctx := context.Background()
//	c := appengine.NewContext(r)

    client, err := datastore.NewClient(ctx, project_name)
    if err != nil {
       http.Error(w, err.Error(), http.StatusInternalServerError)
       return
    }

    id := r.FormValue("id")
//    fmt.Fprintf( w, "water_slope_delete : id %v\n", id )  // デバック

	delidw ,_ := strconv.Atoi(id)
	delid := int64(delidw)

//    fmt.Fprintf( w, "water_slope_delete : delidw %v\n", delidw )  // デバック
//    fmt.Fprintf( w, "water_slope_delete : delid %v\n", delid )  // デバック



///
///   バケット名・ファイル名をセット
///

    key := datastore.IDKey("Water_Slope", delid, nil)

    if err := client.Get(ctx, key , &g ) ; err != nil {
//    if err := datastore.Get(c, key, &g); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

    file_name   := g.File_Name
	bucket_name := "sample-7777"
///
///   導水勾配線情報を削除する
///
    if err := client.Delete(ctx, key ); err != nil {

//    key := datastore.NewKey(c, "Water_Slope", "", delid, nil)
//	if err := datastore.Delete(c, key); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	storage2.File_Delete ( w , r  ,bucket_name ,file_name  )

// モニター　表示 ///

    process2.Water_slope_show(w , r )

//	http.Redirect(w, r, "/", http.StatusFound)
}
