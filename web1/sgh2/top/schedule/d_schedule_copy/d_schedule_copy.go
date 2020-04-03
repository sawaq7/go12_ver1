package d_schedule_copy

import (
//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "net/http"
//	    "fmt"
	    "github.com/sawaq7/go12_ver1/client/sgh/process"
        "github.com/sawaq7/go12_ver1/client/sgh/datastore2"
	    "github.com/sawaq7/go12_ver1/client/sgh/type2"
	    "github.com/sawaq7/go12_ver1/general/type5"
	    "strconv"
//	    "time"

        "cloud.google.com/go/datastore"
        "context"
        "os"
                                                  )

func D_schedule_copy(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "d_schedule_copy start \n" )  // γEγE―

    var idmy int64

	var g type2.D_Schedule

	id := r.FormValue("id")
//    fmt.Fprintf( w, "d_schedule_copy  : id %v\n", id )  // γEγE―

	copyidw ,_ := strconv.Atoi(id)
	copyid := int64(copyidw)

//    fmt.Fprintf( w, "d_schedule_copy  : copyidw %v\n", copyidw )  // γEγE―
//    fmt.Fprintf( w, "d_schedule_copy  : copyid %v\n", copyid )  // γEγE―

    project_name := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if project_name == "" {

      project_name = "sample-7777"

	}

    ctx := context.Background()
//	c := appengine.NewContext(r)

    client, err := datastore.NewClient(ctx, project_name)
    if err != nil {
       http.Error(w, err.Error(), http.StatusInternalServerError)
       return
    }

    key := datastore.IDKey("D_Schedule", copyid, nil)

    if err := client.Get(ctx, key , &g ) ; err != nil {
//	key := datastore.NewKey(c, "D_Schedule", "", copyid, nil)
//	if  err := datastore.Get(c, key,  &g); err != nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}

    new_key := datastore.IncompleteKey("D_Schedule", nil)

    if _, err = client.Put(ctx, new_key, &g ); err != nil {
//    if _, err := datastore.Put(c, datastore.NewIncompleteKey(c, "D_Schedule", nil), &g); err != nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}
//   	fmt.Fprintf( w, "d_schedule_copy : g.Course_332 %v\n", g.Course_332 )  // γEγE―

//  ε°εΊζE ±γGET

    flexible_out := datastore2.Datastore_sgh( "D_District_Temp" ,"check" ,idmy , w , r  )

    value2, _ := flexible_out.([]type5.General_Work)

    district_no := value2[0].Int64_Work

/// γ’γγΏγΌγεθ‘¨η€Ί

	process.D_schedule_showall(w , r ,district_no)

//	fmt.Fprintf( w, "d_schedule_copy : normal end \n" )  // γEγE―




}
