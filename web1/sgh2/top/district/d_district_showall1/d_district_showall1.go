package d_district_showall1

import (

	    "net/http"
//	    "fmt"
	    "github.com/sawaq7/go12_ver1/client/sgh/process"
	    "github.com/sawaq7/go12_ver1/client/sgh/type2"
	    "strconv"
//	    "time"

        "cloud.google.com/go/datastore"
        "context"
        "os"
                                                  )

///
/// 　　   show district inf.
///

func D_district_showall1(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "d_district_showall1 start \n" )

    var g type2.D_District

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

	g.District_Name = r.FormValue("district_name")

	district_no := r.FormValue("district_no")
//	fmt.Fprintf( w, "d_district_showall1 : district_no %v\n", district_no )
//	fmt.Fprintf( w, "d_district_showall1 : district_name %v\n", g.District_Name )

	district_now ,err := strconv.Atoi(district_no)
	if err != nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)
//       fmt.Fprintf( w, "d_district_showall1 : a number must be half-width characters %v\n"  )
		return
	}

	g.District_No = int64(district_now)

    ///  put  new record in d.s.
    new_key := datastore.IncompleteKey("D_District", nil)

    if _, err = client.Put(ctx, new_key, &g ); err != nil {

		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}

    //    show district inf.  on web
	process.D_district_showall1(w , r )

//	fmt.Fprintf( w, "d_district_showall1 : normal end \n" )




}
