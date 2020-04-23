package deliver_delete

import (

	"net/http"
	"strconv"
//	"fmt"
	"github.com/sawaq7/go12_ver1/client/sgh/process"

	"cloud.google.com/go/datastore"
    "context"
    "os"
                                            )

///
/// 　　   delete delivery inf.
///

func Deliver_delete(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "deliver_delete start \n" )

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

    id := r.FormValue("id")
//    fmt.Fprintf( w, "deliver_delete : id %v\n", id )

	delidw ,_ := strconv.Atoi(id)
	delid := int64(delidw)

//    fmt.Fprintf( w, "deliver_delete : delidw %v\n", delidw )
//    fmt.Fprintf( w, "deliver_delete : delid %v\n", delid )

    key := datastore.IDKey("Deliver", delid, nil)

    if err := client.Delete(ctx, key ); err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

    ///     show deliver inf. on web
	process.Deliver_showall1(w , r )

}
