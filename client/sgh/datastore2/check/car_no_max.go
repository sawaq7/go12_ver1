package check

import (
//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "net/http"
//	    "fmt"
//	    "html/template"
//	    "web/htmls/sgh"
	    "github.com/sawaq7/go12_ver1/client/sgh/type2"

        "cloud.google.com/go/datastore"
	    "context"
	    "os"

                                                )

///
/// 縲縲縲縲縲逋ｻ骭ｲ蜿ｷ霆頑焚繧偵ご繝・ヨ縺吶ｋ
///

func Car_no_max(w http.ResponseWriter, r *http.Request ,district_no int64)  (car_number int64 ) {

//     IN    w      縲縲: 繝ｬ繧ｹ繝昴Φ繧ｹ繝ｩ繧､繧ｿ繝ｼ
//     IN    r      縲縲: 繝ｪ繧ｯ繧ｨ繧ｹ繝医ヱ繝ｩ繝｡繝ｼ繧ｿ
//     IN 縲district_no : 蝨ｰ蛹ｺNo

//    fmt.Fprintf( w, "car_no_max start \n" )  // 繝・ヰ繝・け

    var idmy int64

    projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if projectID == "" {

      projectID = "sample-7777"

	}

//	c := appengine.NewContext(r)
    ctx := context.Background()

    client, err := datastore.NewClient(ctx, projectID)

    query := datastore.NewQuery("Car").Order("Car_No")
//	q := datastore.NewQuery("Car").Order("Car_No")

    count, err := client.Count(ctx, query)
//	count, err := q.Count(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return idmy
	}

	car := make([]type2.Car, 0, count)

    if _, err := client.GetAll(ctx, query , &car) ; err != nil {
//	if _, err := q.GetAll(c, &car); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return idmy
	} else {

	  car_number = 0

	  for _, carw := range car {
        if district_no == carw.District_No {

		  car_number ++

        }
	  }
	}
//    fmt.Fprintf( w, "car_no_max car_number \n" , car_number)  // 繝・ヰ繝・け
	return car_number
}

