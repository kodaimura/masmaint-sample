<?php

declare(strict_types=1);

namespace App\Application\Models\Entities;

use JsonSerializable;

class Employee implements JsonSerializable
{
    private int $id;

    private string $firstName;

    private ?string $lastName;

    private ?string $email;

    private ?string $phoneNumber;

    private ?string $address;

    private ?string $hireDate;

    private ?string $jobTitle;

    private ?string $departmentCode;

    private ?float $salary;

    private string $createdAt;

    private string $updatedAt;

    public function getId(): int
    {
        return $this->id;
    }

    public function getFirstName(): string
    {
        return $this->firstName;
    }

    public function getLastName(): ?string
    {
        return $this->lastName;
    }

    public function getEmail(): ?string
    {
        return $this->email;
    }

    public function getPhoneNumber(): ?string
    {
        return $this->phoneNumber;
    }

    public function getAddress(): ?string
    {
        return $this->address;
    }

    public function getHireDate(): ?string
    {
        return $this->hireDate;
    }

    public function getJobTitle(): ?string
    {
        return $this->jobTitle;
    }

    public function getDepartmentCode(): ?string
    {
        return $this->departmentCode;
    }

    public function getSalary(): ?float
    {
        return $this->salary;
    }

    public function getCreatedAt(): string
    {
        return $this->createdAt;
    }

    public function getUpdatedAt(): string
    {
        return $this->updatedAt;
    }

    public function setId($id)
    {
        if ($id === null || $id === "") {
            throw new \InvalidArgumentException("error: setId");
        } else if (filter_var($id, FILTER_VALIDATE_INT) !== false) {
            $this->id = (int) $id;
        } else {
            throw new \InvalidArgumentException("error: setId");
        }
    }

    public function setFirstName($firstName)
    {
        $this->firstName = $firstName;
    }

    public function setLastName($lastName)
    {
        $this->lastName = $lastName;
    }

    public function setEmail($email)
    {
        $this->email = $email;
    }

    public function setPhoneNumber($phoneNumber)
    {
        $this->phoneNumber = $phoneNumber;
    }

    public function setAddress($address)
    {
        $this->address = $address;
    }

    public function setHireDate($hireDate)
    {
        if ($hireDate === null || $hireDate === "") {
            $this->hireDate = null;
        } else {
            $this->hireDate = $hireDate;
        }
    }

    public function setJobTitle($jobTitle)
    {
        $this->jobTitle = $jobTitle;
    }

    public function setDepartmentCode($departmentCode)
    {
        $this->departmentCode = $departmentCode;
    }

    public function setSalary($salary)
    {
        if ($salary === null || $salary === "") {
            $this->salary = null;
        } else if (is_numeric($salary)) {
            $this->salary = (float) $salary;
        } else {
            throw new \InvalidArgumentException("error: setSalary");
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
            'id' => $this->id,
            'first_name' => $this->firstName,
            'last_name' => $this->lastName,
            'email' => $this->email,
            'phone_number' => $this->phoneNumber,
            'address' => $this->address,
            'hire_date' => $this->hireDate,
            'job_title' => $this->jobTitle,
            'department_code' => $this->departmentCode,
            'salary' => $this->salary,
            'created_at' => $this->createdAt,
            'updated_at' => $this->updatedAt,
        ];
    }
}
