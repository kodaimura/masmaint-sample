<?php

declare(strict_types=1);

namespace App\Application\Models\DaoImpls;

use App\Application\Models\Daos\DepartmentDao;
use App\Application\Models\Entities\Department;

use \PDOException;
use \PDO;
use Psr\Log\LoggerInterface;

class DepartmentDaoImpl implements DepartmentDao
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
            $x = new Department();
            $x->setCode($row['code']);
            $x->setName($row['name']);
            $x->setDescription($row['description']);
            $x->setManagerId($row['manager_id']);
            $x->setLocation($row['location']);
            $x->setBudget($row['budget']);
            $x->setCreatedAt($row['created_at']);
            $x->setUpdatedAt($row['updated_at']);

            $ret[] = $x;
        }

        return $ret;
    }

    public function create(Department $department): Department 
    {
        $query = 
            "INSERT INTO department (
                code
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
        
        try {
            $stmt = $this->db->prepare($query);
            $stmt->bindValue(':code', $department->getCode());
            $stmt->bindValue(':name', $department->getName());
            $stmt->bindValue(':description', $department->getDescription());
            $stmt->bindValue(':managerId', $department->getManagerId());
            $stmt->bindValue(':location', $department->getLocation());
            $stmt->bindValue(':budget', $department->getBudget());
            $stmt->execute();
        } catch (PDOException $e) {
            $this->logger->error($e->getMessage());
        }

        $result = $stmt->fetch(PDO::FETCH_ASSOC);

        $ret = new Department();
        $ret->setCode($result['code']);
        $ret->setName($result['name']);
        $ret->setDescription($result['description']);
        $ret->setManagerId($result['manager_id']);
        $ret->setLocation($result['location']);
        $ret->setBudget($result['budget']);
        $ret->setCreatedAt($result['created_at']);
        $ret->setUpdatedAt($result['updated_at']);

        return $ret;
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
            RETURNING
                code
                ,name
                ,description
                ,manager_id
                ,location
                ,budget
                ,created_at
                ,updated_at";

        try {
            $stmt = $this->db->prepare($query);
            $stmt->bindValue(':name', $department->getName());
            $stmt->bindValue(':description', $department->getDescription());
            $stmt->bindValue(':managerId', $department->getManagerId());
            $stmt->bindValue(':location', $department->getLocation());
            $stmt->bindValue(':budget', $department->getBudget());
            $stmt->bindValue(':code', $department->getCode());
            $stmt->execute();
        } catch (PDOException $e) {
            $this->logger->error($e->getMessage());
        }

        $result = $stmt->fetch(PDO::FETCH_ASSOC);

        $ret = new Department();
        $ret->setCode($result['code']);
        $ret->setName($result['name']);
        $ret->setDescription($result['description']);
        $ret->setManagerId($result['manager_id']);
        $ret->setLocation($result['location']);
        $ret->setBudget($result['budget']);
        $ret->setCreatedAt($result['created_at']);
        $ret->setUpdatedAt($result['updated_at']);

        return $ret;
    }

    public function delete(Department $department) 
    {
        $query = "DELETE FROM department WHERE code = :code";

        try {
            $stmt = $this->db->prepare($query);
            $stmt->bindValue(':code', $department->getCode());
            $stmt->execute();
        } catch (PDOException $e) {
            $this->logger->error($e->getMessage());
        }

        return;
    }
}