<?php

declare(strict_types=1);

namespace App\Infrastructure\Persistence\Department;

use \PDO;
use Psr\Log\LoggerInterface;
use App\Domain\Department\Department;

class DepartmentRepositoryImpl implements DepartmentRepository
{

    private LoggerInterface $logger;
    private PDO $db;

    public function __construct(LoggerInterface $logger, PDO $db){
        $this->logger = $logger;
        $this->db = $db;
    }

    public function findAll(): array
    {
        $query = 
            "SELECT
                code
                ,name
                ,description
                ,manager_id
                ,location
                ,budget
                ,created_at
                ,updated_at
            FROM department
            ORDER BY code ASC";

        $stmt = $this->db->prepare($query);
        $stmt->execute();
        $result = $stmt->fetchAll(PDO::FETCH_ASSOC);

        $ret = [];

        foreach ($row as $result) {
            $x = new Department(
                $x['code'], 
                $x['name'],
                $x['description'],
                $x['manager_id'],
                $x['location'],
                $x['budget'],
                $x['created_at'],
                $x['updated_at'],
            );
            $ret[] = $x;
        }
        return $ret;
    }

    public function create(Department $department) {
        $query = 
            "INSERT INTO department (
                ,code
                ,name
                ,description
                ,manager_id
                ,location
                ,budget
            ) VALUES (
                :code 
                ,:name 
                ,:description
                ,:managerId
                ,:location
                ,:budget
            ) RETURNING
                code
                ,name
                ,description
                ,manager_id
                ,location
                ,budget
                ,created_at
                ,updated_at";

        $stmt = $this->db->prepare($query);
        $stmt->bindValue(':code', $department->getCode());
        $stmt->bindValue(':name', $department->getName());
        $stmt->bindValue(':description', $department->getDescription());
        $stmt->bindValue(':managerId', $department->getManagerId());
        $stmt->bindValue(':location', $department->getLocation());
        $stmt->bindValue(':budget', $department->getBudget());

        $result = $stmt->fetch(PDO::FETCH_ASSOC);

        return new Department(
                $result['code'], 
                $result['name'],
                $result['description'],
                $result['manager_id'],
                $result['location'],
                $result['budget'],
                $result['created_at'],
                $result['updated_at'],
            );
    }

    public function update(Department $department) {
        $query = 
            "UPDATE department SET
                name = :name
                ,description = :description
                ,manager_id = :managerId
                ,location = :location
                ,budget = :budget
            WHERE code = :code
            ) RETURNING
                code
                ,name
                ,description
                ,manager_id
                ,location
                ,budget
                ,created_at
                ,updated_at";

        $stmt = $this->db->prepare($query);
        $stmt->bindValue(':name', $department->getName());
        $stmt->bindValue(':description', $department->getDescription());
        $stmt->bindValue(':managerId', $department->getManagerId());
        $stmt->bindValue(':location', $department->getLocation());
        $stmt->bindValue(':budget', $department->getBudget());
        $stmt->bindValue(':code', $department->getCode());

        $result = $stmt->fetch(PDO::FETCH_ASSOC);

        return new Department(
                $result['code'], 
                $result['name'],
                $result['description'],
                $result['manager_id'],
                $result['location'],
                $result['budget'],
                $result['created_at'],
                $result['updated_at'],
            );
    }

    public function delete(Department $department) {
        $query = "DELETE FROM department WHERE code = :code";

        $stmt = $this->db->prepare($query);
        $stmt->bindValue(':code', $department->getCode());
        $stmt->execute(PDO::FETCH_ASSOC);
    }
}