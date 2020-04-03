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

//    fmt.Fprintf( w, "car_show2 start \n" )  // 繝・ヰ繝・け

	var car type2.Car

	var idmy int64

///  temporary-file繧医ｊ縲∝慍蛹ｺNO繝ｻ蝨ｰ蛹ｺ蜷阪ｒGET

//    flexible_out := datastore2.D_store( "D_District_Temp" ,"check" ,idmy , w , r  )
    flexible_out := datastore2.Datastore_sgh( "D_District_Temp" ,"check" ,idmy , w , r  )

    value2, _ := flexible_out.([]type5.General_Work)

    car.District_No = value2[0].Int64_Work
    car.District_Name = value2[0].String_Work


//	fmt.Fprintf( w, "car_show2 : g.District_No %v\n", g.District_No )  // 繝・ヰ繝・け
//	fmt.Fprintf( w, "car_show2 : g.District_Name %v\n", g.District_Name )  // 繝・ヰ繝・け

//  繧ｫ繝ｼNO縲縺ｮ譌｢蟄倥・MAX蛟､繧竪ET

	count := datastore2.Datastore_sgh( "Car" ,"check" ,car.District_No , w , r  )

     // 遨ｺ繧､繝ｳ繧ｿ繝ｼ繝輔ぉ繧､繧ｹ螟画焚繧医ｊ繝舌Μ繝･繝ｼ蛟､繧偵ご繝・ヨ

     value, _ := count.(int64)

//	fmt.Fprintf( w, "car_show2 value %v   \n" , value  )  // 繝・ヰ繝・け
//	fmt.Fprintf( w, "car_show2 district_no %v   \n" , district_no  )  // 繝・ヰ繝・け

    car.Car_No     = car.Car_No + int64(value + 1)
    car.Car_Name   = r.FormValue("car_name")
	car.Car_Explain = r.FormValue("car_explain")

//    fmt.Fprintf( w, "car_show2 : car.Car_No %v\n", car.Car_No )  // 繝・ヰ繝・け
//	fmt.Fprintf( w, "car_show2 : car.Car_Name %v\n", car.Car_Name )  // 繝・ヰ繝・け
//	fmt.Fprintf( w, "car_show2 : car.Car_Explain %v\n", car.Car_Explain )  // 繝・ヰ繝・け

/// 繝・・繧ｿ繧ｹ繝医い繝ｼ縺ｫ繝・・繧ｿ繧偵そ繝・ヨ ///

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

/// 蜿ｷ霆頑ュ蝣ｱ繧定｡ｨ遉ｺ ///

	process.Car_show( w , r ,car.District_No )

//	fmt.Fprintf( w, "car_show2 : normal end \n" )  // 繝・ヰ繝・け

}
