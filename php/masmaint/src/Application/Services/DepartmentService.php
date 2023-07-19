<?php

declare(strict_types=1);

namespace App\Application\Services;

use App\Domain\Department\DepartmentRepository;
use App\Domain\Department\Department;
use Psr\Log\LoggerInterface;
use Psr\Container\ContainerInterface;
use Psr\Http\Message\ResponseInterface as Response;
use Slim\Views\Twig;

class DepartmentService extends BaseService
{

    protected DepartmentRepository $departmentRep;

    public function __construct(ContainerInterface $container, LoggerInterface $logger, DepartmentRepository $departmentRepository)
    {
        parent::__construct($container, $logger);
        $this->departmentRepository = $departmentRepository;
    }

    public function getAll(): array
    {
        return $this->departmentRepository->findAll();
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

        return $this->departmentRepository->create($department);
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

        return $this->departmentRepository->update($department);
    }

}