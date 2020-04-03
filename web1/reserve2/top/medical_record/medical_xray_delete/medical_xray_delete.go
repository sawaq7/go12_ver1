package medical_xray_delete

import (

//	"google.golang.org/appengine"
//	"google.golang.org/appengine/datastore"
	"net/http"
	"strconv"

	"storage2"
//	"fmt"

	"github.com/sawaq7/go12_ver1/client/reserve/process4"
    "github.com/sawaq7/go12_ver1/client/reserve/datastore6/check5"
    "github.com/sawaq7/go12_ver1/client/reserve/type6"

    "cloud.google.com/go/datastore"
	"context"
	"os"

                                            )

func Medical_xray_delete(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "medical_xray_delete start \n" )  // 繝・ヰ繝・け

    var guest_medical_xray type6.Guest_Medical_Xray

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

    id := r.FormValue("id")
//    fmt.Fprintf( w, "medical_xray_delete : id %v\n", id )  // 繝・ヰ繝・け

	delidw ,_ := strconv.Atoi(id)
	delid := int64(delidw)

//    fmt.Fprintf( w, "medical_xray_delete : delidw %v\n", delidw )  // 繝・ヰ繝・け
//    fmt.Fprintf( w, "medical_xray_delete : delid %v\n", delid )  // 繝・ヰ繝・け

///
///     xray 繝輔ぃ繧､繝ｫ繧貞炎髯､
///

    key := datastore.IDKey("Guest_Medical_Xray", delid, nil)           ///    xray諠・ｱ繧偵ご繝・ヨ

    if err := client.Get(ctx, key, &guest_medical_xray ); err != nil {
//    key := datastore.NewKey(c, "Guest_Medical_Xray", "", delid, nil)
//    if err := datastore.Get(c, key, &guest_medical_xray); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

///    xray繝輔ぃ繧､繝ｫ繧貞炎髯､

    bucket := "sample-7777"

	storage2.File_Delete ( w , r  ,bucket ,guest_medical_xray.File_Name  )

///
///     xray 諠・ｱ繧貞炎髯､
///

    if err := client.Delete(ctx, key ); err != nil {
//	if err := datastore.Delete(c, key); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

/// 繝｢繝九ち繝ｼ縲蜀崎｡ｨ遉ｺ ///

    general_work := check5.Guest_temp (w , r )

//    value2, _ := flexible_out.([]type5.General_Work)

    process4.Medical_xray_show(w , r ,general_work[0].Int64_Work)

}
