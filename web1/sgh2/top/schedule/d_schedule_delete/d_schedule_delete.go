package d_schedule_delete

import (

//	"google.golang.org/appengine"
//	"google.golang.org/appengine/datastore"
	"net/http"
	"strconv"
//	"fmt"
	"github.com/sawaq7/go12_ver1/client/sgh/process"
	"github.com/sawaq7/go12_ver1/client/sgh/datastore2"
	"github.com/sawaq7/go12_ver1/general/type5"

	"cloud.google.com/go/datastore"
	"context"
	"os"
                                            )

func D_schedule_delete(w http.ResponseWriter, r *http.Request) {

    var idmy int64

//    fmt.Fprintf( w, "d_schedule_delete start \n" )  // チE��チE��

    id := r.FormValue("id")
//    fmt.Fprintf( w, "d_schedule_delete : id %v\n", id )  // チE��チE��

	delidw ,_ := strconv.Atoi(id)
	delid := int64(delidw)

//    fmt.Fprintf( w, "d_schedule_delete : delidw %v\n", delidw )  // チE��チE��

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

    key := datastore.IDKey("D_Schedule", delid, nil)

    if err := client.Delete(ctx, key ); err != nil {
//	key := datastore.NewKey(c, "D_Schedule"       , "", delid, nil)
//	if err := datastore.Delete(c, key); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

//地区惁E��をGET

    flexible_out := datastore2.Datastore_sgh( "D_District_Temp" ,"check" ,idmy , w , r  )

    value2, _ := flexible_out.([]type5.General_Work)

    district_no := value2[0].Int64_Work

/// モニター　再表示 ///

	process.D_schedule_showall(w , r ,district_no)

}