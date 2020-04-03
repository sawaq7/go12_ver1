package payment_register_excute

import (

//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "net/http"
//	    "fmt"

	    "github.com/sawaq7/go12_ver1/client/reserve/process4"
        "github.com/sawaq7/go12_ver1/client/reserve/datastore6/check5"
	    "github.com/sawaq7/go12_ver1/client/reserve/type6"

	    "strconv"

	    "cloud.google.com/go/datastore"
	    "context"
	    "os"
                                                  )

func Payment_register_excute(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "payment_register_excute start \n" )  // ãEãE¯

	var guest_payment type6.Guest_Payment

///  temporary-fileãããå°åºNOã»å°åºåãGET

//    flexible_out := datastore2.Datastore_sgh( "D_District_Temp" ,"check" ,idmy , w , r  )

    general_work := check5.Guest_temp (w , r )

//    value2, _ := flexible_out.([]type5.General_Work)

    guest_payment.Guest_No   = general_work[0].Int64_Work
    guest_payment.Guest_Name = general_work[0].String_Work

//	fmt.Fprintf( w, "payment_register_excute : guest_payment.Guest_No %v\n", guest_payment.Guest_No )  // ãEãE¯
//	fmt.Fprintf( w, "payment_register_excute : guest_payment.Guest_Name %v\n", guest_payment.Guest_Name )  // ãEãE¯

// ç©ºã¤ã³ã¿ã¼ãã§ã¤ã¹å¤æ°ããããªã¥ã¼å¤ãã²ãE

//    value, _ := count.(int64)

//	fmt.Fprintf( w, "payment_register_excute count %v   \n" , count  )  // ãEãE¯
//	fmt.Fprintf( w, "payment_register_excute district_no %v   \n" , district_no  )  // ãEãE¯

    guest_payment.Date   = r.FormValue("date")

    guest_payment.Item   = r.FormValue("item")

	amount               := r.FormValue("amount")
	amountw ,err := strconv.Atoi(amount)  // æE­ãEæ´æ°åE	if err != nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)

		return
	}

	guest_payment.Amount = int64(amountw)   // æ´æ°ã®64ãããå

//	fmt.Fprintf( w, "payment_register_excute : guest_payment.Item %v\n", guest_payment.Item )  // ãEãE¯
//	fmt.Fprintf( w, "payment_register_excute : guest_payment.Amount %v\n", guest_payment.Amount )  // ãEãE¯

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

    new_key := datastore.IncompleteKey("Guest_Payment", nil)

    if _, err = client.Put(ctx, new_key, &guest_payment ); err != nil {
//	if _, err := datastore.Put(c, datastore.NewIncompleteKey(c, "Guest_Payment", nil), &guest_payment); err != nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}

/// ã¢ãã¿ã¼ãåè¡¨ç¤º ///

	process4.Payment_register(w , r ,guest_payment.Guest_No)

//	fmt.Fprintf( w, "payment_register_excute : normal end \n" )  // ãEãE¯

}
