<?php

declare(strict_types=1);

namespace App\Application\Controllers;

use App\Application\Controllers\BaseController;
use App\Application\Services\EmployeeService ;
use Psr\Log\LoggerInterface;
use Psr\Container\ContainerInterface;
use Psr\Http\Message\ResponseInterface as Response;
use Slim\Views\Twig;

class EmployeeController extends BaseController
{

    private Twig $twig;
    protected EmployeeService $employeeService;

    public function __construct(ContainerInterface $container, LoggerInterface $logger, Twig $twig, EmployeeService $employeeService)
    {
        parent::__construct($container, $logger);
        $this->twig = $twig;
        $this->employeeService = $employeeService;
    }

    public function employeePage($request, $response, $args): Response
    {
        $response = $this->twig->render($response, 'employee.html', []);
        return $response;
    }

    public function getEmployee($request, $response, $args): Response
    {
        $employees = $this->employeeService->getAll();
        $response->getBody()->write(json_encode($employees));
        return $response->withHeader('Content-Type', 'application/json');
    }

    public function postEmployee($request, $response, $args): Response
    {
        $data = $request->getParsedBody();
        try {
            $employee = $this->employeeService->create($data);
            $response->getBody()->write(json_encode($employee));

        } catch (\InvalidArgumentException $e) {
            $this->logger->debug($e->getMessage());
            return $response
            ->withHeader('Content-Type', 'application/json')
            ->withStatus(400);

        } catch (\Exception $e) {
            $this->logger->error($e->getMessage());
            return $response
            ->withHeader('Content-Type', 'application/json')
            ->withStatus(500);
        }
        return $response->withHeader('Content-Type', 'application/json');
    }

    public function putEmployee($request, $response, $args): Response
    {
        $data = $request->getParsedBody();
        try {
            $employee = $this->employeeService->update($data);
            $response->getBody()->write(json_encode($employee));

        } catch (\InvalidArgumentException $e) {
            $this->logger->debug($e->getMessage());
            return $response
            ->withHeader('Content-Type', 'application/json')
            ->withStatus(400);

        } catch (\Exception $e) {
            $this->logger->error($e->getMessage());
            return $response
            ->withHeader('Content-Type', 'application/json')
            ->withStatus(500);
        }
        return $response->withHeader('Content-Type', 'application/json');
    }

    public function deleteEmployee($request, $response, $args): Response
    {
        $data = $request->getParsedBody();
        try {
            $this->employeeService->delete($data);
            
        } catch (\InvalidArgumentException $e) {
            $this->logger->debug($e->getMessage());
            return $response
            ->withHeader('Content-Type', 'application/json')
            ->withStatus(400);

        } catch (\Exception $e) {
            $this->logger->error($e->getMessage());
            return $response
            ->withHeader('Content-Type', 'application/json')
            ->withStatus(500);
        }
        return $response->withHeader('Content-Type', 'application/json');
    }
}