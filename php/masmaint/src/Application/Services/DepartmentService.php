<?php

declare(strict_types=1);

namespace App\Application\Services;

use App\Application\Models\Daos\DepartmentDao;
use App\Application\Models\Entities\Department;
use Psr\Log\LoggerInterface;
use Psr\Container\ContainerInterface;
use Psr\Http\Message\ResponseInterface as Response;
use Slim\Views\Twig;

class DepartmentService extends BaseService
{

    protected DepartmentDao $departmentDao;

    public function __construct(ContainerInterface $container, LoggerInterface $logger, DepartmentDao $departmentDao)
    {
        parent::__construct($container, $logger);
        $this->departmentDao = $departmentDao;
    }

    public function getAll(): array
    {
        return $this->departmentDao->findAll();
    }

    public function create($data): Department
    {
        $department = new Department();
        $department->setCode($data['code']);
        $department->setName($data['name']);
        $department->setDescription($data['description']);
        $department->setManagerId($data['manager_id']);
        $department->setLocation($data['location']);
        $department->setBudget($data['budget']);

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

        return $this->departmentDao->update($department);
    }

    public function delete($data)
    {
        $department = new Department();
        $department->setCode($data['code']);

        $this->departmentDao->delete($department);
        return;
    }

}