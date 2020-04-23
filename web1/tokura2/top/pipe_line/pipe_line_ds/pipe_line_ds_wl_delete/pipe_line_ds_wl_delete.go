package pipe_line_ds_wl_delete

import (

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

//    fmt.Fprintf( w, "pipe_line_ds_wl_delete start \n" )

    var g  type4.Water_Line

    project_name := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if project_name == "" {

      project_name = "sample-7777"

	}

    ctx := context.Background()

    client, err := datastore.NewClient(ctx, project_name)
    if err != nil {
       http.Error(w, err.Error(), http.StatusInternalServerError)
       return
    }

    id := r.FormValue("id")
//    fmt.Fprintf( w, "pipe_line_ds_wl_delete : id %v\n", id )

	delidw ,_ := strconv.Atoi(id)
	delid := int64(delidw)

//    fmt.Fprintf( w, "pipe_line_ds_wl_delete : delidw %v\n", delidw )
//    fmt.Fprintf( w, "pipe_line_ds_wl_delete : delid %v\n", delid )

    key := datastore.IDKey("Water_Line", delid, nil)

    if err := client.Get(ctx, key , &g ) ; err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

    if err := client.Delete(ctx, key ); err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

///     show water-line inf.

	process2.Pipe_line_ds_wl_show(1 ,g.Name ,w , r )

}
