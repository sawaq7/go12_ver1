package html2

   const Deliver_showcar = `
   <!DOCTYPE html>
   <html>
     <head>
        <meta charset="UTF-8">
        <title>car-list</title>
        <link rel="stylesheet" href="css/member1_show.css" type="text/css">
     </head>
     <body>
       <article>
       </article>
       <section id="main">


           <table border="2" cellpadding="8" align="center" bgcolor="#00ced1">
           <h2 align="center">List Of Deliverly Situation</h2>


           <tr> <th>car-no</th> <th>date</th> <th>deliverly-area</th>   <th>private-no</th> <th>deliverly-no</th>   </tr>
           {{range .}}
             <tr>

             <form method="GET"  >
               <td>
                  <input type="text" name="car_no" size="5" align="center" value="{{.Car_No |html}}" />
               </td>
               <td>
                  <input type="text" name="date" size="12" value="{{.Date |html}}" />
               </td>

               <td>
                  <input type="text" name="area_name" size="5" align="center" value="{{.Area_Name|html}}" />
                  <input type="hidden" name="id"  value="{{.Id|html}}"/>
               </td>

               <td>
                  <input type="text" name="private_no" size="5" align="center" value="{{.Private_No |html}}" />
               </td>

               <td>
                  <input type="text" name="number" size="2" align="center" value="{{.Number |html}}" />
               </td>

             </form>

           </tr>
           {{end}}
           </table>

        </section>
     </body>
   </html>
   `
