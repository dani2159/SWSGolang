<form method="POST" action="">
    <label>Siapa Nama Anda : </label>
    <input type="text" name="nama" placeholder="Masukan Nama Anda">
    <br>
    <label>Masukan E-Mail : </label>
    <input type="email" name="email" placeholder="Masukan E-Mail">
    <button type="Submit">Kirim</button>
</form>

<?php

if (isset($_POST['nama'])) {
    $nama = htmlspecialchars($_POST['nama']);
    echo "Nama Anda : " . $nama. "<br>";
    echo "E-Mail Anda : " . $_POST['email'];
}
?>