<html>
<head>
<title>登录</title>
</head>
<body>
<form action="/login" method="post">
    用户名:<input type="text" name="username"><br/><br/>
    密码:<input type="password" name="password"><br/><br/>
        <input type="hidden" name="token" value="{{.}}">
    <input type="submit" value="登录">
</form>
</body>
</html>