<?php

declare(strict_types=1);

namespace App\Application\Controllers;

use App\Application\Controllers\BaseController;
use App\Application\Services\DepartmentService ;
use Psr\Log\LoggerInterface;
use Psr\Container\ContainerInterface;
use Psr\Http\Message\ResponseInterface as Response;
use Slim\Views\Twig;

class DepartmentController extends BaseController
{

    private Twig $twig;
    protected DepartmentService $departmentService;

    public function __construct(ContainerInterface $container, LoggerInterface $logger, Twig $twig, DepartmentService $departmentService)
    {
        parent::__construct($container, $logger);
        $this->twig = $twig;
        $this->departmentService = $departmentService;
    }

    public function departmentPage($request, $response, $args): Response
    {
        $response = $this->twig->render($response, 'department.html', []);
        return $response;
    }

    public function getDepartments($request, $response, $args): Response
    {
        $departments = $this->departmentService->getAll();
        $response->getBody()->write(json_encode($departments));
        return $response->withHeader('Content-Type', 'application/json');
    }

    public function postDepartment($request, $response, $args): Response
    {
        $data = $request->getParsedBody();
        try {
            $department = $this->departmentService->create($data);
            $response->getBody()->write(json_encode($department));

        } catch (Exception $e) {
            $this->logger->error($e->getMessage());
            return $response
            ->withHeader('Content-Type', 'application/json')
            ->withStatus(500);
        }
        $this->logger->error("adajla");
        return $response->withHeader('Content-Type', 'application/json');
    }

}