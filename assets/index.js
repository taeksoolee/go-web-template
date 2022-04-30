console.log('hello assets!');

document.getElementById('post-user').addEventListener('click', function() {
  const fd = {
    "firstName": "abccc",
    "lastName": "defff",
    "rollId": 2
  }

  fetch('/api/v1/users', {
    method: 'POST',
    body: JSON.stringify(fd)
  })
  .then(res => res.json())
  .then((data) => {
    console.log(data);
  })

  return false;
});