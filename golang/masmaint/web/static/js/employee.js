window.addEventListener('DOMContentLoaded', (event) => {
	fetch('api/employee')
	.then(response => response.json())
	.then(data  => setData(data));
});


const setData = async (data) => {
	await renderTbody(data);
	addChangedAction('first_name');
	addChangedAction('last_name');
	addChangedAction('email');
	addChangedAction('phone_number');
	addChangedAction('address');
	addChangedAction('hire_date');
	addChangedAction('job_title');
	addChangedAction('department');
	addChangedAction('salary');
}

const renderTbody = (data) => {
	let tbody= '';
	for (const elem of data) {
		tbody += `<tr><td>${elem.id}</td>`
		+ `<td><input type="text" name="first_name" value=${elem.first_name}><input type="hidden" name="first_name_bk" value=${elem.first_name}></td>`
		+ `<td><input type="text" name="last_name" value=${elem.last_name}><input type="hidden" name="last_name_bk" value=${elem.last_name}></td>`
		+ `<td><input type="text" name="email" value=${elem.email}><input type="hidden" name="email_bk" value=${elem.email}></td>`
		+ `<td><input type="text" name="phone_number" value=${elem.phone_number}><input type="hidden" name="phone_number_bk" value=${elem.phone_number}></td>`
		+ `<td><input type="text" name="address" value=${elem.address}><input type="hidden" name="address_bk" value=${elem.address}></td>`
		+ `<td><input type="text" name="hire_date" value=${elem.hire_date}><input type="hidden" name="hire_date_bk" value=${elem.hire_date}></td>`
		+ `<td><input type="text" name="job_title" value=${elem.job_title}><input type="hidden" name="job_title_bk" value=${elem.job_title}></td>`
		+ `<td><input type="text" name="department_id" value=${elem.department_id}><input type="hidden" name="department_id_bk" value=${elem.department_id}></td>`
		+ `<td><input type="text" name="salary" value=${elem.salary}><input type="hidden" name="salary_bk" value=${elem.salary}></td></tr>`
	}

	document.getElementById('list-body').innerHTML = tbody;
}

const addChangedAction = (columnName) => {
	const elems = document.getElementsByName(columnName);
	const elems_bk = document.getElementsByName(`${columnName}_bk`);
	for (let i = 0; i < elems.length; i++) {
		elems[i].addEventListener('change', () => {
			if (elems[i].value !== elems_bk[i].value) {
				elems[i].classList.add('changed');
			} else {
				elems[i].classList.remove('changed');
			}
		});
	}
} 


