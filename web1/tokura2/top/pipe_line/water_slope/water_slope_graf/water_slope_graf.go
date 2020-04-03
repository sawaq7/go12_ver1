package water_slope_graf

import (

	    "strconv"
//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "net/http"
//	    "fmt"

	    "storage2"

	    "github.com/sawaq7/go12_ver1/client/tokura/suiri/type4"
        "github.com/sawaq7/go12_ver1/general/type5"

        "cloud.google.com/go/datastore"
        "context"
        "os"

                                                   )

func Water_slope_graf(w http.ResponseWriter, r *http.Request) {

//	fmt.Fprintf( w, "sky_water_slope_graf start %v\n" )  // 繝・ヰ繝・け

	var water_slope  type4.Water_Slope

	var image_show type5.Image_Show

	var idmy int64

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

/// 謖・ｮ壹＠縺溘ョ繝ｼ繧ｿid繧竪ET ///

    select_idw , err := strconv.Atoi(r.FormValue("id"))
	if err  != nil {

//	   fmt.Fprintf( w, "sky_water_slope_graf :error select_idw %v\n", select_idw )  // 繝・ヰ繝・け

	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
	}
    select_id := int64(select_idw)

    key := datastore.IDKey("Water_Slope", select_id, nil)

    if err := client.Get(ctx, key , &water_slope ) ; err != nil {

//	key := datastore.NewKey(c, "Water_Slope", "", select_id, nil)
//	if err := datastore.Get(c, key, &water_slope); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

///
/// 繝｢繝九ち繝ｼ縺ｫ繧ｰ繝ｩ繝戊｡ｨ遉ｺ
///
    image_show.File_Name = water_slope.File_Name
    image_show.Url       = water_slope.Url

    storage2.Storage_basic( "show2" ,image_show ,idmy , w , r  )

//    get.Image_file_show2( w  ,r  ,image_show  )


}
