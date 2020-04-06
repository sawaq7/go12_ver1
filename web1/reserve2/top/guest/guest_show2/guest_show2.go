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

//    fmt.Fprintf( w, "guest_show2 start \n" )  // チE��チE��



	var guest type6.Guest

	guest.Guest_Name = r.FormValue("guest_name")  // 地区名をゲチE��

	guest_no := r.FormValue("guest_no")         // 地区No.をゲチE��
//	fmt.Fprintf( w, "guest_show2 : guest_no %v\n", guest_no )  // チE��チE��

	guest_now ,err := strconv.Atoi(guest_no)  // 斁E���E整数匁E	if err != nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)
//       fmt.Fprintf( w, "guest_show2 : a number must be half-width characters %v\n"  )
		return
	}

	guest.Guest_No = int64(guest_now)   // 整数の64ビット化

/// チE�EタストアーにチE�EタをセチE�� ///

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

/// モニター　再表示 ///

    process4.Guest_show(w , r )

//	fmt.Fprintf( w, "guest_show2 : normal end \n" )  // チE��チE��




}