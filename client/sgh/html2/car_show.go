package html2

   const Car_show = `
   <!DOCTYPE html>
   <html>
     <head>
        <meta charset="UTF-8">
        <title>datastore update/delete</title>
        <link rel="stylesheet" href="css/sgh/d_district_area.css" type="text/css">
     </head>
     <body>
       <header>



       </header>
       <nav>


	   </nav>
       <article>
         <table border="2" cellpadding="12" align="center" bgcolor="#00ced1">
         <h2 align="center">District Name</h2>

           <tr> <th>district-no</th> <th>district-name</th> </tr>
           {{range .}}
             {{ if eq .Car_No 1}}

             <tr>

             <form method="GET" action="/d_district_area_update" >


               <td>
                  <input type="text" name="district_no" size="5" align="center" value="{{.District_No|html}}" />
                  <input type="hidden" name="id"  value="{{.Id|html}}"/>

               </td>
               <td>
                  <input type="text" name="district_name" size="10" align="center" value="{{.District_Name|html}}" />
               </td>
              </form>

            </tr>
            {{end}}
           {{end}}
           </table>


       </article>
       <section id="main">
           <table border="2" cellpadding="8" align="center" bgcolor="#00ced1">
           <h2 align="center">Car Information</h2>

           <tr> <th>car-no</th><th>car-name  </th> <th>car-explain</th> <th>number-total</th><th>time-total</th> <th>productivity</th>

                                                <th>access1</th> <th>access2</th> </tr>
           {{range .}}

             <tr>

             <form method="GET" action="/car_update" >

               <td>
                  <input type="text" name="car_no" size="5" align="center" value="{{.Car_No|html}}" />
               </td>
               <td>
                  <input type="text" name="car_name" size="15" align="center" value="{{.Car_Name|html}}" />
               </td>
               <td>
                  <input type="text" name="car_explain" size="30" align="center" value="{{.Car_Explain|html}}" />
               </td>
               <td>
                  <input type="text" name="car_explain" size="5" align="center" value="{{.Number_Total|html}}" />
               </td>
               <td>
                  <input type="text" name="car_explain" size="5" align="center" value="{{.Time_Total|html}}" />
               </td>
               <td>
                  <input type="text" name="car_explain" size="5" align="center" value="{{.Productivity|html}}" />
               </td>
               <td>
                  <input type="submit"  size="2" value="変更"  />
                  <input type="hidden" name="id"  value="{{.Id|html}}"/>
               </td>
             </form>

             <form method="GET" action="/car_delete" >

               <td>
                  <input type="hidden" name="id"  value="{{.Id|html}}"/>
                  <input type="submit"  size="2" value="削除"  />
               </td>
             </form>


           </tr>

           {{end}}
          </table>
          <table>
           <h2 align="center">Input New Data ,please</h2>
          <form method="GET" action="/car_show2" >

              <table border="2" cellpadding="8" align="center" bgcolor="#00ced1">

              <tr>   <th>car-name</th> <th>car-explain</th>  <th>access</th> </tr>

              <td > <input type="text" name="car_name" size="15" />            </td>
              <td > <input type="text" name="car_explain" size="30" />            </td>
              <td > <input type="submit" size="2"  value="登録" />           </td>
          </form>
         </table>
        </section>


        <aside>

        </aside>

        <footer>


        </footer>
     </body>
   </html>
   `
