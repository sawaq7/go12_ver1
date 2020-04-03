package deliver_copy

import (
//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "net/http"
//	    "fmt"
	    "github.com/sawaq7/go12_ver1/client/sgh/process"

	    "github.com/sawaq7/go12_ver1/client/sgh/type2"
	    "strconv"
//	    "time"

        "cloud.google.com/go/datastore"
        "context"
        "os"
                                                  )

func Deliver_copy(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "deliver_copy start \n" )  // 繝・ヰ繝・け

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
//    fmt.Fprintf( w, "deliver_copy  : id %v\n", id )  // 繝・ヰ繝・け

	delidw ,_ := strconv.Atoi(id)
	delid := int64(delidw)

//    fmt.Fprintf( w, "deliver_copy  : delidw %v\n", delidw )  // 繝・ヰ繝・け
//    fmt.Fprintf( w, "deliver_copy  : delid %v\n", delid )  // 繝・ヰ繝・け

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

/// 繝｢繝九ち繝ｼ縲蜀崎｡ｨ遉ｺ ///

	process.Deliver_showall1(w , r )

//	fmt.Fprintf( w, "deliver_copy : normal end \n" )  // 繝・ヰ繝・け

}
