package private_update

import (

	    "strconv"

	    "net/http"
//	    "fmt"
	    "github.com/sawaq7/go12_ver1/client/sgh/process"

        "github.com/sawaq7/go12_ver1/client/sgh/type2"

        "cloud.google.com/go/datastore"
	    "context"
	    "os"
                                                   )

///
/// 　　   show update-private inf. on web
///

func Private_update(w http.ResponseWriter, r *http.Request) {

	var private type2.Private

//    fmt.Fprintf( w, "private_update start \n" )

	updidw , err := strconv.Atoi(r.FormValue("id"))
	if err  != nil {

//	   fmt.Fprintf( w, "private_update :error updidw %v\n", updidw )

	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
	}
    updid := int64(updidw)

//    fmt.Fprintf( w, "private_update : updidw %v\n", updidw )
//    fmt.Fprintf( w, "private_update : updid %v\n", updid )

    projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if projectID == "" {

      projectID = "sample-7777"

	}

    ctx := context.Background()

    client, err := datastore.NewClient(ctx, projectID)
    if err != nil {
       http.Error(w,err.Error(), http.StatusInternalServerError)
		return
    }

    key := datastore.IDKey("Private", updid, nil)

    if err := client.Get(ctx, key , &private ) ; err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

    private.Worker_Name = r.FormValue("worker_name")
//	fmt.Fprintf( w, "private_update : worker_name %v\n", private.Worker_Name )

	worker_no := r.FormValue("worker_no")
//	fmt.Fprintf( w, "private_update : worker_no %v\n", worker_no )


	worker_now ,err := strconv.Atoi(worker_no)
	if err != nil {

//       fmt.Fprintf( w, "private_update : a number must be half-width characters %v\n"  )
		return
	}

	private.Worker_No = int64(worker_now)

	private.Worker_Type = r.FormValue("worker_type")

	worker_salary_str  := r.FormValue("worker_salary")

	private.Worker_Salary , _ = strconv.ParseFloat( worker_salary_str,64 )

	private.Worker_Twh  = 50.0 * 52.14
	private.Worker_H_Pay  = private.Worker_Salary * 10000. / private.Worker_Twh

    if _, err = client.Put(ctx, key, &private ); err != nil {

		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}

///  show private inf. on web

	process.Private_showall1(w , r )

}
