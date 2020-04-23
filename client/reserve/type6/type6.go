package type6

import (

//          "time"

                  )
//
// the collection of struct information for reservation
//

//
//      Guest information　
//

type Guest struct {            ///    guest inf. for d.s.

       Id            int64           //　id
       Guest_No      int64           // guest no
	   Guest_Name    string          // guest name


   }

type Guest_Temp struct {         ///    guest inf. for temp.

       Id          int64           //　id
       Guest_No    int64           // guest no
	   Guest_Name  string          // guest name


   }
//
// reservation information　
//

type Guest_Reserve_Minor struct {               ///  reserve inf.

       Id             int64          //
       Line_No        int64          // line no
       Date           string         // reservation date
       Guest_No       int64          // guest no
	   Guest_Name     string         // guest name
	   Start_Time     string         // start time
	   End_Time       string         // end time
   }
//
// reservation information　
//

type Guest_Reserve_View struct {               ///  reserve inf. for temp.

       Id             int64          //
       Line_No        int64          // line no
       Date           string         // reservation date
       Guest_No       int64          // guest no
	   Guest_Name     string         // guest name
	   Start_Time     string         // start time
	   End_Time       string         // end time
	   File_Name      string         // file name
	   Url            string         // url

   }
//
// reservation information　
//

type Guest_Medical_Record struct {   ///  medical record inf.

       Id             int64          //
       Line_No        int64          // line no
       Date           string         // consultation day
       Guest_No       int64          // guest no
	   Guest_Name     string         // guest name
	   Text           string         // text
   }

//
// reservation information　
//

type Guest_Medical_Xray struct {   ///  medical xray inf.

       Id             int64          //
       Line_No        int64          // line no
       Date           string         // filming date
       Guest_No       int64          // guest no
	   Guest_Name     string         // guest name
	   File_Name      string         // file name
	   Url            string         // url

	 }

//
// reservation information　
//

type Guest_Payment struct {       ///  payment inf.

       Id             int64          //
       Line_No        int64          // line no
       Date           string         // pay day
       Guest_No       int64          // guest no
	   Guest_Name     string         // ゲスト名
	   Item           string         // guest name
	   Amount         int64          // amount paid
   }
