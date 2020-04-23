package guest_delete

import (

	"net/http"
	"strconv"
//	"fmt"

	"github.com/sawaq7/go12_ver1/client/reserve/process4"

	"cloud.google.com/go/datastore"
	"context"
	"os"
                                            )

///
///     delete guest inf.
///

func Guest_delete(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "guest_delete start \n" )

    id := r.FormValue("id")
//    fmt.Fprintf( w, "guest_delete : id %v\n", id )

	delidw ,_ := strconv.Atoi(id)
	delid := int64(delidw)

//    fmt.Fprintf( w, "guest_delete : delidw %v\n", delidw )
//    fmt.Fprintf( w, "guest_delete : delid %v\n", delid )

    projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if projectID == "" {

      projectID = "sample-7777"

	}

    ctx := context.Background()

    client, err := datastore.NewClient(ctx, projectID)
    if err != nil {
       http.Error(w,err.Error(), http.StatusInternalServerError)
		return
    }

    key := datastore.IDKey("Guest", delid, nil)

    if err := client.Delete(ctx, key ); err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

///
///      show guest inf. on web
///

    process4.Guest_show(w , r )

}
