<?php

declare(strict_types=1);

namespace App\Domain\Department;

use JsonSerializable;

class Department implements JsonSerializable
{
    private string $code;

    private string $name;

    private ?string $description;

    private ?int $managerId;

    private ?string $location;

    private string $budget;

    private ?string $createdAt;

    private ?string $updatedAt;

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
        return $this->managerId;
    }

    public function getLocation(): ?string
    {
        return $this->location;
    }

    public function getBudget(): string
    {
        return $this->budget;
    }

    public function getCreatedAt(): ?string
    {
        return $this->createdAt;
    }

    public function getUpdatedAt(): ?string
    {
        return $this->updatedAt;
    }

    //json_encode()でエンコードされるときに呼ばれる
    #[\ReturnTypeWillChange]
    public function jsonSerialize(): array
    {
        return [
            'code' => $this->code,
            'name' => $this->name,
            'description' => $this->description,
            'manager_id' => $this->managerId,
            'location' => $this->location,
            'budget' => $this->budget,
            'created_at' => $this->createdAt,
            'updated_at' => $this->updatedAt,
        ];
    }
}
