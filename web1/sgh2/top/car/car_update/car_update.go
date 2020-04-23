package car_update

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

func Car_update(w http.ResponseWriter, r *http.Request) {

	var car type2.Car

//    fmt.Fprintf( w, "car_update start \n" )

	updidw , err := strconv.Atoi(r.FormValue("id"))
	if err  != nil {

//	   fmt.Fprintf( w, "car_update :error updidw %v\n", updidw )

	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
	}
    updid := int64(updidw)

//    fmt.Fprintf( w, "car_update : updidw %v\n", updidw )
//    fmt.Fprintf( w, "car_update : updid %v\n", updid )

    projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if projectID == "" {

      projectID = "sample-7777"

	}

    ctx := context.Background()

    client, err := datastore.NewClient(ctx, projectID)
    if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

    key := datastore.IDKey("Car", updid, nil)

    if err := client.Get(ctx, key , &car ) ; err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

    car.Car_Name   = r.FormValue("car_name")
	car.Car_Explain = r.FormValue("car_explain")

//	fmt.Fprintf( w, "car_update : car.Car_Name %v\n", car.Car_Name )
//	fmt.Fprintf( w, "car_update : car.Car_Explain %v\n", car.Car_Explain )


    if _, err = client.Put(ctx, key, &car ); err != nil {     //  put one record

		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}

///
///     show car inf. on web
///

	process.Car_show( w , r ,car.District_No )

}
