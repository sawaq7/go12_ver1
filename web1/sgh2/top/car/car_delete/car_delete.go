package car_delete

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
///     delete  car inf.
///

func Car_delete(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "car_delete start \n" )
    var car type2.Car

    id := r.FormValue("id")
//    fmt.Fprintf( w, "car_delete : id %v\n", id )

	delidw ,_ := strconv.Atoi(id)
	delid := int64(delidw)

//    fmt.Fprintf( w, "car_delete : delidw %v\n", delidw )
//    fmt.Fprintf( w, "car_delete : delid %v\n", delid )

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

    key := datastore.IDKey("Car", delid, nil)

    if err := client.Get(ctx, key, &car ); err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	district_no := car.District_No

    if err := client.Delete(ctx, key ); err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

///
///      show car inf. on web
///

	process.Car_show( w , r ,district_no )

}
