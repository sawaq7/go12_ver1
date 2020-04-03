package private_update

import (

	    "strconv"
//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "net/http"
//	    "fmt"
	    "github.com/sawaq7/go12_ver1/client/sgh/process"

        "github.com/sawaq7/go12_ver1/client/sgh/type2"

        "cloud.google.com/go/datastore"
	    "context"
	    "os"
                                                   )

func Private_update(w http.ResponseWriter, r *http.Request) {

	var private type2.Private

//    fmt.Fprintf( w, "private_update start \n" )  // ãEãE¯

	updidw , err := strconv.Atoi(r.FormValue("id"))
	if err  != nil {

//	   fmt.Fprintf( w, "private_update :error updidw %v\n", updidw )  // ãEãE¯

	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
	}
    updid := int64(updidw)

//    fmt.Fprintf( w, "private_update : updidw %v\n", updidw )  // ãEãE¯
//    fmt.Fprintf( w, "private_update : updid %v\n", updid )  // ãEãE¯

    projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if projectID == "" {

      projectID = "sample-7777"

	}

//	c := appengine.NewContext(r)
    ctx := context.Background()

    client, err := datastore.NewClient(ctx, projectID)
    if err != nil {
       http.Error(w,err.Error(), http.StatusInternalServerError)
		return
    }

    key := datastore.IDKey("Private", updid, nil)

    if err := client.Get(ctx, key , &private ) ; err != nil {
//	key := datastore.NewKey(c, "Private", "", updid, nil)
//	if err := datastore.Get(c, key, &private); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

    private.Worker_Name = r.FormValue("worker_name")  /// åäººåãã²ãE
//	fmt.Fprintf( w, "private_update : worker_name %v\n", private.Worker_Name )  // ãEãE¯

	worker_no := r.FormValue("worker_no")             /// åäººNo.ãã²ãE
//	fmt.Fprintf( w, "private_update : worker_no %v\n", worker_no )  // ãEãE¯


	worker_now ,err := strconv.Atoi(worker_no)           // æE­ãEæ´æ°åE	if err != nil {

//       fmt.Fprintf( w, "private_update : a number must be half-width characters %v\n"  )
		return
	}

	private.Worker_No = int64(worker_now)                // æ´æ°ã®64ãããå

	private.Worker_Type = r.FormValue("worker_type")   /// ã¯ã¼ã«ã¼ã¿ã¤ããã²ãE

	worker_salary_str  := r.FormValue("worker_salary") /// ã¯ã¼ã«ã¼ãµã©ãªã¼ãã²ãE

	private.Worker_Salary , _ = strconv.ParseFloat( worker_salary_str,64 )  // ã¯ã¼ã«ã¼ãµã©ãªã¼ãfloat64ã«å¤æ

	private.Worker_Twh  = 50.0 * 52.14                 /// å¹´éç·å´åæéãè¨ç®E
	private.Worker_H_Pay  = private.Worker_Salary * 10000. / private.Worker_Twh  /// æçµ¦ãè¨ç®ã

    if _, err = client.Put(ctx, key, &private ); err != nil {
//	if _, err := datastore.Put(c, key, &private); err != nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}

/// ã¢ãã¿ã¼ãåè¡¨ç¤º ///

	process.Private_showall1(w , r )

}
