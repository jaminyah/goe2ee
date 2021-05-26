let rsaPubKey = '';


$(document).ready(function() {
    getPublicKey();
});

function getPublicKey() {
    console.log("Public Key");

    const url = 'api/public';

    fetch(url) 
        .then(function(response) {
            console.log("Response");
            return response.json();
        })
        .then(function(data) {
            console.log(data.pubkey)
            console.log(data.msg);
            rsaPubKey = data.pubkey;
        })
        .catch(function(error) {
            console.log(error);
        });
}

function publicKeyEncrypt() {

    let encoder = new JSEncrypt({
        default_key_size: 2048
        //default_public_exponent:"010001"
    });
 
    //rsaPubKey = "-----BEGIN PUBLIC KEY-----" + rsaPubKey + "-----END PUBLIC KEY-----";
    encoder.setPublicKey(rsaPubKey);

    let encoded = encoder.encrypt($('#input-text').val());
    document.getElementById("encoding").value = encoded;
}