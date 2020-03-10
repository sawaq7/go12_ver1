package pipe_line_ds_wl_delete

import (

//	"google.golang.org/appengine"
//	"google.golang.org/appengine/datastore"
	"net/http"
	"strconv"
//	"fmt"
	"client/tokura/suiri/process2"
	"client/tokura/suiri/type4"

	"cloud.google.com/go/datastore"
    "context"
    "os"

                                            )

func Pipe_line_ds_wl_delete(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "pipe_line_ds_wl_delete start \n" )  // デバック

    var g  type4.Water_Line

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
//    fmt.Fprintf( w, "pipe_line_ds_wl_delete : id %v\n", id )  // デバック

	delidw ,_ := strconv.Atoi(id)
	delid := int64(delidw)

//    fmt.Fprintf( w, "pipe_line_ds_wl_delete : delidw %v\n", delidw )  // デバック
//    fmt.Fprintf( w, "pipe_line_ds_wl_delete : delid %v\n", delid )  // デバック

    key := datastore.IDKey("Water_Line", delid, nil)

    if err := client.Get(ctx, key , &g ) ; err != nil {

//	key := datastore.NewKey(c, "Water_Line", "", delid, nil)
//	if err := datastore.Get(c, key, &g); err != nil {              //　　水路名をGET
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

    if err := client.Delete(ctx, key ); err != nil {
//	if err := datastore.Delete(c, key); err != nil {               //   水路ラインを削除
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

/// モニター　再表示 ///

	process2.Pipe_line_ds_wl_show(1 ,g.Name ,w , r )


//	http.Redirect(w, r, "/", http.StatusFound)
}
