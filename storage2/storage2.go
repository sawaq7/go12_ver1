///////////////////////////////////////////////////////////////////////////////////////////
///                                                                                     ///
///   any routine which access file for App Engine app using the Google Cloud Storage  ///
///                                                                                     ///
///////////////////////////////////////////////////////////////////////////////////////////
package storage2


import (

	"fmt"
	"io"
	"net/http"

	"strings"
    "cloud.google.com/go/storage"

    "context"
	"google.golang.org/api/iterator"
	"general/type5"

                                        )


///                                                                      ///
///   Bucket_Handler_Get : get bucket handler for Google Cloud Storage.  ///
///                                              type-2                  ///

func Bucket_Handler_Get(w http.ResponseWriter ,r *http.Request ,bucket string) (*storage.BucketHandle) {

//     IN    w      : レスポンスライター
//     IN    r      : リクエストパラメータ
//     IN  bucket   : バケット名
//     OUT  one     : バケットハンドラー

//    fmt.Fprintf( w, "Bucket_Handler_Get start \n" )  // デバック

    ctx := context.Background()

	client, err := storage.NewClient(ctx)
	if err != nil {
	   http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil
	}
	return client.Bucket(bucket)
}

///                                                                 ///
///    File_Open : open file in Google Cloud Storage.             ///
///                                                                 ///

func File_Open(w http.ResponseWriter ,r *http.Request ,bucket string ,filename string)(io.ReadCloser) {

//     IN    w      : レスポンスライター
//     IN    r      : リクエストパラメータ
//     IN  bucket     : バケット名
//     IN  filename   : ファイル名
//     OUT  one       : ストレッジ用リーダー

//    fmt.Fprintf( w, "File_Open start \n" )  // デバック

/// get bucket handler for storage

    bh := Bucket_Handler_Get( w ,r ,bucket )

    ctx := context.Background()

    rc, err := bh.Object(filename).NewReader(ctx)
    if err != nil {
	   http.Error(w, err.Error(), http.StatusInternalServerError)
//		return nil
	}

    return rc
}

///                                                               ///
///   File_Create : create a file in Google Cloud Storage.  ///
///                        type-2                                 ///

func File_Create ( w http.ResponseWriter ,r *http.Request ,bucket string ,filename string ) ( *storage.Writer ){

//     IN    w      : レスポンスライター
//     IN    r      : リクエストパラメータ
//     IN  bucket     : バケット名
//     IN  filename   : ファイル名
//     OUT  one       : ストレッジ用ライター

//    fmt.Fprintf( w, "File_Create start \n" )  // デバック

/// get bucket handler for storage

    bh := Bucket_Handler_Get( w ,r ,bucket )

    ctx := context.Background()

    wc := bh.Object(filename).NewWriter(ctx)

	wc.ContentType = "text/plain; charset=utf-8"
	wc.ACL = []storage.ACLRule{{Entity: storage.AllUsers, Role: storage.RoleReader}}
//	wc.ContentType = fh.Header.Get("Content-Type")
	wc.CacheControl = "public, max-age=86400"  // Entries are immutable, be aggressive about caching (1 day).

//	fmt.Fprintf(gcs_gae.W, "StorageCreate: ファイルアドレス　%d\n", wc  )        // デバック

	return wc
}

///                                                               ///
///   File_Create : create a file in Google Cloud Storage.  ///
///                        type-2                                 ///

func File_Create2 ( w http.ResponseWriter ,r *http.Request ,bucket string ,filename string ,content_type string  ) ( *storage.Writer ){

//     IN    w      : レスポンスライター
//     IN    r      : リクエストパラメータ
//     IN  bucket     : バケット名
//     IN  filename   : ファイル名
//     IN  content_type   : コンテキストタイプ
//     OUT  one       : ストレッジ用ライター

//    fmt.Fprintf( w, "File_Create2 start \n" )  // デバック

/// get bucket handler for storage

    bh := Bucket_Handler_Get( w ,r ,bucket )

    ctx := context.Background()

    wc := bh.Object(filename).NewWriter(ctx)

	wc.ContentType = content_type
	wc.ACL = []storage.ACLRule{{Entity: storage.AllUsers, Role: storage.RoleReader}}
//	wc.ContentType = fh.Header.Get("Content-Type")
	wc.CacheControl = "public, max-age=86400"  // Entries are immutable, be aggressive about caching (1 day).

//	fmt.Fprintf(gcs_gae.W, "StorageCreate: ファイルアドレス　%d\n", wc  )        // デバック

	return wc
}

///                                                           ///
///   StorageCopy : copy a file in Google Cloud Storage.      ///
///                                                           ///

