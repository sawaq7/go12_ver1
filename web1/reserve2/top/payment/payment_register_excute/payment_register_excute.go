package payment_register_excute

import (

//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "net/http"
//	    "fmt"

	    "client/reserve/process4"
        "client/reserve/datastore6/check5"
	    "client/reserve/type6"

	    "strconv"

	    "cloud.google.com/go/datastore"
	    "context"
	    "os"
                                                  )

func Payment_register_excute(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "payment_register_excute start \n" )  // デバック

	var guest_payment type6.Guest_Payment

///  temporary-fileより、地区NO・地区名をGET

//    flexible_out := datastore2.Datastore_sgh( "D_District_Temp" ,"check" ,idmy , w , r  )

    general_work := check5.Guest_temp (w , r )

//    value2, _ := flexible_out.([]type5.General_Work)

    guest_payment.Guest_No   = general_work[0].Int64_Work
    guest_payment.Guest_Name = general_work[0].String_Work

//	fmt.Fprintf( w, "payment_register_excute : guest_payment.Guest_No %v\n", guest_payment.Guest_No )  // デバック
//	fmt.Fprintf( w, "payment_register_excute : guest_payment.Guest_Name %v\n", guest_payment.Guest_Name )  // デバック

// 空インターフェイス変数よりバリュー値をゲット

//    value, _ := count.(int64)

//	fmt.Fprintf( w, "payment_register_excute count %v   \n" , count  )  // デバック
//	fmt.Fprintf( w, "payment_register_excute district_no %v   \n" , district_no  )  // デバック

    guest_payment.Date   = r.FormValue("date")

    guest_payment.Item   = r.FormValue("item")

	amount               := r.FormValue("amount")
	amountw ,err := strconv.Atoi(amount)  // 文字の整数化
	if err != nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)

		return
	}

	guest_payment.Amount = int64(amountw)   // 整数の64ビット化

//	fmt.Fprintf( w, "payment_register_excute : guest_payment.Item %v\n", guest_payment.Item )  // デバック
//	fmt.Fprintf( w, "payment_register_excute : guest_payment.Amount %v\n", guest_payment.Amount )  // デバック

/// データストアーにデータをセット ///

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

/// モニター　再表示 ///

	process4.Payment_register(w , r ,guest_payment.Guest_No)

//	fmt.Fprintf( w, "payment_register_excute : normal end \n" )  // デバック

}
