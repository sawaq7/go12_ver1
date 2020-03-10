package car_delete

import (

//	     "google.golang.org/appengine"
//	     "google.golang.org/appengine/datastore"
	     "net/http"
	     "strconv"
//	"fmt"
	     "client/sgh/process"
	     "client/sgh/type2"

	     "cloud.google.com/go/datastore"
	     "context"
	     "os"
                                            )

func Car_delete(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "car_delete start \n" )  // デバック
    var car type2.Car

    id := r.FormValue("id")
//    fmt.Fprintf( w, "car_delete : id %v\n", id )  // デバック

	delidw ,_ := strconv.Atoi(id)
	delid := int64(delidw)

//    fmt.Fprintf( w, "car_delete : delidw %v\n", delidw )  // デバック
//    fmt.Fprintf( w, "car_delete : delid %v\n", delid )  // デバック

    projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if projectID == "" {

      projectID = "sample-7777"

	}

//	c := appengine.NewContext(r)
    ctx := context.Background()

    client, err := datastore.NewClient(ctx, projectID)
    if err != nil {
       http.Error(w,err.Error(), http.StatusInternalServerError)
		return
    }

    key := datastore.IDKey("Car", delid, nil)

    if err := client.Get(ctx, key, &car ); err != nil {
//	key := datastore.NewKey(c, "Car", "", delid, nil)
//    if err := datastore.Get(c, key, &car); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	district_no := car.District_No

    if err := client.Delete(ctx, key ); err != nil {
//	if err := datastore.Delete(c, key); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

/// 号車情報を表示 ///

	process.Car_show( w , r ,district_no )

}
