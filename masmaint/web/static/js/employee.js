import { api } from '/js/api.js';
import { parseFloatOrReturnOriginal, parseIntOrReturnOriginal } from './script.js';

/* 初期設定 */
window.addEventListener('DOMContentLoaded', (event) => {
    getRows();
});

/* リロードボタン押下 */
document.getElementById('reload').addEventListener('click', (event) => {
    clearMessage();
    getRows();
})

/* 保存モーダル確定ボタン押下 */
document.getElementById('modal-save-ok').addEventListener('click', (event) => {
    clearMessage();
    putRows();
    postRow();
})

/* 削除モーダル確定ボタン押下 */
document.getElementById('modal-delete-dk').addEventListener('click', (event) => {
    clearMessage();
    deleteRows();
})

/* チェックボックスの選択一覧取得 */
const getDeleteTarget = () => {
    let dels = document.getElementsByName('del');
    let ret = [];

    for (let x of dels) {
        if (x.checked) {
            ret.push(JSON.parse(x.value));
        }
    }
    return ret
}

const renderMessage = (msg, count, isSuccess) => {
    if (count !== 0) {
        let message = document.createElement('div');
        message.textContent = `${count}件の${msg}に${isSuccess ? '成功' : '失敗'}しました。`
        message.className = `alert alert-${isSuccess ? 'success' : 'danger'} alert-custom my-1`;
        document.getElementById('message').appendChild(message);
    }
}

const clearMessage = () => {
    document.getElementById('message').innerHTML = '';
}

