package deliver_delete

import (

//	"google.golang.org/appengine"
//	"google.golang.org/appengine/datastore"
	"net/http"
	"strconv"
//	"fmt"
	"github.com/sawaq7/go12_ver1/client/sgh/process"

	"cloud.google.com/go/datastore"
    "context"
    "os"
                                            )

func Deliver_delete(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "deliver_delete start \n" )  // γEγE―

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
//    fmt.Fprintf( w, "deliver_delete : id %v\n", id )  // γEγE―

	delidw ,_ := strconv.Atoi(id)
	delid := int64(delidw)

//    fmt.Fprintf( w, "deliver_delete : delidw %v\n", delidw )  // γEγE―
//    fmt.Fprintf( w, "deliver_delete : delid %v\n", delid )  // γEγE―

    key := datastore.IDKey("Deliver", delid, nil)

    if err := client.Delete(ctx, key ); err != nil {
//	key := datastore.NewKey(c, "Deliver", "", delid, nil)
//	if err := datastore.Delete(c, key); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

/// γ’γγΏγΌγεθ‘¨η€Ί ///

	process.Deliver_showall1(w , r )

}
