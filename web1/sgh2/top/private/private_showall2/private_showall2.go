package private_showall2

import (

	    "net/http"
//	    "fmt"
	    "github.com/sawaq7/go12_ver1/client/sgh/process"
	    "github.com/sawaq7/go12_ver1/client/sgh/type2"
	    "strconv"

        "cloud.google.com/go/datastore"
	    "context"
	    "os"

                                                  )

///
/// 　　   show private inf. on web
///

func Private_showall2(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "private_showall2 start \n" )

	var private type2.Private

	private.Worker_Name = r.FormValue("worker_name")
//	fmt.Fprintf( w, "private_showall2 : worker_name %v\n", private.Worker_Name )

	worker_no := r.FormValue("worker_no")
//	fmt.Fprintf( w, "private_showall2 : worker_no %v\n", worker_no )

	worker_now ,err := strconv.Atoi(worker_no)
	if err != nil {

//       fmt.Fprintf( w, "private_showall2 : a number must be half-width characters %v\n"  )
		return
	}

	private.Worker_No = int64(worker_now)                //    make an integer

	private.Worker_Type = r.FormValue("worker_type")

	worker_salary_str  := r.FormValue("worker_salary")

	private.Worker_Salary , _ = strconv.ParseFloat( worker_salary_str,64 )  // make a integer

	private.Worker_Twh  = 50.0 * 52.14                 //    calculate annual total working hours

	private.Worker_H_Pay  = private.Worker_Salary * 10000. / private.Worker_Twh  // calculate payment by hour　

    project_name := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if project_name == "" {

      project_name = "sample-7777"

	}

    ctx := context.Background()

    client, err := datastore.NewClient(ctx, project_name)
    if err != nil {
       http.Error(w,err.Error(), http.StatusInternalServerError)
		return
    }



///   put new private inf. in d.s.

    new_key := datastore.IncompleteKey("Private", nil)

    if _, err = client.Put(ctx, new_key, &private ); err != nil {

		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}

///  show private inf. on web

	process.Private_showall1(w , r )

//	fmt.Fprintf( w, "private_showall2 : normal end \n" )

}