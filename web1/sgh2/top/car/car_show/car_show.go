package car_show

import (

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

///                       　
///     show car inf.
///

func Car_show(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintln( w, "car_show start \n" )

    var d_district      type2.D_District

    var d_district_temp type2.D_District_Temp

    var idmy int64

    select_idw , err := strconv.Atoi(r.FormValue("id"))

	if err  != nil {
//	   fmt.Fprintf( w, "car_show :error select_idw %v\n", select_idw )
	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
	}
    select_id := int64(select_idw)

//    fmt.Fprintf( w, "car_show : select_idw %v\n", select_idw )
//    fmt.Fprintf( w, "car_show : select_id %v\n", select_id )

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

    key := datastore.IDKey("D_District", select_id, nil)

    if err := client.Get(ctx, key , &d_district ) ; err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}


//     initialize district temp. inf.

//    initialize.D_district_temp (w , r )

//    _ = datastore2.D_store( "D_District_Temp" ,"initialize" ,idmy , w , r  )
    _ = datastore2.Datastore_sgh( "D_District_Temp" ,"initialize" ,idmy , w , r  )

    d_district_temp.District_No   = d_district.District_No
    d_district_temp.District_Name = d_district.District_Name

    new_key := datastore.IncompleteKey("D_District_Temp", nil)

    if _, err = client.Put(ctx, new_key, &d_district_temp ); err != nil {

		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}

///
///      show car inf. on web
///

	process.Car_show(w , r ,d_district.District_No)

//	fmt.Fprintf( w, "d_district_area : normal end \n" )

}
