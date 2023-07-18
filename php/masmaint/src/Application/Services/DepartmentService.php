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
        $this->logger->error("1");
        $department->setCode($data['code']);
        $this->logger->error("1");
        $department->setName($data['name']);
        $this->logger->error("2");
        $department->setDescription($data['description']);
        $this->logger->error("3");
        $department->setManagerId($data['manager_id']);
        $this->logger->error("4");
        $department->setLocation($data['location']);
        $this->logger->error("5");
        $department->setBudget($data['budget']);
        $this->logger->error("6");

        return $this->departmentRepository->create($department);
    }

}