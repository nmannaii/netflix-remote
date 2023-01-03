if (require('electron-squirrel-startup')) return;

const express = require('express');
const bodyParse = require('body-parser');
const cors = require('cors');
const bot = require('robotjs')
const path = require('path');
const app = express();
const {app: electronApp, BrowserWindow, Tray, Menu} = require('electron');
let tray = null;
let mainWindow = null;
const port = 9876;

app.use(cors("*"));
app.use(bodyParse.urlencoded({extended: false}));
app.use(bodyParse.json());

app.use('/', express.static(path.join(__dirname, 'statics/netflix-remote')));


app.post('/api/keyboard-press', (req, res) => {
    const key = req.body.key;
    bot.keyTap(key);
    res.json({success: true});
});


app.listen(port, () => {
    console.log('SERVER IS UP!');
});


electronApp.whenReady().then(async () => {
    mainWindow = new BrowserWindow({
        width: 600,
        height: 600,
        icon: 'remote-ctrl-icon.png',
        resizable: false,
        autoHideMenuBar: true,
        webPreferences: {
            nodeIntegration: true,
            contextIsolation: false
        }
    });
    await mainWindow.loadFile('welcome-page/index.html');
    mainWindow.on('minimize', (ev) => {
        ev.preventDefault();
        mainWindow.hide();
        tray = createTray();
    });

    mainWindow.on('restore', function (event) {
        mainWindow.show();
        tray.destroy();
    });
})

function createTray() {
    let appIcon = new Tray(path.join(__dirname, 'remote-ctrl-icon.png'));
    const contextMenu = Menu.buildFromTemplate([
        {
            label: 'Show', click: function () {
                mainWindow.show();
            }
        },
        {
            label: 'Exit', click: function () {
                electronApp.isQuiting = true;
                electronApp.quit();
            }
        }
    ]);

    appIcon.on('double-click', function (event) {
        mainWindow.show();
    });
    appIcon.setToolTip('Netflix remote');
    appIcon.setContextMenu(contextMenu);
    return appIcon;
}
