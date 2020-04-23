package d_district_area

import (

	    "net/http"
//	    "fmt"
//	    "github.com/sawaq7/go12_ver1/client/sgh"
        "github.com/sawaq7/go12_ver1/client/sgh/process"
        "github.com/sawaq7/go12_ver1/client/sgh/type2"
	    "strconv"
	    "github.com/sawaq7/go12_ver1/client/sgh/datastore2"
//	    "time"

        "cloud.google.com/go/datastore"
        "context"
        "os"
                                                  )

///
/// 　　   show area inf. for each district
///

func D_district_area(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "d_district_area start \n" )

    var g type2.D_District

    var g2 type2.D_District_Temp

    var idmy int64

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

    updidw , err := strconv.Atoi(r.FormValue("id"))

	if err  != nil {
//	   fmt.Fprintf( w, "d_district_area :error updidw %v\n", updidw )
	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
	}
    updid := int64(updidw)

//    fmt.Fprintf( w, "d_district_area : updidw %v\n", updidw )
//    fmt.Fprintf( w, "d_district_area : updid %v\n", updid )

    key := datastore.IDKey("D_District", updid, nil)

    if err := client.Get(ctx, key , &g ) ; err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}


//     ini. and set temp.-file

    _ = datastore2.Datastore_sgh( "D_District_Temp" ,"initialize" ,idmy , w , r  )

    g2.District_No   = g.District_No
    g2.District_Name = g.District_Name

    new_key := datastore.IncompleteKey("D_District_Temp", nil)

    if _, err = client.Put(ctx, new_key, &g2 ); err != nil {

		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}

    //    show area inf. on web
	process.D_district_area(w , r ,g.District_No)

//	fmt.Fprintf( w, "d_district_area : normal end \n" )

}
