import { html, render } from 'https://unpkg.com/lit-html@0.7.1/lit-html.js'
import { getUsers } from '/assets/service.mjs'

const UsersTemplate = (users) => html`
  <div>
    ${users.map((user) => (html`
      <div>${user.firstName}, ${user.lastName}</div>
    `))}
  </div>
`;

async function bootstrap() {
  const app = document.getElementById('app');

  const users = await getUsers();
  render(
    UsersTemplate(users),
    app,  
  );
}

bootstrap();