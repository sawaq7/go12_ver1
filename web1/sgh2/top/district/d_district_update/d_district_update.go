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

//    fmt.Fprintf( w, "d_district_update start \n" )  // 繝・ヰ繝・け

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

//	   fmt.Fprintf( w, "d_district_update :error updidw %v\n", updidw )  // 繝・ヰ繝・け

	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
	}
    updid := int64(updidw)

//    fmt.Fprintf( w, "d_district_update : updidw %v\n", updidw )  // 繝・ヰ繝・け
//    fmt.Fprintf( w, "d_district_update : updid %v\n", updid )  // 繝・ヰ繝・け

//	key := datastore.NewKey(c, "D_District", "", updid, nil)
    key := datastore.IDKey("D_District", updid, nil)

    if err := client.Get(ctx, key , &g ) ; err != nil {
//	if err := datastore.Get(c, key, &g); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

    g.District_Name = r.FormValue("district_name")  // 蝨ｰ蛹ｺ蜷阪ｒ繧ｲ繝・ヨ

	district_no := r.FormValue("district_no")         // 蝨ｰ蛹ｺNo.繧偵ご繝・ヨ
//	fmt.Fprintf( w, "d_district_update : district_no %v\n", district_no )  // 繝・ヰ繝・け
//	fmt.Fprintf( w, "d_district_update : district_name %v\n", g.District_Name )  // 繝・ヰ繝・け

	district_now ,err := strconv.Atoi(district_no)  // 譁・ｭ励・謨ｴ謨ｰ蛹・	if err != nil {
//		http.Error(w,err.Error(), http.StatusInternalServerError)
//       fmt.Fprintf( w, "d_district_update : a number must be half-width characters %v\n"  )
		return
	}

	g.District_No = int64(district_now)   // 謨ｴ謨ｰ縺ｮ64繝薙ャ繝亥喧

//	fmt.Fprintf( w, "d_district_update : g.District_Name %v\n", g.District_Name )  // 繝・ヰ繝・け
//	fmt.Fprintf( w, "d_district_update : g.District_No %v\n", g.District_No )  // 繝・ヰ繝・け

    if _, err := client.Put(ctx, key, &g ); err != nil {
//	if _, err := datastore.Put(c, key, &g); err != nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}

/// 繝｢繝九ち繝ｼ縲蜀崎｡ｨ遉ｺ ///

	process.D_district_showall1(w , r )

//	http.Redirect(w, r, "/", http.StatusFound)


}
