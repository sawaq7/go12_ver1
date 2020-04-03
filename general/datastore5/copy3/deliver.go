package copy3

import (
//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "net/http"
//	    "fmt"

        "github.com/sawaq7/go12_ver1/client/sgh/type2"

        "os"
        "cloud.google.com/go/datastore"
        "context"

                                                )

///
/// ãEEã¿ã¹ãã¢ãã³ããEãã  (ãããEã·ãE¯åãEãDeliverãEE///

func Deliver( w http.ResponseWriter, r *http.Request ,basic_name string ,copy_file string ,new_file string ) {

//     IN    w        : ã¬ã¹ãã³ã¹ã©ã¤ã¿ã¼
//     IN    r        : ãªã¯ã¨ã¹ããã©ã¡ã¼ã¿
//     IN  basic_name : åºæ¬ã®ãEEã¿ã¹ãã¢åE//     IN  copy_file  : ã³ããEåEEãEEã¿ã¹ãã¢åE//     IN  new_file   : ãã¥ã¼ãEEã¿ã¹ãã¢åE//    OUT  err        : ã¨ã©ã¼ã¡ãE»ã¼ã¸

//    fmt.Fprintf( w, "copy3.deliver start \n" )  // ãEãE¯
//    fmt.Fprintf( w, "copy3.deliver basic_name %v\n" ,basic_name)  // ãEãE¯

///
///  ãã­ã¸ã§ã¯ãEDãã²ãE
///
    projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if projectID == "" {

      projectID = "sample-7777"

	}

//	c := appengine.NewContext(r)
    ctx := context.Background()

    client, err := datastore.NewClient(ctx, projectID)

	query := datastore.NewQuery(copy_file)
//	q := datastore.NewQuery(copy_file)

//	count, err := q.Count(c)
    count, err := client.Count(ctx, query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

//    fmt.Fprintf( w, "copy3.deliver count %v\n" ,count)  // ãEãE¯

///
/// ã¯ã­ã¼ã³æE ±ãSETããã¯ã¼ã¯ã¨ãªã¢ãç¢ºä¿E///

    ds_data := make([]type2.Deliver, 0, count)

    if _, err := client.GetAll(ctx, query , &ds_data);  err != nil {
//	if _, err := q.GetAll(c, &ds_data);  err != nil {         // ã¯ã­ã¼ã³æE ±ãGET

	  http.Error(w, err.Error(), http.StatusInternalServerError)
	  return

	} else{
      for _, ds_dataw := range ds_data {                       //ããã¥ã¼ãã¡ã¤ã«ã«ã¯ã­ã¼ã³æE ±ãã»ãE

//	    if _, err := datastore.Put(c, datastore.NewIncompleteKey(c, new_file, nil), &ds_dataw); err != nil {
        new_key := datastore.IncompleteKey( new_file , nil)

        if _, err = client.Put(ctx, new_key, &ds_dataw ); err != nil {

		  http.Error(w,err.Error(), http.StatusInternalServerError)
		  return

	    }

	  }
	}

//	fmt.Fprintf( w, "copy3.deliver normal end \n" )  // ãEãE¯

    return
}

