<?php

declare(strict_types=1);

namespace App\Application\Models\Daos;

use App\Application\Models\Entities\Employee;

interface EmployeeDao
{

    public function findAll(): array;

    public function create(Employee $employee): Employee;

    public function update(Employee $employee): Employee;

    public function delete(Employee $employee);
}
