package d_district_area_update

import (

	    "strconv"

	    "net/http"
//	    "fmt"
	    "github.com/sawaq7/go12_ver1/client/sgh/process"

        "github.com/sawaq7/go12_ver1/client/sgh/type2"

        "cloud.google.com/go/datastore"
	    "context"
	    "os"
                                                   )

///
/// 　　   update area inf.
///

func D_district_area_update(w http.ResponseWriter, r *http.Request) {

	var g type2.D_Area

//    fmt.Fprintf( w, "d_district_area_update start \n" )

	 projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if projectID == "" {

      projectID = "sample-7777"

	}

    ctx := context.Background()

    client, err := datastore.NewClient(ctx, projectID)
    if err  != nil {

	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
	}

	updidw , err := strconv.Atoi(r.FormValue("id"))
	if err  != nil {

//	   fmt.Fprintf( w, "d_district_area_update :error updidw %v\n", updidw )

	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
	}
    updid := int64(updidw)

//    fmt.Fprintf( w, "d_district_area_update : updidw %v\n", updidw )
//    fmt.Fprintf( w, "d_district_area_update : updid %v\n", updid )

    key := datastore.IDKey("D_Area", updid, nil)

    if err := client.Get(ctx, key , &g ) ; err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

    g.Area_Name = r.FormValue("area_name")
    g.Area_Detail = r.FormValue("area_detail")

//	fmt.Fprintf( w, "d_district_area_update : g.Area_Name %v\n", g.Area_Name )
//	fmt.Fprintf( w, "d_district_area_update : g.Area_Detail %v\n", g.Area_Detail )

    ///  put  new record in d.s.
    if _, err = client.Put(ctx, key, &g ); err != nil {

		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}

    //    show area inf. for each district-no  on web
	process.D_district_area_show(w , r  ,g.District_No)

}
