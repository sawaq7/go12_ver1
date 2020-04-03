package water_slope_delete

import (

//	"google.golang.org/appengine"
//	"google.golang.org/appengine/datastore"
	"net/http"
	"strconv"
//	"fmt"
	"github.com/sawaq7/go12_ver1/client/tokura/suiri/process2"
	"github.com/sawaq7/go12_ver1/client/tokura/suiri/type4"
	"storage2"

	"cloud.google.com/go/datastore"
    "context"
    "os"
                                            )

///
///   蟆取ｰｴ蜍ｾ驟咲ｷ壹ｒ蜑企勁縺吶ｋ
///

func Water_slope_delete(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "water_slope_delete start \n" )  // 繝・ヰ繝・け

    var g type4.Water_Slope

    project_name := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if project_name == "" {

      project_name = "sample-7777"

	}

    ctx := context.Background()
//	c := appengine.NewContext(r)

    client, err := datastore.NewClient(ctx, project_name)
    if err != nil {
       http.Error(w, err.Error(), http.StatusInternalServerError)
       return
    }

    id := r.FormValue("id")
//    fmt.Fprintf( w, "water_slope_delete : id %v\n", id )  // 繝・ヰ繝・け

	delidw ,_ := strconv.Atoi(id)
	delid := int64(delidw)

//    fmt.Fprintf( w, "water_slope_delete : delidw %v\n", delidw )  // 繝・ヰ繝・け
//    fmt.Fprintf( w, "water_slope_delete : delid %v\n", delid )  // 繝・ヰ繝・け



///
///   繝舌こ繝・ヨ蜷阪・繝輔ぃ繧､繝ｫ蜷阪ｒ繧ｻ繝・ヨ
///

    key := datastore.IDKey("Water_Slope", delid, nil)

    if err := client.Get(ctx, key , &g ) ; err != nil {
//    if err := datastore.Get(c, key, &g); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

    file_name   := g.File_Name
	bucket_name := "sample-7777"
///
///   蟆取ｰｴ蜍ｾ驟咲ｷ壽ュ蝣ｱ繧貞炎髯､縺吶ｋ
///
    if err := client.Delete(ctx, key ); err != nil {

//    key := datastore.NewKey(c, "Water_Slope", "", delid, nil)
//	if err := datastore.Delete(c, key); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	storage2.File_Delete ( w , r  ,bucket_name ,file_name  )

// 繝｢繝九ち繝ｼ縲陦ｨ遉ｺ ///

    process2.Water_slope_show(w , r )

//	http.Redirect(w, r, "/", http.StatusFound)
}
