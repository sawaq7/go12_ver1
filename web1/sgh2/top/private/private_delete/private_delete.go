package private_delete

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

func Private_delete(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "private_delete start \n" )  // デバック


    id := r.FormValue("id")
//    fmt.Fprintf( w, "private_delete : id %v\n", id )  // デバック

	delidw ,_ := strconv.Atoi(id)
	delid := int64(delidw)

//    fmt.Fprintf( w, "private_delete : delidw %v\n", delidw )  // デバック
//    fmt.Fprintf( w, "private_delete : delid %v\n", delid )  // デバック

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

    key := datastore.IDKey("Private", delid, nil)

    if err := client.Delete(ctx, key ); err != nil {
//	key := datastore.NewKey(c, "Private", "", delid, nil)
//	if err := datastore.Delete(c, key); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

/// モニター　再表示 ///

	process.Private_showall1(w , r )


//	http.Redirect(w, r, "/", http.StatusFound)
}
