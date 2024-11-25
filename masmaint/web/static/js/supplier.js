import { api } from '/js/api.js';
import { nullToEmpty, emptyToNull, parseFloatOrReturnOriginal, parseIntOrReturnOriginal } from './script.js';

/* 初期設定 */
window.addEventListener('DOMContentLoaded', (event) => {
    getRows();
});

/* リロードボタン押下 */
document.getElementById('reload').addEventListener('click', (event) => {
    clearMessage();
    getRows();
})

/* 保存モーダル確定押下 */
document.getElementById('modal-save-ok').addEventListener('click', (event) => {
    clearMessage();
    putRows();
    postRow();
})

/* 削除モーダル確定押下 */
document.getElementById('modal-delete-dk').addEventListener('click', (event) => {
    clearMessage();
    deleteRows();
})

/* チェックボックスの選択一覧取得 */
const getDeleteTargetRows = () => {
    const elems = document.getElementsByName('del');
    let ret = [];

    for (let elem of elems) {
        if (elem.checked) {
            ret.push(JSON.parse(elem.value));
        }
    }
    return ret
}

const renderMessage = (msg, count, isSuccess) => {
    if (count !== 0) {
        const message = document.createElement('div');
        message.textContent = `${count}件の${msg}に${isSuccess ? '成功' : '失敗'}しました。`
        message.className = `alert alert-${isSuccess ? 'success' : 'danger'} alert-custom my-1`;
        document.getElementById('message').appendChild(message);
    }
}

const clearMessage = () => {
    document.getElementById('message').innerHTML = '';
}

/* changeイベントハンドラ */
const handleChange = (event) => {
    const target = event.target;
    const target_bk = target.nextElementSibling;

    if (target_bk == null) return

    if (target.value !== target_bk.value) {
        target.classList.add('changed');
    } else {
        target.classList.remove('changed');
    }
}

/* <tbody></tbody>内のレコードにチェンジアクション追加 */
const addChangeEvent = (columnName) => {
    const elems = document.getElementsByName(columnName);
    for (const elem of elems) {
        elem.addEventListener('change', handleChange);
    }
}

/* <tbody></tbody>レンダリング */
const renderTbody = (data) => {
    const tbody = document.getElementById('records');
    if (data != null) {
        for (const elem of data) {
            tbody.appendChild(createTr(elem));
        }
    }
    tbody.appendChild(createTrNew());
}

/* <tr></tr>を作成 （tbody末尾の新規登録用レコード）*/
const createTrNew = (elem) => {
	const tr = document.createElement('tr');
	tr.id = 'new';
	tr.innerHTML = `
		<td></td>
		<td><input type='text' disabled></td>
		<td><input type='text' id='name_new'></td>
		<td><input type='text' id='contact_person_new'></td>
		<td><input type='text' id='phone_new'></td>
		<td><input type='text' id='email_new'></td>
		<td><input type='text' id='address_new'></td>
		<td><input type='text' id='is_active_new'></td>
		<td><input type='text' disabled></td>
		<td><input type='text' disabled></td>`;
	return tr;
}

/* <tr></tr>を作成 */
const createTr = (elem) => {
	const tr = document.createElement('tr');
	tr.innerHTML = `
		<td><input class='form-check-input' type='checkbox' name='del' value='${JSON.stringify(elem)}'></td>
		<td><input type='text' name='id' value='${nullToEmpty(elem.id)}' disabled></td>
		<td><input type='text' name='name' value='${nullToEmpty(elem.name)}'><input type='hidden' name='name_bk' value='${nullToEmpty(elem.name)}'></td>
		<td><input type='text' name='contact_person' value='${nullToEmpty(elem.contact_person)}'><input type='hidden' name='contact_person_bk' value='${nullToEmpty(elem.contact_person)}'></td>
		<td><input type='text' name='phone' value='${nullToEmpty(elem.phone)}'><input type='hidden' name='phone_bk' value='${nullToEmpty(elem.phone)}'></td>
		<td><input type='text' name='email' value='${nullToEmpty(elem.email)}'><input type='hidden' name='email_bk' value='${nullToEmpty(elem.email)}'></td>
		<td><input type='text' name='address' value='${nullToEmpty(elem.address)}'><input type='hidden' name='address_bk' value='${nullToEmpty(elem.address)}'></td>
		<td><input type='text' name='is_active' value='${nullToEmpty(elem.is_active)}'><input type='hidden' name='is_active_bk' value='${nullToEmpty(elem.is_active)}'></td>
		<td><input type='text' name='created_at' value='${nullToEmpty(elem.created_at)}' disabled></td>
		<td><input type='text' name='updated_at' value='${nullToEmpty(elem.updated_at)}' disabled></td>`;
	return tr;
}


/* セットアップ */
const getRows = async () => {
	document.getElementById('records').innerHTML = '';
	const rows = await api.get('supplier');
	renderTbody(rows);
	addChangeEvent('name');
	addChangeEvent('contact_person');
	addChangeEvent('phone');
	addChangeEvent('email');
	addChangeEvent('address');
	addChangeEvent('is_active');
}


