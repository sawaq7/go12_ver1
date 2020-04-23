package html2

   const Private_showall1 = `
   <!DOCTYPE html>
   <html>
     <head>
        <meta charset="UTF-8">
        <title>Private Information</title>
        <link rel="stylesheet" href="css/member1_show.css" type="text/css">
     </head>
     <body>
       <article>
       </article>
       <section>
            <form method="GET" action="/private_showall2" >

              <table border="2" cellpadding="8" align="center" bgcolor="#00ced1">
              <h1  align="center">Register Your deliverly-number ,please</h1>
              <tr> <th>worker-no</th> <th>worker-name</th> <th>worker-type</th> <th>worker-salary</th>
                                                                                           <th>access</th> </tr>

              <td > <input type="text" name="worker_no"  size="5" />         </td>

              <td > <input type="text" name="worker_name" size="10" />       </td>

              <td > <input type="text" name="worker_type" size="3" />        </td>

              <td > <input type="text" name="worker_salary" size="5" />      </td>

              <td > <input type="submit" size="2"  value="register" />            </td>
            </form>
       </section>
       <section id="main">


           <table border="2" cellpadding="8" align="center" bgcolor="#00ced1">
           <h2 align="center">List Of Private Information</h2>

           <tr> <th>worker-no</th> <th>worker-name</th> <th>worker-type</th> <th>worker-salary</th>
                                                        <th>hourly-pay</th> <th>access1</th> <th>access2</th>  </tr>

           {{range .}}
             <tr>

             <form method="GET" action="/private_update" >

               <td>
                  <input type="text" name="worker_no" size="5" align="center" value="{{.Worker_No|html}}" />
                  <input type="hidden" name="id"  value="{{.Id|html}}"/>
               </td>

               <td>
                  <input type="text" name="worker_name" size="10" align="center" value="{{.Worker_Name|html}}" />
               </td>

               <td>
                  <input type="text" name="worker_type" size="3" align="center" value="{{.Worker_Type|html}}" />
               </td>

               <td>
                  <input type="text" name="worker_salary" size="5" align="center" value="{{.Worker_Salary|html}}" />
               </td>
               <td>
                  <input type="text" name="hourly_pay" size="5" align="center" value="{{.Worker_H_Pay|html}}" />
               </td>
               <td>
                  <input type="submit"  size="2" value="change"  />
               </td>
             </form>

             <form method="GET" action="/private_delete" >

               <td>
                  <input type="hidden" name="id"  value="{{.Id|html}}"/>
                  <input type="submit"  size="2" value="delete"  />
               </td>
             </form>


           </tr>
           {{end}}
           </table>

        </section>
     </body>
   </html>
   `
