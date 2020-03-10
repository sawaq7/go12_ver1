package payment_register

import (
//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "net/http"
//	    "fmt"
//	    "client/sgh"

        "client/reserve/process4"
        "client/reserve/type6"
	    "strconv"

        "client/reserve/datastore6/initialize3"

        "cloud.google.com/go/datastore"
        "context"
        "os"

                                                  )

func Payment_register(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "payment_register start \n" )  // デバック

    var guest type6.Guest

    var guest2 type6.Guest_Temp

    updidw , err := strconv.Atoi(r.FormValue("id"))

	if err  != nil {
//	   fmt.Fprintf( w, "payment_register :error updidw %v\n", updidw )  // デバック
	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
	}
    updid := int64(updidw)

//    fmt.Fprintf( w, "payment_register : updidw %v\n", updidw )  // デバック
//    fmt.Fprintf( w, "payment_register : updid %v\n", updid )  // デバック

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

	key := datastore.IDKey("Guest", updid, nil)

    if err := client.Get(ctx, key , &guest ) ; err != nil {
//	key := datastore.NewKey(c, "Guest", "", updid, nil)
//	if err := datastore.Get(c, key, &guest); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}


// temporary-fileをイニシャライズ  & セット//

//    _ = datastore2.Datastore_sgh( "D_District_Temp" ,"initialize" ,idmy , w , r  )

    initialize3.Guest_temp (w ,r )

    guest2.Guest_No   = guest.Guest_No
    guest2.Guest_Name = guest.Guest_Name

    new_key := datastore.IncompleteKey("Guest_Temp", nil)

    if _, err = client.Put(ctx, new_key, &guest2 ); err != nil {
//	if _, err := datastore.Put(c, datastore.NewIncompleteKey(c, "Guest_Temp", nil), &guest2); err != nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}

/// モニター　表示 ///

	process4.Payment_register(w , r ,guest.Guest_No)

//	fmt.Fprintf( w, "reserve_register : normal end \n" )  // デバック




}
