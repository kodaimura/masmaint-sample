<?php

declare(strict_types=1);

use App\Application\Models\Daos\DepartmentDao;
use App\Application\Models\Daos\EmployeeDao;
use App\Application\Models\DaoImpls\DepartmentDaoImpl2;
use App\Application\Models\DaoImpls\EmployeeDaoImpl2;
use DI\ContainerBuilder;

return function (ContainerBuilder $containerBuilder) {

    $containerBuilder->addDefinitions([
        DepartmentDao::class => \DI\autowire(DepartmentDaoImpl2::class),
        EmployeeDao::class => \DI\autowire(EmployeeDaoImpl2::class),
    ]);
};
