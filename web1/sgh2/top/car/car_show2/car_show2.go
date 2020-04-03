package car_show2

import (
//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "net/http"
//	    "fmt"
	    "github.com/sawaq7/go12_ver1/client/sgh/process"
//	    "github.com/sawaq7/go12_ver1/client/sgh/datastore2/check"
	    "github.com/sawaq7/go12_ver1/client/sgh/datastore2"
	    "github.com/sawaq7/go12_ver1/client/sgh/type2"
	    "github.com/sawaq7/go12_ver1/general/type5"

        "cloud.google.com/go/datastore"
	    "context"
	    "os"
                                                  )

func Car_show2(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "car_show2 start \n" )  // ãEãE¯

	var car type2.Car

	var idmy int64

///  temporary-fileãããå°åºNOã»å°åºåãGET

//    flexible_out := datastore2.D_store( "D_District_Temp" ,"check" ,idmy , w , r  )
    flexible_out := datastore2.Datastore_sgh( "D_District_Temp" ,"check" ,idmy , w , r  )

    value2, _ := flexible_out.([]type5.General_Work)

    car.District_No = value2[0].Int64_Work
    car.District_Name = value2[0].String_Work


//	fmt.Fprintf( w, "car_show2 : g.District_No %v\n", g.District_No )  // ãEãE¯
//	fmt.Fprintf( w, "car_show2 : g.District_Name %v\n", g.District_Name )  // ãEãE¯

//  ã«ã¼NOãã®æ¢å­ãEMAXå¤ãGET

	count := datastore2.Datastore_sgh( "Car" ,"check" ,car.District_No , w , r  )

     // ç©ºã¤ã³ã¿ã¼ãã§ã¤ã¹å¤æ°ããããªã¥ã¼å¤ãã²ãE

     value, _ := count.(int64)

//	fmt.Fprintf( w, "car_show2 value %v   \n" , value  )  // ãEãE¯
//	fmt.Fprintf( w, "car_show2 district_no %v   \n" , district_no  )  // ãEãE¯

    car.Car_No     = car.Car_No + int64(value + 1)
    car.Car_Name   = r.FormValue("car_name")
	car.Car_Explain = r.FormValue("car_explain")

//    fmt.Fprintf( w, "car_show2 : car.Car_No %v\n", car.Car_No )  // ãEãE¯
//	fmt.Fprintf( w, "car_show2 : car.Car_Name %v\n", car.Car_Name )  // ãEãE¯
//	fmt.Fprintf( w, "car_show2 : car.Car_Explain %v\n", car.Car_Explain )  // ãEãE¯

/// ãEEã¿ã¹ãã¢ã¼ã«ãEEã¿ãã»ãE ///

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

    new_key := datastore.IncompleteKey("Car", nil)

    if _, err = client.Put(ctx, new_key, &car ); err != nil {
//	if _, err := datastore.Put(c, datastore.NewIncompleteKey(c, "Car", nil), &car); err != nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}

/// å·è»æå ±ãè¡¨ç¤º ///

	process.Car_show( w , r ,car.District_No )

//	fmt.Fprintf( w, "car_show2 : normal end \n" )  // ãEãE¯

}
