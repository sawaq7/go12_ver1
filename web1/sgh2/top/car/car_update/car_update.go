package car_update

import (

	    "strconv"
//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "net/http"
//	    "fmt"
	    "client/sgh/process"

        "client/sgh/type2"

        "cloud.google.com/go/datastore"
	    "context"
	    "os"
                                                   )

func Car_update(w http.ResponseWriter, r *http.Request) {

	var car type2.Car

//    fmt.Fprintf( w, "car_update start \n" )  // デバック

	updidw , err := strconv.Atoi(r.FormValue("id"))
	if err  != nil {

//	   fmt.Fprintf( w, "car_update :error updidw %v\n", updidw )  // デバック

	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
	}
    updid := int64(updidw)

//    fmt.Fprintf( w, "car_update : updidw %v\n", updidw )  // デバック
//    fmt.Fprintf( w, "car_update : updid %v\n", updid )  // デバック

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

    car.Car_Name   = r.FormValue("car_name")      // 号車名をゲット
	car.Car_Explain = r.FormValue("car_explain")  // 号車説明をゲット

//	fmt.Fprintf( w, "car_update : car.Car_Name %v\n", car.Car_Name )  // デバック
//	fmt.Fprintf( w, "car_update : car.Car_Explain %v\n", car.Car_Explain )  // デバック

// データストアから1レコードアップデート

    if _, err = client.Put(ctx, key, &car ); err != nil {
//	if _, err := datastore.Put( c, key, &car ); err != nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}

/// 号車情報を表示 ///

	process.Car_show( w , r ,car.District_No )

}
