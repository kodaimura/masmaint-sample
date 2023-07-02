/* 初期設定 */
window.addEventListener('DOMContentLoaded', (event) => {
	setUp();
});

/* リロードボタン押下 */
document.getElementById('reload').addEventListener('click', (event) => {
	clearMessage();
	document.getElementById('records').innerHTML = '';
	setUp();
})

/* 保存モーダル確定ボタン押下 */
document.getElementById('ModalSaveAllOk').addEventListener('click', (event) => {
	clearMessage();
	doPutAll();
	doPost();
})

/* 削除モーダル確定ボタン押下 */
document.getElementById('ModalDeleteAllOk').addEventListener('click', (event) => {
	clearMessage();
	doDeleteAll();
})

/* チェックボックスの選択一覧取得 */
const getDeleteTarget = () => {
	let dels = document.getElementsByName('del');
	let ret = [];

	for (let x of dels) {
		if (x.checked) {
			ret.push(x.value);
		}
	}
	return ret
}

const renderMessage = (msg, count, isSuccess) => {
	if (count !== 0) {
		let message = document.createElement('div');
		message.textContent = `${count}件の${msg}に${isSuccess? '成功' : '失敗'}しました。`
		message.className = `alert alert-${isSuccess? 'success' : 'danger'} alert-custom my-1`;
		document.getElementById('message').appendChild(message);
	}
}

const clearMessage = () => {
	document.getElementById('message').innerHTML = '';
}

const nullToEmpty = (s) => {
	return (s == null)? '' : s;
}

/* チェンジアクション */
const changeAction = (event) => {
	let target = event.target;
	let target_bk = target.nextElementSibling;

	if (target_bk == null) return

	if (target.value !== target_bk.value) {
		target.classList.add('changed');
	} else {
		target.classList.remove('changed');
	}
}

/* <tbody></tbody>内のレコードにチェンジアクション追加 */
const addChangedAction = (columnName) => {
	let elems = document.getElementsByName(columnName);
	for (const elem of elems) {
		elem.addEventListener('change', changeAction);
	}
}

/* <tbody></tbody>レンダリング */
const renderTbody = (data) => {
	let tbody= '';
	for (const elem of data) {
		tbody += createTr(elem);
	}
	tbody += createTrNew();

	document.getElementById('records').innerHTML = tbody;
}

/* <tr></tr>を作成 （tbody末尾の新規登録用レコード）*/
const createTrNew = (elem) => {
	return `<tr id='new'><td></td>`
		+ `<td><input type='text' disabled></td>`
		+ `<td><input type='text' id='first_name_new'></td>`
		+ `<td><input type='text' id='last_name_new'></td>`
		+ `<td><input type='text' id='email_new'></td>`
		+ `<td><input type='text' id='phone_number_new'></td>`
		+ `<td><input type='text' id='address_new'></td>`
		+ `<td><input type='text' id='hire_date_new'></td>`
		+ `<td><input type='text' id='job_title_new'></td>`
		+ `<td><input type='text' id='department_code_new'></td>`
		+ `<td><input type='text' id='salary_new'></td></tr>`;
} 

