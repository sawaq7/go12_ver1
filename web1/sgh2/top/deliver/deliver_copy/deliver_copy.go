package deliver_copy

import (
//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "net/http"
//	    "fmt"
	    "client/sgh/process"

	    "client/sgh/type2"
	    "strconv"
//	    "time"

        "cloud.google.com/go/datastore"
        "context"
        "os"
                                                  )

func Deliver_copy(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "deliver_copy start \n" )  // デバック

	var g type2.Deliver

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

	id := r.FormValue("id")
//    fmt.Fprintf( w, "deliver_copy  : id %v\n", id )  // デバック

	delidw ,_ := strconv.Atoi(id)
	delid := int64(delidw)

//    fmt.Fprintf( w, "deliver_copy  : delidw %v\n", delidw )  // デバック
//    fmt.Fprintf( w, "deliver_copy  : delid %v\n", delid )  // デバック

    key := datastore.IDKey("Deliver", delid, nil)

    if err := client.Get(ctx, key , &g ) ; err != nil {

//	key := datastore.NewKey(c, "Deliver", "", delid, nil)
//	if  err := datastore.Get(c, key,  &g); err != nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}

    new_key := datastore.IncompleteKey("Deliver", nil)

    if _, err = client.Put(ctx, new_key, &g ); err != nil {
//    if _, err := datastore.Put(c, datastore.NewIncompleteKey(c, "Deliver", nil), &g); err != nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}

/// モニター　再表示 ///

	process.Deliver_showall1(w , r )

//	fmt.Fprintf( w, "deliver_copy : normal end \n" )  // デバック

}
