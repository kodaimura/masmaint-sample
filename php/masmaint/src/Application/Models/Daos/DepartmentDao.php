<?php

declare(strict_types=1);

namespace App\Application\Models\Daos;

use App\Application\Models\Entities\Department;

interface DepartmentDao
{

    public function findAll(): array;

    public function create(Department $department): Department;

    public function update(Department $department): Department;

    public function delete(Department $department);
}
