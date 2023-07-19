<?php

declare(strict_types=1);

use App\Application\Models\Daos\DepartmentDao;
use App\Application\Models\Daos\EmployeeDao;
use App\Application\Models\DaoImpls\DepartmentDaoImpl;
use App\Application\Models\DaoImpls\EmployeeDaoImpl;
use DI\ContainerBuilder;

return function (ContainerBuilder $containerBuilder) {
    // Here we map our UserRepository interface to its in memory implementation
    $containerBuilder->addDefinitions([
        DepartmentDao::class => \DI\autowire(DepartmentDaoImpl::class),
        EmployeeDao::class => \DI\autowire(EmployeeDaoImpl::class),
    ]);
};
