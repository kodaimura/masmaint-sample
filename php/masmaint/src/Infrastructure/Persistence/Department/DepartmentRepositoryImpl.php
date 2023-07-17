<?php

declare(strict_types=1);

namespace App\Infrastructure\Persistence\Department;
use App\Domain\Department\DepartmentRepository;

use \PDOException;
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

        try {
            $stmt = $this->db->prepare($query);
            $stmt->execute();
        } catch (PDOException $e) {
            $this->logger->error($e->getMessage());
        }

        $result = $stmt->fetchAll(PDO::FETCH_ASSOC);
        $ret = [];
        foreach ($result as $row) {
            $x = new Department(
                $row['code'], 
                $row['name'],
                $row['description'],
                $row['manager_id'],
                $row['location'],
                $row['budget'],
                $row['created_at'],
                $row['updated_at']
            );
            
            $ret[] = $x;
        }

        return $ret;
    }

    public function create(Department $department): Department 
    {
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

        $stmt->execute();
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

    public function update(Department $department): Department 
    {
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

        $stmt->execute();
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

    public function delete(Department $department) 
    {
        $query = "DELETE FROM department WHERE code = :code";

        $stmt = $this->db->prepare($query);
        $stmt->bindValue(':code', $department->getCode());
        $stmt->execute(PDO::FETCH_ASSOC);
    }
}