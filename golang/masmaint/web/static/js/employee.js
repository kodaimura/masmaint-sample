window.addEventListener('DOMContentLoaded', (event) => {
	fetch('api/employee')
	.then(response => response.json())
	.then(data  => renderTbody(data));
});

const renderTbody = (data) => {
	let tbody= '';
	for (const elem of data) {
		tbody += `<tr><td>${elem.id}</td>`
		+ `<td>${elem.first_name}</td>`
		+ `<td>${elem.last_name}</td>`
		+ `<td>${elem.email}</td>`
		+ `<td>${elem.phone_number}</td>`
		+ `<td>${elem.address}</td>`
		+ `<td>${elem.hire_date}</td>`
		+ `<td>${elem.job_title}</td>`
		+ `<td>${elem.department_id}</td>`
		+ `<td>${elem.salary}</td></tr>`
	}

	document.getElementById('list-body').innerHTML = tbody;
}