<?php

declare(strict_types=1);

use App\Application\Controllers\IndexController;
use App\Application\Controllers\EmployeeController;
use App\Application\Controllers\DepartmentController;
use Psr\Http\Message\ResponseInterface as Response;
use Psr\Http\Message\ServerRequestInterface as Request;
use Slim\App;
use Slim\Interfaces\RouteCollectorProxyInterface as Group;

return function (App $app) {
    $app->options('/{routes:.*}', function (Request $request, Response $response) {
        // CORS Pre-Flight OPTIONS Request Handler
        return $response;
    });

    $app->get('/', function (Request $request, Response $response) {
        $response->getBody()->write('Hello world!');
        return $response;
    });

    $app->group('/mastertables', function (Group $group) {
        $group->get('', IndexController::class. ':indexPage');
        $group->get('/', IndexController::class. ':indexPage');

        $group->get('/employee', EmployeeController::class. ':employeePage');
        $group->get('/api/employee', EmployeeController::class. ':getEmployee');
        $group->post('/api/employee', EmployeeController::class. ':postEmployee');
        $group->put('/api/employee', EmployeeController::class. ':putEmployee');
        $group->delete('/api/employee', EmployeeController::class. ':deleteEmployee');

        $group->get('/department', DepartmentController::class. ':departmentPage');
        $group->get('/api/department', DepartmentController::class. ':getDepartment');
        $group->post('/api/department', DepartmentController::class. ':postDepartment');
        $group->put('/api/department', DepartmentController::class. ':putDepartment');
        $group->delete('/api/department', DepartmentController::class. ':deleteDepartment');
    });
};
