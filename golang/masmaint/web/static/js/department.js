window.addEventListener('DOMContentLoaded', (event) => {
	fetch('api/department')
	.then(response => response.json())
	.then(data  => renderTbody(data));
});

const renderTbody = (data) => {
	let tbody= '';
	for (const elem of data) {
		tbody += `<tr><td>${elem.id}</td>`
		+ `<td>${elem.name}</td>`
		+ `<td>${elem.description}</td>`
		+ `<td>${elem.manager_id}</td>`
		+ `<td>${elem.location}</td>`
		+ `<td>${elem.budget}</td>`
		+ `<td>${elem.created_at}</td>`
		+ `<td>${elem.updated_at}</td></tr>`
	}

	document.getElementById('list-body').innerHTML = tbody;
}