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

//    fmt.Fprintf( w, "car_update start \n" )  // 繝・ヰ繝・け

	updidw , err := strconv.Atoi(r.FormValue("id"))
	if err  != nil {

//	   fmt.Fprintf( w, "car_update :error updidw %v\n", updidw )  // 繝・ヰ繝・け

	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
	}
    updid := int64(updidw)

//    fmt.Fprintf( w, "car_update : updidw %v\n", updidw )  // 繝・ヰ繝・け
//    fmt.Fprintf( w, "car_update : updid %v\n", updid )  // 繝・ヰ繝・け

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

    car.Car_Name   = r.FormValue("car_name")      // 蜿ｷ霆雁錐繧偵ご繝・ヨ
	car.Car_Explain = r.FormValue("car_explain")  // 蜿ｷ霆願ｪｬ譏弱ｒ繧ｲ繝・ヨ

//	fmt.Fprintf( w, "car_update : car.Car_Name %v\n", car.Car_Name )  // 繝・ヰ繝・け
//	fmt.Fprintf( w, "car_update : car.Car_Explain %v\n", car.Car_Explain )  // 繝・ヰ繝・け

// 繝・・繧ｿ繧ｹ繝医い縺九ｉ1繝ｬ繧ｳ繝ｼ繝峨い繝・・繝・・繝・
    if _, err = client.Put(ctx, key, &car ); err != nil {
//	if _, err := datastore.Put( c, key, &car ); err != nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}

/// 蜿ｷ霆頑ュ蝣ｱ繧定｡ｨ遉ｺ ///

	process.Car_show( w , r ,car.District_No )

}