/* <tr></tr>を作成 */
const createTr = (elem) => {
	return `<tr><td><input class='form-check-input' type='checkbox' name='del' value=${JSON.stringify(elem)}></td>`
		+ `<td><input type='text' name='id' value='${nullToEmpty(elem.id)}' disabled></td>`
		+ `<td><input type='text' name='first_name' value='${nullToEmpty(elem.first_name)}'><input type='hidden' name='first_name_bk' value='${nullToEmpty(elem.first_name)}'></td>`
		+ `<td><input type='text' name='last_name' value='${nullToEmpty(elem.last_name)}'><input type='hidden' name='last_name_bk' value='${nullToEmpty(elem.last_name)}'></td>`
		+ `<td><input type='text' name='email' value='${nullToEmpty(elem.email)}'><input type='hidden' name='email_bk' value='${nullToEmpty(elem.email)}'></td>`
		+ `<td><input type='text' name='phone_number' value='${nullToEmpty(elem.phone_number)}'><input type='hidden' name='phone_number_bk' value='${nullToEmpty(elem.phone_number)}'></td>`
		+ `<td><input type='text' name='address' value='${nullToEmpty(elem.address)}'><input type='hidden' name='address_bk' value='${nullToEmpty(elem.address)}'></td>`
		+ `<td><input type='text' name='hire_date' value='${nullToEmpty(elem.hire_date)}'><input type='hidden' name='hire_date_bk' value='${nullToEmpty(elem.hire_date)}'></td>`
		+ `<td><input type='text' name='job_title' value='${nullToEmpty(elem.job_title)}'><input type='hidden' name='job_title_bk' value='${nullToEmpty(elem.job_title)}'></td>`
		+ `<td><input type='text' name='department_code' value='${nullToEmpty(elem.department_code)}'><input type='hidden' name='department_code_bk' value='${nullToEmpty(elem.department_code)}'></td>`
		+ `<td><input type='text' name='salary' value='${nullToEmpty(elem.salary)}'><input type='hidden' name='salary_bk' value='${nullToEmpty(elem.salary)}'></td></tr>`;
} 


/* セットアップ */
const setUp = () => {
	fetch('api/employee')
	.then(response => response.json())
	.then(data  => renderTbody(data))
	.then(() => {
		addChangedAction('first_name');
		addChangedAction('last_name');
		addChangedAction('email');
		addChangedAction('phone_number');
		addChangedAction('address');
		addChangedAction('hire_date');
		addChangedAction('job_title');
		addChangedAction('department_code');
		addChangedAction('salary');
	});
}


/* 一括更新 */
const doPutAll = async () => {
	let successCount = 0;
	let errorCount = 0;

	let id = document.getElementsByName('id');

	let first_name = document.getElementsByName('first_name');
	let first_name_bk = document.getElementsByName('first_name_bk');
	let last_name = document.getElementsByName('last_name');
	let last_name_bk = document.getElementsByName('last_name_bk');
	let email = document.getElementsByName('email');
	let email_bk = document.getElementsByName('email_bk');
	let phone_number = document.getElementsByName('phone_number');
	let phone_number_bk = document.getElementsByName('phone_number_bk');
	let address = document.getElementsByName('address');
	let address_bk = document.getElementsByName('address_bk');
	let hire_date = document.getElementsByName('hire_date');
	let hire_date_bk = document.getElementsByName('hire_date_bk');
	let job_title = document.getElementsByName('job_title');
	let job_title_bk = document.getElementsByName('job_title_bk');
	let department_code = document.getElementsByName('department_code');
	let department_code_bk = document.getElementsByName('department_code_bk');
	let salary = document.getElementsByName('salary');
	let salary_bk = document.getElementsByName('salary_bk');

	for (let i = 0; i < first_name.length; i++) {
		if ((first_name[i].value !== first_name_bk[i].value) 
			|| (last_name[i].value !== last_name_bk[i].value) 
			|| (email[i].value !== email_bk[i].value)
			|| (phone_number[i].value !== phone_number_bk[i].value)
			|| (address[i].value !== address_bk[i].value)
			|| (hire_date[i].value !== hire_date_bk[i].value)
			|| (job_title[i].value !== job_title_bk[i].value)
			|| (department_code[i].value !== department_code_bk[i].value)
			|| (salary[i].value !== salary_bk[i].value)) {

			let requestBody = {
				id: id[i].value,
				first_name: first_name[i].value,
				last_name: last_name[i].value,
				email: email[i].value,
				phone_number: phone_number[i].value,
				address: address[i].value,
				hire_date: hire_date[i].value,
				job_title: job_title[i].value,
				department_code: department_code[i].value,
				salary: salary[i].value,
			}

			await fetch('api/employee', {
				method: 'PUT',
				headers: {'Content-Type': 'application/json'},
				body: JSON.stringify(requestBody)
			})
			.then(response => {
				if (!response.ok){
					throw new Error(response.statusText);
				}
  				return response.json();
  			})
			.then(data => {
				id[i].value = data.id;
				first_name[i].value = data.first_name;
				first_name_bk[i].value = data.first_name;
				last_name[i].value = data.last_name;
				last_name_bk[i].value = data.last_name;
				email[i].value = data.email;
				email_bk[i].value = data.email;
				phone_number[i].value = data.phone_number;
				phone_number_bk[i].value = data.phone_number;
				address[i].value = data.address;
				address_bk[i].value = data.address;
				hire_date[i].value = data.hire_date;
				hire_date_bk[i].value = data.hire_date;
				job_title[i].value = data.job_title;
				job_title_bk[i].value = data.job_title;
				department_code[i].value = data.department_code;
				department_code_bk[i].value = data.department_code;
				salary[i].value = data.salary;
				salary_bk[i].value = data.salary;

				first_name[i].classList.remove('changed');
				last_name[i].classList.remove('changed');
				email[i].classList.remove('changed');
				phone_number[i].classList.remove('changed');
				address[i].classList.remove('changed');
				hire_date[i].classList.remove('changed');
				job_title[i].classList.remove('changed');
				department_code[i].classList.remove('changed');
				salary[i].classList.remove('changed');

				successCount += 1;
			}).catch(error => {
				errorCount += 1;				
			})
		}
	}

	renderMessage('更新', successCount, true);
	renderMessage('更新', errorCount, false);
} 


