package storage_object_copy_excute

import (
	    "github.com/sawaq7/go12_ver1/storage2"
//	    "fmt"
	    "net/http"
//	    "google.golang.org/appengine"
//	    "google.golang.org/appengine/datastore"
	    "github.com/sawaq7/go12_ver1/general/type5"
	    "github.com/sawaq7/go12_ver1/general/process3"
        "os"
	    "log"
        "cloud.google.com/go/datastore"
        "context"

                                                  )

func Storage_object_copy_excute(w http.ResponseWriter, r *http.Request) {

//    fmt.Fprintf( w, "storage_object_copy_excute start \n" )  // チE��チE��

    var project_name ,bucket_name ,basic_file_name string

    var db_access_list2 type5.Db_Access_List2      //D.B. アクセスリスト�E用のワークエリアを確俁E
///
/// 入力データをGET 　
///

//    basic_file_name := r.FormValue("basic_file_name")  // ベ�EシチE��ファイル名をゲチE��

    new_file_name := r.FormValue("new_file_name")  // ニューファイル名をゲチE��

///
///   バケチE��名をゲチE��
///
    project_name = os.Getenv("GOOGLE_CLOUD_PROJECT")

    if project_name == "" {
//      fmt.Fprintf( w, "storage_bucket_list :  projectID unset \n"  )  // チE��チE��

      project_name = "sample-7777"

	}
    ctx := context.Background()
//    c := appengine.NewContext(r)

	query := datastore.NewQuery("Storage_B_O_Temp")

    client, err := datastore.NewClient(ctx, project_name)
    if err != nil {
       log.Fatalf("Failed to create client: %v", err)
    }

//	count, err := q.Count(c)
    count, err := client.Count(ctx, query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	storage_b_o_temp      := make([]type5.Storage_B_O_Temp, 0, count)

    if _, err = client.GetAll(ctx, query , &storage_b_o_temp) ; err != nil {
//	if _, err := q.GetAll(c, &storage_b_o_temp);  err != nil {

	  http.Error(w, err.Error(), http.StatusInternalServerError)
	  return

	} else{

      for pos, storage_b_o_tempw := range storage_b_o_temp {

        if pos == 0 {

           project_name       = storage_b_o_tempw.Project_Name    // プロジェクト名をゲチE��
           bucket_name        = storage_b_o_tempw.Bucket_Name    // バケチE��名をゲチE��
           basic_file_name    = storage_b_o_tempw.Object_Name    // ベ�EシチE��ファイル名をゲチE��

        }
	  }
	}

///
/// ストレチE��ファイルをコピ�Eする　
///

//    general_work := make([]type5.General_Work, 2 )    // ワークエリア確俁E
//    general_work[0].String_Work = basic_file_name     //　ベ�EシチE��ファイルネ�EムセチE��
//    general_work[1].String_Work = new_file_name       //　ニューファイルネ�EムセチE��

//    _ ,_ = storage2.Storage_basic( "copy" ,bucket_name ,general_work , w , r  )

//    fmt.Fprintf(w, "storage_object_copy_excute: bucket_name %v\n", bucket_name )  // チE��チE��
//    fmt.Fprintf(w, "storage_object_copy_excute: basic_file_name %v\n", basic_file_name )  // チE��チE��
//    fmt.Fprintf(w, "storage_object_copy_excute: new_file_name %v\n", new_file_name )  // チE��チE��

    storage2.File_Copy ( w , r  ,bucket_name ,basic_file_name ,new_file_name  )


///
/// アクセスリストに登録　
///

///  吁E��アクセス惁E��をセチE��

    db_access_list2.Db_Type = "sr"
    db_access_list2.Access_Type = "copy"
    db_access_list2.Project_Name = project_name
    db_access_list2.Bucket_Name = bucket_name
    db_access_list2.Basic_File_Name = basic_file_name
    db_access_list2.New_File_Name = new_file_name

/// チE�Eタストアに、Eレコードを追加

    new_key := datastore.IncompleteKey("Db_Access_List2", nil)

//    fmt.Fprintf(w, "storage_object_copy_excute: new_key %v\n", new_key )  // チE��チE��

    _, err = client.Put(ctx, new_key, &db_access_list2 )

//    fmt.Fprintf(w, "storage_object_copy_excute: key2 %v\n", key2 )  // チE��チE��

    if err != nil {
//   if _, err := datastore.Put(c, datastore.NewIncompleteKey(c, "Db_Access_List2", nil) , &db_access_list2); err != nil {
      http.Error(w,err.Error(), http.StatusInternalServerError)
      return
    }

//    key := datastore.IDKey( "Storage_B_O_Temp",int64(1) , nil)

//      if err := client.Delete(ctx, key); err != nil {
//	  if err := datastore.Delete(c, key  ); err != nil {

//		 http.Error(w, err.Error(), http.StatusInternalServerError)
//		 return
//	  }




///
/// オブジェクトリストを表示する　
///
    process3.Storage_object_show ( w , r ,project_name  ,bucket_name )


//	fmt.Fprintf( w, "storage_object_copy_excute normal end \n" )  // チE��チE��
}
