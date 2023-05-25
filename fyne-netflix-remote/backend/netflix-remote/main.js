import './assets/hammer.min.js';

let socket = io();

const mousePad = document.querySelector('.mouse-touchpad');
const closeBtn = document.querySelector('.close');
const closeIcon = document.querySelector('#closeIcon');

socket.on('connect', () => {
    Snackbar.show({text: 'Connection established 🚀', backgroundColor: '#138636', textColor: '#FFFFFF', showAction: false})
})
socket.on('disconnect', () => {
    Snackbar.show({text: 'Connection is gone 🤯', backgroundColor: '#c91432', textColor: '#FFFFFF', showAction: false})
})

let hammerEl = new Hammer(mousePad, {
    touchAction: 'none',
});

let buttonClick = async (key) => {
    await fetch('/key-press', {
        method: 'POST',
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({key})
    });
}
let showTouchpad = () => {
    mousePad.style.display = 'block';
}
window.buttonClick = buttonClick;
window.showTouchpad = showTouchpad;

closeBtn.addEventListener('click', (ev) => {
    mousePad.style.display = 'none';
});

let dx = 0, dy = 0, dt = 0;
hammerEl.on('pan', (ev) => {
    let scale = 1.2;
    dx = ev.deltaX - dx;
    dy = ev.deltaY - dy;
    dt = ev.deltaTime - dt;
    socket.emit('move-mouse', JSON.stringify({x: dx * scale, y: dy * scale}));
    dx = ev.deltaX;
    dy = ev.deltaY;
    dt = ev.deltaTime;
    if (ev.isFinal) {
        dx = 0;
        dy = 0;
        dt = 0;
    }
});

hammerEl.on('tap', async (ev) => {
    if (ev.target === closeBtn || ev.target === closeIcon) {
        return;
    }
    await fetch('/mouse-click', {
        method: 'POST',
        headers: {
            "Content-Type": "application/json"
        }
    });
})
