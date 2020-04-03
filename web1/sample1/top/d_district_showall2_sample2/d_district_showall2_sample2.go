package d_district_showall2_sample2

import (

//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"

	    "net/http"
//	    "fmt"
        "html/template"
        "github.com/sawaq7/go12_ver1/temp/type1000"
	    "github.com/sawaq7/go12_ver1/temp/html1000"

//	    "strconv"

        "cloud.google.com/go/datastore"
	    "context"
	    "os"
                                                  )

func D_district_showall2_sample2(w http.ResponseWriter, r *http.Request) {

//   fmt.Fprintf( w, "d_district_showall2_sample2 start \n" )  // 繝・ヰ繝・け

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
	d_district_view := make([]type1000.D_District, 0)

    keys, err := client.GetAll(ctx, query , &d_district_sample)
//	keys, err := q.GetAll(c, &d_district_sample)
    if err != nil {
       http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	keys_wk := make([]int64, count)

	for ii, keysw := range keys {

       keys_wk[ii] = keysw.ID

    }

	for pos, d_district_samplew := range d_district_sample {

         d_district_view = append(d_district_view, type1000.D_District { keys_wk[pos]                        ,
                                                                         d_district_samplew.District_No      ,
                                                                         d_district_samplew.District_Name    ,
                                                                         d_district_samplew.D_Area_Slice     ,
                                                                         d_district_samplew.D_Area_Small_Slice })

	}

// 繝・Φ繝励Ξ繝ｼ繝医・繝倥ャ繝繝ｼ繧竪ET

   monitor := template.Must(template.New("html").Parse(html1000.D_district_showall1_sample))

// 繝｢繝九ち繝ｼ縺ｫ陦ｨ遉ｺ

	err = monitor.Execute(w, d_district_view)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

//	fmt.Fprintf( w, "d_district_showall2_sample2 normal end \n" )  // 繝・ヰ繝・け

}
