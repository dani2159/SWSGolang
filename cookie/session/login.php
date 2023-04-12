<html>

<body>
    <?php 
    session_start();
    if($_SESSION['login'] == true){     //Jika session login bernilai true
        //maka akan diarahkan ke halaman home.php
        header('Location: /session/home.php');
        exit();
    }

    if(isset($_SESSION['error'])){ ?>
    <h2><?= $_SESSION['error'] ?></h2>
    <?php } ?>
    <h1>Login</h1>
    <form action="/session/cek_login.php" method="post">
        <label>Username :
            <input type="text" name="username">
        </label>
        <br />
        <label>Password :
            <input type="password" name="password">
        </label>
        <br />
        <button type="submit" name="login">Login</button>
    </form>
</body>

</html>