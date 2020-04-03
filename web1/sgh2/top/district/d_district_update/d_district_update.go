package d_district_update

import (

	    "strconv"
//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "net/http"
//	    "fmt"
	    "github.com/sawaq7/go12_ver1/client/sgh/process"

        "github.com/sawaq7/go12_ver1/client/sgh/type2"

        "os"
        "cloud.google.com/go/datastore"
        "context"
                                                   )

func D_district_update(w http.ResponseWriter, r *http.Request) {

	var g type2.D_District

//    fmt.Fprintf( w, "d_district_update start \n" )  // ใEใEฏ

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

	updidw , err := strconv.Atoi(r.FormValue("id"))
	if err  != nil {

//	   fmt.Fprintf( w, "d_district_update :error updidw %v\n", updidw )  // ใEใEฏ

	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
	}
    updid := int64(updidw)

//    fmt.Fprintf( w, "d_district_update : updidw %v\n", updidw )  // ใEใEฏ
//    fmt.Fprintf( w, "d_district_update : updid %v\n", updid )  // ใEใEฏ

//	key := datastore.NewKey(c, "D_District", "", updid, nil)
    key := datastore.IDKey("D_District", updid, nil)

    if err := client.Get(ctx, key , &g ) ; err != nil {
//	if err := datastore.Get(c, key, &g); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

    g.District_Name = r.FormValue("district_name")  // ๅฐๅบๅใใฒใE

	district_no := r.FormValue("district_no")         // ๅฐๅบNo.ใใฒใE
//	fmt.Fprintf( w, "d_district_update : district_no %v\n", district_no )  // ใEใEฏ
//	fmt.Fprintf( w, "d_district_update : district_name %v\n", g.District_Name )  // ใEใEฏ

	district_now ,err := strconv.Atoi(district_no)  // ๆEญใEๆดๆฐๅE	if err != nil {
//		http.Error(w,err.Error(), http.StatusInternalServerError)
//       fmt.Fprintf( w, "d_district_update : a number must be half-width characters %v\n"  )
		return
	}

	g.District_No = int64(district_now)   // ๆดๆฐใฎ64ใใใๅ

//	fmt.Fprintf( w, "d_district_update : g.District_Name %v\n", g.District_Name )  // ใEใEฏ
//	fmt.Fprintf( w, "d_district_update : g.District_No %v\n", g.District_No )  // ใEใEฏ

    if _, err := client.Put(ctx, key, &g ); err != nil {
//	if _, err := datastore.Put(c, key, &g); err != nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}

/// ใขใใฟใผใๅ่กจ็คบ ///

	process.D_district_showall1(w , r )

//	http.Redirect(w, r, "/", http.StatusFound)


}