/* 新規登録 */
const doPost = () => {
	let first_name = document.getElementById('first_name_new').value;
	let last_name = document.getElementById('last_name_new').value;
	let email = document.getElementById('email_new').value;
	let phone_number = document.getElementById('phone_number_new').value;
	let address = document.getElementById('address_new').value;
	let hire_date = document.getElementById('hire_date_new').value;
	let job_title = document.getElementById('job_title_new').value;
	let department_code = document.getElementById('department_code_new').value;
	let salary = document.getElementById('salary_new').value;

	if ((first_name !== '') 
		|| (last_name !== '') 
		|| (email !== '') 
		|| (phone_number !== '') 
		|| (address !== '')
		|| (hire_date !== '')
		|| (job_title !== '') 
		|| (department_code !== '')
		|| (salary !== ''))
	{
		let requestBody = {
			first_name: first_name,
			last_name: last_name,
			email: email,
			phone_number: phone_number ,
			address: address,
			hire_date: hire_date,
			job_title: job_title,
			department_code: department_code,
			salary: salary
		}

		fetch('api/employee', {
			method: 'POST',
			headers: {'Content-Type': 'application/json'},
			body: JSON.stringify(requestBody)
		})
		.then(response => {
			if (!response.ok){
				throw new Error(response.statusText);
			}
  			return response.json();
  		})
		.then(data => {
			document.getElementById('new').remove();

			let tmpElem = document.createElement('tbody');
			tmpElem.innerHTML = createTr(data);
			tmpElem.firstChild.addEventListener('change', changeAction);
			document.getElementById('records').appendChild(tmpElem.firstChild);

			tmpElem = document.createElement('tbody');
			tmpElem.innerHTML = createTrNew();
			document.getElementById('records').appendChild(tmpElem.firstChild);

			renderMessage('登録', 1, true);
		}).catch(error => {
			renderMessage('登録', 1, false);
		})
	}
}


/* 一括削除 */
const doDeleteAll = async () => {
	let ls = getDeleteTarget();
	let successCount = 0;
	let errorCount = 0;

	for (let x of ls) {
		await fetch('api/employee', {
			method: 'DELETE',
			headers: {'Content-Type': 'application/json'},
			body: x
		})
		.then(response => {
			if (!response.ok){
				throw new Error(response.statusText);
			}
			successCount += 1;
  		}).catch(error => {
			errorCount += 1;
		});
	}

	setUp();

	renderMessage('削除', successCount, true);
	renderMessage('削除', errorCount, false);
}