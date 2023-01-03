module.exports = {
    packagerConfig: {
        icon: './icons/win&mac/icon'
    },
    rebuildConfig: {},
    makers: [
        {
            name: '@electron-forge/maker-squirrel',
            config: {
                setupIcon: './icons/win&mac/icon.ico',
                loadingGif: './remote-ctrl-icon.png',
            },
        },
        {
            name: '@electron-forge/maker-zip',
            platforms: ['darwin'],
        },
        {
            name: '@electron-forge/maker-deb',
            config: {},
        },
        {
            name: '@electron-forge/maker-rpm',
            config: {},
        },
    ],
    publishers: [
        {
            name: '@electron-forge/publisher-github',
            config: {
                repository: {
                    owner: 'nmannaii',
                    name: 'netflix-remote'
                },
                prerelease: true
            }
        }
    ],
};
