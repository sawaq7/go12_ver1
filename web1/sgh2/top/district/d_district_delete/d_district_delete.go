package d_district_delete

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
/// 　　   delete district inf. in d.s.
///

func D_district_delete(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "d_district_delete start \n" )

	project_name := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if project_name == "" {

      project_name = "sample-7777"

	}

    ctx := context.Background()

    client, err := datastore.NewClient(ctx, project_name)
    if err != nil {
       http.Error(w,err.Error(), http.StatusInternalServerError)
		return
    }

    id := r.FormValue("id")
//    fmt.Fprintf( w, "d_district_delete : id %v\n", id )

	delidw ,_ := strconv.Atoi(id)
	delid := int64(delidw)

//    fmt.Fprintf( w, "d_district_delete : delidw %v\n", delidw )
//    fmt.Fprintf( w, "d_district_delete : delid %v\n", delid )

    if err := client.Delete(ctx, datastore.IDKey("D_District", delid, nil)); err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

    //    show district inf. on web
	process.D_district_showall1(w , r )

}
