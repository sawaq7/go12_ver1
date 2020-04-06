package html5

   const Storage_object_list = `
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


           {{range .}}

             {{ if eq .Line_No 1}}
              <tr> <th>project-name</th> </tr>
              <tr>
               <td>
                  <input type="text" name="project_name" size="10" align="center" value="{{.Project_Name|html}}" />
                  <input type="hidden" name="line_no"  value="{{.Line_No|html}}"/>
               </td>
              </tr>
               <tr> <th>bucket-name</th> </tr>
               <tr>
               <td>
                  <input type="text" name="object_name" size="10" align="center" value="{{.Bucket_Name|html}}" />
                  <input type="hidden" name="line_no"  value="{{.Line_No|html}}"/>
               </td>
              </tr>

             {{end}}
           {{end}}

          </table>
       </article>


       <section id="main">


         <table border="2" cellpadding="8" align="center" bgcolor="#00ced1">

         <h1  align="center">Select image file which you want to upload ,please</h1>

         <tr>  <th>image-file</th>  <th>access1</th> <th>access2</th></tr>

          <form method="POST" enctype="multipart/form-data" action="/storage_object_upload" >

            <td >
              <input type="file" name="image" id="image" size="10" align="center">
            </td>

            <td >
              <input type="submit" size="2"  value="upload" />
            </td>

          </form>

          <form method="GET" action="/db_access_list" >

            <td>
              <input type="submit"  size="2" value="access list"  />
            </td>

          </form>

           <table border="2" cellpadding="8" align="center" bgcolor="#00ced1">
           <h2 align="center">Object List Of Storage</h2>

           <tr> <th>bucket-name</th>  <th>created-time</th> <th>access1</th> <th>access2</th><th>access3</th> <th>access4</th></tr>
           {{range .}}
             <tr>

             <form method="GET" action="/storage_object_show" >
               <td>
                  <input type="text" name="object_name" size="10" align="center" value="{{.Object_Name|html}}" />

                  <input type="hidden" name="line_no"  value="{{.Line_No|html}}"/>
               </td>
               <td>
                  <input type="text" name="created_time" size="15" align="center" value="{{.Created|html}}" />

                  <input type="hidden" name="line_no"  value="{{.Line_No|html}}"/>
               </td>
               <td>

                  <input type="submit"  size="2" value="object show"  />
                  <input type="hidden" name="line_no"  value="{{.Line_No|html}}"/>
               </td>

             </form>

             <form method="GET" action="/storage_object_delete" >
               <td>
                  <input type="hidden" name="line_no"  value="{{.Line_No|html}}"/>
                  <input type="submit"  size="2" value="削除"  />
               </td>
             </form>

             <form method="GET" action="/storage_object_copy_keyin" >
               <td>
                  <input type="hidden" name="line_no"  value="{{.Line_No|html}}"/>
                  <input type="submit"  size="2" value="copy"  />
               </td>
             </form>

             <form method="GET" action="/storage_object_rename_keyin" >
               <td>
                  <input type="hidden" name="line_no"  value="{{.Line_No|html}}"/>
                  <input type="submit"  size="2" value="rename"  />
               </td>
             </form>

             <form method="GET" action="/csv_show" >
          <!--   <form method="GET" action="/csv_show_test1" >  -->

               <td>
                  <input type="hidden" name="line_no"  value="{{.Line_No|html}}"/>
                  <input type="submit"  size="2" value="storage access"  />
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
