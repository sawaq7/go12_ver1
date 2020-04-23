package html6

   const Payment_register = `
   <!DOCTYPE html>
   <html>
     <head>
        <meta charset="UTF-8">
        <title>uodate/delete in D.S.</title>
        <link rel="stylesheet" href="css/sgh/d_district_area.css" type="text/css">
     </head>
     <body>
       <header>



       </header>
       <nav>


	   </nav>
       <article>
         <table border="2" cellpadding="12" align="center" bgcolor="#00ced1">
         <h2 align="center">Current Guest</h2>

           <tr> <th>guest-no</th> <th>guest-name</th> </tr>
           {{range .}}

            {{ if eq .Line_No 1}}

             <tr>

             <form method="GET" action="/d_district_area_update" >


               <td>
                  <input type="text" name="district_no" size="5" align="center" value="{{.Guest_No|html}}" />
                  <input type="hidden" name="id"  value="{{.Id|html}}"/>

               </td>
               <td>
                  <input type="text" name="district_name" size="10" align="center" value="{{.Guest_Name|html}}" />
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

           <tr> <th>payment-date</th> <th>item</th><th>amount</th> <th>access1</th> <th>access2</th> </tr>
           {{range .}}

             <tr>

             <form method="GET" action="/d_district_area_update" >

               <td>
                   <input type="text" name="date" size="12" value="{{.Date |html}}" />
              </td>

               <td>
                  <input type="text" name="item" size="10" align="center" value="{{.Item|html}}" />
               </td>
               <td>
                  <input type="text" name="amount" size="10" align="center" value="{{.Amount|html}}" />
               </td>

               <td>
                  <input type="submit"  size="2" value="change"  />
               </td>
             </form>

             <form method="GET" action="/payment_delete" >

               <td>
                  <input type="hidden" name="id"  value="{{.Id|html}}"/>
                  <input type="submit"  size="2" value="delete"  />
               </td>
             </form>

           </tr>

           {{end}}
          </table>
          <table>
           <h2 align="center">Input New Data ,please</h2>
          <form method="GET" action="/payment_register_excute" >

              <table border="2" cellpadding="8" align="center" bgcolor="#00ced1">

              <tr> <th>payment-date</th> <th>item</th><th>amount</th> <th>access</th> </tr>

              <td>
                   <input type="text" name="date" size="12"  />
              </td>

              <td >
                    <input type="text" name="item" size="10" />
              </td>

              <td >
                     <input type="text" name="amount" size="10" />
              </td>

              <td >
                     <input type="submit" size="2"  value="register" />
              </td>

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
