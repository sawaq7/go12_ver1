package d_district_update

import (

	    "strconv"

	    "net/http"
//	    "fmt"
	    "github.com/sawaq7/go12_ver1/client/sgh/process"

        "github.com/sawaq7/go12_ver1/client/sgh/type2"

        "os"
        "cloud.google.com/go/datastore"
        "context"
                                                   )

///
/// 　　   update district inf.
///

func D_district_update(w http.ResponseWriter, r *http.Request) {

	var g type2.D_District

//    fmt.Fprintf( w, "d_district_update start \n" )

    project_name := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if project_name == "" {

      project_name = "sample-7777"

	}


    ctx := context.Background()

    client, err := datastore.NewClient(ctx, project_name)
    if err != nil {
       http.Error(w,err.Error(), http.StatusInternalServerError)
		return
    }

	updidw , err := strconv.Atoi(r.FormValue("id"))
	if err  != nil {

//	   fmt.Fprintf( w, "d_district_update :error updidw %v\n", updidw )

	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
	}
    updid := int64(updidw)

//    fmt.Fprintf( w, "d_district_update : updidw %v\n", updidw )
//    fmt.Fprintf( w, "d_district_update : updid %v\n", updid )

    key := datastore.IDKey("D_District", updid, nil)

    if err := client.Get(ctx, key , &g ) ; err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

    g.District_Name = r.FormValue("district_name")

	district_no := r.FormValue("district_no")
//	fmt.Fprintf( w, "d_district_update : district_no %v\n", district_no )
//	fmt.Fprintf( w, "d_district_update : district_name %v\n", g.District_Name )

	district_now ,err := strconv.Atoi(district_no)
	if err != nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)
//       fmt.Fprintf( w, "d_district_update : a number must be half-width characters %v\n"  )
		return
	}

	g.District_No = int64(district_now)

//	fmt.Fprintf( w, "d_district_update : g.District_Name %v\n", g.District_Name )
//	fmt.Fprintf( w, "d_district_update : g.District_No %v\n", g.District_No )

    if _, err := client.Put(ctx, key, &g ); err != nil {

		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}

    //    show district inf.  on web
	process.D_district_showall1(w , r )


}
