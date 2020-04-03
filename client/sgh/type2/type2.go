package type2

import (
          "time"
                  )
//
// the collection of struct information for sagawa express
//

//
// deliver district informationãEå°åºæE ±EE//

type D_District struct {               /// ãEEã¿ã¹ãã¢ç¨

       Id               int64           //ããEEã¿id
       District_No      int64           // ééå°åNO.
	   District_Name    string          // ééå°åå


   }

type D_District_Temp struct {         /// ä¸æãã¡ã¤ã«ç¨

       Id             int64           //ããEEã¿id
       District_No    int64           // ééå°åNO.
	   District_Name  string          // ééå°åå


   }

type D_District_View struct {               /// è¡¨ç¤ºç¨

       Id               int64           //ããEEã¿id
       District_No      int64           // ééå°åNO.
	   District_Name    string          // ééå°åå
       D_Area_Slice     []D_Area        //ãã¨ãªã¢æE ±ã®ã¹ã©ã¤ã¹

   }

//
// deliver area informationãEã¨ãªã¢æE ±EE//

type D_Area struct {               /// å¸¸ç¨ãã¡ã¤ã«ç¨

       Id             int64           // ãEEã¿id
       Course_No      int64           // ã³ã¼ã¹NO.
       District_No    int64           // ééå°åNO.
       District_Name  string          // ééå°åå
       Area_No        int64           // ééã¨ãªã¢NO.
	   Area_Name      string          // ééã¨ãªã¢åE	   Area_Detail    string          // ééã¨ãªã¢ã®è©³ç´°
	   Number_Total   int64           // å®EEééç·æ°
	   Time_Total     float64         // å®EEééç·æéE	   Productivity   float64         // å®EEçç£æ§
       Car_No         int64           // ã¬ã®ã¥ã©ã¼å·è»E   }

//
// deliver area informationãEã¨ãªã¢æE ±EE//

type D_Area_Temp struct {           /// ä¸æãã¡ã¤ã«ç¨

       Id           int64           // ãEEã¿id
       Course_No    int64           // ã³ã¼ã¹NO.
       District_No    int64         // ééå°åNO.
       District_Name  string          // ééå°åå
       Area_No      int64           // ééã¨ãªã¢NO.
	   Area_Name    string          // ééã¨ãªã¢åE	   Area_Detail  string          // ééã¨ãªã¢ã®è©³ç´°
	   Number_Total   int64           // å®EEééç·æ°
	   Time_Total     float64         // å®EEééç·æéE	   Productivity   float64         // å®EEçç£æ§
       Car_No       int64           // ã¬ã®ã¥ã©ã¼å·è»E   }

//
// deliver informationãEéEéï¼E//

type Deliver struct {

       Id           int64           // ãEEã¿id
       Line_No      int64           // è¡NO.
       Course_No    int64           // ã³ã¼ã¹NO.
       District_No    int64         // ééå°åNO.
       District_Name  string          // ééå°åå
       Area_No      int64           // ééã¨ãªã¢NO.
	   Area_Name    string          // ééã¨ãªã¢åE       Date         string         // ééæ¥
       Date_Real    time.Time      // ééæ¥(ã¿ã¤ã ãEEã¿EE	   Car_No       int64           // å·è»E	   Private_No   int64           // åäººçªå·
	   Car_Division int64           // è»ä¸¡åºåEEï¼è»½èªåè»ãEï¼ã¯ã´ã³ãEï¼ãã©ãE¯
	   Number       int64           // å®EEééæ°

   }


//
// deliver's Schedule informationãEéEéEã¹ã±ã¸ã¥ã¼ã«EE//

type Private struct {

       Id             int64           //ããEEã¿id
       Worker_No      int64           // ã¯ã¼ã«ã¼NO.
       Worker_Name    string          // ã¯ã¼ã«ã¼åE       Worker_Type    string          // ã¯ã¼ã«ã¼ã¿ã¤ãããSD : ã»ã¼ã«ã¹ãã©ã¤ããEãDD : å®EEãã©ã¤ããE
                                      // ãããããããããOD : å¤æ³¨ãã©ã¤ããE
       Worker_Salary  float64         // ã¯ã¼ã«ã¼ãµã©ãªã¼Eå¹´åï¼E       Worker_Twh     float64         // ã¯ã¼ã«ã¼ç·å´åæéï¼å¹´éï¼Twh : total working hours
       Worker_H_Pay   float64         // ã¯ã¼ã«ã¼æçµ¦ãH_Pay : hourlyãpay
       Number_Total   int64           // å®EEééç·æ°
	   Time_Total     float64         // å®EEééç·æéE	   Productivity   float64         // å´åçç£æ§

   }


//
//   car informationãEå·è»æå ±EE//

type Car struct {

       Id                int64           //ããEEã¿id
       District_No       int64           // ééå°åNO.
       District_Name     string          // ééå°åå
       Car_No            int64           // å·è»NO.(ã·ã¼ã±ã³ã·ã£ã«ãªNO)
	   Car_Name          string          // å·è»å
	   Car_Explain       string          // å·è»ãEèµ°è¡ã«ã¼ãç­ãEèª¬æE	   Number_Total      int64           // å·è»ãEå®EEééç·æ°
	   Time_Total        float64         // å®EEééç·æéE	   Productivity      float64         // å´åçç£æ§

   }



//
// deliver's Schedule informationãEéEéEã¹ã±ã¸ã¥ã¼ã«EE//

type D_Schedule struct {

       Id             int64           //ããEEã¿id
       District_No    int64           // ééå°åNO.
       District_Name  string          // ééå°åå
       Date           string          // ééæ¥
       Date_Real    time.Time         // ééæ¥(ã¿ã¤ã ãEEã¿EE       Expected_Num   float64         // è·ç©ã®äºæ³åæ°
       Judge          string          // ééã®å¤å®E       Course_Num     int64           // ã³ã¼ã¹æ°
	   Course_01      string          // ã³ã¼ã¹1ãä¹è»ãããã©ã¤ããE
	   Car_Name_01     string          // ã³ã¼ã¹1ã®å·è»å
	   Course_02      string          // ã³ã¼ã¹2ãä¹è»ãããã©ã¤ããE
	   Car_Name_02     string          // ã³ã¼ã¹1ã®å·è»å
	   Course_03      string          // ã³ã¼ã¹3ãä¹è»ãããã©ã¤ããE
	   Car_Name_03     string          // ã³ã¼ã¹1ã®å·è»å
	   Course_04      string          // ã³ã¼ã¹4ãä¹è»ãããã©ã¤ããE
	   Car_Name_04     string          // ã³ã¼ã¹1ã®å·è»å
       Course_05      string          // ã³ã¼ã¹5ãä¹è»ãããã©ã¤ããE
       Car_Name_05     string          // ã³ã¼ã¹1ã®å·è»å
	   Course_06      string          // ã³ã¼ã¹6ãä¹è»ãããã©ã¤ããE
	   Car_Name_06     string          // ã³ã¼ã¹1ã®å·è»å
	   Course_07      string          // ã³ã¼ã¹7ãä¹è»ãããã©ã¤ããE
	   Car_Name_07     string          // ã³ã¼ã¹1ã®å·è»å
	   Course_08      string          // ã³ã¼ã¹8ãä¹è»ãããã©ã¤ããE
	   Car_Name_08     string          // ã³ã¼ã¹1ã®å·è»å
	   Course_09      string          // ã³ã¼ã¹9ãä¹è»ãããã©ã¤ããE
	   Car_Name_09     string          // ã³ã¼ã¹1ã®å·è»å
	   Course_10      string          // ã³ã¼ã¹10ãä¹è»ãããã©ã¤ããE
	   Car_Name_10     string          // ã³ã¼ã¹1ã®å·è»å

   }
//
// the collection of some condition expressionãEæ¡ä»¶å¼ï¼E//

type Sgh_Ai struct {

       Id              int64          // ãEEã¿id
       Course_No       int64          // ã³ã¼ã¹NO.
       District_No     int64          // ééå°åNO.
       District_Name   string         // ééå°åå
       Area_No         int64          // ééã¨ãªã¢NO.
	   Area_Name       string         // ééã¨ãªã¢åE       Date_Basic      string         // åºæºæ¥(é¢æ°ã®å§ç¹EE       Date_Basic_Real time.Time      // åºæºæ¥(ã¿ã¤ã ãEEã¿EE       Ex_Type         string         // æ¡ä»¶å¼ãEã¿ã¤ãE                                         // 1. function : é¢æ°
       Expression    string         // æ¡ä»¶å¼E	   Item_Num      int64          // é E°
	   Item1_Name    string         // é E1
	   Item1_Factor  float64        // é EEä¿æ°1
	   Item2_Name    string         // é E1
	   Item2_Factor  float64        // é EEä¿æ°1
	   Item3_Name    string         // é E1
	   Item3_Factor  float64        // é EEä¿æ°1
	   Item4_Name    string         // é E1
	   Item4_Factor  float64        // é EEä¿æ°1
	   Item5_Name    string         // é E1
	   Item5_Factor  float64        // é EEä¿æ°1


   }



