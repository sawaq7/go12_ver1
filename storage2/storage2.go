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
	"github.com/sawaq7/go12_ver1/general/type5"

                                        )


///
///   Bucket_Handler_Get : get bucket handler for Google Cloud Storage.
///                                              type-2

func Bucket_Handler_Get(w http.ResponseWriter ,r *http.Request ,bucket string) (*storage.BucketHandle) {

//     IN    w      : response-writer
//     IN    r      : request- parameter
//     IN  bucket   :
//     OUT  one     : bucket handler

//    fmt.Fprintf( w, "Bucket_Handler_Get start \n" )

    ctx := context.Background()

	client, err := storage.NewClient(ctx)
	if err != nil {
	   http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil
	}
	return client.Bucket(bucket)
}

///
///    File_Open : open file in Google Cloud Storage.
///

func File_Open(w http.ResponseWriter ,r *http.Request ,bucket string ,filename string)(io.ReadCloser) {

//     IN    w      : response-writer
//     IN    r      : request- parameter
//     IN  bucket     :
//     IN  filename   :
//     OUT  one       : reader for storage

//    fmt.Fprintf( w, "File_Open start \n" )

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

///
///   File_Create : create a file in Google Cloud Storage.
///                        type-2

func File_Create ( w http.ResponseWriter ,r *http.Request ,bucket string ,filename string ) ( *storage.Writer ){

//     IN    w      : response-writer
//     IN    r      : request- paramete
//     IN  bucket     : bucket-name
//     IN  filename   : file-name

//     OUT  one       : writer for storage

//    fmt.Fprintf( w, "File_Create start \n" )

/// get bucket handler for storage

    bh := Bucket_Handler_Get( w ,r ,bucket )

    ctx := context.Background()

    wc := bh.Object(filename).NewWriter(ctx)

	wc.ContentType = "text/plain; charset=utf-8"
	wc.ACL = []storage.ACLRule{{Entity: storage.AllUsers, Role: storage.RoleReader}}

	wc.CacheControl = "public, max-age=86400"
	   // Entries are immutable, be aggressive about caching (1 day).

	return wc
}

///
///   File_Create : create a file in Google Cloud Storage.
///                        type-2

func File_Create2 ( w http.ResponseWriter ,r *http.Request ,bucket string ,filename string ,content_type string  ) ( *storage.Writer ){

//     IN    w      : response-writer
//     IN    r      : request- paramete
//     IN  bucket     : bucket-name
//     IN  filename   : file-name
//     IN  content_type   :  content-type

//     OUT  one       : writer for storage

//    fmt.Fprintf( w, "File_Create2 start \n" )

/// get bucket handler for storage

    bh := Bucket_Handler_Get( w ,r ,bucket )

    ctx := context.Background()

    wc := bh.Object(filename).NewWriter(ctx)

	wc.ContentType = content_type
	wc.ACL = []storage.ACLRule{{Entity: storage.AllUsers, Role: storage.RoleReader}}

	wc.CacheControl = "public, max-age=86400"
	    // Entries are immutable, be aggressive about caching (1 day).

//	fmt.Fprintf(gcs_gae.W, "StorageCreate: wc　%d\n", wc  )

	return wc
}

///                                                           ///
///   StorageCopy : copy a file in Google Cloud Storage.      ///
///                                                           ///

func File_Copy ( w http.ResponseWriter , r *http.Request ,bucket string ,fileName string ,fileName2 string){

//     IN    w      : response-writer
//     IN    r      : request- paramete
//     IN  bucket   : bucket-name
//     IN  filename    : the basic's file
//     IN  filename2   : the file which is copied

//    fmt.Fprintf( w, "File_Copy start \n" )

	writer := File_Create( w ,r ,bucket ,fileName2 ) //  make the file which is empty
    defer writer.Close()



    reader := File_Open(w ,r ,bucket ,fileName)  // open the file
    defer reader.Close()

    // get file reader

	if _, err := io.Copy(writer, reader); err != nil {
	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return

	}

}

///
///   StorageRename : rename a file in Google Cloud Storage.
///

func File_Rename ( w http.ResponseWriter ,r *http.Request ,bucket string ,fileName1 string ,fileName2 string ) {

//     IN    w      : response-writer
//     IN    r      : request- paramete
//     IN  bucket   : bucket-name
//     IN  filename1   : old file name
//     IN  filename2   : new file name

//    fmt.Fprintf( w, "File_Rename start \n" )

	File_Copy ( w ,r ,bucket , fileName1 ,fileName2  )    // copy the file1

    File_Delete ( w , r ,bucket ,fileName1  )             // delete file1

}

///                                                               ///
///   StorageDelete : delete a file in Google Cloud Storage.      ///
///                                                               ///

func File_Delete ( w http.ResponseWriter , r *http.Request ,bucket string ,fileName string  ) {

//     IN    w      : response-writer
//     IN    r      : request- paramete
//     IN  bucket   : bucket-name
//     IN  filename    : the file which is deleted

//    fmt.Fprintf( w, "File_Delete start \n" )

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

func File_write ( w http.ResponseWriter ,wc *storage.Writer ,ldata []string ) {

//     IN    w      : request- parameter
//     IN    wc       : writer for storage
//     IN  ldata      : one line data

//    fmt.Fprintf( w, "File_Write start \n" )

	count := 0 //　counter initialize

	for  i := 0 ; i < len(ldata) ; i++ {

	    count ++  //　one count

///    write one-line

        fmt.Fprintf(wc ,"%s " ,ldata[i] )

   }

///   line feed

   fmt.Fprintf(wc ,"\n" )

}

///
///    File_Write_Line : write data to file in Google Cloud Storage.
///

func File_Write_Line ( w http.ResponseWriter ,wc *storage.Writer ,ldata string ) {

//     IN    w      : response-writer
//     IN    r      : request- paramete
//     IN  ldata      : line data

//    fmt.Fprintf( w, "File_Write_Line start \n" )

///
///    write one-line
///

        fmt.Fprintf(wc ,"%s" ,ldata )

///     line feed

//   fmt.Fprintf(wc ,"\n" )

}

///
///    File_Write : write data to file in Google Cloud Storage for csv file
///

func File_Write_Csv ( w http.ResponseWriter ,wc *storage.Writer ,ldata []string ) {

//     IN    w      : response-writer
//     IN    wc       :  writer for storage
//     IN  ldata      : line　data

//    fmt.Fprintf( w, "File_Write_Csv start \n" )

//    fmt.Fprintf(w, "File_Write_Csv: ldata %v\n", ldata )

    last_flag := len(ldata) -1

//    fmt.Fprintf(w, "File_Write_Csv: last_flag %v\n", last_flag )

	for  i := 0 ; i < len(ldata) ; i++ {

//		line_break := strings.Count( ldata[i] ,"\n" )
//		fmt.Fprintf(w, "File_Write_Csv: line_break %v\n", line_break )

        if i == last_flag {

          fmt.Fprintf(wc ,"%s" ,ldata[i] )

		} else {

          fmt.Fprintf(wc ,"%s," ,ldata[i] )

        }

     }

///
///    line feed
///


     fmt.Fprintf(wc ,"\n" )

   return

}

///
///    File_Write : write data to file in Google Cloud Storage for csv file
///

func File_Write_Csv2 ( w http.ResponseWriter ,wc *storage.Writer ,ldata []string ) {

//     IN    w      : response-writer
//     IN    wc       : writer for storage
//     IN  ldata      : line data

    var ldata_all string

//    fmt.Fprintf( w, "File_Write_Csv2 start \n" )

//    fmt.Fprintf(w, "File_Write_Csv: ldata %v\n", ldata )

    last_flag := len(ldata) -1

//    fmt.Fprintf(w, "File_Write_Csv: last_flag %v\n", last_flag )

	for  i := 0 ; i < len(ldata) ; i++ {

        strings.Trim( ldata[i], "\n" )

		ldata_all = ldata_all + ldata[i]

        if i == last_flag {


		} else {

          ldata_all = ldata_all + ","

        }

     }

///
///    write  one-line-data
///

     fmt.Fprintf( wc ,"%s" ,ldata_all )  //  one-line-data
     fmt.Fprintf( wc ,"\n" )         //  line feed

   return

}

///
///    File_Write_Struct : write struct data to file in Google Cloud Storage.
///

func File_Write_Struct ( w http.ResponseWriter ,wc *storage.Writer ,lf_flag int64 ,ldata interface{} ) {

//     IN    w     　 : response-writer
//     IN    wc       : writer for storage
//     IN  lf_flag    : flag for line-feed
//                      0 * line feed
//                      1 * don't line feed
//     IN  ldata      : the one record which is written for struct

//    fmt.Fprintf( w, "File_Write_Struct start \n" )

///
///     one line write
///
    if lf_flag == 1 {

      fmt.Fprintf( wc ,"\n" )      // line feed

	}

	fmt.Fprintf( wc ,"%v" ,ldata )  //  write data with format of struct in file


//	fmt.Fprintf(w, "File_Write_Struct: ldata %v\n", ldata )

}

///
///    Bucket_List : list bucket in  project in Google Cloud Storage.
///

func Bucket_List ( w http.ResponseWriter ,r *http.Request, project string) ([]string) {

//     IN    w      : response-writer
//     IN    r      : request- paramete
//     IN  project    : project-name

//     OUT  one       :

//    fmt.Fprintf( w, "Bucket_List start \n" )

//    var buckets []string

    buckets := make([]string, 0)

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

//     IN    w      : response-writer
//     IN    r      : request- paramete
//     IN  bucket   : bucket-name

//     OUT  one     : object list

//	fmt.Fprintf( w, "Object_List start \n" )

	var objects []string

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

//		fmt.Fprintf( w, "Object_List : attrs.Created %v\n", attrs.Created )

	}

	return objects
}

///
///    Object_List_Detail : list detail inf. of object in  bucket in  project in Google Cloud Storage.
///

func Object_List_Detail ( w http.ResponseWriter ,r *http.Request, bucket string) ( objects_inf []type5.Storage_B_O_View ) {

//     IN    w      : response-writer
//     IN    r      : request- paramete
//     IN  bucket   : bucket-name

//     OUT  one       : object list

//	fmt.Fprintf( w, "Object_List_Detail start \n" )

	var idmy int64

	var sdmy1 , sdmy2 string

	objects_inf = make([]type5.Storage_B_O_View, 0)

	ctx := context.Background()

//	fmt.Fprintf( w, "Object_List_Detail ctx: %v\n", ctx)

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

//		fmt.Fprintf( w, "Object_List_Detail : attrs.Created %v\n", attrs.Created )
//		fmt.Fprintf( w, "Object_List_Detail : attrs.ContentType: %v\n", attrs.Name )


	}

	return objects_inf
}
