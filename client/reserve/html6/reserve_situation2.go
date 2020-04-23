package html6

   const Reserve_situation2 = `
   <!DOCTYPE html>
   <html>
     <head>
        <meta charset="UTF-8">
        <title>uodate/delete in D.S.</title>
        <link rel="stylesheet" href="css/medical/reserve_situation.css" type="text/css">
     </head>
     <body>
       <header>



       </header>
       <nav>


	   </nav>
       <article>
         <table border="2" cellpadding="12" align="center" bgcolor="#00ced1">
         <h2 align="center">Current Guest</h2>

           <tr> <th>reserve-date</th> </tr>
           {{range .}}

            {{ if eq .Line_No 1}}

             <tr>

             <form method="GET" action="/d_district_area_update" >

               <td>
                  <input type="text" name="district_no" size="5" align="center" value="{{.Date|html}}" />
                  <input type="hidden" name="id"  value="{{.Id|html}}"/>

               </td>

              </form>

            </tr>

            {{end}}

           {{end}}
           </table>


       </article>
       <section id="main">
           <table border="2" cellpadding="8" align="center" bgcolor="#00ced1">
           <h2 align="center">Reserve Information</h2>

           <tr> <th>guest-name</th> <th>start-time</th><th>end-time</th> </tr>
           {{range .}}

             <tr>

             <form method="GET" action="/d_district_area_update" >

               <td>
                  <input type="text" name="guest_name" size="10" align="center" value="{{.Guest_Name|html}}" />
               </td>

               <td>
                  <input type="text" name="start_time" size="10" align="center" value="{{.Start_Time|html}}" />
               </td>
               <td>
                  <input type="text" name="end_time" size="10" align="center" value="{{.End_Time|html}}" />
               </td>

             </form>

           </tr>

           {{end}}
          </table>

        </section>


        <aside>

          <table border="2" cellpadding="8" align="center" bgcolor="#00ced1">
          <h2 align="center">Reserve Information2</h2>

            {{range .}}

            {{ if eq .Line_No 1}}

             <tr>
              <form  >
               <td>
                 <img src="{{.Url}}" width="400" height="400"      name="List" alt="png"/>
               </td>
             </form>

            </tr>


            {{end}}

           {{end}}

          </table>


        </aside>

        <footer>


        </footer>
     </body>
   </html>
   `
