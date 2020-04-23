package d_district_area_delete

import (

	"net/http"
	"strconv"
//	"fmt"
	"github.com/sawaq7/go12_ver1/client/sgh/process"
	"github.com/sawaq7/go12_ver1/client/sgh/type2"

	"cloud.google.com/go/datastore"
    "context"
    "os"
                                            )

///
/// 　　   delete area inf.
///

func D_district_area_delete(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "d_district_area_delete start \n" )
    var g type2.D_Area

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

//    fmt.Fprintf( w, "d_district_area_delete : id %v\n", id )

	delidw ,_ := strconv.Atoi(id)
	delid := int64(delidw)

//    fmt.Fprintf( w, "d_district_area_delete : delidw %v\n", delidw )
//    fmt.Fprintf( w, "d_district_area_delete : delid %v\n", delid )

    key := datastore.IDKey("D_Area", delid, nil)

    if err := client.Get(ctx, key , &g ) ; err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	district_no := g.District_No

    //  delete one record in d.s.
    if err := client.Delete(ctx, key); err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

    //    show area inf. on web
	process.D_district_area_show(w , r  ,district_no)

}
