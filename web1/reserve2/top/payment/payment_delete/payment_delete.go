package payment_delete

import (

//	"google.golang.org/appengine"
//	"google.golang.org/appengine/datastore"
	"net/http"
	"strconv"
//	"fmt"

	"github.com/sawaq7/go12_ver1/client/reserve/process4"

    "github.com/sawaq7/go12_ver1/client/reserve/datastore6/check5"

    "cloud.google.com/go/datastore"
	"context"
	"os"

                                            )

func Payment_delete(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "payment_delete start \n" )  // チE��チE��

    id := r.FormValue("id")
//    fmt.Fprintf( w, "payment_delete : id %v\n", id )  // チE��チE��

	delidw ,_ := strconv.Atoi(id)
	delid := int64(delidw)

//    fmt.Fprintf( w, "payment_delete : delidw %v\n", delidw )  // チE��チE��
//    fmt.Fprintf( w, "payment_delete : delid %v\n", delid )  // チE��チE��

    projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if projectID == "" {

      projectID = "sample-7777"

	}

//	c := appengine.NewContext(r)
    ctx := context.Background()

    client, err := datastore.NewClient(ctx, projectID)
    if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

    key := datastore.IDKey("Guest_Payment", delid, nil)           ///    xray惁E��をゲチE��

    if err := client.Delete(ctx, key ); err != nil {
//	key := datastore.NewKey(c, "Guest_Payment", "", delid, nil)
//	if err := datastore.Delete(c, key); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

    general_work := check5.Guest_temp (w , r )

    guest_no   := general_work[0].Int64_Work

///
/// モニター　再表示
///

    process4.Payment_register( w , r ,guest_no )

}
