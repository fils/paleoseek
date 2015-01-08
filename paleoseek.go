package paleoseek

import (
	"fmt"
	"net/http"
	//"net/url"
)

const page = `
<!DOCTYPE html>
<html>
<head>
    <title>paleoseek.net</title>
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
<link type="application/opensearchdescription+xml" rel="search" title="paleoseek" href="http://paleoseek.net/opensearch.xml" />
     <link rel="stylesheet" href="./static/customBootstrap.css">

<!-- jQuery -->
    <script type="text/javascript" charset="utf8" src="http://ajax.aspnetcdn.com/ajax/jQuery/jquery-1.8.2.min.js"></script>

    <!-- Latest compiled and minified JavaScript -->
    <script src="http://netdna.bootstrapcdn.com/bootstrap/3.0.2/js/bootstrap.min.js"></script>

    <style type="text/css">
    body {
        padding-top: 50px;
    }
    .starter-template {
        padding: 40px 15px;
        text-align: center;
    }
    </style>

    <script>
    var renderSearchElement = function() {
        google.search.cse.element.render({
            div: "default",
            tag: 'search'
        });
        google.search.cse.element.render({
            div: "test",
            attributes: {
                disableWebSearch: false,
                enableHistory: true
            },
            tag: 'search'
        });
    };
    var myCallback = function() {
        if (document.readyState == 'complete') {
            renderSearchElement();
        } else {
            google.setOnLoadCallback(renderSearchElement, true);
        }
    };

    // Insert it before the CSE code snippet so that cse.js can take the script
    // parameters, like parsetags, callbacks.
    window.__gcse = {
        parsetags: 'explicit',
        callback: myCallback
    };

    (function() {
        var cx = '010146164659624730016:d6ieoekzj1u';
        var gcse = document.createElement('script');
        gcse.type = 'text/javascript';
        gcse.async = true;
        gcse.src = (document.location.protocol == 'https:' ? 'https:' : 'http:') +
            '//www.google.com/cse/cse.js?cx=' + cx;
        var s = document.getElementsByTagName('script')[0];
        s.parentNode.insertBefore(gcse, s);
    })();
    </script>

 <script>
  (function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){
  (i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),
  m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)
  })(window,document,'script','//www.google-analytics.com/analytics.js','ga');

  ga('create', 'UA-76153-8', 'auto');
  ga('send', 'pageview');

</script>
<!--
    <script type="application/ld+json">
    {
        "@context": "http://schema.org",
        "@type": "WebSite",
        "name": "paleo-search.appspot.com",
        "potentialAction": {
            "@type": "SearchAction",
            "target": "http://paleo-search.appspot.com/#gsc.q={q}",
            "query-input": "required maxlength=500 name=gsc.q"
        }
    }
    </script>
    -->
</head>
<body itemscope itemtype="http://schema.org/WebSite">
    <meta itemprop="url" content="http://paleoseek.net" />


    <div class="navbar navbar-default navbar-fixed-top" role="navigation">
        <div class="container">
            <div class="navbar-header">
                <button type="button" class="navbar-toggle" data-toggle="collapse" data-target=".navbar-collapse">
                    <span class="sr-only">Toggle navigation</span>
                    <span class="icon-bar"></span>
                    <span class="icon-bar"></span>
                    <span class="icon-bar"></span>
                </button>

                <a class="navbar-brand"  itemprop="url" href="http://paleoseek.net">paleoseek.net</a>

            </div>
            <div class="collapse navbar-collapse">
                <ul class="nav navbar-nav">
                    <li class="active"><a href="http://paleoseek.net">Home</a>
                    </li>
                    <li><a href="/data">Search Data Sets</a>
                    </li>
                    <li><a href="/static/details.html">About, Usage, History</a>
                    </li>
                    <li><a href="/static/resources.html">Resource Files</a>
                    </li>
                </ul>
            </div>
            <!--/.nav-collapse -->
        </div>
    </div>

   <div class="container">
<div class="row">
	<div class="col-md-8 col-md-offset-2">  <h1> <p class="text-center">paleoseek.net </p></h1> </div>

</div>
<div class="row">

<p class="text-center">

            <span style='color:green'>Searching 600+ paleogeoscience related sites and growing</span>
    <br/>
     <small>examples: paleoclimate eocene , calcium carbonate deposition , U-Pb radiometric dating</small>
    
    </p>
</div>
<div class="row">
  <div class="col-md-8 col-md-offset-2"  itemprop="potentialAction" itemscope itemtype="http://schema.org/SearchAction">
	  <meta itemprop="target" content="http://paleoseek.net/#gsc.q={q}"/>
      <meta itemprop="query-input" content="required maxlength=500 name=gsc.q"/>
    <div id="test"></div>
	</div>

</div>



    </div>

    </body>
</html >
`

const sbpage = `
<html>
  <head>
    <title>JSON/Atom Custom Search API Example</title>
  </head>
  <body>
  <p>test</p>
    <div id="content"></div>
    <script>
      function hndlr(response) {
      for (var i = 0; i < response.items.length; i++) {
        var item = response.items[i];
        // in production code, item.htmlTitle should have the HTML entities escaped.
        document.getElementById("content").innerHTML += "<br>" + item.htmlTitle;
      }
    }
    </script>
    <script src="https://www.googleapis.com/customsearch/v1?key=AIzaSyCdW90P-cnp9yk6ILeQU_ypLNSuhEqu3r0&amp;cx=010146164659624730016:d6ieoekzj1u&amp;q=eocene&amp;callback=hndlr">
    </script>
  </body>
</html>
`

func init() {
	http.HandleFunc("/", root)
	http.HandleFunc("/sandbox", sandbox)
}

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, page)
}

func sandbox(w http.ResponseWriter, r *http.Request) {
	// read in the URL and parse out the q= value
	// spit that back in prelude to creating a new URL object and forwarding the user to that with
	// the JSON call to the GCSE

	//qmap, _ := url.ParseQuery(r.URL.RawQuery)
	//querystring := qmap["q"][0]

	// if no q then load the page to accept a query...
	// if there is a q, then make the call the load the results...  and ready for new calls....

	// fmt.Fprint(w, "Forward to %s with extension query %s", r.URL.String(), querystring)
	fmt.Fprint(w, sbpage)
}
