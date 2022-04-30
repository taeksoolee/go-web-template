export async function getUsers() {
  return fetch('/api/v1/users').then(res => res.json());
}

export async function postUser(fd) {
  return fetch('/api/v1/users', {
    method: 'POST',
    body: JSON.stringify(fd)
  }).then(res => res.json())
}