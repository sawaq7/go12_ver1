package pipe_line_ds_delete

import (

	"net/http"
	"strconv"
//	"fmt"
	"github.com/sawaq7/go12_ver1/client/tokura/suiri/process2"

	"cloud.google.com/go/datastore"
    "context"
    "os"
                                            )

func Pipe_line_ds_delete(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "pipe_line_ds_delete start \n" )

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
//    fmt.Fprintf( w, "pipe_line_ds_delete : id %v\n", id )

	delidw ,_ := strconv.Atoi(id)
	delid := int64(delidw)

//    fmt.Fprintf( w, "pipe_line_ds_delete : delidw %v\n", delidw )
//    fmt.Fprintf( w, "pipe_line_ds_delete : delid %v\n", delid )

    if err := client.Delete(ctx, datastore.IDKey("Water2", delid, nil)); err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

///  show water-line inf.

	process2.Pipe_line_ds_show(w , r )

}
