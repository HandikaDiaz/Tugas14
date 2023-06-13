// console.log("Hello Student");
// alert("You Are So Pretty");
// document.write("I Love You");

// variable

// var
// let
// const

// var piring = "Nasi Goreng";
// var piring = "Telur Ayam";
// console.log(piring);

// let gelas = "Susu";
// gelas = "Air putih";
// console.log(gelas);

// const pisau = "Sendok";
// // pisau = "Garpu";
// console.log(pisau);

// gender = "Laki-laki";
// console.log(gender);

// // data type
// let nama = "Handika"; //string
// let umur = 15; //number
// let isMarried = false; //boolean

// // nama saya Handika umur saya 15 tahun
// console.log("nama saya Handika umur saya 15 tahun");
// console.log(`nama saya ${nama} umur saya ${umur} tahun`);
// console.log("nama saya", nama, "umur saya", umur, "tahun");
// console.log("nama saya " + nama + " umur saya " + umur + " tahun")

// // operator
// let x = 48
// let y = 4

// let result = x / y

// console.log(result);

// condition

// let nilai = 77;
// if (nilai >= 75) {
//     console.log("Kamu Lulus!");
// } else {
//     console.log("Kamu Tidak Lulus!")
// }

// // funcition
// function Hello() {
//     let x = 5;
//     let y = 10;

//     let result = x * y;
//     console.log(result);
// }

// Hello();

// function Hello2(x, y) {
//     console.log(x);
//     console.log(y);

//     let result = x * y;
//     console.log(result);
// }

// Hello2(5, 10);

// // camlecase = namaSayaAdalah

// function namaSaya(name) {
//     console.log(name);
// }

// namaSaya("Handika");



function submitData() {
    let name = document.getElementById("input-name").value;
    let email = document.getElementById("input-email").value;
    let number = document.getElementById("input-number").value;
    let subject = document.getElementById("input-subject").value;
    let message = document.getElementById("input-message").value;

    if (name == "") {
        return alert ("Nama Harus diisi")
    } else if(email == "") {
        return alert ("Email harus diisi")
    } else if (number == "") {
        return alert ("Number harus diisi")
    } else if (subject == "") {
        return alert ("Subject harus dipilih")
    } else if (message == "") {
        return alert ("Message harus diisi")
    }

    let emailReceiver = "asd@gmail.com"

    

    let a = document.createElement("a")
        a.href = 'mailto:${emailReceiver}?subject=${subject}&body=Halo, nama saya ${name}, ${message}, silahkan kontak saya di nomor ${number}, terima kasi.';
        a.click();

    console.log(name);
    console.log(email);
    console.log(number);
    console.log(subject);
    console.log(message);
}

