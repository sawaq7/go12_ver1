package guest_show2

import (

//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "net/http"
//	    "fmt"

	    "github.com/sawaq7/go12_ver1/client/reserve/process4"
	    "github.com/sawaq7/go12_ver1/client/reserve/type6"
	    "strconv"

	    "cloud.google.com/go/datastore"
	    "context"
	    "os"
                                                  )

func Guest_show2(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "guest_show2 start \n" )  // ãEãE¯



	var guest type6.Guest

	guest.Guest_Name = r.FormValue("guest_name")  // å°åºåãã²ãE

	guest_no := r.FormValue("guest_no")         // å°åºNo.ãã²ãE
//	fmt.Fprintf( w, "guest_show2 : guest_no %v\n", guest_no )  // ãEãE¯

	guest_now ,err := strconv.Atoi(guest_no)  // æE­ãEæ´æ°åE	if err != nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)
//       fmt.Fprintf( w, "guest_show2 : a number must be half-width characters %v\n"  )
		return
	}

	guest.Guest_No = int64(guest_now)   // æ´æ°ã®64ãããå

/// ãEEã¿ã¹ãã¢ã¼ã«ãEEã¿ãã»ãE ///

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

    new_key := datastore.IncompleteKey("Guest", nil)

    if _, err = client.Put(ctx, new_key, &guest ); err != nil {
//	if _, err := datastore.Put(c, datastore.NewIncompleteKey(c, "Guest", nil), &guest); err != nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}

/// ã¢ãã¿ã¼ãåè¡¨ç¤º ///

    process4.Guest_show(w , r )

//	fmt.Fprintf( w, "guest_show2 : normal end \n" )  // ãEãE¯




}
