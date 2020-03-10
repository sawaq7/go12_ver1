package d_district_area

import (

//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "net/http"
//	    "fmt"
//	    "client/sgh"
        "client/sgh/process"
        "client/sgh/type2"
	    "strconv"
	    "client/sgh/datastore2"
//	    "time"

        "cloud.google.com/go/datastore"
        "context"
        "os"
                                                  )

func D_district_area(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "d_district_area start \n" )  // デバック

    var g type2.D_District

    var g2 type2.D_District_Temp

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

    updidw , err := strconv.Atoi(r.FormValue("id"))

	if err  != nil {
//	   fmt.Fprintf( w, "d_district_area :error updidw %v\n", updidw )  // デバック
	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
	}
    updid := int64(updidw)

//    fmt.Fprintf( w, "d_district_area : updidw %v\n", updidw )  // デバック
//    fmt.Fprintf( w, "d_district_area : updid %v\n", updid )  // デバック

    key := datastore.IDKey("D_District", updid, nil)

    if err := client.Get(ctx, key , &g ) ; err != nil {
//	key := datastore.NewKey(c, "D_District", "", updid, nil)
//	if err := datastore.Get(c, key, &g); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}


// temporary-fileをイニシャライズ  & セット//

//    _ = datastore2.D_store( "D_District_Temp" ,"initialize" ,idmy , w , r  )
    _ = datastore2.Datastore_sgh( "D_District_Temp" ,"initialize" ,idmy , w , r  )

    g2.District_No   = g.District_No
    g2.District_Name = g.District_Name

    new_key := datastore.IncompleteKey("D_District_Temp", nil)

    if _, err = client.Put(ctx, new_key, &g2 ); err != nil {
//	if _, err := datastore.Put(c, datastore.NewIncompleteKey(c, "D_District_Temp", nil), &g2); err != nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}

/// モニター　表示 ///

	process.D_district_area(w , r ,g.District_No)

//	fmt.Fprintf( w, "d_district_area : normal end \n" )  // デバック




}
