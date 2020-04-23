package html2

   const Ai_sgh_ex_show = `
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



       </article>
       <section id="main">
           <table border="2" cellpadding="8" align="center" bgcolor="#00ced1">
           <h2 align="center">Area Information</h2>

           <tr> <th>course_no</th><th>expression</th><th>item1-name</th> <th>item1-factor</th>
                                                      <th>item2-name</th> <th>item2-factor</th> <th>access1</th> </tr>
           {{range .}}

             <tr>

             <form method="GET" action="/ai_sgh_ex_delete" >

               <td>
                  <input type="text" name="area_no" size="7" align="center" value="{{.Course_No|html}}" />
               </td>
               <td>
                  <input type="text" name="area_no" size="15" align="center" value="{{.Expression|html}}" />
               </td>
               <td>
                  <input type="text" name="area_name" size="3" align="center" value="{{.Item1_Name|html}}" />
               </td>
               <td>
                  <input type="text" name="area_detail" size="7" align="center" value="{{.Item1_Factor|html}}" />
               </td>
               <td>
                  <input type="text" name="area_name" size="3" align="center" value="{{.Item2_Name|html}}" />
               </td>
               <td>
                  <input type="text" name="area_detail" size="7" align="center" value="{{.Item2_Factor|html}}" />
               </td>
               <td>
                  <input type="hidden" name="id"  value="{{.Id|html}}"/>
                  <input type="submit"  size="2" value="delete"  />
               </td>
             </form>


           </tr>

           {{end}}
          </table>

        </section>

        <aside>

        </aside>

        <footer>


        </footer>
     </body>
   </html>
   `
