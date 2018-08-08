function main()
{

    console.log("here");

    var qrcode = new QRCode(document.getElementById("qrcode"), {
        text: "void",
        width: 400,
        height: 400,
        colorDark : "#000000",
        colorLight : "#ffffff",
        correctLevel : QRCode.CorrectLevel.H
    });

    var bc = new BroadcastChannel('qrcode_channel');
    bc.onmessage = function (ev) {
        console.log(ev); 
        qrcode.makeCode(JSON.stringify(ev.data)) 
    } /* receive */
}