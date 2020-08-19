window.println()

document.addEventListener('DOMContentLoaded', function () {
    if (!Notification) {
        alert('Desktop notifications not available in your browser. Try Chromium.');
        return;
    }

    if (Notification.permission !== 'granted')
        Notification.requestPermission();
});

function notifyMe() {
    if (Notification.permission !== 'granted')
        Notification.requestPermission();
    else {
        new Notification('muonの練習', {
            icon: './img/logo.svg',
            body: `デスクトップ通知を出してるよ～`,
        });
    }
}
