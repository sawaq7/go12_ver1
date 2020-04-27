package vaccine1

import (
//	      "fmt"
	      "io"
	      "strings"
	      "net/http"
	      "bufio"

	      "github.com/sawaq7/go12_ver1/storage2"

	      "cloud.google.com/go/storage"

                                         )
///
///   StoragePack : pack a file in Google Cloud Storage.
///

func File_Pack ( w http.ResponseWriter , r *http.Request ,bucket_name string ,file_name string ){

//     IN    w      : response-writer
//     IN    r      : request- paramete
//     IN  bucket     : bucket-name
//     IN  filename   : file-name

    file_name2 := file_name + "temp"

    writer_minor , _ := storage2.Storage_basic( "create" ,bucket_name ,file_name2 , w , r  )

    writer, _ := writer_minor.(*storage.Writer)

	defer writer.Close()

//      open file which needs packing

    reader_minor , _ := storage2.Storage_basic( "open" ,bucket_name ,file_name , w , r  )

    reader, _ := reader_minor.(io.ReadCloser)

    defer reader.Close()

///    get file-reader

    sreader := bufio.NewReaderSize(reader, 4096)

    index := 0

    for {

      index ++

//      fmt.Fprintf(w, "File_Pack : lndex %v\n", index )

///    read one-record

      line ,err  := sreader.ReadString('\n')

      if err == io.EOF {

	      break

	  }

	  if err != nil {

		 http.Error(w, err.Error(), http.StatusInternalServerError)
	     return

	  }

//	  line = strings.Replace( line, ",", " ", -1)

      column := strings.Count( line ,",") + 1

//      fmt.Fprintf(w, "File_Pack : column %v\n", column )

      if  column > 1 {      //   if a record isn't space ,write in file

          line2 := strings.Trim(line, " ")           //   trim at both ends
//          fmt.Fprintf(w, "File_Pack :line2 [%s]\n", line2 )

          storage2.File_Write_Line ( w ,writer ,line2 )

      }

   }

///
/// 　　　　change file-name
///

   storage2.File_Delete ( w , r ,bucket_name ,file_name  )    //   delete old file

   storage2.File_Rename ( w , r  ,bucket_name ,file_name2 ,file_name ) // rename new file


//	fmt.Fprintf(w, " File_Pack : Calculate succeeded.\n" )

	return
}
