package water_slope_delete

import (

	"net/http"
	"strconv"
//	"fmt"
	"github.com/sawaq7/go12_ver1/client/tokura/suiri/process2"
	"github.com/sawaq7/go12_ver1/client/tokura/suiri/type4"
	"github.com/sawaq7/go12_ver1/storage2"

	"cloud.google.com/go/datastore"
    "context"
    "os"
                                            )

///
///     delete water-slope-line
///

func Water_slope_delete(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "water_slope_delete start \n" )

    var g type4.Water_Slope

    project_name := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if project_name == "" {

      project_name = "sample-7777"

	}

    ctx := context.Background()

    client, err := datastore.NewClient(ctx, project_name)
    if err != nil {
       http.Error(w, err.Error(), http.StatusInternalServerError)
       return
    }

    id := r.FormValue("id")
//    fmt.Fprintf( w, "water_slope_delete : id %v\n", id )

	delidw ,_ := strconv.Atoi(id)
	delid := int64(delidw)

//    fmt.Fprintf( w, "water_slope_delete : delidw %v\n", delidw )
//    fmt.Fprintf( w, "water_slope_delete : delid %v\n", delid )



///
///   バケチE��名�Eファイル名をセチE��
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
///   delete water-slope-line
///

    if err := client.Delete(ctx, key ); err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	storage2.File_Delete ( w , r  ,bucket_name ,file_name  )

///
/// 　　　　　show water-slope inf. on web
///

    process2.Water_slope_show(w , r )

}
