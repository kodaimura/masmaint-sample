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

    private float $budget;

    private ?string $createdAt;

    private ?string $updatedAt;

    public function __construct(string $code, string $name, string $description, ?int $managerId, ?string $location, float $budget)
    {
        $this->code = $code;
        $this->name = $name;
        $this->description = $description;
        $this->managerId = $managerId;
        $this->location = $location;
        $this->budget = $budget;
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

    public function getBudget(): float
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

    #[\ReturnTypeWillChange]
    public function jsonSerialize(): array
    {
        return [
            'code' => $this->code,
            'name' => $this->name,
            'description' => $this->description,
            'managerId' => $this->managerId,
            'location' => $this->location,
            'budget' => $this->budget,
            'createdAt' => $this->createdAt,
            'updatedAt' => $this->updatedAt,
        ];
    }
}
