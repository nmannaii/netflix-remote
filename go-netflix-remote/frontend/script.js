import { GetIpAddressQrCode, GetLocalIpAddress } from './wailsjs/go/main/App';

document.querySelector('#host').textContent = `http://${await GetLocalIpAddress()}:3698`;

document.querySelector('#qrCode').setAttribute('src', await GetIpAddressQrCode())

document.addEventListener("DOMContentLoaded", function() {
    alert('VISIBLE')
  });