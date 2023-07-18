<?php

declare(strict_types=1);

namespace App\Domain\Department;

use JsonSerializable;

class Department implements JsonSerializable
{
    private string $code;

    private string $name;

    private ?string $description;

    private ?int $manager_id;

    private ?string $location;

    private float $budget;

    private ?string $created_at;

    private ?string $updated_at;

/*
    public function __construct($code, $name, $description, $managerId, $location, $budget, $createdAt, $updatedAt)
    {
        $this->code = $code;
        $this->name = $name;
        $this->description = $description;
        $this->managerId = $managerId;
        $this->location = $location;
        $this->budget = $budget;
        $this->createdAt = $createdAt;
        $this->updatedAt = $updatedAt;
    }
*/
    public function getCode(): string
    {
        return $this->code;
    }

    public function getName(): string
    {
        return $this->name;
    }

    public function getDescription(): ?string
    {
        return $this->description;
    }

    public function getManagerId(): ?int
    {
        return $this->manager_id;
    }

    public function getLocation(): ?string
    {
        return $this->location;
    }

    public function getBudget(): float
    {
        return $this->budget;
    }

    public function getCreatedAt(): ?string
    {
        return $this->created_at;
    }

    public function getUpdatedAt(): ?string
    {
        return $this->updated_at;
    }

    //json_encode()でエンコードされるときに呼ばれる
    #[\ReturnTypeWillChange]
    public function jsonSerialize(): array
    {
        return [
            'code' => $this->code,
            'name' => $this->name,
            'description' => $this->description,
            'manager_id' => $this->manager_id,
            'location' => $this->location,
            'budget' => $this->budget,
            'created_at' => $this->created_at,
            'updated_at' => $this->updated_at,
        ];
    }
}
