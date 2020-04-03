package medical_xray_upload

import (

//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"

	    "net/http"
	    "fmt"
	    "io"
        "github.com/sawaq7/go12_ver1/client/reserve/process4"
	    "storage2"

	    "github.com/sawaq7/go12_ver1/client/reserve/type6"
	    "github.com/sawaq7/go12_ver1/client/reserve/datastore6/trans5"
	    "time"

	    "cloud.google.com/go/datastore"
	    "context"
	    "os"

                                                  )

func Medical_xray_upload(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "medical_xray_upload start \n" )  // 繝・ヰ繝・け

    var bucket  string

    var guest_medical_xray type6.Guest_Medical_Xray

    bucket    = "sample-7777"    // 繝舌こ繝・ヨ蜷阪ｒ繧ｲ繝・ヨ

///
///     謖・ｮ壹＠縺溘ヵ繧｡繧､繝ｫ諠・ｱ繧偵ご繝・ヨ
///

	file_data, fh, err := r.FormFile("image")

	if err != nil {
		return

	}

//	fmt.Fprintf( w, "medical_xray_upload : fh %v\n", fh )  // 繝・ヰ繝・け


//	st_writer_minor , _ := storage2.Storage_basic( "create" ,bucket ,fh.Filename , w , r  )

//    st_writer, _ := st_writer_minor.(*storage.Writer)  // 繧､繝ｳ繧ｿ繝ｼ繝輔ぉ繧､繧ｹ蝙九ｒ蝙句､画鋤

    f_name   :=  fh.Filename

//    st_writer := storage2.File_Create ( w ,r ,bucket  ,f_name )

//    content_type := "image/png; charset=utf-8"

    content_type := fh.Header.Get("Content-Type")

    st_writer := storage2.File_Create2 ( w ,r ,bucket  ,f_name ,content_type )

///
/// 繧ｹ繝医Ξ繝・ず繝輔ぃ繧､繝ｫ縺ｫ譌｢蟄倥・繝輔ぃ繧､繝ｫ縺ｮ諠・ｱ繧偵さ繝斐・
///

	if _, err := io.Copy(st_writer, file_data); err != nil {
	    http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := st_writer.Close(); err != nil {
	   http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

///
///     繧ｫ繝ｬ繝ｳ繝医・繧ｲ繧ｹ繝亥錐繧偵ご繝・ヨ
///

    guest_temp_slice := trans5.Guest_temp (  w , r  )

//    guest_name := guest_temp_slice[0].Guest_Name
//        _ = guest_temp_slice[0].Guest_Name

    date_w := time.Now()        // 譌･莉倥ｒ繧ｻ繝・ヨ

    guest_medical_xray.Date   = fmt.Sprintf("%04d/%02d/%02d/%02d/%02d/%02d",date_w.Year(), date_w.Month(),date_w.Day(), date_w.Hour(), date_w.Minute(), date_w.Second())

    guest_medical_xray.Guest_No   = guest_temp_slice[0].Guest_No

    guest_medical_xray.Guest_Name   = guest_temp_slice[0].Guest_Name

    guest_medical_xray.File_Name   = f_name

	const publicURL = "https://storage.googleapis.com/%s/%s"
	guest_medical_xray.Url = fmt.Sprintf(publicURL, bucket , f_name )

///
///         繝・・繧ｿ繧ｹ繝医い繝ｼ縺ｫ繝・・繧ｿ繧偵そ繝・ヨ
///

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

    new_key := datastore.IncompleteKey("Guest_Medical_Xray", nil)

    if _, err = client.Put(ctx, new_key, &guest_medical_xray ); err != nil {
//	if _, err := datastore.Put(c, datastore.NewIncompleteKey(c, "Guest_Medical_Xray", nil), &guest_medical_xray); err != nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}

///
///    繝｢繝九ち繝ｼ 陦ｨ遉ｺ
///

    process4.Medical_xray_show(w , r ,guest_medical_xray.Guest_No)

//	fmt.Fprintf( w, "medical_xray_upload : normal end \n" )  // 繝・ヰ繝・け

}