const nullToEmpty = (s) => {
    return (s == null) ? '' : s;
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
		<td><input type='text' id='first_name_new'></td>
		<td><input type='text' id='last_name_new'></td>
		<td><input type='text' id='email_new'></td>
		<td><input type='text' id='phone_number_new'></td>
		<td><input type='text' id='address_new'></td>
		<td><input type='text' id='hire_date_new'></td>
		<td><input type='text' id='job_title_new'></td>
		<td><input type='text' id='department_code_new'></td>
		<td><input type='text' id='salary_new'></td>
		<td><input type='text' disabled></td>
		<td><input type='text' disabled></td>
    `;
    return tr;
}

/* <tr></tr>を作成 */
const createTr = (elem) => {
    const tr = document.createElement('tr');
    tr.innerHTML = `
        <td><input class='form-check-input' type='checkbox' name='del' value='${JSON.stringify(elem)}'></td>
		<td><input type='text' name='id' value='${nullToEmpty(elem.id)}' disabled></td>
		<td><input type='text' name='first_name' value='${nullToEmpty(elem.first_name)}'><input type='hidden' name='first_name_bk' value='${nullToEmpty(elem.first_name)}'></td>
		<td><input type='text' name='last_name' value='${nullToEmpty(elem.last_name)}'><input type='hidden' name='last_name_bk' value='${nullToEmpty(elem.last_name)}'></td>
		<td><input type='text' name='email' value='${nullToEmpty(elem.email)}'><input type='hidden' name='email_bk' value='${nullToEmpty(elem.email)}'></td>
		<td><input type='text' name='phone_number' value='${nullToEmpty(elem.phone_number)}'><input type='hidden' name='phone_number_bk' value='${nullToEmpty(elem.phone_number)}'></td>
		<td><input type='text' name='address' value='${nullToEmpty(elem.address)}'><input type='hidden' name='address_bk' value='${nullToEmpty(elem.address)}'></td>
		<td><input type='text' name='hire_date' value='${nullToEmpty(elem.hire_date)}'><input type='hidden' name='hire_date_bk' value='${nullToEmpty(elem.hire_date)}'></td>
		<td><input type='text' name='job_title' value='${nullToEmpty(elem.job_title)}'><input type='hidden' name='job_title_bk' value='${nullToEmpty(elem.job_title)}'></td>
		<td><input type='text' name='department_code' value='${nullToEmpty(elem.department_code)}'><input type='hidden' name='department_code_bk' value='${nullToEmpty(elem.department_code)}'></td>
		<td><input type='text' name='salary' value='${nullToEmpty(elem.salary)}'><input type='hidden' name='salary_bk' value='${nullToEmpty(elem.salary)}'></td>
		<td><input type='text' name='created_at' value='${nullToEmpty(elem.created_at)}' disabled></td>
		<td><input type='text' name='updated_at' value='${nullToEmpty(elem.updated_at)}' disabled></td>
    `;
    return tr;
}


/* セットアップ */
const getRows = async () => {
    document.getElementById('records').innerHTML = '';
    const rows = await api.get('employee');
    renderTbody(rows);
    addChangedAction('first_name');
    addChangedAction('last_name');
    addChangedAction('email');
    addChangedAction('phone_number');
    addChangedAction('address');
    addChangedAction('hire_date');
    addChangedAction('job_title');
    addChangedAction('department_code');
    addChangedAction('salary');
}


/* 一括更新 */
const putRows = async () => {
    let successCount = 0;
    let errorCount = 0;

    const id = document.getElementsByName('id');
    const first_name = document.getElementsByName('first_name');
    const last_name = document.getElementsByName('last_name');
    const email = document.getElementsByName('email');
    const phone_number = document.getElementsByName('phone_number');
    const address = document.getElementsByName('address');
    const hire_date = document.getElementsByName('hire_date');
    const job_title = document.getElementsByName('job_title');
    const department_code = document.getElementsByName('department_code');
    const salary = document.getElementsByName('salary');
    const created_at = document.getElementsByName('created_at');
    const updated_at = document.getElementsByName('updated_at');

    const first_name_bk = document.getElementsByName('first_name_bk');
    const last_name_bk = document.getElementsByName('last_name_bk');
    const email_bk = document.getElementsByName('email_bk');
    const phone_number_bk = document.getElementsByName('phone_number_bk');
    const address_bk = document.getElementsByName('address_bk');
    const hire_date_bk = document.getElementsByName('hire_date_bk');
    const job_title_bk = document.getElementsByName('job_title_bk');
    const department_code_bk = document.getElementsByName('department_code_bk');
    const salary_bk = document.getElementsByName('salary_bk');

    for (let i = 0; i < id.length; i++) {
        const rowMap = {
            'first_name': first_name[i],
            'last_name': last_name[i],
            'email': email[i],
            'phone_number': phone_number[i],
            'address': address[i],
            'hire_date': hire_date[i],
            'job_title': job_title[i],
            'department_code': department_code[i],
            'salary': salary[i],
        }

        const rowBkMap = {
            'first_name': first_name_bk[i],
            'last_name': last_name_bk[i],
            'email': email_bk[i],
            'phone_number': phone_number_bk[i],
            'address': address_bk[i],
            'hire_date': hire_date_bk[i],
            'job_title': job_title_bk[i],
            'department_code': department_code_bk[i],
            'salary': salary_bk[i],
        }

        //差分がある行のみ更新
        if (Object.keys(rowMap).some(key => rowMap[key].value !== rowBkMap[key].value)) {
            const requestBody = {
                id: parseIntOrReturnOriginal(id[i].value),
                first_name: first_name[i].value,
                last_name: last_name[i].value,
                email: email[i].value,
                phone_number: phone_number[i].value,
                address: address[i].value,
                hire_date: hire_date[i].value,
                job_title: job_title[i].value,
                department_code: department_code[i].value,
                salary: parseFloatOrReturnOriginal(salary[i].value),
            }

            try {
                const data = await api.put('employee', requestBody);
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
                created_at[i].value = data.created_at;
                updated_at[i].value = data.updated_at;

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
        'first_name': document.getElementById('first_name_new'),
        'last_name': document.getElementById('last_name_new'),
        'email': document.getElementById('email_new'),
        'phone_number': document.getElementById('phone_number_new'),
        'address': document.getElementById('address_new'),
        'hire_date': document.getElementById('hire_date_new'),
        'job_title': document.getElementById('job_title_new'),
        'department_code': document.getElementById('department_code_new'),
        'salary': document.getElementById('salary_new'),
    }

    if (Object.keys(rowMap).some(key => rowMap[key].value !== '')) {
        const requestBody = {
            first_name: rowMap.first_name.value,
            last_name: rowMap.last_name.value,
            email: rowMap.email.value,
            phone_number: rowMap.phone_number.value,
            address: rowMap.address.value,
            hire_date: rowMap.hire_date.value,
            job_title: rowMap.job_title.value,
            department_code: rowMap.department_code.value,
            salary: parseFloatOrReturnOriginal(rowMap.salary.value),
        }

        try {
            const data = await api.post('employee', requestBody);

            document.getElementById('new').remove();
            const tr = createTr(data);
            tr.addEventListener('change', changeAction);
            document.getElementById('records').appendChild(tr);
            document.getElementById('records').appendChild(createTrNew());

            renderMessage('登録', 1, true);
        } catch (e) {
            Object.keys(rowMap).forEach(key => {
                rowMap[key].classList.toggle('error', key === e.details.field || `employee.${key}` === e.details.column);
            });
            renderMessage('登録', 1, false);
        }
    }
}


/* 一括削除 */
const deleteRows = async () => {
    let rows = getDeleteTarget();
    let successCount = 0;
    let errorCount = 0;

    for (let row of rows) {
        try {
            await api.delete('employee', row);
            successCount += 1;
        } catch (e) {
            errorCount += 1;
        }
    }

    getRows();

    renderMessage('削除', successCount, true);
    renderMessage('削除', errorCount, false);
}
