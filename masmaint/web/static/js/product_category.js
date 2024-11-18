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
		<td><input type='text' id='parent_category_id_new'></td>
		<td><input type='text' id='description_new'></td>
		<td><input type='text' id='sort_order_new'></td>
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
		<td><input type='text' name='parent_category_id' value='${nullToEmpty(elem.parent_category_id)}'><input type='hidden' name='parent_category_id_bk' value='${nullToEmpty(elem.parent_category_id)}'></td>
		<td><input type='text' name='description' value='${nullToEmpty(elem.description)}'><input type='hidden' name='description_bk' value='${nullToEmpty(elem.description)}'></td>
		<td><input type='text' name='sort_order' value='${nullToEmpty(elem.sort_order)}'><input type='hidden' name='sort_order_bk' value='${nullToEmpty(elem.sort_order)}'></td>
		<td><input type='text' name='is_active' value='${nullToEmpty(elem.is_active)}'><input type='hidden' name='is_active_bk' value='${nullToEmpty(elem.is_active)}'></td>
		<td><input type='text' name='created_at' value='${nullToEmpty(elem.created_at)}' disabled></td>
		<td><input type='text' name='updated_at' value='${nullToEmpty(elem.updated_at)}' disabled></td>`;
	return tr;
}


/* セットアップ */
const getRows = async () => {
	document.getElementById('records').innerHTML = '';
	const rows = await api.get('product_category');
	renderTbody(rows);
	addChangeEvent('name');
	addChangeEvent('parent_category_id');
	addChangeEvent('description');
	addChangeEvent('sort_order');
	addChangeEvent('is_active');
}


/* 一括更新 */
const putRows = async () => {
	let successCount = 0;
	let errorCount = 0;

	const id = document.getElementsByName('id');
	const name = document.getElementsByName('name');
	const parent_category_id = document.getElementsByName('parent_category_id');
	const description = document.getElementsByName('description');
	const sort_order = document.getElementsByName('sort_order');
	const is_active = document.getElementsByName('is_active');
	const created_at = document.getElementsByName('created_at');
	const updated_at = document.getElementsByName('updated_at');

	const name_bk = document.getElementsByName('name_bk');
	const parent_category_id_bk = document.getElementsByName('parent_category_id_bk');
	const description_bk = document.getElementsByName('description_bk');
	const sort_order_bk = document.getElementsByName('sort_order_bk');
	const is_active_bk = document.getElementsByName('is_active_bk');

	for (let i = 0; i < id.length; i++) {
		const rowMap = {
			'name': name[i],
			'parent_category_id': parent_category_id[i],
			'description': description[i],
			'sort_order': sort_order[i],
			'is_active': is_active[i],
		}

		const rowBkMap = {
			'name': name_bk[i],
			'parent_category_id': parent_category_id_bk[i],
			'description': description_bk[i],
			'sort_order': sort_order_bk[i],
			'is_active': is_active_bk[i],
		}

		//差分がある行のみ更新
		if (Object.keys(rowMap).some(key => rowMap[key].value !== rowBkMap[key].value)) {
			const requestBody = {
				id: parseIntOrReturnOriginal(id[i].value),
				name: name[i].value,
				parent_category_id: parseIntOrReturnOriginal(parent_category_id[i].value),
				description: emptyToNull(description[i].value),
				sort_order: parseIntOrReturnOriginal(sort_order[i].value),
				is_active: parseIntOrReturnOriginal(is_active[i].value),
				created_at: emptyToNull(created_at[i].value),
				updated_at: emptyToNull(updated_at[i].value),
			}

			try {
				const data = await api.put('product_category', requestBody);

				id[i].value = data.id;
				name[i].value = data.name;
				parent_category_id[i].value = data.parent_category_id;
				description[i].value = data.description;
				sort_order[i].value = data.sort_order;
				is_active[i].value = data.is_active;
				created_at[i].value = data.created_at;
				updated_at[i].value = data.updated_at;
				name_bk[i].value = data.name;
				parent_category_id_bk[i].value = data.parent_category_id;
				description_bk[i].value = data.description;
				sort_order_bk[i].value = data.sort_order;
				is_active_bk[i].value = data.is_active;

				Object.values(rowMap).forEach(element => {
					element.classList.remove('changed');
					element.classList.remove('error');
				});

				successCount += 1;
			} catch (e) {
				Object.keys(rowMap).forEach(key => {
					rowMap[key].classList.toggle('error', key === e.details.field);
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
		'parent_category_id': document.getElementById('parent_category_id_new'),
		'description': document.getElementById('description_new'),
		'sort_order': document.getElementById('sort_order_new'),
		'is_active': document.getElementById('is_active_new'),
	}

	if (Object.keys(rowMap).some(key => rowMap[key].value !== '')) {
		const requestBody = {
			name: rowMap.name.value,
			parent_category_id: parseIntOrReturnOriginal(rowMap.parent_category_id.value),
			description: emptyToNull(rowMap.description.value),
			sort_order: parseIntOrReturnOriginal(rowMap.sort_order.value),
			is_active: parseIntOrReturnOriginal(rowMap.is_active.value),
		}

		try {
			const data = await api.post('product_category', requestBody);

			document.getElementById('new').remove();
			const tr = createTr(data);
			tr.addEventListener('change', handleChange);
			document.getElementById('records').appendChild(tr);
			document.getElementById('records').appendChild(createTrNew());

			renderMessage('登録', 1, true);
		} catch (e) {
			Object.keys(rowMap).forEach(key => {
				rowMap[key].classList.toggle('error', key === e.details.field || `product_category.${key}` === e.details.column);
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
			await api.delete('product_category', row);
			successCount += 1;
		} catch (e) {
			errorCount += 1;
		}
	}

	getRows();

	renderMessage('削除', successCount, true);
	renderMessage('削除', errorCount, false);
}