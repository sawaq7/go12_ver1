package medical_record_delete

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
///     delete medical record inf. in d.s.
///

func Medical_record_delete(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "medical_record_delete start \n" )

    id := r.FormValue("id")
//    fmt.Fprintf( w, "medical_record_delete : id %v\n", id )

	delidw ,_ := strconv.Atoi(id)
	delid := int64(delidw)

//    fmt.Fprintf( w, "medical_record_delete : delidw %v\n", delidw )
//    fmt.Fprintf( w, "medical_record_delete : delid %v\n", delid )

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

    key := datastore.IDKey("Guest_Medical_Record", delid, nil)

    if err := client.Delete(ctx, key ); err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

///
///     show guest tmp. inf. on web
///

    general_work := check5.Guest_temp (w , r )

//    value2, _ := flexible_out.([]type5.General_Work)

	process4.Medical_record_show(w , r ,general_work[0].Int64_Work)

}
