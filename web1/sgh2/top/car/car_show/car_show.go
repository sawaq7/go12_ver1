package car_show

import (
//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "net/http"
//	    "fmt"

        "github.com/sawaq7/go12_ver1/client/sgh/process"
        "github.com/sawaq7/go12_ver1/client/sgh/type2"
	    "strconv"
	    "github.com/sawaq7/go12_ver1/client/sgh/datastore2"

        "cloud.google.com/go/datastore"
	    "context"
	    "os"
                                                  )

/// main 蜿ｷ霆頑ュ蝣ｱ繧定｡ｨ遉ｺ縺吶ｋ ///

func Car_show(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintln( w, "car_show start \n" )  // 繝・ヰ繝・け

    var d_district      type2.D_District

    var d_district_temp type2.D_District_Temp

    var idmy int64

    select_idw , err := strconv.Atoi(r.FormValue("id"))

	if err  != nil {
//	   fmt.Fprintf( w, "car_show :error select_idw %v\n", select_idw )  // 繝・ヰ繝・け
	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
	}
    select_id := int64(select_idw)

//    fmt.Fprintf( w, "car_show : select_idw %v\n", select_idw )  // 繝・ヰ繝・け
//    fmt.Fprintf( w, "car_show : select_id %v\n", select_id )  // 繝・ヰ繝・け

    project_name := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if project_name == "" {

      project_name = "sample-7777"

	}

//	c := appengine.NewContext(r)
    ctx := context.Background()

    client, err := datastore.NewClient(ctx, project_name)
    if err != nil {
       http.Error(w,err.Error(), http.StatusInternalServerError)
		return
    }

    key := datastore.IDKey("D_District", select_id, nil)

    if err := client.Get(ctx, key , &d_district ) ; err != nil {
//	key := datastore.NewKey(c, "D_District", "", select_id, nil)
//	if err := datastore.Get(c, key, &d_district); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}


// temporary-file繧偵う繝九す繝｣繝ｩ繧､繧ｺ  & 繧ｻ繝・ヨ (驕ｸ謚槭＠縺溷慍蝓滓ュ蝣ｱ繧偵そ繝・ヨ

//    initialize.D_district_temp (w , r ) // temporary-file繧偵う繝九す繝｣繝ｩ繧､繧ｺ

//    _ = datastore2.D_store( "D_District_Temp" ,"initialize" ,idmy , w , r  )
    _ = datastore2.Datastore_sgh( "D_District_Temp" ,"initialize" ,idmy , w , r  )

    d_district_temp.District_No   = d_district.District_No
    d_district_temp.District_Name = d_district.District_Name

    new_key := datastore.IncompleteKey("D_District_Temp", nil)

    if _, err = client.Put(ctx, new_key, &d_district_temp ); err != nil {
//	if _, err := datastore.Put(c, datastore.NewIncompleteKey(c, "D_District_Temp", nil), &d_district_temp); err != nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}

/// 蜿ｷ霆頑ュ蝣ｱ繧定｡ｨ遉ｺ ///

	process.Car_show(w , r ,d_district.District_No)

//	fmt.Fprintf( w, "d_district_area : normal end \n" )  // 繝・ヰ繝・け




}
