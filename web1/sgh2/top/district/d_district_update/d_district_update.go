package d_district_update

import (

	    "strconv"
//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "net/http"
//	    "fmt"
	    "client/sgh/process"

        "client/sgh/type2"

        "os"
        "cloud.google.com/go/datastore"
        "context"
                                                   )

func D_district_update(w http.ResponseWriter, r *http.Request) {

	var g type2.D_District

//    fmt.Fprintf( w, "d_district_update start \n" )  // デバック

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

//	   fmt.Fprintf( w, "d_district_update :error updidw %v\n", updidw )  // デバック

	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
	}
    updid := int64(updidw)

//    fmt.Fprintf( w, "d_district_update : updidw %v\n", updidw )  // デバック
//    fmt.Fprintf( w, "d_district_update : updid %v\n", updid )  // デバック

//	key := datastore.NewKey(c, "D_District", "", updid, nil)
    key := datastore.IDKey("D_District", updid, nil)

    if err := client.Get(ctx, key , &g ) ; err != nil {
//	if err := datastore.Get(c, key, &g); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

    g.District_Name = r.FormValue("district_name")  // 地区名をゲット

	district_no := r.FormValue("district_no")         // 地区No.をゲット
//	fmt.Fprintf( w, "d_district_update : district_no %v\n", district_no )  // デバック
//	fmt.Fprintf( w, "d_district_update : district_name %v\n", g.District_Name )  // デバック

	district_now ,err := strconv.Atoi(district_no)  // 文字の整数化
	if err != nil {
//		http.Error(w,err.Error(), http.StatusInternalServerError)
//       fmt.Fprintf( w, "d_district_update : a number must be half-width characters %v\n"  )
		return
	}

	g.District_No = int64(district_now)   // 整数の64ビット化

//	fmt.Fprintf( w, "d_district_update : g.District_Name %v\n", g.District_Name )  // デバック
//	fmt.Fprintf( w, "d_district_update : g.District_No %v\n", g.District_No )  // デバック

    if _, err := client.Put(ctx, key, &g ); err != nil {
//	if _, err := datastore.Put(c, key, &g); err != nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}

/// モニター　再表示 ///

	process.D_district_showall1(w , r )

//	http.Redirect(w, r, "/", http.StatusFound)


}
