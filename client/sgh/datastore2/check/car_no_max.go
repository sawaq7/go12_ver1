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
/// 　　　　　登録号車数をゲチE��する
///

func Car_no_max(w http.ResponseWriter, r *http.Request ,district_no int64)  (car_number int64 ) {

//     IN    w      　　: レスポンスライター
//     IN    r      　　: リクエストパラメータ
//     IN 　district_no : 地区No

//    fmt.Fprintf( w, "car_no_max start \n" )  // チE��チE��

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
//    fmt.Fprintf( w, "car_no_max car_number \n" , car_number)  // チE��チE��
	return car_number
}

