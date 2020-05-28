package date1

import (

	    "net/http"
	    "strconv"
//	    "fmt"
	    "strings"

	    "time"
                                                )
///
///     change date data type ( from string to time )
///

func Date_realdata_get(w http.ResponseWriter ,date string) (date_real time.Time ){

//     IN    w       : response-writer
//     IN  date　    : string type
//    OUT  date_real : time type

//    fmt.Fprintf( w, "date_realdata_get start \n" )


///
///    make time-data
///

    strings_slice := strings.Split( date, "/" )   //　split string data by mark "/"

    loc, _ := time.LoadLocation("Local")            //   get local time zone
    year_int64 ,_ :=strconv.ParseInt( strings_slice[0],10 ,0 )   //  year data type change int type
    month_int64,_  :=strconv.ParseInt( strings_slice[1],10 ,0 )   // month data type change int type
    date_int64 ,_ :=strconv.ParseInt( strings_slice[2],10 ,0 )   // date data type change int type

    switch month_int64 {
       case 1 :

             date_real = time.Date( int(year_int64) ,1 ,int(date_int64) ,0 ,0 ,0 ,0 ,loc)


             break;

       case 2 :

             date_real = time.Date( int(year_int64) ,2 ,int(date_int64) ,0 ,0 ,0 ,0 ,loc )

             break;

       case 3 :

             date_real = time.Date( int(year_int64) ,3 ,int(date_int64) ,0 ,0 ,0 ,0 ,loc )

          break;

       case 4 :

            date_real = time.Date( int(year_int64) ,4 ,int(date_int64) ,0 ,0 ,0 ,0 ,loc )

          break;

       case 5 :

             date_real = time.Date( int(year_int64) ,5 ,int(date_int64) ,0 ,0 ,0 ,0 ,loc )

          break;

       case 6:

             date_real = time.Date( int(year_int64) ,6 ,int(date_int64) ,0 ,0 ,0 ,0 ,loc )

          break;

       case 7 :

             date_real = time.Date( int(year_int64) ,7 ,int(date_int64) ,0 ,0 ,0 ,0 ,loc )

          break;

       case 8 :

             date_real = time.Date( int(year_int64) ,8 ,int(date_int64) ,0 ,0 ,0 ,0 ,loc )

          break;

       case 9 :

             date_real = time.Date( int(year_int64) ,9 ,int(date_int64) ,0 ,0 ,0 ,0 ,loc )

          break;

       case 10:

            date_real = time.Date( int(year_int64) ,10 ,int(date_int64) ,0 ,0 ,0 ,0 ,loc )

          break;

       case 11 :

            date_real = time.Date( int(year_int64) ,11 ,int(date_int64) ,0 ,0 ,0 ,0 ,loc )


          break;

       case 12 :

             date_real = time.Date( int(year_int64) ,12 ,int(date_int64) ,0 ,0 ,0 ,0 ,loc )

          break;

          }
//    fmt.Fprintf( w, "date_realdata_get : date_real %v\n", date_real )
//    fmt.Fprintf( w, "date_realdata_get normal end \n" )

    return date_real


}
