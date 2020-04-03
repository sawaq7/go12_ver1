package html2

   const D_district_area_show = `
   <!DOCTYPE html>
   <html>
     <head>
        <meta charset="UTF-8">
        <title>チE�Eタストアの更新・削除</title>
        <link rel="stylesheet" href="css/member1_show.css" type="text/css">
     </head>
     <body>
       <article>
       </article>
       <section id="main">


           <table border="2" cellpadding="8" align="center" bgcolor="#00ced1">
           <h2 align="center">List Of Deliverly Situation</h2>

           <tr> <th>district-no</th> <th>district-name</th> <th>area-no</th> <th>area-name</th><th>area-detail</th><th>access1</th> <th>access2</th> </tr>
           {{range .}}
             <tr>

             <form method="GET" action="/d_district_area_update" >


               <td>
                  <input type="text" name="district_no" size="5" align="center" value="{{.District_No|html}}" />
                  <input type="hidden" name="id"  value="{{.Id|html}}"/>

               </td>
               <td>
                  <input type="text" name="district_name" size="10" align="center" value="{{.District_Name|html}}" />
               </td>
               <td>
                  <input type="text" name="area_no" size="10" align="center" value="{{.Area_No|html}}" />
               </td>
               <td>
                  <input type="text" name="area_name" size="10" align="center" value="{{.Area_Name|html}}" />
               </td>
               <td>
                  <input type="text" name="area_detail" size="30" align="center" value="{{.Area_Detail|html}}" />
               </td>

               <td>
                  <input type="submit"  size="2" value="変更"  />
               </td>
             </form>

             <form method="GET" action="/d_district_area_delete" >

               <td>
                  <input type="hidden" name="id"  value="{{.Id|html}}"/>
                  <input type="submit"  size="2" value="削除"  />
               </td>
             </form>

             <form method="GET"  align="center" action="/deliver_showall2">
               <td >
                <input type="submit" value="配達日報" />

                                                                             </td>
             </form>

           </tr>
           {{end}}
           </table>

        </section>
        <section >
          <h2 align="center">Input New Data ,please</h2>
          <form method="GET" action="/d_district_area_show" >

              <table border="2" cellpadding="8" align="center" bgcolor="#00ced1">
    /*          <h1  align="center">Register Your deliverly-number ,please</h1>   */
              <tr>   <th>area_name</th> <th>area-detail</th>  <th>access</th> </tr>
                                                                                                 </td>
              <td > <input type="text" name="area_name" size="10" />            </td>
              <td > <input type="text" name="area_detail" size="30" />            </td>
              <td > <input type="submit" size="2"  value="登録" />           </td>
          </form>

        </section>
     </body>
   </html>
   `
