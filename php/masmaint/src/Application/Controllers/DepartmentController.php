<?php

declare(strict_types=1);

namespace App\Application\Controllers;

use App\Application\Controllers\BaseController;
use App\Domain\Department\DepartmentRepository;
use Psr\Log\LoggerInterface;
use Psr\Container\ContainerInterface;
use Psr\Http\Message\ResponseInterface as Response;
use Slim\Views\Twig;

class DepartmentController extends BaseController
{

    private Twig $twig;
    protected DepartmentRepository $departmentRep;

    public function __construct(ContainerInterface $container, LoggerInterface $logger, Twig $twig, DepartmentRepository $departmentRepository)
    {
        parent::__construct($container, $logger);
        $this->twig = $twig;
        $this->departmentRepository = $departmentRepository;
    }

    public function departmentPage($request, $response, $args): Response
    {
        $response = $this->twig->render($response, 'department.html', []);
        return $response;
    }

    public function getDepartments($request, $response, $args): Response
    {

        $departments = $this->departmentRepository->findAll();
        $response->getBody()->write(json_encode($departments));
        return $response->withHeader('Content-Type', 'application/json');
    }

}