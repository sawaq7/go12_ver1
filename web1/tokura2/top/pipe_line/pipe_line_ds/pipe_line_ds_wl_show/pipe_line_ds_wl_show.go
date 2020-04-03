package pipe_line_ds_wl_show

import (

//        "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "net/http"
	    "github.com/sawaq7/go12_ver1/client/tokura/suiri/process2"
	    "github.com/sawaq7/go12_ver1/client/tokura/suiri/type4"
	    "strconv"
//	    "fmt"

	    "cloud.google.com/go/datastore"
        "context"
        "os"

                                                  )

func Pipe_line_ds_wl_show(w http.ResponseWriter, r *http.Request) {

//   fmt.Fprintf( w, "sky/pipe_line_ds_wl_show start \n"  )  // γEγE―

/// key-in γEEγΏγGET ///

   var g type4.Water_Line

   project_name := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if project_name == "" {

      project_name = "sample-7777"

	}

    ctx := context.Background()
//    c := appengine.NewContext(r)

    client, err := datastore.NewClient(ctx, project_name)
    if err != nil {
       http.Error(w, err.Error(), http.StatusInternalServerError)
       return
    }

/// γE³γγ©γͺγΌγγ‘γ€γ«γγγζ°΄θ·―εγγ²γE

    query := datastore.NewQuery("Water2_Temp").Order("Name")
//    q2 := datastore.NewQuery("Water2_Temp").Order("Name")

//    count, err := q2.Count(c)
    count, err := client.Count(ctx, query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

//    fmt.Fprintf( w, "sky/pipe_line_ds_wl_show  \n" ,count )  // γEγE―

	water2_temp     := make([]type4.Water2_Temp, 0, count)


//	keys, err := q2.GetAll(c, &water2_temp )
	keys, err := client.GetAll(ctx, query , &water2_temp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
    }

    keys_wk := make([]int64, count)

	for ii, keysw := range keys {

       keys_wk[ii] = keysw.ID

    }

//    fmt.Fprintf( w, "sky/pipe_line_ds_wl_show : len(water2_temp) %v\n", len(water2_temp) )  // γEγE―

	for pos2, water2_tempw := range water2_temp {

       g.Name = water2_tempw.Name  /// ζ°΄θ·―εγEγ»γE
//       g.Id   = keys[pos2].IntID()
       g.Id   = keys_wk[pos2]

    }

//	g.Name = r.FormValue("water_name")  // ζ°΄θ·―εγγ²γE

	g.Section = r.FormValue("section")  // εΊιεγγ²γE

	f_facter := r.FormValue("f_facter")                   // ζ©ζ¦δΏζ°γγ²γE
	g.Friction_Factor,_ =strconv.ParseFloat(f_facter,64)  //γfloat64γγ«ε€ζ

	velocity := r.FormValue("velocity")                   // ιεΊ¦γγ²γE
	g.Velocity,_ =strconv.ParseFloat(velocity,64)         //γfloat64γγ«ε€ζ

	p_diameter := r.FormValue("p_diameter")      // ζ©ζ¦δΏζ°γγ²γE
	g.Pipe_Diameter,_ =strconv.ParseFloat(p_diameter,64)  //γfloat64γγ«ε€ζ

	p_length := r.FormValue("p_length")      // ζ©ζ¦δΏζ°γγ²γE
	g.Pipe_Length,_ =strconv.ParseFloat(p_length,64)  //γfloat64γγ«ε€ζ

//	fmt.Fprintf( w, "sky/pipe_line_ds_wl_show : g.Name %v\n", g.Name )  // γEγE―
//	fmt.Fprintf( w, "sky/pipe_line_ds_wl_show : g.Section %v\n", g.Section )  // γEγE―
//	fmt.Fprintf( w, "sky/pipe_line_ds_wl_show : g.Friction_Factor %v\n", g.Friction_Factor )  // γEγE―
//	fmt.Fprintf( w, "sky/pipe_line_ds_wl_show : g.Velocity %v\n", g.Velocity )  // γEγE―
//	fmt.Fprintf( w, "sky/pipe_line_ds_wl_show : g.Pipe_Diameter %v\n", g.Pipe_Diameter )  // γEγE―
//	fmt.Fprintf( w, "sky/pipe_line_ds_wl_show : g.Pipe_Length %v\n", g.Pipe_Length )  // γEγE―

/// γEEγΏγΉγγ’γΌγ«γEEγΏγγ»γE ///

    new_key := datastore.IncompleteKey("Water_Line", nil)

    if _, err = client.Put(ctx, new_key, &g ); err != nil {
//	if _, err := datastore.Put(c, datastore.NewIncompleteKey(c, "Water_Line", nil), &g); err != nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}

/// γ’γγΏγΌθ‘¨η€Ί ///

   process2.Pipe_line_ds_wl_show(1 ,g.Name ,w , r )


}

