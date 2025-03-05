<?php

declare(strict_types=1);

namespace App\Core;

final class InMemoryUsersClient implements UsersClientInterface
{
    /** @var array<int, UserDetails> */
    private array $users;

    public function getUserDetails(int $id): UserDetails
    {
       return $this->users[$id];
    }

    public function addUser(int $id, UserDetails $userDetails): void
    {
        $this->users[$id] = $userDetails;
    }
}
