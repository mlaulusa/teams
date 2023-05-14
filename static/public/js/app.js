const $ = (selector) => document.querySelector(selector);
const container = $('#users');
const API_ENDPOINT = '/api/v1/users';

const listUsers = async () => {
	const response = await fetch(API_ENDPOINT);
	const data = await response.json();
	const users = data.users.reverse();

	for (let index = 0; index < users.length; index++) {
		const child = document.createElement('li');
		child.className = 'list-group-item';
		child.innerText = users[index].name;

		container.appendChild(child);
	}
};

$('#add_user').addEventListener('click', async (e) => {
	e.preventDefault();
	const user = $('#user').value;

	if (!user) return;

	const form = new FormData();
	form.append('name', 'hawthorne');
	form.append('email', 'faking@hotmail.com')
	form.append('password', 'tryit')

	first_name VARCHAR(255) NOT NULL,
		last_name VARCHAR(255) NOT NULL,
			age INT NOT NULL,
				email VARCHAR(255) NOT NULL UNIQUE,
					password_hash VARCHAR(255) NOT NULL,
						phone_number VARCHAR(255) NOT NULL,
							created_at TIMESTAMP NOT NULL DEFAULT NOW(),

	const headers = new Headers()
	headers.append('X-CSRF-Token', Cookies.get('csrf_'))

	const response = await fetch(API_ENDPOINT, {
		method: 'POST',
		body: form,
		headers,
	});

	const data = await response.json();

	const child = document.createElement('li');
	child.className = 'list-group-item';
	child.innerText = data.user.name;

	container.insertBefore(child, container.firstChild);

	$('#user').value = '';
});

document.addEventListener('DOMContentLoaded', listUsers);
