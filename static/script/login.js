const button = document.getElementsByClassName('button')[0];


button.addEventListener('click', function (event) {
  event.preventDefault();

  const form_email = document.getElementById('form_email');
  const form_password = document.getElementById('form_password');
  const changeMainForm  = document.querySelector('.main_form'); 


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
    
    },
    form_email.value = '',
    form_password.value = '',)
    .then(response => {
      if (!response.ok) {
        throw new Error('Network response was not ok.');
      }
      return response.text(); // Или response.json() если сервер возвращает JSON
    })
    .then(data => {
      if (data === 'Succes') {
        // Действие при успешном ответе
        // Например, вывод сообщения об успешном выполнении операции
        console.log('Операция выполнена успешно!');
        // Здесь можете добавить логику для обновления страницы или отображения сообщения об успехе
        const mainForm = document.getElementsByClassName('main_form');
        const newElement = document.createElement('p');
        newElement.id = 'testid';
        newElement.textContent = 'Hello';
        mainForm.parentNode.replaceChild(newElement, mainForm);
      } else {
        // Действие при другом ответе
        console.log('Произошла ошибка!');
        // Здесь можете добавить логику для отображения сообщения об ошибке
      }
    })
    .catch(error => {
      // Обработка ошибок запроса
      console.error('There has been a problem with your fetch operation:', error);
    })  
 }
});

