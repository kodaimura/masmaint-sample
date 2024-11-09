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
document.getElementById('modal-save-ok').addEventListener('click', (event) => {
	clearMessage();
	doPutAll();
	doPost();
})

/* 削除モーダル確定ボタン押下 */
document.getElementById('modal-delete-dk').addEventListener('click', (event) => {
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
	if (data != null) {
		for (const elem of data) {
			tbody += createTr(elem);
		}
	}
	tbody += createTrNew();

	document.getElementById('records').innerHTML = tbody;
}

/* <tr></tr>を作成 （tbody末尾の新規登録用レコード）*/
const createTrNew = (elem) => {
	return `<tr id='new'><td></td>`
		+ `<td><input type='text' id='code_new'></td>`
		+ `<td><input type='text' id='name_new'></td>`
		+ `<td><input type='text' id='description_new'></td>`
		+ `<td><input type='text' id='manager_id_new'></td>`
		+ `<td><input type='text' id='location_new'></td>`
		+ `<td><input type='text' id='budget_new'></td>`
		+ `<td><input type='text' disabled></td>`
		+ `<td><input type='text' disabled></td>`;
} 

/* <tr></tr>を作成 */
const createTr = (elem) => {
	return `<tr><td><input class='form-check-input' type='checkbox' name='del' value='${JSON.stringify(elem)}'></td>`
		+ `<td><input type='text' name='code' value='${nullToEmpty(elem.code)}' disabled></td>`
		+ `<td><input type='text' name='name' value='${nullToEmpty(elem.name)}'><input type='hidden' name='name_bk' value='${nullToEmpty(elem.name)}'></td>`
		+ `<td><input type='text' name='description' value='${nullToEmpty(elem.description)}'><input type='hidden' name='description_bk' value='${nullToEmpty(elem.description)}'></td>`
		+ `<td><input type='text' name='manager_id' value='${nullToEmpty(elem.manager_id)}'><input type='hidden' name='manager_id_bk' value='${nullToEmpty(elem.manager_id)}'></td>`
		+ `<td><input type='text' name='location' value='${nullToEmpty(elem.location)}'><input type='hidden' name='location_bk' value='${nullToEmpty(elem.location)}'></td>`
		+ `<td><input type='text' name='budget' value='${nullToEmpty(elem.budget)}'><input type='hidden' name='budget_bk' value='${nullToEmpty(elem.budget)}'></td>`
		+ `<td><input type='text' name='created_at' value='${nullToEmpty(elem.created_at)}' disabled></td>`
		+ `<td><input type='text' name='updated_at' value='${nullToEmpty(elem.updated_at)}' disabled></td>`;
} 


/* セットアップ */
const setUp = () => {
	fetch('api/department')
	.then(response => response.json())
	.then(data  => renderTbody(data))
	.then(() => {
		addChangedAction('name');
		addChangedAction('description');
		addChangedAction('manager_id');
		addChangedAction('location');
		addChangedAction('budget');
	});
}


/* 一括更新 */
const doPutAll = async () => {
	let successCount = 0;
	let errorCount = 0;
	
	let code = document.getElementsByName('code');
	let name = document.getElementsByName('name');
	let name_bk = document.getElementsByName('name_bk');
	let description = document.getElementsByName('description');
	let description_bk = document.getElementsByName('description_bk');
	let manager_id = document.getElementsByName('manager_id');
	let manager_id_bk = document.getElementsByName('manager_id_bk');
	let location = document.getElementsByName('location');
	let location_bk = document.getElementsByName('location_bk');
	let budget = document.getElementsByName('budget');
	let budget_bk = document.getElementsByName('budget_bk');
	let created_at = document.getElementsByName('created_at');
	let updated_at = document.getElementsByName('updated_at');

	for (let i = 0; i < code.length; i++) {
		if ((name[i].value !== name_bk[i].value)
			|| (description[i].value !== description_bk[i].value)
			|| (manager_id[i].value !== manager_id_bk[i].value)
			|| (location[i].value !== location_bk[i].value)
			|| (budget[i].value !== budget_bk[i].value)) {

			let requestBody = {
				code: code[i].value,
				name: name[i].value,
				description: description[i].value,
				manager_id: manager_id[i].value,
				location: location[i].value,
				budget: budget[i].value,
				created_at: created_at[i].value,
				updated_at: updated_at[i].value,
			}

			await fetch('api/department', {
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
				code[i].value = data.code;
				name[i].value = data.name;
				name_bk[i].value = data.name;
				description[i].value = data.description;
				description_bk[i].value = data.description;
				manager_id[i].value = data.manager_id;
				manager_id_bk[i].value = data.manager_id;
				location[i].value = data.location;
				location_bk[i].value = data.location;
				budget[i].value = data.budget;
				budget_bk[i].value = data.budget;
				created_at[i].value = data.created_at;
				updated_at[i].value = data.updated_at;

				name[i].classList.remove('changed');
				description[i].classList.remove('changed');
				manager_id[i].classList.remove('changed');
				location[i].classList.remove('changed');
				budget[i].classList.remove('changed');

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
	let code = document.getElementById('code_new').value;
	let name = document.getElementById('name_new').value;
	let description = document.getElementById('description_new').value;
	let manager_id = document.getElementById('manager_id_new').value;
	let location = document.getElementById('location_new').value;
	let budget = document.getElementById('budget_new').value;

	if ((code !== '')
		|| (name !== '')
		|| (description !== '')
		|| (manager_id !== '')
		|| (location !== '')
		|| (budget !== ''))
	{
		let requestBody = {
			code: code,
			name: name,
			description: description,
			manager_id: manager_id,
			location: location,
			budget: budget,
		}

		fetch('api/department', {
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
		await fetch('api/department', {
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