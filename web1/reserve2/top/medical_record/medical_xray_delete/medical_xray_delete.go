package medical_xray_delete

import (

	"net/http"
	"strconv"

	"github.com/sawaq7/go12_ver1/storage2"
//	"fmt"

	"github.com/sawaq7/go12_ver1/client/reserve/process4"
    "github.com/sawaq7/go12_ver1/client/reserve/datastore6/check5"
    "github.com/sawaq7/go12_ver1/client/reserve/type6"

    "cloud.google.com/go/datastore"
	"context"
	"os"

                                            )

func Medical_xray_delete(w http.ResponseWriter, r *http.Request) {

//     IN    w      : response-writer
//     IN    r      : request- paramete

//    fmt.Fprintf( w, "medical_xray_delete start \n" )

    var guest_medical_xray type6.Guest_Medical_Xray

    projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if projectID == "" {

      projectID = "sample-7777"

	}

    ctx := context.Background()

    client, err := datastore.NewClient(ctx, projectID)
    if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

    id := r.FormValue("id")
//    fmt.Fprintf( w, "medical_xray_delete : id %v\n", id )

	delidw ,_ := strconv.Atoi(id)
	delid := int64(delidw)

//    fmt.Fprintf( w, "medical_xray_delete : delidw %v\n", delidw )
//    fmt.Fprintf( w, "medical_xray_delete : delid %v\n", delid )

///
///     delete x-ray-file
///

    key := datastore.IDKey("Guest_Medical_Xray", delid, nil)           ///    get id for delete

    if err := client.Get(ctx, key, &guest_medical_xray ); err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

    bucket := "sample-7777"

	storage2.File_Delete ( w , r  ,bucket ,guest_medical_xray.File_Name  )

///
///     delete xray file in d.s.
///

    if err := client.Delete(ctx, key ); err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

    general_work := check5.Guest_temp (w , r )

//    value2, _ := flexible_out.([]type5.General_Work)

///
/// show medical recode inf. on web
///

    process4.Medical_xray_show(w , r ,general_work[0].Int64_Work)

}
