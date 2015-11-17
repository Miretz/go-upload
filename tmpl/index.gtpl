<!doctype HTML>
<html>
<head>
	<title>Upload File</title>
	<style>
	div.content {
	    display: table;
	    width: 100%;
	    border: 1px solid #ddd;
	}
	div.content header, div.content main section {
		display: table-row;
		border: 1px solid #ddd;
	}
	main section span.date {
		font-size: 70%;
		font-style: italic;
	}
	span.hd {
	    display: table-cell;
	    width: 25%;
		border: 1px solid #ddd;	
	}
	span.date,span.comment,span.file_link,span.delete_link {
	    display: table-cell;
	    width: 25%;
	    border: 0px;
	}
	main { display: table-row-group }
	</style>
	<script>
	function del(id, name){
		var r = confirm("Delete " + name + "?");
		if(r == true){
			var url = "/delete?id=" + id;
			var xhr = new XMLHttpRequest();
			xhr.open("POST", url, true);
			xhr.send();
		}
		location.reload();
	}
	</script>
</head>
<h1>Upload File</h1>
<div class="content">
<header>
	<span class="hd">Creation Date</span>
	<span class="hd">Comment</span>
	<span class="hd">Download</span>
	<span class="hd">Actions</span>
</header>
<main>
{{range .}}
	<section id="item-{{.Id}}">
		<span class="date">{{ .Create_date}}</span>
		<span class="comment">{{ .Comment}}</span>
		<span class="file_link"><a href="{{.Path}}">{{.Name}}</a></span> 
		<span class="delete_link"><button onclick="del('{{.Id}}','{{.Name}}')">Delete</button></span>
	</section> 
{{end}}
</main>
</div>
<br/>
<a href='/upload'>Upload File</a>
</html>
