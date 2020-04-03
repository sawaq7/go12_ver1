package private_showall2

import (
//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "net/http"
//	    "fmt"
	    "github.com/sawaq7/go12_ver1/client/sgh/process"
	    "github.com/sawaq7/go12_ver1/client/sgh/type2"
	    "strconv"

        "cloud.google.com/go/datastore"
	    "context"
	    "os"

                                                  )

func Private_showall2(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "private_showall2 start \n" )  // ãEãE¯

	var private type2.Private

	private.Worker_Name = r.FormValue("worker_name")  /// åäººåãã²ãE
//	fmt.Fprintf( w, "private_showall2 : worker_name %v\n", private.Worker_Name )  // ãEãE¯

	worker_no := r.FormValue("worker_no")             /// åäººNo.ãã²ãE
//	fmt.Fprintf( w, "private_showall2 : worker_no %v\n", worker_no )  // ãEãE¯


	worker_now ,err := strconv.Atoi(worker_no)           // æE­ãEæ´æ°åE	if err != nil {

//       fmt.Fprintf( w, "private_showall2 : a number must be half-width characters %v\n"  )
		return
	}

	private.Worker_No = int64(worker_now)                // æ´æ°ã®64ãããå

	private.Worker_Type = r.FormValue("worker_type")   /// ã¯ã¼ã«ã¼ã¿ã¤ããã²ãE

	worker_salary_str  := r.FormValue("worker_salary") /// ã¯ã¼ã«ã¼ãµã©ãªã¼ãã²ãE

	private.Worker_Salary , _ = strconv.ParseFloat( worker_salary_str,64 )  // ã¯ã¼ã«ã¼ãµã©ãªã¼ãfloat64ã«å¤æ

	private.Worker_Twh  = 50.0 * 52.14                 /// å¹´éç·å´åæéãè¨ç®E
	private.Worker_H_Pay  = private.Worker_Salary * 10000. / private.Worker_Twh  /// æçµ¦ãè¨ç®ã

    project_name := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if project_name == "" {

      project_name = "sample-7777"

	}

//	c := appengine.NewContext(r)
    ctx := context.Background()

    client, err := datastore.NewClient(ctx, project_name)
    if err != nil {
       http.Error(w,err.Error(), http.StatusInternalServerError)
		return
    }



/// ãEEã¿ã¹ãã¢ã¼ã«ãEEã¿ãã»ãE ///

    new_key := datastore.IncompleteKey("Private", nil)

    if _, err = client.Put(ctx, new_key, &private ); err != nil {
//	if _, err := datastore.Put(c, datastore.NewIncompleteKey(c, "Private", nil), &private); err != nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}

/// ã¢ãã¿ã¼ãåè¡¨ç¤º ///

	process.Private_showall1(w , r )

//	fmt.Fprintf( w, "private_showall2 : normal end \n" )  // ãEãE¯




}
