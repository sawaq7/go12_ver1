package html4

   const Water_slope_show = `
   <!DOCTYPE html>
   <html>
     <head>
        <meta charset="UTF-8">
        <title>update/delete Datastore</title>
        <link rel="stylesheet" href="css/member1_show.css" type="text/css">
     </head>
     <body>
       <article>
       </article>

          <section id="main">

           <table border="2" cellpadding="8" align="center" bgcolor="#00ced1">
           <h2 align="center">water-slope list</h2>

           <tr> <th>file-name</th> <th>access1</th> <th>access2</th> </tr>
           {{range .}}

             <tr>

             <form method="GET" >

               <td>
                  <input type="text" name="basic_name" size="15" align="center" value="{{.File_Name|html}}" />
               </td>


             </form>

             <form method="GET" action="/water_slope_delete" >

               <td>
                  <input type="hidden" name="id"  value="{{.Id|html}}"/>
                  <input type="submit"  size="2" value="delete"  />
               </td>
             </form>

             <form method="GET" action="/water_slope_graf" >

               <td>
                  <input type="hidden" name="id"  value="{{.Id|html}}"/>
                  <input type="submit"  size="2" value="show the graf"  />
               </td>
             </form>


           </tr>

           {{end}}
           </table>

          </section>



     </body>
   </html>
   `
