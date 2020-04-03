package datastore_copy_excute

import (
//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "net/http"
//	    "fmt"
	    "strconv"
        "errors"
	    "github.com/sawaq7/go12_ver1/general/datastore5/copy3"
	    "github.com/sawaq7/go12_ver1/general/type5"

	    "os"

        "cloud.google.com/go/datastore"
        "context"

                                                  )

func Datastore_copy_excute(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "datastore_copy_excute start \n" )  // チE��チE��

///
///  エラーメチE��ージ　セチE��
///

var (

	  Err1 = errors.New("can't find this datastore's file (/datastore_copy_excute)")

	                                                                        )

var g type5.Ds_Copy_List /// チE�Eタストアのコピ�Eリスト�Eワークエリア確俁E
///                      ///
/// 持E��したデータidをGET ///
///                      ///

    copyidw , err := strconv.Atoi(r.FormValue("id"))
	if err  != nil {
//	   fmt.Fprintf( w, "datastore_copy_excute :error copyidw %v\n", copyidw )  // チE��チE��
	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
	}
    copyid := int64(copyidw)

///
///   プロジェクト名をゲチE��
///
    project_name := os.Getenv("GOOGLE_CLOUD_PROJECT")

    if project_name == "" {
//      fmt.Fprintf( w, "datastore_copy_excute :  projectID unset \n"  )  // チE��チE��

      project_name = "sample-7777"

	}

//	c := appengine.NewContext(r)
    ctx := context.Background()

    client, err := datastore.NewClient(ctx, project_name)
    if err != nil {
       http.Error(w,err.Error(), http.StatusInternalServerError)
		return
    }

///
///       コピ�Eリスト情報をGET
///

    key := datastore.IDKey("Ds_Copy_List", copyid, nil)
//    key := datastore.NewKey(c, "Ds_Copy_List", "", copyid, nil)

    if err := client.Get(ctx, key , &g ) ; err != nil {
//	if err := datastore.Get(c, key, &g); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

//    fmt.Fprintf( w, "datastore_copy_excut g.Basic_Name %v\n" ,g.Basic_Name)  // チE��チE��
//	fmt.Fprintf( w, "datastore_copy_excut g.New_Name %v\n" ,g.New_Name)  // チE��チE��

///
/// コピ�Eして、ニューファイルを作�E　///
///


    switch g.Basic_Name {

      case "Deliver" :                   // 配達惁E��　�E�Egh )

        copy3.Deliver( w , r  ,g.Basic_Name ,g.Copy_Name ,g.New_Name )

      case "D_District" :                // 地区惁E��　�E�Egh )

        copy3.D_district( w , r  ,g.Basic_Name ,g.Copy_Name ,g.New_Name )

      default :                          // 該当するファイルなぁE        http.Error(w, Err1.Error(), http.StatusInternalServerError)
        return

    }
//    copy3.All( w , r  ,g.Basic_Name ,g.Copy_Name ,g.New_Name )
//    if err := copy3.All( w , r  ,g.Basic_Name ,g.Copy_Name ,g.New_Name ); err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}

//	fmt.Fprintf( w, "datastore_copy_excute normal end \n" )  // チE��チE��
}
