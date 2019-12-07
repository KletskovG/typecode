function setProcessText(text, processBlock) {
  processBlock.innerHTML = text;
}

window.onload = () => {
  const input = document.querySelector('#studentFile');
  const button = document.querySelector('.sendButton');
  const process = document.querySelector('.process');
  const name = document.querySelector('#studentName');

  
  try {
    button.addEventListener('click', () => {
      if (input.files[0] === undefined) return;

      const reader = new FileReader();
      setProcessText('Processing file', process);

      reader.onloadend = e => {

        const body = JSON.stringify({
          data: e.target.result,
          name: name.value,
        })

        console.log(body);

        setProcessText('File was sent to me', process);

        fetch('http://kletskovg.tech/file', {
          method: 'post',
          mode: 'no-cors',
          headers: {
            'Content-type': 'application/json',
          },
          body,
        })
          .then(res => {
            console.log(res.json())
          })
          .catch(err => {
            console.log(err);
          })
      }
      reader.readAsDataURL(input.files[0])
    });
  }
  catch (err) {
    console.log(err);
  }
  
}










