package d_district_area_update

import (

	    "strconv"
//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "net/http"
//	    "fmt"
	    "github.com/sawaq7/go12_ver1/client/sgh/process"

        "github.com/sawaq7/go12_ver1/client/sgh/type2"

        "cloud.google.com/go/datastore"
	    "context"
	    "os"
                                                   )

func D_district_area_update(w http.ResponseWriter, r *http.Request) {

	var g type2.D_Area

//    fmt.Fprintf( w, "d_district_area_update start \n" )  // ăEăEŻ

	 projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if projectID == "" {

      projectID = "sample-7777"

	}

//	c := appengine.NewContext(r)
    ctx := context.Background()

    client, err := datastore.NewClient(ctx, projectID)
    if err  != nil {

	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
	}

	updidw , err := strconv.Atoi(r.FormValue("id"))
	if err  != nil {

//	   fmt.Fprintf( w, "d_district_area_update :error updidw %v\n", updidw )  // ăEăEŻ

	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
	}
    updid := int64(updidw)

//    fmt.Fprintf( w, "d_district_area_update : updidw %v\n", updidw )  // ăEăEŻ
//    fmt.Fprintf( w, "d_district_area_update : updid %v\n", updid )  // ăEăEŻ

    key := datastore.IDKey("D_Area", updid, nil)

    if err := client.Get(ctx, key , &g ) ; err != nil {

//	key := datastore.NewKey(c, "D_Area", "", updid, nil)
//	if err := datastore.Get(c, key, &g); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

    g.Area_Name = r.FormValue("area_name")          // ă¨ăŞă˘ĺăă˛ăE
    g.Area_Detail = r.FormValue("area_detail")      // ă¨ăŞă˘čŠłç´°ăă˛ăE

//	fmt.Fprintf( w, "d_district_area_update : g.Area_Name %v\n", g.Area_Name )  // ăEăEŻ
//	fmt.Fprintf( w, "d_district_area_update : g.Area_Detail %v\n", g.Area_Detail )  // ăEăEŻ

// ăEEăżăšăă˘ăŽ1ăŹăłăźăă˘ăEEăEEăE
    if _, err = client.Put(ctx, key, &g ); err != nil {
//	if _, err := datastore.Put(c, key, &g); err != nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}

/// ă˘ăăżăźăĺčĄ¨ç¤ş ///

	process.D_district_area_show(w , r  ,g.District_No)

}
