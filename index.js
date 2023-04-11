const video = document.getElementById('video');

navigator.mediaDevices.getUserMedia({ video: true })
  .then(stream => {
    video.srcObject = stream;
  })
  .catch(error => {
    console.log(error);
  });

function capture() {
  const canvas = document.createElement('canvas');
  const context = canvas.getContext('2d');
  canvas.width = video.videoWidth;
  canvas.height = video.videoHeight;
  context.drawImage(video, 0, 0, canvas.width, canvas.height);

  const dataUrl = canvas.toDataURL('image/jpeg');
  sendData(dataUrl);
}

function sendData(dataUrl) {
  const xhr = new XMLHttpRequest();
  xhr.open('POST', '/image', true);
  xhr.setRequestHeader('Content-Type', 'application/json;charset=UTF-8');
  xhr.onreadystatechange = function() {
    if (xhr.readyState === 4 && xhr.status === 200) {
      console.log('Image sent to server!');
    }
  };
  xhr.send(JSON.stringify({ dataUrl }));
}
