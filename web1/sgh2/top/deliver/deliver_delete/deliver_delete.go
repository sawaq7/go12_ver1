package deliver_delete

import (

//	"google.golang.org/appengine"
//	"google.golang.org/appengine/datastore"
	"net/http"
	"strconv"
//	"fmt"
	"client/sgh/process"

	"cloud.google.com/go/datastore"
    "context"
    "os"
                                            )

func Deliver_delete(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "deliver_delete start \n" )  // デバック

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

    id := r.FormValue("id")
//    fmt.Fprintf( w, "deliver_delete : id %v\n", id )  // デバック

	delidw ,_ := strconv.Atoi(id)
	delid := int64(delidw)

//    fmt.Fprintf( w, "deliver_delete : delidw %v\n", delidw )  // デバック
//    fmt.Fprintf( w, "deliver_delete : delid %v\n", delid )  // デバック

    key := datastore.IDKey("Deliver", delid, nil)

    if err := client.Delete(ctx, key ); err != nil {
//	key := datastore.NewKey(c, "Deliver", "", delid, nil)
//	if err := datastore.Delete(c, key); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

/// モニター　再表示 ///

	process.Deliver_showall1(w , r )

}