func File_Copy ( w http.ResponseWriter , r *http.Request ,bucket string ,fileName string ,fileName2 string){

//     IN    w      : レスポンスライター
//     IN    r      : リクエストパラメータ
//     IN  bucket   : バケット名
//     IN  filename    : コピー元ファイル名
//     IN  filename2   : コピー先ファイル名

//    fmt.Fprintf( w, "File_Copy start \n" )  // デバック

	writer := File_Create( w ,r ,bucket ,fileName2 ) // コピー先ファイルの枠を作成
    defer writer.Close()



    reader := File_Open(w ,r ,bucket ,fileName)  // コピー元ファイルをオープン
    defer reader.Close()

    // ファイルリーダー　をＧＥＴ

	if _, err := io.Copy(writer, reader); err != nil {
	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return

	}

// end process

//	fmt.Fprintf( w, "\n StorageCopy : Calculate succeeded.\n" )

}

///                                                           ///
///   StorageRename : rename a file in Google Cloud Storage.      ///
///                                                           ///

func File_Rename ( w http.ResponseWriter ,r *http.Request ,bucket string ,fileName1 string ,fileName2 string ) {

//     IN    w      : レスポンスライター
//     IN    r      : リクエストパラメータ
//     IN  bucket   : バケット名
//     IN  filename1   : オールドファイル名
//     IN  filename2   : ニューファイル名

//    fmt.Fprintf( w, "File_Rename start \n" )  // デバック

	File_Copy ( w ,r ,bucket , fileName1 ,fileName2  )    // ファイルをコピー

    File_Delete ( w , r ,bucket ,fileName1  )             // 元ファイルを削除

}

///                                                               ///
///   StorageDelete : delete a file in Google Cloud Storage.      ///
///                                                               ///

func File_Delete ( w http.ResponseWriter , r *http.Request ,bucket string ,fileName string  ) {

//     IN    w      : レスポンスライター
//     IN    r      : リクエストパラメータ
//     IN  bucket   : バケット名
//     IN  filename    : 削除するファイル名

//    fmt.Fprintf( w, "File_Delete start \n" )  // デバック

/// get bucket handler for storage

    bh := Bucket_Handler_Get( w ,r ,bucket )

    ctx := context.Background()

	err := bh.Object(fileName).Delete(ctx)
	if err != nil {
	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
	}

}


///                                                                 ///
///    File_Write : write data to file in Google Cloud Storage.   ///
///                                                                 ///

// func File_Write ( w http.ResponseWriter ,bucket string ,filename string ,wc *storage.Writer ,ldata []string ) {
func File_write ( w http.ResponseWriter ,wc *storage.Writer ,ldata []string ) {

//     IN    w      : レスポンスライター
//     IN    wc       : ストレッジ用ライター
//     IN  ldata      : 1行のデータ

//    fmt.Fprintf( w, "File_Write start \n" )  // デバック

	count := 0 //　デバック

	for  i := 0 ; i < len(ldata) ; i++ {

	    count ++  //　デバック

// 一行、書き込み

        fmt.Fprintf(wc ,"%s " ,ldata[i] )

   }

// 改行する

   fmt.Fprintf(wc ,"\n" )

}

///
///    File_Write_Line : write data to file in Google Cloud Storage.
///

func File_Write_Line ( w http.ResponseWriter ,wc *storage.Writer ,ldata string ) {

//     IN    w      : レスポンスライター
//     IN    wc       : ストレッジ用ライター
//     IN  ldata      : 1行のデータ

//    fmt.Fprintf( w, "File_Write_Line start \n" )  // デバック

///
///    一行、書き込み
///

        fmt.Fprintf(wc ,"%s" ,ldata )

// 改行する

//   fmt.Fprintf(wc ,"\n" )

}

///
///    File_Write : write data to file in Google Cloud Storage for csv file
///

func File_Write_Csv ( w http.ResponseWriter ,wc *storage.Writer ,ldata []string ) {

//     IN    w      : レスポンスライター
//     IN    wc       : ストレッジ用ライター
//     IN  ldata      : 1行のデータ

//    fmt.Fprintf( w, "File_Write_Csv start \n" )  // デバック

//    fmt.Fprintf(w, "File_Write_Csv: ldata %v\n", ldata )  // デバック

    last_flag := len(ldata) -1

//    fmt.Fprintf(w, "File_Write_Csv: last_flag %v\n", last_flag )  // デバック

	for  i := 0 ; i < len(ldata) ; i++ {

//		line_break := strings.Count( ldata[i] ,"\n" )
//		fmt.Fprintf(w, "File_Write_Csv: line_break %v\n", line_break )  // デバック

        if i == last_flag {

          fmt.Fprintf(wc ,"%s" ,ldata[i] )

		} else {

          fmt.Fprintf(wc ,"%s," ,ldata[i] )

        }

     }

///
///    改行する
///


     fmt.Fprintf(wc ,"\n" )

   return

}

///
///    File_Write : write data to file in Google Cloud Storage for csv file
///

