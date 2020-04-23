package d_district_showall2_sample3

import (


	"net/http"
//	"strconv"
//	"fmt"
//	"github.com/sawaq7/go12_ver1/client/sgh/process"
	"github.com/sawaq7/go12_ver1/temp/type1000"

	"cloud.google.com/go/datastore"
	"context"
	"os"
                                            )

func D_district_showall2_sample3(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "d_district_showall2_sample3 start \n" )

    projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if projectID == "" {

      projectID = "sample-7777"

	}

    ctx := context.Background()

    client, err := datastore.NewClient(ctx, projectID)
    if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

    query := datastore.NewQuery("D_District_Sample").Order("District_No")

    count, err := client.Count(ctx, query)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	d_district_sample      := make([]type1000.D_District, 0, count)

    keys, err := client.GetAll(ctx, query , &d_district_sample)

    if err != nil {
       http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

    for _, keysw := range keys {

      if err := client.Delete(ctx, datastore.IDKey("D_District_Sample", keysw.ID, nil)); err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	  }

	}

}
