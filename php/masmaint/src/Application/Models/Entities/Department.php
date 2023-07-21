<?php

declare(strict_types=1);

namespace App\Application\Models\Entities;

use JsonSerializable;

class Department implements JsonSerializable
{
    private string $code;

    private string $name;

    private ?string $description;

    private ?int $managerId;

    private ?string $location;

    private float $budget;

    private string $createdAt;

    private string $updatedAt;

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

    public function getCreatedAt(): string
    {
        return $this->createdAt;
    }

    public function getUpdatedAt(): string
    {
        return $this->updatedAt;
    }

    public function setCode($code)
    {
        $this->code = $code;
    }

    public function setName($name)
    {
        $this->name = $name;
    }

    public function setDescription($description)
    {
        $this->description = $description;
    }

    public function setManagerId($managerId)
    {
        if ($managerId === null || $managerId === "") {
            $this->managerId = null;
        } else if (filter_var($managerId, FILTER_VALIDATE_INT) !== false) {
            $this->managerId = (int) $managerId;
        } else {
            throw new \InvalidArgumentException("error: setManagerId");
        }
    }

    public function setLocation($location)
    {
        $this->location = $location;
    }

    public function setBudget($budget)
    {
        if ($budget === null || $budget === "") {
            throw new \InvalidArgumentException("error: setBudget");
        } else if (is_numeric($budget)) {
            $this->budget = (float) $budget;
        } else {
            throw new \InvalidArgumentException("error: setBudget");
        }
    }

    public function setCreatedAt($createdAt)
    {
        $this->createdAt = $createdAt;
    }

    public function setUpdatedAt($updatedAt)
    {
        $this->updatedAt = $updatedAt;
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
