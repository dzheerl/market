const button = document.getElementsByClassName('button')[0];


button.addEventListener('click', function (event) {
  event.preventDefault();

  const form_email = document.getElementById('form_email');
  const form_password = document.getElementById('form_password');
  const formData = {
    form_email: form_email.value,
    form_password: form_password.value
  };

  if (form_email.value === '' || form_password.value === '') {
    button.innerText = 'ERROR';
    form_email.value = '';
    form_password.value = '';
    setTimeout(function () {
      button.innerText = 'Confirm'
    }, 1500)
  } else {
    fetch('/send', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(formData)
    })
      .then(response => response.json())
      .then(data => {
        // Обработка ответа сервера
        console.log(data);
      })
      .catch(error => {
        console.error("Error:", error);
      });
    console.log("succes");
    form_email.value = '';
    form_password.value = '';
  }
});