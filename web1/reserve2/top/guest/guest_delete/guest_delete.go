package guest_delete

import (

//	"google.golang.org/appengine"
//	"google.golang.org/appengine/datastore"
	"net/http"
	"strconv"
//	"fmt"

	"github.com/sawaq7/go12_ver1/client/reserve/process4"

	"cloud.google.com/go/datastore"
	"context"
	"os"
                                            )

func Guest_delete(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "guest_delete start \n" )  // 繝・ヰ繝・け

    id := r.FormValue("id")
//    fmt.Fprintf( w, "guest_delete : id %v\n", id )  // 繝・ヰ繝・け

	delidw ,_ := strconv.Atoi(id)
	delid := int64(delidw)

//    fmt.Fprintf( w, "guest_delete : delidw %v\n", delidw )  // 繝・ヰ繝・け
//    fmt.Fprintf( w, "guest_delete : delid %v\n", delid )  // 繝・ヰ繝・け

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

    key := datastore.IDKey("Guest", delid, nil)

    if err := client.Delete(ctx, key ); err != nil {
//	key := datastore.NewKey(c, "Guest", "", delid, nil)
//	if err := datastore.Delete(c, key); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

///
/// 繝｢繝九ち繝ｼ縲蜀崎｡ｨ遉ｺ
///

    process4.Guest_show(w , r )

}
