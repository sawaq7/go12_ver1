package copy3

import (
	    "google.golang.org/appengine"
	    "google.golang.org/appengine/datastore"
	    "net/http"
	    "fmt"

//	    "github.com/sawaq7/go12_ver1/general/datastore5/check3"
        "github.com/sawaq7/go12_ver1/client/sgh/type2"

                                                )

///                          ///
/// 繝・・繧ｿ繧ｹ繝医い繧偵さ繝斐・縺吶ｋ  ///
///                         ///

// func All( w http.ResponseWriter, r *http.Request ,basic_name string ,copy_file string ,new_file string ) {
 func All( w http.ResponseWriter, r *http.Request ,basic_name string ,copy_file string ,new_file string )( err error) {

//     IN    w        : 繝ｬ繧ｹ繝昴Φ繧ｹ繝ｩ繧､繧ｿ繝ｼ
//     IN    r        : 繝ｪ繧ｯ繧ｨ繧ｹ繝医ヱ繝ｩ繝｡繝ｼ繧ｿ
//     IN  basic_name : 蝓ｺ譛ｬ縺ｮ繝・・繧ｿ繧ｹ繝医い蜷・//     IN  copy_file  : 繧ｳ繝斐・蜈・・繝・・繧ｿ繧ｹ繝医い蜷・//     IN  new_file   : 繝九Η繝ｼ繝・・繧ｿ繧ｹ繝医い蜷・//    OUT  err        : 繧ｨ繝ｩ繝ｼ繝｡繝・そ繝ｼ繧ｸ

    fmt.Fprintf( w, "copy3.all start \n" )  // 繝・ヰ繝・け
    fmt.Fprintf( w, "copy3.all basic_name %v\n" ,basic_name)  // 繝・ヰ繝・け

///
///  繧ｨ繝ｩ繝ｼ繝｡繝・そ繝ｼ繧ｸ縲繧ｻ繝・ヨ
///

    c := appengine.NewContext(r)       ///縲繧ｳ繝ｳ繝・く繧ｹ繝医ｒ蜿門ｾ励☆繧・
	q := datastore.NewQuery(copy_file) /// 繧ｯ繝ｭ繝ｼ繝ｳ縺ｮ繝ｪ繝ｼ繝繝ｼ繧竪ET

	count, err := q.Count(c)           ///縲繝ｬ繧ｳ繝ｼ繝画焚繧竪ET
	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
		return  err

	}
    fmt.Fprintf( w, "copy3.all count %v\n" ,count)  // 繝・ヰ繝・け


//   if err :=check3.Name( w , basic_name   ) ;  err != nil{
//       http.Error(w, err.Error(), http.StatusInternalServerError)

//	  return err

//	}

///
/// 繧ｯ繝ｭ繝ｼ繝ｳ諠・ｱ繧担ET縺吶ｋ繝ｯ繝ｼ繧ｯ繧ｨ繝ｪ繧｢繧堤｢ｺ菫・///

    ds_data := make([]type2.Deliver, 0, count)



	if _, err := q.GetAll(c, &ds_data);  err != nil {         // 繧ｯ繝ｭ繝ｼ繝ｳ諠・ｱ繧竪ET

//	  http.Error(w, err.Error(), http.StatusInternalServerError)
	  return  err

	} else{
      for _, ds_dataw := range ds_data {                       //縲繝九Η繝ｼ繝輔ぃ繧､繝ｫ縺ｫ繧ｯ繝ｭ繝ｼ繝ｳ諠・ｱ繧偵そ繝・ヨ

	    if _, err := datastore.Put(c, datastore.NewIncompleteKey(c, new_file, nil), &ds_dataw); err != nil {

//		  http.Error(w,err.Error(), http.StatusInternalServerError)
		  return  err

	    }

	  }
	}

	fmt.Fprintf( w, "copy3.all normal end \n" )  // 繝・ヰ繝・け

    return nil
}

