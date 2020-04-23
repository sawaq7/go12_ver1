package payment_register

import (

	    "net/http"
//	    "fmt"
//	    "github.com/sawaq7/go12_ver1/client/sgh"

        "github.com/sawaq7/go12_ver1/client/reserve/process4"
        "github.com/sawaq7/go12_ver1/client/reserve/type6"
	    "strconv"

        "github.com/sawaq7/go12_ver1/client/reserve/datastore6/initialize3"

        "cloud.google.com/go/datastore"
        "context"
        "os"
                                                                                )
///                         　　　　
///     register payment inf. in d.s.
///

func Payment_register(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "payment_register start \n" )

    var guest type6.Guest

    var guest2 type6.Guest_Temp

    updidw , err := strconv.Atoi(r.FormValue("id"))

	if err  != nil {
//	   fmt.Fprintf( w, "payment_register :error updidw %v\n", updidw )
	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
	}
    updid := int64(updidw)

//    fmt.Fprintf( w, "payment_register : updidw %v\n", updidw )
//    fmt.Fprintf( w, "payment_register : updid %v\n", updid )

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

	key := datastore.IDKey("Guest", updid, nil)

    if err := client.Get(ctx, key , &guest ) ; err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}


//    initialize guest tmp. file

//    _ = datastore2.Datastore_sgh( "D_District_Temp" ,"initialize" ,idmy , w , r  )

    initialize3.Guest_temp (w ,r )

    guest2.Guest_No   = guest.Guest_No
    guest2.Guest_Name = guest.Guest_Name

    new_key := datastore.IncompleteKey("Guest_Temp", nil)

    if _, err = client.Put(ctx, new_key, &guest2 ); err != nil {

		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}

///
///       show payment inf. on web
///

	process4.Payment_register(w , r ,guest.Guest_No)

//	fmt.Fprintf( w, "reserve_register : normal end \n" )




}
