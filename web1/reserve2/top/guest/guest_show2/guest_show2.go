package guest_show2

import (

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

//    fmt.Fprintf( w, "guest_show2 start \n" )

	var guest type6.Guest

	guest.Guest_Name = r.FormValue("guest_name")  //    get guest name

	guest_no := r.FormValue("guest_no")         //    get guest no

//	fmt.Fprintf( w, "guest_show2 : guest_no %v\n", guest_no )

	guest_now ,err := strconv.Atoi(guest_no)  // make an integer
	if err != nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)

//       fmt.Fprintf( w, "guest_show2 : a number must be half-width characters %v\n"  )
		return
	}

	guest.Guest_No = int64(guest_now)   //     make an integer64

///     set guest .inf. in d.s

    projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if projectID == "" {

      projectID = "sample-7777"

	}

    ctx := context.Background()

    client, err := datastore.NewClient(ctx, projectID)
    if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

    new_key := datastore.IncompleteKey("Guest", nil)

    if _, err = client.Put(ctx, new_key, &guest ); err != nil {

		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}

///
///     show guest inf. on web
///
    process4.Guest_show(w , r )

//	fmt.Fprintf( w, "guest_show2 : normal end \n" )




}
