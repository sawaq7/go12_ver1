package html5

   const Storage_bucket_list = `
   <!DOCTYPE html>
   <html>
     <head>
        <meta charset="UTF-8">
        <title>datastore's update/delete</title>

        <link rel="stylesheet" href="css/sgh/deliver_showall1.css" type="text/css">
     </head>
     <body>
       <header>


       </header>
       <nav>

       </nav>
       <article>
          <table border="2" cellpadding="8" align="center" bgcolor="#00ced1">
           <h2 align="center">Current Project-Name</h2>

           <tr> <th>project-name</th> </tr>
           {{range .}}

             {{ if eq .Line_No 1}}


               <tr>

               <td>
                  <input type="text" name="project_name" size="10" align="center" value="{{.Project_Name|html}}" />
                  <input type="hidden" name="line_no"  value="{{.Line_No|html}}"/>
               </td>

               </tr>

             {{end}}
           {{end}}

           </table>


       </article>

       <section id="main">

           <table border="2" cellpadding="8" align="center" bgcolor="#00ced1">
           <h2 align="center">List Of Deliverly Situation</h2>

           <tr> <th>bucket-name</th>  <th>access1</th> <th>access2</th><th>access3</th></tr>
           {{range .}}
             <tr>

             <form method="GET" action="/storage_object_list" >
               <td>
                  <input type="text" name="bucket_name" size="10" align="center" value="{{.Bucket_Name|html}}" />

                  <input type="hidden" name="line_no"  value="{{.Line_No|html}}"/>
               </td>
               <td>

                  <input type="submit"  size="2" value="object list"  />
                  <input type="hidden" name="line_no"  value="{{.Line_No|html}}"/>
               </td>

             </form>

             <form method="GET"  >
               <td>
                  <input type="hidden" name="line_no"  value="{{.Line_No|html}}"/>
                  <input type="submit"  size="2" value="削除"  />
               </td>
             </form>
             <form method="GET"  >
               <td>
                  <input type="hidden" name="line_no"  value="{{.Line_No|html}}"/>
                  <input type="submit"  size="2" value="copy"  />
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
