package pipe_line_ds_delete

import (

//	"google.golang.org/appengine"
//	"google.golang.org/appengine/datastore"
	"net/http"
	"strconv"
//	"fmt"
	"github.com/sawaq7/go12_ver1/client/tokura/suiri/process2"

	"cloud.google.com/go/datastore"
    "context"
    "os"
                                            )

func Pipe_line_ds_delete(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "pipe_line_ds_delete start \n" )  // γEγE―

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
//    fmt.Fprintf( w, "pipe_line_ds_delete : id %v\n", id )  // γEγE―

	delidw ,_ := strconv.Atoi(id)
	delid := int64(delidw)

//    fmt.Fprintf( w, "pipe_line_ds_delete : delidw %v\n", delidw )  // γEγE―
//    fmt.Fprintf( w, "pipe_line_ds_delete : delid %v\n", delid )  // γEγE―

    if err := client.Delete(ctx, datastore.IDKey("Water2", delid, nil)); err != nil {

//	key := datastore.NewKey(c, "Water2", "", delid, nil)
//	if err := datastore.Delete(c, key); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

/// γ’γγΏγΌγεθ‘¨η€Ί ///

	process2.Pipe_line_ds_show(w , r )

}
