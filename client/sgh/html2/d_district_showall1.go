package html2

   const D_district_showall1 = `
   <!DOCTYPE html>
   <html>
     <head>
        <meta charset="UTF-8">
        <title>district-information</title>
        <link rel="stylesheet" href="css/member1_show.css" type="text/css">
     </head>
     <body>
       <article>
       </article>
       <section>
            <form method="GET" action="/d_district_showall1" >

              <table border="2" cellpadding="8" align="center" bgcolor="#00ced1">
              <h1  align="center">Register Your deliverly-number ,please</h1>
              <tr> <th>district-no</th> <th>district-name</th>  <th>access</th> </tr>
              <td > <input type="text" name="district_no"  size="5" />               </td>
              <td > <input type="text" name="district_name" size="10" />            </td>

              <td > <input type="submit" size="2"  value="register" />           </td>
            </form>
       </section>
       <section id="main">


         <table border="2" cellpadding="8" align="center" bgcolor="#00ced1">
         <h2 align="center">List Of Deliverly Situation</h2>

           <tr> <th>district-no</th> <th>district-name</th> <th>area in district</th> <th>change district inf.</th>
                                           <th>change district inf.</th> <th>show schedule inf.</th> <th>show schedule inf.</th> </tr>
           {{range .}}
             <tr>

             <form method="GET" action="/d_district_update" >


               <td>
                  <input type="text" name="district_no" size="5" align="center" value="{{.District_No|html}}" />
                  <input type="hidden" name="id"  value="{{.Id|html}}"/>

               </td>
               <td>
                  <input type="text" name="district_name" size="10" align="center" value="{{.District_Name|html}}" />
               </td>
               <td>


                   {{range .D_Area_Slice}}


                    <input type="text" name="district_name" size="10" align="center" value="{{.Area_No|html}} {{.Area_Name|html}}" />

                    <br>

                   {{end}}

               </td>

               <td>
                  <input type="submit"  size="2" value="change"  />
               </td>
             </form>

             <form method="GET" action="/d_district_delete" >

               <td>
                  <input type="hidden" name="id"  value="{{.Id|html}}"/>
                  <input type="submit"  size="2" value="delete"  />
               </td>
             </form>
             <form method="GET" action="/d_district_area" >

               <td>
                  <input type="hidden" name="id"  value="{{.Id|html}}"/>
                  <input type="submit"  size="2" value="area"  />
               </td>
             </form>
             <form method="GET" action="/d_schedule_keyin" >
               <td>
                  <input type="hidden" name="id"  value="{{.Id|html}}"/>
                  <input type="submit"  size="2" value="schedule"  />
               </td>
             </form>
             <form method="GET" action="/car_show" >
               <td>
                  <input type="hidden" name="id"  value="{{.Id|html}}"/>
                  <input type="submit"  size="2" value="car"  />
               </td>
             </form>

           </tr>
          {{end}}
         </table>
        </section>

     </body>
   </html>
   `