func File_Write_Csv2 ( w http.ResponseWriter ,wc *storage.Writer ,ldata []string ) {

//     IN    w      : レスポンスライター
//     IN    wc       : ストレッジ用ライター
//     IN  ldata      : 1行のデータ

    var ldata_all string

//    fmt.Fprintf( w, "File_Write_Csv2 start \n" )  // デバック

//    fmt.Fprintf(w, "File_Write_Csv: ldata %v\n", ldata )  // デバック

    last_flag := len(ldata) -1

//    fmt.Fprintf(w, "File_Write_Csv: last_flag %v\n", last_flag )  // デバック

	for  i := 0 ; i < len(ldata) ; i++ {

        strings.Trim( ldata[i], "\n" )

		ldata_all = ldata_all + ldata[i]

        if i == last_flag {


		} else {

          ldata_all = ldata_all + ","

        }

     }

///
///    write する
///

     fmt.Fprintf( wc ,"%s" ,ldata_all )  //  ラインデータをファイルに書き込む
     fmt.Fprintf( wc ,"\n" )         //  改行する

   return

}

///
///    File_Write_Struct : write struct data to file in Google Cloud Storage.
///

func File_Write_Struct ( w http.ResponseWriter ,wc *storage.Writer ,lf_flag int64 ,ldata interface{} ) {

//     IN    w     　 : レスポンスライター
//     IN    wc       : ストレッジ用ライター
//     IN  lf_flag    : 改行フラグ
//                      0 * 改行しない
//                      1 * 改行する
//     IN  ldata      : 構造体の1行のデータ

//    fmt.Fprintf( w, "File_Write_Struct start \n" )  // デバック

///
///     一行ライト
///
    if lf_flag == 1 {

      fmt.Fprintf( wc ,"\n" )      // 改行する

	}

	fmt.Fprintf( wc ,"%v" ,ldata )  //  構造体をファイルに書き込む


//	fmt.Fprintf(w, "File_Write_Struct: ldata %v\n", ldata )  // デバック

}

///
///    Bucket_List : list bucket in  project in Google Cloud Storage.
///

func Bucket_List ( w http.ResponseWriter ,r *http.Request, project string) ([]string) {

//     IN    w        : レスポンスライター
//     IN    r     　 : リクエストパラメータ
//     IN  project    : プロジェクト名

//     OUT  one       : バケット名（複数）

//    fmt.Fprintf( w, "Bucket_List start \n" )  // デバック

//    var buckets []string

    buckets := make([]string, 0)

//	ctx := appengine.NewContext(r)
	ctx := context.Background()

//	fmt.Fprintf( w, "ctx: %v\n", ctx)

	client, _ := storage.NewClient(ctx)

	it := client.Buckets(ctx, project)

	for {
		battrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil
		}
		buckets = append(buckets, battrs.Name)
	}
	// [END list_buckets]
	return buckets
}

func Bucket_List2 ( w http.ResponseWriter ,r *http.Request ,client *storage.Client, projectID string) ([]string) {

	ctx := context.Background()

	var buckets []string

	it := client.Buckets(ctx, projectID)

	for {
		battrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil
		}
		buckets = append(buckets, battrs.Name)
	}
	// [END list_buckets]
	return buckets
}

///
///    Object_List : list object in  bucket in  project in Google Cloud Storage.
///

func Object_List( w http.ResponseWriter ,r *http.Request, bucket string) ( []string ) {

//     IN    w        : レスポンスライター
//     IN    r     　 : リクエストパラメータ
//     IN  bucket     : バケット名

//     OUT  one       : オブジェクト名（複数）

//	fmt.Fprintf( w, "Object_List start \n" )  // デバック

	var objects []string

//	ctx := appengine.NewContext(r)
	ctx := context.Background()

//	fmt.Fprintf( w, "ctx: %v\n", ctx)

	client, _ := storage.NewClient(ctx)

	it := client.Bucket(bucket).Objects(ctx, nil)

	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil
		}
		objects = append(objects, attrs.Name)

//		fmt.Fprintf( w, "Object_List : attrs.Created %v\n", attrs.Created )  //デバック

	}

	return objects
}

///
///    Object_List_Detail : list detail inf. of object in  bucket in  project in Google Cloud Storage.
///

func Object_List_Detail ( w http.ResponseWriter ,r *http.Request, bucket string) ( objects_inf []type5.Storage_B_O_View ) {

//     IN    w        : レスポンスライター
//     IN    r     　 : リクエストパラメータ
//     IN  bucket     : バケット名

//     OUT  one       : オブジェクト名（複数）

//	fmt.Fprintf( w, "Object_List_Detail start \n" )  // デバック

	var idmy int64

	var sdmy1 , sdmy2 string

	objects_inf = make([]type5.Storage_B_O_View, 0)

//	ctx := appengine.NewContext(r)
	ctx := context.Background()

//	fmt.Fprintf( w, "Object_List_Detail ctx: %v\n", ctx)  // デバック

	client, _ := storage.NewClient(ctx)

	it := client.Bucket(bucket).Objects(ctx, nil)

	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil
		}

		objects_inf = append( objects_inf, type5.Storage_B_O_View { idmy          ,
                                                                    sdmy1         ,
                                                                    sdmy2         ,
                                                                    attrs.Name    ,
//                                                                    attrs.Created
                                                                    attrs.Updated    })

//		fmt.Fprintf( w, "Object_List_Detail : attrs.Created %v\n", attrs.Created )  //デバック
//		fmt.Fprintf( w, "Object_List_Detail : attrs.ContentType: %v\n", attrs.Name )


	}

	return objects_inf
}