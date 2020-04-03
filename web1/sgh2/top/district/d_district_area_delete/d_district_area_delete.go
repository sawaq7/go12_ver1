package d_district_area_delete

import (

//	"google.golang.org/appengine"
//	"google.golang.org/appengine/datastore"
	"net/http"
	"strconv"
//	"fmt"
	"github.com/sawaq7/go12_ver1/client/sgh/process"
	"github.com/sawaq7/go12_ver1/client/sgh/type2"

	"cloud.google.com/go/datastore"
    "context"
    "os"
                                            )

func D_district_area_delete(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "d_district_area_delete start \n" )  // γEγE―
    var g type2.D_Area

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

   id := r.FormValue("id")

//    fmt.Fprintf( w, "d_district_area_delete : id %v\n", id )  // γEγE―

	delidw ,_ := strconv.Atoi(id)
	delid := int64(delidw)

//    fmt.Fprintf( w, "d_district_area_delete : delidw %v\n", delidw )  // γEγE―
//    fmt.Fprintf( w, "d_district_area_delete : delid %v\n", delid )  // γEγE―

    key := datastore.IDKey("D_Area", delid, nil)
//	key := datastore.NewKey(c, "D_Area", "", delid, nil)

    if err := client.Get(ctx, key , &g ) ; err != nil {
// ε°εΊNOγγ²γE
//    if err := datastore.Get(c, key, &g); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	district_no := g.District_No

// γEEγΏγΉγγ’γγ1γ¬γ³γΌγει€
    if err := client.Delete(ctx, key); err != nil {
//	if err := datastore.Delete(c, key); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}


/// γ’γγΏγΌγεθ‘¨η€Ί ///

	process.D_district_area_show(w , r  ,district_no)


//	http.Redirect(w, r, "/", http.StatusFound)
}
