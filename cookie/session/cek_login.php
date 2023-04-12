<?php
session_start();    //Memulai Session didalam file ini

if(isset($_POST['login'])){  //Jika tombol login ditekan 
    //maka akan menjalankan script dibawahnya

    if($_POST['username'] == "tedc" && MD5($_POST['password']) == '0a9634a1c45dde7e222d381e9cbb2a32'){ //password : tedc
        // jika username dan password benar

        // sukses
        $_SESSION['login'] = true;
        $_SESSION['username'] = 'tedc';
        header('Location: /session/home.php');
        exit();
    }else{
        // gagal
        $_SESSION['error'] = "Username or Password is wrong";
        header('Location: /session/login.php');
        exit();
    }
}

?>