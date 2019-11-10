<html>

<head>
    <title>Login</title>
</head>

<body>
    <form action="/login" method="post">
        Username: <input type="text" name="username"><br>
        Password: <input type="password" name="password"><br>
        age: <input type="text" name="age"><br>
        email: <input type="text" name="email"><br>        
        <select name="fruit">
            <option value="apple">Apple</option>
            <option value="banana">Banana</option>
            <option value="pear">Pear</option>
        </select>
        <input type="radio" name="gender" value="1">Male
        <input type="radio" name="gender" value="2">Female
        <input type="submit" value="Login">
    </form>
</body>

</html>