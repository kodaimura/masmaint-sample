<?php

declare(strict_types=1);

namespace App\Domain\Department;

interface DepartmentRepository
{

    public function findAll(): array;

    public function create(Department $department): Department;

    public function update(Department $department): Department;

    public function delete(Department $department);
}
