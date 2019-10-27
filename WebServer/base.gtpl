<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <tittle>{{.Title}}</tiitle>
    </br>
</head>
<body>
    </br>
    {{template "Header"}}
    </br>
    {{range .Items}}
        <div>
            {{ . }}
        </div>
    {{else}}
        <div>
            <strong>Tidak Ada Menu</strong>
        </div>
    {{end}}
    </br>
    {{template "Body"}}
    </br>
    {{template "Footer"}}
</body>
</html>