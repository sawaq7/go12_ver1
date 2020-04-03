package car_update

import (

	    "strconv"
//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
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

//    fmt.Fprintf( w, "car_update start \n" )  // γEγE―

	updidw , err := strconv.Atoi(r.FormValue("id"))
	if err  != nil {

//	   fmt.Fprintf( w, "car_update :error updidw %v\n", updidw )  // γEγE―

	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
	}
    updid := int64(updidw)

//    fmt.Fprintf( w, "car_update : updidw %v\n", updidw )  // γEγE―
//    fmt.Fprintf( w, "car_update : updid %v\n", updid )  // γEγE―

    projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if projectID == "" {

      projectID = "sample-7777"

	}

//	c := appengine.NewContext(r)
    ctx := context.Background()

    client, err := datastore.NewClient(ctx, projectID)
    if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

    key := datastore.IDKey("Car", updid, nil)

    if err := client.Get(ctx, key , &car ) ; err != nil {
//	key := datastore.NewKey(c, "Car", "", updid, nil)
//	if err := datastore.Get(c, key, &car); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

    car.Car_Name   = r.FormValue("car_name")      // ε·θ»εγγ²γE
	car.Car_Explain = r.FormValue("car_explain")  // ε·θ»θͺ¬ζγγ²γE

//	fmt.Fprintf( w, "car_update : car.Car_Name %v\n", car.Car_Name )  // γEγE―
//	fmt.Fprintf( w, "car_update : car.Car_Explain %v\n", car.Car_Explain )  // γEγE―

// γEEγΏγΉγγ’γγ1γ¬γ³γΌγγ’γEEγEEγE
    if _, err = client.Put(ctx, key, &car ); err != nil {
//	if _, err := datastore.Put( c, key, &car ); err != nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}

/// ε·θ»ζε ±γθ‘¨η€Ί ///

	process.Car_show( w , r ,car.District_No )

}