/* 一括更新 */
const putRows = async () => {
	let successCount = 0;
	let errorCount = 0;

	const id = document.getElementsByName('id');
	const name = document.getElementsByName('name');
	const contact_person = document.getElementsByName('contact_person');
	const phone = document.getElementsByName('phone');
	const email = document.getElementsByName('email');
	const address = document.getElementsByName('address');
	const is_active = document.getElementsByName('is_active');
	const created_at = document.getElementsByName('created_at');
	const updated_at = document.getElementsByName('updated_at');

	const name_bk = document.getElementsByName('name_bk');
	const contact_person_bk = document.getElementsByName('contact_person_bk');
	const phone_bk = document.getElementsByName('phone_bk');
	const email_bk = document.getElementsByName('email_bk');
	const address_bk = document.getElementsByName('address_bk');
	const is_active_bk = document.getElementsByName('is_active_bk');

	for (let i = 0; i < id.length; i++) {
		const rowMap = {
			'name': name[i],
			'contact_person': contact_person[i],
			'phone': phone[i],
			'email': email[i],
			'address': address[i],
			'is_active': is_active[i],
		}

		const rowBkMap = {
			'name': name_bk[i],
			'contact_person': contact_person_bk[i],
			'phone': phone_bk[i],
			'email': email_bk[i],
			'address': address_bk[i],
			'is_active': is_active_bk[i],
		}

		//差分がある行のみ更新
		if (Object.keys(rowMap).some(key => rowMap[key].value !== rowBkMap[key].value)) {
			const requestBody = {
				id: parseIntOrReturnOriginal(id[i].value),
				name: name[i].value,
				contact_person: emptyToNull(contact_person[i].value),
				phone: emptyToNull(phone[i].value),
				email: emptyToNull(email[i].value),
				address: emptyToNull(address[i].value),
				is_active: parseIntOrReturnOriginal(is_active[i].value),
				created_at: emptyToNull(created_at[i].value),
				updated_at: emptyToNull(updated_at[i].value),
			}

			try {
				const data = await api.put('supplier', requestBody);

				id[i].value = data.id;
				name[i].value = data.name;
				contact_person[i].value = data.contact_person;
				phone[i].value = data.phone;
				email[i].value = data.email;
				address[i].value = data.address;
				is_active[i].value = data.is_active;
				created_at[i].value = data.created_at;
				updated_at[i].value = data.updated_at;
				name_bk[i].value = data.name;
				contact_person_bk[i].value = data.contact_person;
				phone_bk[i].value = data.phone;
				email_bk[i].value = data.email;
				address_bk[i].value = data.address;
				is_active_bk[i].value = data.is_active;

				Object.values(rowMap).forEach(element => {
					element.classList.remove('changed');
					element.classList.remove('error');
				});

				successCount += 1;
			} catch (e) {
				Object.keys(rowMap).forEach(key => {
					rowMap[key].classList.toggle(
						'error',
						[e.details.field, e.details.column].includes(key) ||
						[e.details.field, e.details.column].includes(`supplier.${key}`)
					);
				});
				errorCount += 1;
			}
		}
	}

	renderMessage('更新', successCount, true);
	renderMessage('更新', errorCount, false);
}


/* 新規登録 */
const postRow = async () => {
	const rowMap = {
		'name': document.getElementById('name_new'),
		'contact_person': document.getElementById('contact_person_new'),
		'phone': document.getElementById('phone_new'),
		'email': document.getElementById('email_new'),
		'address': document.getElementById('address_new'),
		'is_active': document.getElementById('is_active_new'),
	}

	if (Object.keys(rowMap).some(key => rowMap[key].value !== '')) {
		const requestBody = {
			name: rowMap.name.value,
			contact_person: emptyToNull(rowMap.contact_person.value),
			phone: emptyToNull(rowMap.phone.value),
			email: emptyToNull(rowMap.email.value),
			address: emptyToNull(rowMap.address.value),
			is_active: parseIntOrReturnOriginal(rowMap.is_active.value),
		}

		try {
			const data = await api.post('supplier', requestBody);

			document.getElementById('new').remove();
			const tr = createTr(data);
			tr.addEventListener('change', handleChange);
			document.getElementById('records').appendChild(tr);
			document.getElementById('records').appendChild(createTrNew());

			renderMessage('登録', 1, true);
		} catch (e) {
			Object.keys(rowMap).forEach(key => {
				rowMap[key].classList.toggle(
					'error',
					[e.details.field, e.details.column].includes(key) ||
					[e.details.field, e.details.column].includes(`supplier.${key}`)
                );
			});
			renderMessage('登録', 1, false);
		}
	}
}


/* 一括削除 */
const deleteRows = async () => {
	const rows = getDeleteTargetRows();
	let successCount = 0;
	let errorCount = 0;

	for (let row of rows) {
		try {
			await api.delete('supplier', row);
			successCount += 1;
		} catch (e) {
			errorCount += 1;
		}
	}

	getRows();

	renderMessage('削除', successCount, true);
	renderMessage('削除', errorCount, false);
}