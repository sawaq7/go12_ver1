package type2

import (
          "time"
                  )
//
// the collection of struct information for sagawa express
//

//
// deliver district information

type D_District struct {               ///  for datastore

       Id               int64           //
       District_No      int64           // 配達地域NO.
	   District_Name    string          // 配達地域名


   }

type D_District_Temp struct {         /// 一時ファイル用

       Id             int64           //　
       District_No    int64           // 配達地域NO.
	   District_Name  string          // 配達地域名


   }

type D_District_View struct {               ///  for showing

       Id               int64           //　
       District_No      int64           // 配達地域NO.
	   District_Name    string          // 配達地域名
       D_Area_Slice     []D_Area        //　slice for area

   }

//
// deliver area information　

type D_Area struct {               ///    for datastore

       Id             int64           //
       Course_No      int64           // コースNO.
       District_No    int64           // 配達地域NO.
       District_Name  string          // 配達地域名
       Area_No        int64           // 配達エリアNO.
	   Area_Name      string          // 配達エリア吁E
	   Area_Detail    string          // 配達エリアの詳細
	   Number_Total   int64           // total delivery number
	   Time_Total     float64         // total deliver time
	   Productivity   float64         //
       Car_No         int64           // regular car number

    }

//
// deliver area information　�E�エリア惁E���E�E//

type D_Area_Temp struct {           ///   for temp.

       Id           int64           //
       Course_No    int64           // コースNO.
       District_No    int64         // 配達地域NO.
       District_Name  string          // 配達地域名
       Area_No      int64           // 配達エリアNO.
	   Area_Name    string          //
	   Area_Detail  string          // 配達エリアの詳細
	   Number_Total   int64         // total delivery number
	   Time_Total     float64       // total deliver time
	   Productivity   float64       //
       Car_No       int64           // regular car number

   }

//
// deliver information　

type Deliver struct {

       Id           int64           //
       Line_No      int64           // 行NO.
       Course_No    int64           // コースNO.
       District_No    int64         // 配達地域NO.
       District_Name  string          // 配達地域名
       Area_No      int64           // area number
	   Area_Name    string          //
	   Date         string         // 配達日
       Date_Real    time.Time      // 配達日（time type）
       Car_No       int64           // regular car number
       Private_No   int64           // 個人番号
	   Car_Division int64           // 1：light car 2：estate car 3：truck
	   Number       int64           // daily delivery number

   }


//
// deliver's Schedule information
//

type Private struct {

       Id             int64           //　
       Worker_No      int64           //
       Worker_Name    string          //
       Worker_Type    string          // ワーカータイプ　　SD : sales driver　DD : delivery driver
                                      // 　　　　　　　　　OD : outsoucing driver
       Worker_Salary  float64         // waker salary in year
       Worker_Twh     float64         // total working hours in yearTwh
       Worker_H_Pay   float64         // ワーカー時給　H_Pay : hourly　pay
       Number_Total   int64           // total delivery number
	   Time_Total     float64         // total deliver time
	   Productivity   float64         //

   }


//
//   car information　�E�号車情報�E�E//

type Car struct {

       Id                int64           //
       District_No       int64           // 配達地域NO.
       District_Name     string          // 配達地域名
       Car_No            int64           // 号車NO.(シーケンシャルなNO)
	   Car_Name          string          // 号車名
	   Car_Explain       string          // explain cource inf.
	   Number_Total      int64           // total delivery number
	   Time_Total        float64         // total deliver time
	   Productivity      float64         //

   }



//
// deliver's Schedule information　�E��E遁Eスケジュール�E�E//

type D_Schedule struct {

       Id             int64           //
       District_No    int64           // 配達地域NO.
       District_Name  string          // 配達地域名
       Date           string          // 配達日
       Date_Real    time.Time         // 配達日(time type）
       Expected_Num   float64         // 荷物の予想個数
       Judge          string          // judge delivery
       Course_Num     int64           // コース数
	   Course_01      string          // driver name1 who deliver
	   Car_Name_01     string         // car number which is course 1
	   Course_02      string          // driver name2 who deliver
	   Car_Name_02     string         // car number which is course 2
	   Course_03      string          // driver name3 who deliver
	   Car_Name_03     string         // car number which is course 3
	   Course_04      string          // driver name4 who deliver
	   Car_Name_04     string         // car number which is course 4
       Course_05      string          // driver name5 who deliver
       Car_Name_05     string         // car number which is course 5
	   Course_06      string          // driver name6 who deliver
	   Car_Name_06     string         // car number which is course 6
	   Course_07      string          // driver name7 who deliver
	   Car_Name_07     string         // car number which is course 7
	   Course_08      string          // driver name8 who deliver
	   Car_Name_08     string         // car number which is course 8
	   Course_09      string          // driver name9 who deliver
	   Car_Name_09     string         // car number which is course 9
	   Course_10      string          // driver name10 who deliver
	   Car_Name_10     string         // car number which is course 10

   }
//
// the collection of some condition expression　�E�条件式！E//

type Sgh_Ai struct {

       Id              int64          // チE�Eタid
       Course_No       int64          // コースNO.
       District_No     int64          // 配達地域NO.
       District_Name   string         // 配達地域名
       Area_No         int64          // 配達エリアNO.
	   Area_Name       string         // 配達エリア吁E
	   Date_Basic      string         // 基準日(function start point)
	   Date_Basic_Real time.Time      // 基準日(time type)
	   Ex_Type         string         // expression type
	                                  // 1. function
       Expression    string
	   Item_Num      int64          // term number
	   Item1_Name    string         // term1
	   Item1_Factor  float64        // term facter1
	   Item2_Name    string         // term2
	   Item2_Factor  float64        // term facter2
	   Item3_Name    string         // term3
	   Item3_Factor  float64        // term facter3
	   Item4_Name    string         // term4
	   Item4_Factor  float64        // term facter4
	   Item5_Name    string         // term5
	   Item5_Factor  float64        // term facter5

   }



