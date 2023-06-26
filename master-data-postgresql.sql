
INSERT INTO department (
  name,description,manager_id,location,budget 
) VALUES('テスト部署','ああああああ','1','東京都葛飾区',1000.5);
     
INSERT INTO department (
  name,description,manager_id,location,budget
) VALUES('テスト部署２','いいいいいいい','2','東京都江戸川区',9999);
     
     
INSERT INTO employee (
  first_name,last_name,email,phone_number,address,hire_date,job_title,department_id,salary
) VALUES('テスト','太郎','test@gmail.com','123456789','東京都葛飾区','20230626','プログラマ',1,500);
     
INSERT INTO employee (
  first_name,last_name,email,phone_number,address,hire_date,job_title,department_id,salary
) VALUES('山田','テスト','test2@gmail.com','987654321','東京都江戸川区','20230627','SE',2,600);