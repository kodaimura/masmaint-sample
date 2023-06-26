window.addEventListener('DOMContentLoaded', (event) => {
	fetch('api/department')
	.then(response => response.json())
	.then(data  => setData(data));
});


const setData = async (data) => {
	await renderTbody(data);
	addChangedAction('name');
	addChangedAction('description');
	addChangedAction('manager_id');
	addChangedAction('location');
	addChangedAction('budget');
}

const renderTbody = (data) => {
	let tbody= '';
	for (const elem of data) {
		tbody += `<tr><td>${elem.id}</td>`
		+ `<td><input type="text" name="name" value=${elem.name}><input type="hidden" name="name_bk" value=${elem.name}></td>`
		+ `<td><input type="text" name="description" value=${elem.description}><input type="hidden" name="description_bk" value=${elem.description}></td>`
		+ `<td><input type="text" name="manager_id" value=${elem.manager_id}><input type="hidden" name="manager_id_bk" value=${elem.manager_id}></td>`
		+ `<td><input type="text" name="location" value=${elem.location}><input type="hidden" name="location_bk" value=${elem.location}></td>`
		+ `<td><input type="text" name="budget" value=${elem.budget}><input type="hidden" name="budget_bk" value=${elem.budget}></td>`
		+ `<td>${elem.created_at}</td>`
		+ `<td>${elem.updated_at}</td></tr>`
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