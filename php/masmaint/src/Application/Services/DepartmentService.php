<?php

declare(strict_types=1);

namespace App\Application\Services;

use App\Domain\Department\DepartmentRepository;
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

    public function getDepartments(): array
    {
        return $this->departmentRepository->findAll();
    }

}