package pipe_line_ds_wl_delete

import (

//	"google.golang.org/appengine"
//	"google.golang.org/appengine/datastore"
	"net/http"
	"strconv"
//	"fmt"
	"github.com/sawaq7/go12_ver1/client/tokura/suiri/process2"
	"github.com/sawaq7/go12_ver1/client/tokura/suiri/type4"

	"cloud.google.com/go/datastore"
    "context"
    "os"

                                            )

func Pipe_line_ds_wl_delete(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "pipe_line_ds_wl_delete start \n" )  // 繝・ヰ繝・け

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
//    fmt.Fprintf( w, "pipe_line_ds_wl_delete : id %v\n", id )  // 繝・ヰ繝・け

	delidw ,_ := strconv.Atoi(id)
	delid := int64(delidw)

//    fmt.Fprintf( w, "pipe_line_ds_wl_delete : delidw %v\n", delidw )  // 繝・ヰ繝・け
//    fmt.Fprintf( w, "pipe_line_ds_wl_delete : delid %v\n", delid )  // 繝・ヰ繝・け

    key := datastore.IDKey("Water_Line", delid, nil)

    if err := client.Get(ctx, key , &g ) ; err != nil {

//	key := datastore.NewKey(c, "Water_Line", "", delid, nil)
//	if err := datastore.Get(c, key, &g); err != nil {              //縲縲豌ｴ霍ｯ蜷阪ｒGET
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

    if err := client.Delete(ctx, key ); err != nil {
//	if err := datastore.Delete(c, key); err != nil {               //   豌ｴ霍ｯ繝ｩ繧､繝ｳ繧貞炎髯､
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

/// 繝｢繝九ち繝ｼ縲蜀崎｡ｨ遉ｺ ///

	process2.Pipe_line_ds_wl_show(1 ,g.Name ,w , r )


//	http.Redirect(w, r, "/", http.StatusFound)
}
