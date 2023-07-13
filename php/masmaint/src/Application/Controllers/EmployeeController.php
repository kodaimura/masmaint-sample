<?php

declare(strict_types=1);

namespace App\Application\Controllers;

use App\Application\Controllers\BaseController;
use Psr\Log\LoggerInterface;
use Psr\Container\ContainerInterface;
use Psr\Http\Message\ResponseInterface as Response;
use Slim\Views\Twig;

class EmployeeController extends BaseController
{

    private Twig $twig;

    public function __construct(ContainerInterface $container, LoggerInterface $logger, Twig $twig)
    {
        parent::__construct($container, $logger);
        $this->twig = $twig;
    }

    public function employeePage($request, $response, $args): Response
    {
        echo "aaaaa";
        $response = $this->twig->render($response, 'employee.html', []);
        return $response;
    }

}