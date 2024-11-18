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
		<td><input type='text' id='description_new'></td>
		<td><input type='text' id='price_new'></td>
		<td><input type='text' id='stock_quantity_new'></td>
		<td><input type='text' id='category_id_new'></td>
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
		<td><input type='text' name='description' value='${nullToEmpty(elem.description)}'><input type='hidden' name='description_bk' value='${nullToEmpty(elem.description)}'></td>
		<td><input type='text' name='price' value='${nullToEmpty(elem.price)}'><input type='hidden' name='price_bk' value='${nullToEmpty(elem.price)}'></td>
		<td><input type='text' name='stock_quantity' value='${nullToEmpty(elem.stock_quantity)}'><input type='hidden' name='stock_quantity_bk' value='${nullToEmpty(elem.stock_quantity)}'></td>
		<td><input type='text' name='category_id' value='${nullToEmpty(elem.category_id)}'><input type='hidden' name='category_id_bk' value='${nullToEmpty(elem.category_id)}'></td>
		<td><input type='text' name='created_at' value='${nullToEmpty(elem.created_at)}' disabled></td>
		<td><input type='text' name='updated_at' value='${nullToEmpty(elem.updated_at)}' disabled></td>`;
	return tr;
}


/* セットアップ */
const getRows = async () => {
	document.getElementById('records').innerHTML = '';
	const rows = await api.get('product');
	renderTbody(rows);
	addChangeEvent('name');
	addChangeEvent('description');
	addChangeEvent('price');
	addChangeEvent('stock_quantity');
	addChangeEvent('category_id');
}


/* 一括更新 */
const putRows = async () => {
	let successCount = 0;
	let errorCount = 0;

	const id = document.getElementsByName('id');
	const name = document.getElementsByName('name');
	const description = document.getElementsByName('description');
	const price = document.getElementsByName('price');
	const stock_quantity = document.getElementsByName('stock_quantity');
	const category_id = document.getElementsByName('category_id');
	const created_at = document.getElementsByName('created_at');
	const updated_at = document.getElementsByName('updated_at');

	const name_bk = document.getElementsByName('name_bk');
	const description_bk = document.getElementsByName('description_bk');
	const price_bk = document.getElementsByName('price_bk');
	const stock_quantity_bk = document.getElementsByName('stock_quantity_bk');
	const category_id_bk = document.getElementsByName('category_id_bk');

	for (let i = 0; i < id.length; i++) {
		const rowMap = {
			'name': name[i],
			'description': description[i],
			'price': price[i],
			'stock_quantity': stock_quantity[i],
			'category_id': category_id[i],
		}

		const rowBkMap = {
			'name': name_bk[i],
			'description': description_bk[i],
			'price': price_bk[i],
			'stock_quantity': stock_quantity_bk[i],
			'category_id': category_id_bk[i],
		}

		//差分がある行のみ更新
		if (Object.keys(rowMap).some(key => rowMap[key].value !== rowBkMap[key].value)) {
			const requestBody = {
				id: parseIntOrReturnOriginal(id[i].value),
				name: name[i].value,
				description: emptyToNull(description[i].value),
				price: parseFloatOrReturnOriginal(price[i].value),
				stock_quantity: parseIntOrReturnOriginal(stock_quantity[i].value),
				category_id: parseIntOrReturnOriginal(category_id[i].value),
				created_at: emptyToNull(created_at[i].value),
				updated_at: emptyToNull(updated_at[i].value),
			}

			try {
				const data = await api.put('product', requestBody);

				id[i].value = data.id;
				name[i].value = data.name;
				description[i].value = data.description;
				price[i].value = data.price;
				stock_quantity[i].value = data.stock_quantity;
				category_id[i].value = data.category_id;
				created_at[i].value = data.created_at;
				updated_at[i].value = data.updated_at;
				name_bk[i].value = data.name;
				description_bk[i].value = data.description;
				price_bk[i].value = data.price;
				stock_quantity_bk[i].value = data.stock_quantity;
				category_id_bk[i].value = data.category_id;

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
		'description': document.getElementById('description_new'),
		'price': document.getElementById('price_new'),
		'stock_quantity': document.getElementById('stock_quantity_new'),
		'category_id': document.getElementById('category_id_new'),
	}

	if (Object.keys(rowMap).some(key => rowMap[key].value !== '')) {
		const requestBody = {
			name: rowMap.name.value,
			description: emptyToNull(rowMap.description.value),
			price: parseFloatOrReturnOriginal(rowMap.price.value),
			stock_quantity: parseIntOrReturnOriginal(rowMap.stock_quantity.value),
			category_id: parseIntOrReturnOriginal(rowMap.category_id.value),
		}

		try {
			const data = await api.post('product', requestBody);

			document.getElementById('new').remove();
			const tr = createTr(data);
			tr.addEventListener('change', handleChange);
			document.getElementById('records').appendChild(tr);
			document.getElementById('records').appendChild(createTrNew());

			renderMessage('登録', 1, true);
		} catch (e) {
			Object.keys(rowMap).forEach(key => {
				rowMap[key].classList.toggle('error', key === e.details.field || `product.${key}` === e.details.column);
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
			await api.delete('product', row);
			successCount += 1;
		} catch (e) {
			errorCount += 1;
		}
	}

	getRows();

	renderMessage('削除', successCount, true);
	renderMessage('削除', errorCount, false);
}