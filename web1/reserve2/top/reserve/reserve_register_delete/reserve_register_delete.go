package reserve_register_delete

import (

	"net/http"
	"strconv"
//	"fmt"

	"github.com/sawaq7/go12_ver1/client/reserve/process4"
    "github.com/sawaq7/go12_ver1/client/reserve/datastore6/check5"

    "cloud.google.com/go/datastore"
	"context"
	"os"

                                            )

///
///     delete reservation inf.
///

func Reserve_register_delete(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "reserve_register_delete start \n" )

    id := r.FormValue("id")
//    fmt.Fprintf( w, "reserve_register_delete : id %v\n", id )

	delidw ,_ := strconv.Atoi(id)
	delid := int64(delidw)

//    fmt.Fprintf( w, "reserve_register_delete : delidw %v\n", delidw )
//    fmt.Fprintf( w, "reserve_register_delete : delid %v\n", delid )

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

    key := datastore.IDKey("Guest_Reserve_Minor", delid, nil)

    if err := client.Delete(ctx, key ); err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

///
///     register reservation inf.
///

    general_work := check5.Guest_temp (w , r )

//    value2, _ := flexible_out.([]type5.General_Work)

	process4.Reserve_register(w , r ,general_work[0].Int64_Work)

}
