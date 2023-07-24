<?php

declare(strict_types=1);

namespace App\Application\Models\DaoImpls;

use App\Application\Models\Daos\EmployeeDao;
use App\Application\Models\Entities\Employee;

use \PDOException;
use \PDO;
use Psr\Log\LoggerInterface;

// RETURNING が使えない場合
class EmployeeDaoImpl2 implements EmployeeDao
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
            throw $e;
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

    public function findOne(Employee $employee): Employee
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
            WHERE id = :id";

        try {
            $stmt = $this->db->prepare($query);
            $stmt->bindValue(':id', $employee->getId());
            $stmt->execute();
        } catch (PDOException $e) {
            $this->logger->error($e->getMessage());
            throw $e;
        }

        $result = $stmt->fetch(PDO::FETCH_ASSOC);

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
            )";
        
        try {
            $stmt = $this->db->prepare($query);
            $stmt->bindValue(':firstName', $employee->getFirstName());
            $stmt->bindValue(':lastName', $employee->getLastName());
            $stmt->bindValue(':email', $employee->getEmail());
            $stmt->bindValue(':phoneNumber', $employee->getPhoneNumber());
            $stmt->bindValue(':address', $employee->getAddress());
            $stmt->bindValue(':hireDate', $employee->getHireDate());
            $stmt->bindValue(':jobTitle', $employee->getJobTitle());
            $stmt->bindValue(':departmentCode', $employee->getDepartmentCode());
            $stmt->bindValue(':salary', $employee->getSalary());
            $stmt->execute();
        } catch (PDOException $e) {
            $this->logger->error($e->getMessage());
            throw $e;
        }

        $lastInsertedId = $this->db->lastInsertId();
        $employee->setId($lastInsertedId);

        return $this->findOne($employee);
    }

    public function update(Employee $employee): Employee 
    {
        $query = 
            "UPDATE employee
             SET
                first_name = :firstName
                ,last_name = :lastName
                ,email = :email
                ,phone_number = :phoneNumber
                ,address = :address
                ,hire_date = :hireDate
                ,job_title = :jobTitle
                ,department_code = :departmentCode
                ,salary = :salary
             WHERE id = :id";

        try {
            $stmt = $this->db->prepare($query);
            $stmt->bindValue(':firstName', $employee->getFirstName());
            $stmt->bindValue(':lastName', $employee->getLastName());
            $stmt->bindValue(':email', $employee->getEmail());
            $stmt->bindValue(':phoneNumber', $employee->getPhoneNumber());
            $stmt->bindValue(':address', $employee->getAddress());
            $stmt->bindValue(':hireDate', $employee->getHireDate());
            $stmt->bindValue(':jobTitle', $employee->getJobTitle());
            $stmt->bindValue(':departmentCode', $employee->getDepartmentCode());
            $stmt->bindValue(':salary', $employee->getSalary());
            $stmt->bindValue(':id', $employee->getId());
            $stmt->execute();
        } catch (PDOException $e) {
            $this->logger->error($e->getMessage());
            throw $e;
        }

        return $this->findOne($employee);
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
            throw $e;
        }

        return;
    }
}