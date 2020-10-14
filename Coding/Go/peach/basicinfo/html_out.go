package basicinfo

import (
	"html/template"
	"os"
)

/*type htmlFileOutput struct {
	CommandLine string
	Time        string
	Keys        []string
	Results     []Result
}*/

const (
	htmlTemplate = `
<!DOCTYPE html>
<html>
  <head>
    <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
    <meta
      name="viewport"
      content="width=device-width, initial-scale=1, maximum-scale=1.0"
    />
    <title>Scan Report - </title>

    <!-- CSS  -->
    <link
      href="https://fonts.googleapis.com/icon?family=Material+Icons"
      rel="stylesheet"
    />
    <link
      rel="stylesheet"
      href="https://cdnjs.cloudflare.com/ajax/libs/materialize/1.0.0/css/materialize.min.css"
	/>
	<link 
	  rel="stylesheet" 
	  type="text/css" 
	  href="https://cdn.datatables.net/1.10.20/css/jquery.dataTables.css"
	/>
  
  </head>

  <body>
    <nav>
      <div class="nav-wrapper">
        <a href="#" class="brand-logo">Scan</a>
        <ul id="nav-mobile" class="right hide-on-med-and-down">
        </ul>
      </div>
    </nav>

    <main class="section no-pad-bot" id="index-banner">
      <div class="container">
        <br /><br />
        <h1 class="header center ">Scan Report</h1>
        <div class="row center">

   <table id="report">
        <thead>
        <div style="display:none">
|result_raw|StatusCode|port|URL|Server|Ip|Location|Title|
        </div>
          <tr>
              <th>Status</th>
			  <th>port</th>
			  <th>url</th>
			  <th>Server</th>
              <th>Title</th>
              <th>Location</th>
              <th>ip</th>
          </tr>
        </thead>
        <tbody>
			{{range .Infos}}
                <div style="display:none">
					|result_raw|{{ .Statuscode }}|{{ .Port }}|{{.Url }}|{{ .Server }}|{{ .Ip }}|{{ .Location }}|{{ .Title }}|
                </div>
                <tr>
					<td>{{ .Statuscode }}</td>
                    <td>{{ .Port }}</td>
                    <td><a href="{{ .Url }}">{{ .Url }}</a></td>
                    <td>{{ .Server }}</td>
                    <td>{{ .Title }}</td>
                    <td><a href="{{ .Location }}">{{ .Location }}</a></td>
					<td>{{ .Ip }}</td>
                </tr>
            {{ end }}
        </tbody>
      </table>

        </div>
        <br /><br />
      </div>
    </main>

    <!--JavaScript at end of body for optimized loading-->
	<script src="https://code.jquery.com/jquery-3.4.1.min.js" integrity="sha256-CSXorXvZcTkaix6Yvo6HppcZGetbYMGWSFlBw8HfCJo=" crossorigin="anonymous"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/materialize/1.0.0/js/materialize.min.js"></script>
    <script type="text/javascript" charset="utf8" src="https://cdn.datatables.net/1.10.20/js/jquery.dataTables.js"></script>
    <script>
    $(document).ready(function() {
        $('#report').DataTable(
            {
                "aLengthMenu": [
                    [250, 250, 1000, 1000, 500,1000,all],
                    [250, 250, 1000, 1000, 500,1000,all],
                ]
            }
        )
        $('select').formSelect();
        });
    </script>
    <style>
      body {
        display: flex;
        min-height: 100vh;
        flex-direction: column;
      }

      main {
        flex: 1 0 auto;
      }
    </style>
  </body>
</html>

	`
)

// colorizeResults returns a new slice with HTMLColor attribute


func writeHTML(results saveres,filename string) error {
	outHTML := results
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	templateName := "output.html"
	t := template.New(templateName).Delims("{{", "}}")
	_, err = t.Parse(htmlTemplate)
	if err != nil {
		return err
	}
	err = t.Execute(f, outHTML)
	return err
}