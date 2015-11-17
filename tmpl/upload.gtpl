<!doctype HTML>
<h1>File Upload</h1>
<form enctype="multipart/form-data" action="/upload" method="post">
      <input type="text" name="comment" placeholder="File Comment" />
      <input type="file" name="uploadfile" />
	  <input type="hidden" name="token" value="{{.}}"/>
      <input type="submit" value="upload" />
</form>
