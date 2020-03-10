package process

import (
//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "net/http"
	    "strconv"
//	    "fmt"

	    "basic/date1"
	    "client/sgh/type2"

	    "cloud.google.com/go/datastore"
	    "context"
	    "os"

                                                )


func Deliver_update_single(w http.ResponseWriter, r *http.Request ,updid int64) {

//     IN    w      : レスポンスライター
//     IN    r      : リクエストパラメータ
//     IN  updid　  : 修正するデータストアのレコードのid

//    fmt.Fprintf( w, "deliver_update_single start \n" )  // デバック
    var g type2.Deliver

    projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if projectID == "" {

      projectID = "sample-7777"

	}

//	c := appengine.NewContext(r)
    ctx := context.Background()

    client, err := datastore.NewClient(ctx, projectID)

    key := datastore.IDKey("Deliver", updid, nil)

    if err := client.Get(ctx, key , &g ) ; err != nil {
//	key := datastore.NewKey(c, "Deliver", "", updid, nil)
//	if err := datastore.Get(c, key, &g); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

    g.Area_Name = r.FormValue("area_name")

    g.Date   = r.FormValue("date")

    g.Date_Real = date1.Date_realdata_get( w  ,g.Date )   // タイムデータ作成

	number := r.FormValue("number")

//	fmt.Fprintf( w, "deliver_update_single : number %v\n", number )  // デバック

	numberw ,err := strconv.Atoi(number)
	if err != nil {
//		http.Error(w,err.Error(), http.StatusInternalServerError)
//       fmt.Fprintf( w, "deliver_update_single : a number must be half-width characters %v\n"  )
//		return
	}

	g.Number = int64(numberw)

	private_no := r.FormValue("private_no")         // 個人Noをゲット
//	fmt.Fprintf( w, "deliver_update_single : private_no %v\n", private_no )  // デバック

	private_now ,err := strconv.Atoi(private_no)  // 個人Noの整数化
	if err != nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)

	}

	g.Private_No = int64(private_now)   // 整数の64ビット化

	car_no := r.FormValue("car_no")         // 個人Noをゲット
//	fmt.Fprintf( w, "deliver_update_single : car_no %v\n", car_no )  // デバック

	car_now ,err := strconv.Atoi(car_no)  // 個人Noの整数化
	if err != nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)

	}

	g.Car_No = int64(car_now)   // 整数の64ビット化

//	fmt.Fprintf( w, "deliver_update_single : g.Area_Name %v\n", g.Area_Name )  // デバック
//	fmt.Fprintf( w, "deliver_update_single : g.Course_No %v\n", g.Course_No )  // デバック
//	fmt.Fprintf( w, "deliver_update_single : g.Private_No %v\n", g.Private_No )  // デバック
//	fmt.Fprintf( w, "deliver_update_single : g.Car_No %v\n", g.Car_No )  // デバック
//	fmt.Fprintf( w, "deliver_update_single : g.Number %v\n", g.Number )  // デバック

    if _, err = client.Put(ctx, key, &g ); err != nil {
//	if _, err := datastore.Put(c, key, &g); err != nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}

}

