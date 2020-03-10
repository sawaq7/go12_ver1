package d_district_showall2_sample3

import (

//	"google.golang.org/appengine"
//	"google.golang.org/appengine/datastore"
	"net/http"
//	"strconv"
//	"fmt"
//	"client/sgh/process"
	"temp/type1000"

	"cloud.google.com/go/datastore"
	"context"
	"os"
                                            )

func D_district_showall2_sample3(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "d_district_showall2_sample3 start \n" )  // デバック

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

    query := datastore.NewQuery("D_District_Sample").Order("District_No")
//	q := datastore.NewQuery("D_District_Sample").Order("District_No")

    count, err := client.Count(ctx, query)
//	count, err := q.Count(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	d_district_sample      := make([]type1000.D_District, 0, count)

    keys, err := client.GetAll(ctx, query , &d_district_sample)
//	keys, err := q.GetAll(c, &d_district_sample)
    if err != nil {
       http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

//	for pos, _ := range d_district_sample {
    for _, keysw := range keys {

      if err := client.Delete(ctx, datastore.IDKey("D_District_Sample", keysw.ID, nil)); err != nil {
//	  key := datastore.NewKey(c, "D_District_Sample", "", keys[pos].IntID(), nil)
//	  if err := datastore.Delete(c, key); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	  }

	}

}
