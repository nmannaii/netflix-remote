const {networkInterfaces} = require('os');

module.exports = () => {
    const wifi = networkInterfaces()['Wi-Fi'];
    if (wifi) {
        return wifi.filter(u => u.family === 'IPv4')[0]?.address;
    }

    return null;
}
