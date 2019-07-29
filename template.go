var htmltemplate=`<!DOCTYPE html>
<html lang="en">

<head>
	<meta charset="utf-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<link rel="stylesheet" media="screen" href="../assets/css/bootstrap.css">
	<link rel="stylesheet" media="screen" href="../assets/css/dataTables.bootstrap.css">
	<link rel="stylesheet" media="screen" href="../assets/css/dataTables.tableTools.css">
	<link rel="stylesheet" media="screen" href="../assets/css/style.css">

	<title>Simple CRUD with Golang</title>

	<!--  -->

	<style>
	</style>

</head>

<body>
   <nav class="navbar navbar-default">
    <div class="container-fluid">
        <div class="navbar-header">
            <button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target="#bs-example-navbar-collapse-1" aria-expanded="false">
                <span class="sr-only">Toggle navigation</span>
                <span class="icon-bar"></span>
                <span class="icon-bar"></span>
                <span class="icon-bar"></span>
            </button>
            <a class="navbar-brand" href="#">Simple CRUD with Golang</a>
        </div>

        <div class="collapse navbar-collapse" id="bs-example-navbar-collapse-1">
            <ul class="nav navbar-nav navbar-right">
                <li><a href="home">Home</a></li>
            </ul>
        </div>
    </div>
    </nav>

	<div class="container">
<div class="row">
    <div class="col-md-4 col-sm-4 col-xs-4">
       <div class="panel panel-default">
            <div class="panel-heading">Form User</div>
            <div class="panel-body">
                <form method="POST" action="/update/" accept-charset="UTF-8" class="test-form">
                    {{range .Userdata}}
                    <input name="id" type="hidden" value="{{.ID}}">
                        <div class="form-group">
                            <label for="name">Username</label>
                            <input required="required" class="form-control" placeholder="Username" name="name" type="text" id="username" value="{{.Name}}">
                        </div>
                        <div class="form-group">
                            <label for="email">Email</label>
                            <input required="required" class="form-control" placeholder="Email" name="email" type="text" id="email" value="{{.Email}}">
                        </div>
                        <div class="form-group">
                            <label for="Address">Address</label>
                            <textarea required="required" class="form-control" placeholder="Address" name="address">{{.Address}}</textarea>
                        </div>
                        <div class="form-group">
                            <input class="btn btn-primary" type="submit" value="Update User">
                        </div>
                    {{end}}
                </form>
            </div>
        </div>
    </div>
    
    <div class="col-md-8 col-sm-8 col-xs-8">

        <div class="panel panel-default">
            <div class="panel-heading">List User</div>
            <div class="panel-body table-responsive">
                <table class="table table-striped table-bordered datatable" cellspacing="0" width="100%">
                    <thead>
                        <tr>
                            <th>No</th>
                            <th>Username</th>
                            <th>Email</th>
                            <th width="100px">Address</th>
                            <th>Action</th>
                        </tr>
                    </thead>

                    <tbody>
                        {{range .Datasql}}
                            <tr>
                                <td>{{.Number}}</td>
                                <td>{{.Name}}</td>
                                <td>{{.Email}}</td>
                                <td>{{.Address}}</td>
                                <td><a href="/edit/{{.ID}}">Edit</a> | <a href="/delete/{{.ID}}" onclick="return confirm('Are You Sure Want To Delete ?')">Delete</a></td>
                            </tr>
                        {{end}}
                    </tbody>
                </table>
            </div>
        </div>
    </div>
</div>
    
</div>
<footer class="footer">
    <div class="container">
        <p class="muted credit">Simple CRUD with Golang By <a href="https://github.com/Aris-haryanto/Simple-CRUD-Web-Applications-With-Golang-Mysql">Aris Haryanto</a></p>
    </div>
</footer>    
    <script type="text/javascript" src="../assets/js/jquery-2.1.3.min.js"></script>
	<script type="text/javascript" src="../assets/js/bootstrap.min.js"></script>
	<script type="text/javascript" src="../assets/js/jquery.dataTables.js"></script>
    <script type="text/javascript" src="../assets/js/main.js"></script>
</body>

</html>

<!DOCTYPE html>
<html lang="en">

<head>
	<meta charset="utf-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<link rel="stylesheet" media="screen" href="../assets/css/bootstrap.css">
	<link rel="stylesheet" media="screen" href="../assets/css/dataTables.bootstrap.css">
	<link rel="stylesheet" media="screen" href="../assets/css/dataTables.tableTools.css">
	<link rel="stylesheet" media="screen" href="../assets/css/style.css">

	<title>Simple CRUD with Golang</title>

	<!--  -->

	<style>
	</style>

</head>

<body>
   <nav class="navbar navbar-default">
    <div class="container-fluid">
        <div class="navbar-header">
            <button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target="#bs-example-navbar-collapse-1" aria-expanded="false">
                <span class="sr-only">Toggle navigation</span>
                <span class="icon-bar"></span>
                <span class="icon-bar"></span>
                <span class="icon-bar"></span>
            </button>
            <a class="navbar-brand" href="#">Simple CRUD with Golang</a>
        </div>

        <div class="collapse navbar-collapse" id="bs-example-navbar-collapse-1">
            <ul class="nav navbar-nav navbar-right">
                <li><a href="home">Home</a></li>
            </ul>
        </div>
    </div>
    </nav>

	<div class="container">
<div class="row">
    <div class="col-md-4 col-sm-4 col-xs-4">
       <div class="panel panel-default">
            <div class="panel-heading">Form User</div>
            <div class="panel-body">
                <form method="POST" action="/save/" accept-charset="UTF-8" class="test-form">

                    <div class="form-group">
                        <label for="name">Username</label>
                        <input required="required" class="form-control" placeholder="Username" name="name" type="text" id="username">
                    </div>
                    <div class="form-group">
                        <label for="email">Email</label>
                        <input required="required" class="form-control" placeholder="Email" name="email" type="text" id="email">
                    </div>
                    <div class="form-group">
                        <label for="Address">Address</label>
                        <textarea required="required" class="form-control" placeholder="Address" name="address"></textarea>
                    </div>
                    <div class="form-group">
                        <input class="btn btn-primary" type="submit" value="Insert User">
                    </div>
                </form>
            </div>
        </div>
    </div>
    
    <div class="col-md-8 col-sm-8 col-xs-8">

        <div class="panel panel-default">
            <div class="panel-heading">List User</div>
            <div class="panel-body table-responsive">
                <table class="table table-striped table-bordered datatable" cellspacing="0" width="100%">
                    <thead>
                        <tr>
                            <th>No</th>
                            <th>Username</th>
                            <th>Email</th>
                            <th width="100px">Address</th>
                            <th>Action</th>
                        </tr>
                    </thead>

                    <tbody>
                    	{{range .Datasql}}
                            <tr>
                                <td>{{.Number}}</td>
                                <td>{{.Name}}</td>
                                <td>{{.Email}}</td>
                                <td>{{.Address}}</td>
                                <td><a href="/edit/{{.ID}}">Edit</a> | <a href="/delete/{{.ID}}" onclick="return confirm('Are You Sure Want To Delete ?')">Delete</a></td>
                            </tr>
                        {{end}}
                    </tbody>
                </table>
            </div>
        </div>
    </div>
</div>
    
</div>
<footer class="footer">
    <div class="container">
        <p class="muted credit">Simple CRUD with Golang By <a href="https://github.com/Aris-haryanto/Simple-CRUD-Web-Applications-With-Golang-Mysql">Aris Haryanto</a></p>
    </div>
</footer>    
    <script type="text/javascript" src="../assets/js/jquery-2.1.3.min.js"></script>
	<script type="text/javascript" src="../assets/js/bootstrap.min.js"></script>
	<script type="text/javascript" src="../assets/js/jquery.dataTables.js"></script>
    <script type="text/javascript" src="../assets/js/main.js"></script>
</body>

</html>
`