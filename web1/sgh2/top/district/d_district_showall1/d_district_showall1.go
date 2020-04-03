package d_district_showall1

import (

//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "net/http"
//	    "fmt"
	    "github.com/sawaq7/go12_ver1/client/sgh/process"
	    "github.com/sawaq7/go12_ver1/client/sgh/type2"
	    "strconv"
//	    "time"

        "cloud.google.com/go/datastore"
        "context"
        "os"
                                                  )

func D_district_showall1(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "d_district_showall1 start \n" )  // ãEãE¯

    var g type2.D_District

    project_name := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if project_name == "" {

      project_name = "sample-7777"

	}
    ctx := context.Background()
//    c := appengine.NewContext(r)

    client, err := datastore.NewClient(ctx, project_name)
    if err != nil {
       http.Error(w, err.Error(), http.StatusInternalServerError)
       return
    }

	g.District_Name = r.FormValue("district_name")  // å°åºåãã²ãE

	district_no := r.FormValue("district_no")         // å°åºNo.ãã²ãE
//	fmt.Fprintf( w, "d_district_showall1 : district_no %v\n", district_no )  // ãEãE¯
//	fmt.Fprintf( w, "d_district_showall1 : district_name %v\n", g.District_Name )  // ãEãE¯

	district_now ,err := strconv.Atoi(district_no)  // æE­ãEæ´æ°åE	if err != nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)
//       fmt.Fprintf( w, "d_district_showall1 : a number must be half-width characters %v\n"  )
		return
	}

	g.District_No = int64(district_now)   // æ´æ°ã®64ãããå

/// ãEEã¿ã¹ãã¢ã¼ã«ãEEã¿ãã»ãE ///

    new_key := datastore.IncompleteKey("D_District", nil)

    if _, err = client.Put(ctx, new_key, &g ); err != nil {
//	if _, err := datastore.Put(c, datastore.NewIncompleteKey(c, "D_District", nil), &g); err != nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}

/// ã¢ãã¿ã¼ãåè¡¨ç¤º ///

	process.D_district_showall1(w , r )

//	fmt.Fprintf( w, "d_district_showall1 : normal end \n" )  // ãEãE¯




}
