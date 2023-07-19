<?php

declare(strict_types=1);

namespace App\Application\Services;

use App\Application\Models\Daos\EmployeeDao;
use App\Application\Models\Entities\Employee;
use Psr\Log\LoggerInterface;
use Psr\Container\ContainerInterface;
use Psr\Http\Message\ResponseInterface as Response;
use Slim\Views\Twig;

class EmployeeService extends BaseService
{

    protected EmployeeDao $employeeDao;

    public function __construct(ContainerInterface $container, LoggerInterface $logger, EmployeeDao $employeeDao)
    {
        parent::__construct($container, $logger);
        $this->employeeDao = $employeeDao;
    }

    public function getAll(): array
    {
        return $this->employeeDao->findAll();
    }

    public function create($data): Department
    {
        $employee = new Employee();
        $employee->setFirstName($data['first_name']);
        $employee->setLastName($data['last_name']);
        $employee->setEmail($data['email']);
        $employee->setPhoneNumber($data['phone_number']);
        $employee->setAddress($data['address']);
        $employee->setHireDate($data['hire_date']);
        $employee->setJobTitle($data['job_title']);
        $employee->setDepartmentCode($data['department_code']);
        $employee->setSalary($data['salary']);

        return $this->departmentDao->create($department);
    }

    public function update($data): Department
    {
        $department = new Department();
        $department->setCode($data['code']);
        $department->setName($data['name']);
        $department->setDescription($data['description']);
        $department->setManagerId($data['manager_id']);
        $department->setLocation($data['location']);
        $department->setBudget($data['budget']);

        return $this->employeeDao->update($employee);
    }

    public function delete($data)
    {
        $employee = new Employee();
        $employee->setId($data['id']);

        $this->employeeDao->delete($employee);
        return;
    }

}