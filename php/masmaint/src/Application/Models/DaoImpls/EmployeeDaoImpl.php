<?php

declare(strict_types=1);

namespace App\Application\Models\DaoImpls;

use App\Application\Models\Daos\EmployeeDao;
use App\Application\Models\Entities\Employee;

use \PDOException;
use \PDO;
use Psr\Log\LoggerInterface;

class EmployeeDaoImpl implements EmployeeDao
{

    private LoggerInterface $logger;
    private PDO $db;

    public function __construct(LoggerInterface $logger, PDO $db){
        $this->logger = $logger;
        $this->db = $db;
    }

    public function findAll(): array
    {
        $query = 
            "SELECT
                id
                ,first_name
                ,last_name
                ,email
                ,phone_number
                ,address
                ,hire_date
                ,job_title
                ,department_code
                ,salary
                ,created_at
                ,updated_at
            FROM employee
            ORDER BY id ASC";

        try {
            $stmt = $this->db->prepare($query);
            $stmt->execute();
        } catch (PDOException $e) {
            $this->logger->error($e->getMessage());
        }

        $result = $stmt->fetchAll(PDO::FETCH_ASSOC);
        $ret = [];
        foreach ($result as $row) {
            $x = new Employee();
            $x->setId($row['id']);
            $x->setFirstName($row['first_name']);
            $x->setLastName($row['last_name']);
            $x->setEmail($row['email']);
            $x->setPhoneNumber($row['phone_number']);
            $x->setAddress($row['address']);
            $x->setHireDate($row['hire_date']);
            $x->setJobTitle($row['job_title']);
            $x->setDepartmentCode($row['department_code']);
            $x->setSalary($row['salary']);
            $x->setCreatedAt($row['created_at']);
            $x->setUpdatedAt($row['updated_at']);

            $ret[] = $x;
        }

        return $ret;
    }

    public function create(Employee $employee): Employee
    {
        $query = 
            "INSERT INTO employee (
                first_name
                ,last_name
                ,email
                ,phone_number
                ,address
                ,hire_date
                ,job_title
                ,department_code
                ,salary
            ) VALUES (
                :firstName
                ,:lastName
                ,:email
                ,:phoneNumber
                ,:address
                ,:hireDate
                ,:jobTitle
                ,:departmentCode
                ,:salary
            ) RETURNING
                id
                ,first_name
                ,last_name
                ,email
                ,phone_number
                ,address
                ,hire_date
                ,job_title
                ,department_code
                ,salary
                ,created_at
                ,updated_at";
        
        try {
            $stmt = $this->db->prepare($query);
            $stmt->bindValue(':firstName', $employee->getFirstName());
            $stmt->bindValue(':lastName', $employee->getLastName(), PDO::PARAM_NULL);
            $stmt->bindValue(':email', $employee->getEmail(), PDO::PARAM_NULL);
            $stmt->bindValue(':phoneNumber', $employee->getPhoneNumber(), PDO::PARAM_NULL);
            $stmt->bindValue(':address', $employee->getAddress(), PDO::PARAM_NULL);
            $stmt->bindValue(':hireDate', $employee->getHireDate(), PDO::PARAM_NULL);
            $stmt->bindValue(':jobTitle', $employee->getJobTitle(), PDO::PARAM_NULL);
            $stmt->bindValue(':departmentCode', $employee->getDepartmentCode(), PDO::PARAM_NULL);
            $stmt->bindValue(':salary', $employee->getSalary());
            $stmt->execute();
        } catch (PDOException $e) {
            $this->logger->error($e->getMessage());
        }

        $result = $stmt->fetch(PDO::FETCH_ASSOC);
        $ret = [];

        $ret = new Employee();
        $ret->setId($result['id']);
        $ret->setFirstName($result['first_name']);
        $ret->setLastName($result['last_name']);
        $ret->setEmail($result['email']);
        $ret->setPhoneNumber($result['phone_number']);
        $ret->setAddress($result['address']);
        $ret->setHireDate($result['hire_date']);
        $ret->setJobTitle($result['job_title']);
        $ret->setDepartmentCode($result['department_code']);
        $ret->setSalary($result['salary']);
        $ret->setCreatedAt($result['created_at']);
        $ret->setUpdatedAt($result['updated_at']);

        return $ret;
    }

    public function update(Department $department): Department 
    {
        $query = 
            "UPDATE employee
             SET
                first_name = ?
                ,last_name = ?
                ,email = ?
                ,phone_number = ?
                ,address = ?
                ,hire_date = ?
                ,job_title = ?
                ,department_code = ?
                ,salary = ?
             WHERE id = ?
             RETURNING 
                id
                ,first_name
                ,last_name
                ,email
                ,phone_number
                ,address
                ,hire_date
                ,job_title
                ,department_code
                ,salary
                ,created_at
                ,updated_at";

        try {
            $stmt = $this->db->prepare($query);
            $stmt->bindValue(':firstName', $employee->getFirstName());
            $stmt->bindValue(':lastName', $employee->getLastName(), PDO::PARAM_NULL);
            $stmt->bindValue(':email', $employee->getEmail(), PDO::PARAM_NULL);
            $stmt->bindValue(':phoneNumber', $employee->getPhoneNumber(), PDO::PARAM_NULL);
            $stmt->bindValue(':address', $employee->getAddress(), PDO::PARAM_NULL);
            $stmt->bindValue(':hireDate', $employee->getHireDate(), PDO::PARAM_NULL);
            $stmt->bindValue(':jobTitle', $employee->getJobTitle(), PDO::PARAM_NULL);
            $stmt->bindValue(':departmentCode', $employee->getDepartmentCode(), PDO::PARAM_NULL);
            $stmt->bindValue(':salary', $employee->getSalary());
            $stmt->execute();
        } catch (PDOException $e) {
            $this->logger->error($e->getMessage());
        }

        $result = $stmt->fetch(PDO::FETCH_ASSOC);
        $ret = [];

        $ret = new Employee();
        $ret->setId($result['id']);
        $ret->setFirstName($result['first_name']);
        $ret->setLastName($result['last_name']);
        $ret->setEmail($result['email']);
        $ret->setPhoneNumber($result['phone_number']);
        $ret->setAddress($result['address']);
        $ret->setHireDate($result['hire_date']);
        $ret->setJobTitle($result['job_title']);
        $ret->setDepartmentCode($result['department_code']);
        $ret->setSalary($result['salary']);
        $ret->setCreatedAt($result['created_at']);
        $ret->setUpdatedAt($result['updated_at']);

        return $ret;
    }

    public function delete(Employee $employee) 
    {
        $query = "DELETE FROM employee WHERE id = :id";

        try {
            $stmt = $this->db->prepare($query);
            $stmt->bindValue(':id', $employee->getId());
            $stmt->execute();
        } catch (PDOException $e) {
            $this->logger->error($e->getMessage());
        }

        return;
    }
}