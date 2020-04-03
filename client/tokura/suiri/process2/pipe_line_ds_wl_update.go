package process2

import (
//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "net/http"
	    "strconv"
//	    "fmt"
	    "github.com/sawaq7/go12_ver1/client/tokura/suiri/type4"
//	    "time"

        "os"
        "cloud.google.com/go/datastore"
        "context"
                                                )


func Pipe_line_ds_wl_update(w http.ResponseWriter, r *http.Request ,updid int64) {

//    fmt.Fprintf( w, "pipe_line_ds_wl_update start \n" )  // γEγE―
//    fmt.Fprintf( w, "pipe_line_ds_wl_update : updid %v\n", updid )  // γEγE―


    var g type4.Water_Line

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

    key := datastore.IDKey("Water_Line", updid, nil)

    if err := client.Get(ctx, key , &g ) ; err != nil {
//	key := datastore.NewKey(c, "Water_Line", "", updid, nil)
//	if err := datastore.Get(c, key, &g); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	g.Name = r.FormValue("water_name")  // ζ°΄θ·―εγγ²γE
	g.Section = r.FormValue("section")  // εΊιεγγ²γE

	f_facter := r.FormValue("f_facter")                   // ζ©ζ¦δΏζ°γγ²γE
	g.Friction_Factor,_ =strconv.ParseFloat(f_facter,64)  //γfloat64γγ«ε€ζ

	velocity := r.FormValue("velocity")                   // ιεΊ¦γγ²γE
	g.Velocity,_ =strconv.ParseFloat(velocity,64)         //γfloat64γγ«ε€ζ

	p_diameter := r.FormValue("p_diameter")      // ζ©ζ¦δΏζ°γγ²γE
	g.Pipe_Diameter,_ =strconv.ParseFloat(p_diameter,64)  //γfloat64γγ«ε€ζ

	p_length := r.FormValue("p_length")      // ζ©ζ¦δΏζ°γγ²γE
	g.Pipe_Length,_ =strconv.ParseFloat(p_length,64)  //γfloat64γγ«ε€ζ

//	fmt.Fprintf( w, "pipe_line_ds_wl_update : g.Name %v\n", g.Name )  // γEγE―
//	fmt.Fprintf( w, "pipe_line_ds_wl_update : g.Section %v\n", g.Section )  // γEγE―
//	fmt.Fprintf( w, "pipe_line_ds_wl_update : g.Friction_Factor %v\n", g.Friction_Factor )  // γEγE―
//	fmt.Fprintf( w, "pipe_line_ds_wl_update : g.Velocity %v\n", g.Velocity )  // γEγE―
//	fmt.Fprintf( w, "pipe_line_ds_wl_update : g.Pipe_Diameter %v\n", g.Pipe_Diameter )  // γEγE―
//	fmt.Fprintf( w, "pipe_line_ds_wl_update : g.Pipe_Length %v\n", g.Pipe_Length )  // γEγE―

    if _, err := client.Put(ctx, key, &g ); err != nil {
//	if _, err := datastore.Put(c, key, &g); err != nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}

}
