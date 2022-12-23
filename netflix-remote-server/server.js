const express = require('express');
const bodyParse = require('body-parser');
const cors = require('cors');
const bot = require('robotjs')
const app = express();
const ip = require('ip');
const chalk = require('chalk');
const boxen = require('boxen');
const banner = require('./consts/banner');
const port = 9876;
app.use(cors("*"));
app.use(bodyParse.urlencoded({extended: false}));
app.use(bodyParse.json());

app.use('/', express.static('statics/netflix-remote'));

app.post('/api/keyboard-press', (req, res) => {
    const key = req.body.key;
    bot.keyTap(key);
    res.json({success: true});
});


app.listen(port, () => {
    console.log(chalk.red(banner));
    console.log(
        boxen('Welcome to Netflix remote control.\n' +
            'An opensource free , simple and user friendly app, that you can use it on IOS or Android.' +
            ' You need only three things:\n' +
            '- Your favorite web browser (preferred Chrome)\n' +
            '- Netflix navigator plugin: https://www.netflixnavigator.com/ to navigate using arrows\n' +
            '- Being connected to the same network',
            {borderStyle: 'double'}
        ))
    console.log(`Now open your browser and access this url: http://${ip.address()}:${port}`);
});
